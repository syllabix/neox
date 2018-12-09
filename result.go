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
// of a single result instance should be considered unsafe and will provided unstable results
func (r *Result) ToStruct(dest interface{}) error {
	if r.Err() != nil {
		return r.Err()
	}

	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return ErrInvalidArg
	}

	e := v.Elem()

	if !r.set {
		r.m = make(rcache, e.NumField())

		if e.Kind() != reflect.Struct {
			return ErrInvalidArg
		}

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
			return fmt.Errorf("neo4j record does contain a value labeled %s", name)
		}
		if ok {
			field := e.Field(cache.index)

			if field.CanSet() {
				value := reflect.ValueOf(r)
				if value.Kind() == cache.kind {
					field.Set(value)
					continue
				} else {
					return fmt.Errorf("cannot set struct field \"%s\" of type %s with record %s of type %s",
						e.Type().Field(cache.index).Name,
						e.Type().Field(cache.index).Type.Name(),
						name,
						value.Type().Name())
				}
			} else {
				return fmt.Errorf("struct field \"%s\" cannot be set", e.Type().Field(cache.index).Name)
			}
		}
	}

	return nil
}
