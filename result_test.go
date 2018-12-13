package neox

import (
	"testing"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/stretchr/testify/mock"
)

type user struct {
	Name     string  `db:"user_name"`
	Age      uint    `db:"user_age"`
	Power    float32 `db:"user_strength"`
	IsActive bool    `db:"is_active"`
	Avatar   rune    `db:"avatar_icon"`
}

var (
	// test one
	t1 = user{
		Name:     "Yolanda Erasmus",
		Age:      17,
		Power:    65.234,
		IsActive: true,
		Avatar:   0x1F607,
	}

	t1assert = func(u *user) func() (*user, bool) {
		return func() (*user, bool) {
			passed := (u.Name == t1.Name &&
				u.Age == t1.Age &&
				u.Power == t1.Power &&
				u.IsActive == t1.IsActive &&
				u.Avatar == t1.Avatar)
			return &t1, passed
		}
	}

	t1mock = func() neo4j.Result {
		record := new(mrec)
		record.On("Get", "user_name").Return(t1.Name, true)
		record.On("Get", "user_age").Return(t1.Age, true)
		record.On("Get", "user_strength").Return(t1.Power, true)
		record.On("Get", "is_active").Return(t1.IsActive, true)
		record.On("Get", "avatar_icon").Return(t1.Avatar, true)

		result := new(mres)
		result.On("Record").Return(record)
		result.On("Err").Return(nil)
		return result
	}

	t2 = user{
		Name:     "",
		Age:      17,
		Power:    0,
		IsActive: true,
		Avatar:   0x1F607,
	}

	t2mock = func() neo4j.Result {
		record := new(mrec)
		record.On("Get", "user_name").Return(0x02f, true)
		record.On("Get", "user_age").Return(t2.Age, true)
		record.On("Get", "user_strength").Return(12312, true)
		record.On("Get", "is_active").Return(t2.IsActive, true)
		record.On("Get", "avatar_icon").Return(t2.Avatar, true)

		result := new(mres)
		result.On("Record").Return(record)
		result.On("Err").Return(nil)
		return result
	}

	t2assert = func(u *user) func() (*user, bool) {
		return func() (*user, bool) {
			passed := (u.Name == t2.Name &&
				u.Age == t2.Age &&
				u.Power == t2.Power &&
				u.IsActive == t2.IsActive &&
				u.Avatar == t2.Avatar)
			return &t2, passed
		}
	}
)

func TestResult_ToStruct(t *testing.T) {
	u := new(user)
	u2 := new(user)

	type fields struct {
		Result neo4j.Result
		m      rcache
		set    bool
	}
	type args struct {
		dest interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		assertion func() (*user, bool)
		wantErr   bool
	}{
		{
			name: "Successfully marshals into fields that are primitive type",
			fields: fields{
				Result: t1mock(),
			},
			args:      args{u},
			assertion: t1assert(u),
			wantErr:   false,
		},
		{
			name: "Handles case were destination is not a struct",
			fields: fields{
				Result: t1mock(),
			},
			args:    args{&[]string{"foo", "fun", "fire"}},
			wantErr: true,
		},
		{
			name: "Handles case were destination is not a pointer",
			fields: fields{
				Result: t1mock(),
			},
			args:    args{123912},
			wantErr: true,
		},
		{
			name: "Attempts to map all compatible values skipping those that are not",
			fields: fields{
				Result: t2mock(),
			},
			args:      args{u2},
			assertion: t2assert(u2),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Result{
				Result: tt.fields.Result,
				m:      tt.fields.m,
				set:    tt.fields.set,
			}
			if err := r.ToStruct(tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("Result.ToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			expected, passed := tt.assertion()
			if passed == false {
				t.Errorf("Result.ToStruct() expected = %+v, got %+v", expected, tt.args.dest)
			}
		})
	}
}

func BenchmarkResult_ToStruct(b *testing.B) {
	u := new(user)
	result := t1mock()

	type fields struct {
		Result neo4j.Result
		m      rcache
		set    bool
	}
	type args struct {
		dest interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Primitve Types Only",
			fields: fields{
				Result: result,
			},
			args: args{u},
		},
	}

	for _, tt := range tests {
		r := &Result{
			Result: tt.fields.Result,
			m:      tt.fields.m,
			set:    tt.fields.set,
		}

		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r.ToStruct(tt.args.dest)
			}
		})
	}
}

// mocked neo4j.Result
type mres struct {
	neo4j.Result
	mock.Mock
}

func (m *mres) Err() error {
	args := m.Called()
	return args.Error(0)
}

func (m *mres) Record() neo4j.Record {
	args := m.Called()
	record, _ := args.Get(0).(neo4j.Record)
	return record
}
