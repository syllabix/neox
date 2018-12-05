package neox

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

const neotag = "neo"

var (
	// ErrInvalidArg is returned when provided arguments are invalid
	ErrInvalidArg = errors.New("the provided argument is invalid")
)

// A Result is returned from successful
// calls to a Session.Runx, it exposes all of the standard
// driver interface methods as well as its various extensions
type Result struct {
	neo4j.Result
	m   map[string]int
	set bool
}

// ToStruct attempts to assign the values of the current
// result record to fields of the provided struct.
// The argument must be a pointer to a struct or an
// neox.ErrInvalidArg will be returned
func (r *Result) ToStruct(dest interface{}) error {
	if r.Err() != nil {
		return r.Err()
	}

	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return errors.New("destination is not a pointer")
	}

	e := v.Elem()

	if !r.set {
		r.m = make(map[string]int)
		if e.Kind() != reflect.Struct {
			return errors.New("destination is not a struct")
		}

		for i := 0; i < e.NumField(); i++ {
			fieldType := e.Type().Field(i)
			r.m[fieldType.Tag.Get(neotag)] = i
		}
		r.set = true
	}

	record := r.Record()
	for name, idx := range r.m {
		r, ok := record.Get(name)
		if !ok {
			return errors.New("neo4j record does contain a value labeled " + name)
		}
		if ok {
			field := e.Field(idx)

			if field.CanSet() {
				recVal := reflect.ValueOf(r)
				if recVal.Kind() == field.Kind() {
					field.Set(reflect.ValueOf(r))
					continue
				} else {
					return fmt.Errorf("cannot set struct field \"%s\" of type %s with record %s of type %s",
						e.Type().Field(idx).Name,
						e.Type().Field(idx).Type.Name(),
						name,
						recVal.Type().Name())
				}
			} else {
				return fmt.Errorf("struct field \"%s\" cannot be set", e.Type().Field(idx).Name)
			}
		}
	}

	return nil
}
