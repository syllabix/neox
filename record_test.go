package neox

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func TestRecord_GetIntAtIndex(t *testing.T) {
	t.Parallel()
	m := new(mrec)
	m.On("GetByIndex", 0).Return(12)
	m.On("GetByIndex", 1).Return("hello")

	type fields struct {
		Record neo4j.Record
	}
	type args struct {
		index int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue int
		wantOk    bool
	}{
		{
			name:      "Should return expected value as integer",
			fields:    fields{m},
			args:      args{index: 0},
			wantValue: 12,
			wantOk:    true,
		},
		{
			name:      "Should handle when type is not integer",
			fields:    fields{m},
			args:      args{index: 1},
			wantValue: 0,
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Record{
				Record: tt.fields.Record,
			}
			gotValue, gotOk := r.GetIntAtIndex(tt.args.index)
			if gotValue != tt.wantValue {
				t.Errorf("Record.GetIntAtIndex() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Record.GetIntAtIndex() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestRecord_GetInt(t *testing.T) {
	t.Parallel()
	m := new(mrec)
	m.On("Get", "foo_bar").Return(239, true)
	m.On("Get", "buzz_baz").Return("notanint", false)
	m.On("Get", "razz_fuzz").Return(0, false)

	type fields struct {
		Record neo4j.Record
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue int
		wantOk    bool
	}{
		{
			name:      "Should return an int for valid key",
			fields:    fields{m},
			args:      args{"foo_bar"},
			wantValue: 239,
			wantOk:    true,
		},
		{
			name:      "Should return false and zero value when the return type is not compatible",
			fields:    fields{m},
			args:      args{"buzz_baz"},
			wantValue: 0,
			wantOk:    false,
		},
		{
			name:      "Should return false and zero value for non existent key",
			fields:    fields{m},
			args:      args{"razz_fuzz"},
			wantValue: 0,
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Record{
				Record: tt.fields.Record,
			}
			gotValue, gotOk := r.GetInt(tt.args.key)
			if gotValue != tt.wantValue {
				t.Errorf("Record.GetInt() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Record.GetInt() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestRecord_GetStringAtIndex(t *testing.T) {
	t.Parallel()
	m := new(mrec)
	m.On("GetByIndex", 0).Return("excellent@cool.com")
	m.On("GetByIndex", 1).Return("")

	type fields struct {
		Record neo4j.Record
	}
	type args struct {
		index int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue string
		wantOk    bool
	}{
		{
			name:      "Should return a string at valid index",
			fields:    fields{m},
			args:      args{0},
			wantValue: "excellent@cool.com",
			wantOk:    true,
		},
		{
			name:      "Should return an empty string and false for ok",
			fields:    fields{m},
			args:      args{1},
			wantValue: "",
			wantOk:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Record{
				Record: tt.fields.Record,
			}
			gotValue, gotOk := r.GetStringAtIndex(tt.args.index)
			if gotValue != tt.wantValue {
				t.Errorf("Record.GetStringAtIndex() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Record.GetStringAtIndex() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestRecord_GetString(t *testing.T) {
	t.Parallel()
	m := new(mrec)
	m.On("Get", "foo_bar").Return("noname@fun.net", true)
	m.On("Get", "buzz_baz").Return(123.12, false)
	m.On("Get", "razz_fuzz").Return("", false)

	type fields struct {
		Record neo4j.Record
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue string
		wantOk    bool
	}{
		{
			name:      "Should return an string for valid key",
			fields:    fields{m},
			args:      args{"foo_bar"},
			wantValue: "noname@fun.net",
			wantOk:    true,
		},
		{
			name:      "Should return false and zero value when the return type is not compatible",
			fields:    fields{m},
			args:      args{"buzz_baz"},
			wantValue: "",
			wantOk:    false,
		},
		{
			name:      "Should return false and zero value for non existent key",
			fields:    fields{m},
			args:      args{"razz_fuzz"},
			wantValue: "",
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Record{
				Record: tt.fields.Record,
			}
			gotValue, gotOk := r.GetString(tt.args.key)
			if gotValue != tt.wantValue {
				t.Errorf("Record.GetString() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Record.GetString() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestRecord_GetFloatAtIndex(t *testing.T) {
	t.Parallel()
	m := new(mrec)
	m.On("GetByIndex", 0).Return(123.34)
	m.On("GetByIndex", 1).Return(0)

	type fields struct {
		Record neo4j.Record
	}
	type args struct {
		index int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue float64
		wantOk    bool
	}{
		{
			name:      "Should return a float at valid index",
			fields:    fields{m},
			args:      args{0},
			wantValue: 123.34,
			wantOk:    true,
		},
		{
			name:      "Should return a zero value for float and false for ok",
			fields:    fields{m},
			args:      args{1},
			wantValue: 0,
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Record{
				Record: tt.fields.Record,
			}
			gotValue, gotOk := r.GetFloatAtIndex(tt.args.index)
			if gotValue != tt.wantValue {
				t.Errorf("Record.GetFloatAtIndex() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Record.GetFloatAtIndex() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestRecord_GetFloat(t *testing.T) {
	t.Parallel()
	m := new(mrec)
	m.On("Get", "foo_bar").Return(289.23, true)
	m.On("Get", "buzz_baz").Return(false, false)
	m.On("Get", "razz_fuzz").Return(0, false)

	type fields struct {
		Record neo4j.Record
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue float64
		wantOk    bool
	}{
		{
			name:      "Should return an string for valid key",
			fields:    fields{m},
			args:      args{"foo_bar"},
			wantValue: 289.23,
			wantOk:    true,
		},
		{
			name:      "Should return false and zero value when the return type is not compatible",
			fields:    fields{m},
			args:      args{"buzz_baz"},
			wantValue: 0,
			wantOk:    false,
		},
		{
			name:      "Should return false and zero value for non existent key",
			fields:    fields{m},
			args:      args{"razz_fuzz"},
			wantValue: 0,
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Record{
				Record: tt.fields.Record,
			}
			gotValue, gotOk := r.GetFloat(tt.args.key)
			if gotValue != tt.wantValue {
				t.Errorf("Record.GetFloat() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Record.GetFloat() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestRecord_GetBoolAtIndex(t *testing.T) {
	t.Parallel()
	m := new(mrec)
	m.On("GetByIndex", 0).Return(true)
	m.On("GetByIndex", 1).Return("not a bool")

	type fields struct {
		Record neo4j.Record
	}
	type args struct {
		index int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue bool
		wantOk    bool
	}{
		{
			name:      "Should return a bool at valid index",
			fields:    fields{m},
			args:      args{0},
			wantValue: true,
			wantOk:    true,
		},
		{
			name:      "Should return a zero value for a bool and false for ok",
			fields:    fields{m},
			args:      args{1},
			wantValue: false,
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Record{
				Record: tt.fields.Record,
			}
			gotValue, gotOk := r.GetBoolAtIndex(tt.args.index)
			if gotValue != tt.wantValue {
				t.Errorf("Record.GetBoolAtIndex() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Record.GetBoolAtIndex() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestRecord_GetBool(t *testing.T) {
	t.Parallel()
	m := new(mrec)
	m.On("Get", "foo_bar").Return(true, true)
	m.On("Get", "buzz_baz").Return(123.343, false)
	m.On("Get", "razz_fuzz").Return(false, false)

	type fields struct {
		Record neo4j.Record
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue bool
		wantOk    bool
	}{
		{
			name:      "Should return a bool for valid key",
			fields:    fields{m},
			args:      args{"foo_bar"},
			wantValue: true,
			wantOk:    true,
		},
		{
			name:      "Should return false and zero value when the return type is not compatible",
			fields:    fields{m},
			args:      args{"buzz_baz"},
			wantValue: false,
			wantOk:    false,
		},
		{
			name:      "Should return false and zero value for non existent key",
			fields:    fields{m},
			args:      args{"razz_fuzz"},
			wantValue: false,
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Record{
				Record: tt.fields.Record,
			}
			gotValue, gotOk := r.GetBool(tt.args.key)
			if gotValue != tt.wantValue {
				t.Errorf("Record.GetBool() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Record.GetBool() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

type mrec struct {
	mock.Mock
	neo4j.Record
}

func (m *mrec) GetByIndex(index int) interface{} {
	args := m.Called(index)
	return args.Get(0)
}

func (m *mrec) Get(key string) (interface{}, bool) {
	args := m.Called(key)
	return args.Get(0), args.Bool(1)
}
