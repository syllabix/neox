package neox

import (
	"errors"
	"reflect"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

const neotag = "db"

var (
	// ErrInvalidArg is returned when provided arguments are invalid
	ErrInvalidArg = errors.New("the provided destination is not a pointer to a struct")
)

type rprops struct {
	index int
	kind  reflect.Kind
}

type rcache map[string]rprops

// A Result is returned from successful
// calls to a Session.Runx, it exposes all of the standard
// driver interface methods as well as its various extensions
type Result struct {
	neo4j.Result
	m   rcache
	set bool
}

// Recordx returns a neox.Record at the current index in the
// the result stream
func (r *Result) Recordx() *Record {
	return &Record{r.Record()}
}

// ToStruct attempts to assign the values of the current result record to fields of
// the provided struct. The argument must be a pointer to a struct or an ErrInvalidArg will be returned.
// ToStruct will cache results of reflecting on the provided destination type to improve performance
// on every subsequent call for an instance of a Result. That said, using varying struct types through the lifetime
// of a single result instance should be considered unsafe and will yield unstable results
func (r *Result) ToStruct(dest interface{}) error {
	if r.Err() != nil {
		return r.Err()
	}

	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return ErrInvalidArg
	}

	e := v.Elem()
	if e.Kind() != reflect.Struct {
		return ErrInvalidArg
	}

	if !r.set {
		r.m = make(rcache, e.NumField())
		for i := 0; i < e.NumField(); i++ {
			fieldType := e.Type().Field(i)
			r.m[fieldType.Tag.Get(neotag)] = rprops{
				index: i,
				kind:  e.Field(i).Kind(),
			}
		}
		r.set = true
	}

	record := r.Record()
	for name, cache := range r.m {
		r, ok := record.Get(name)
		if !ok {
			continue
		}

		field := e.Field(cache.index)
		if field.CanSet() {
			value := reflect.ValueOf(r)
			if value.Kind() == cache.kind {
				field.Set(value)
			}
		}
	}

	return nil
}
