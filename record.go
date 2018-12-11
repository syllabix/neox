package neox

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Record wraps the standard implementation of a neo4j.Record
// adding some useful utlities
type Record struct {
	neo4j.Record
}

// GetIntAtIndex retrieves the value for the record at the provided index
// asserting it as an integer, returning the value and boolean indicating
// whether the type was asserted to be an integer
func (r *Record) GetIntAtIndex(index int) (value int, ok bool) {
	value, ok = r.GetByIndex(index).(int)
	return
}

// GetInt attempts to retrieve an integer value for the provided key
// If the provided key does not exist, or the value is not an integer, the method
// returns the zero value and false
func (r *Record) GetInt(key string) (value int, ok bool) {
	v, ok := r.Get(key)
	if !ok {
		return
	}
	value, ok = v.(int)
	if !ok {
		return
	}
	return
}

// GetStringAtIndex retrieves the value for the record at the provided index
// asserting it as a string, returning the value and boolean indicating
// whether the type was asserted correctly
func (r *Record) GetStringAtIndex(index int) (value string, ok bool) {
	value, ok = r.GetByIndex(index).(string)
	return
}

// GetString attempts to retrieve a string value for the provided key
// If the provided key does not exist, or the value is not a string, the method
// returns the zero value and false
func (r *Record) GetString(key string) (value string, ok bool) {
	v, ok := r.Get(key)
	if !ok {
		return "", false
	}
	value, ok = v.(string)
	if !ok {
		return "", false
	}
	return
}

// GetFloatAtIndex retrieves the value for the record at the provided index
// asserting it as a float, returning the value and boolean indicating
// whether the type was asserted correctly
func (r *Record) GetFloatAtIndex(index int) (value float64, ok bool) {
	value, ok = r.GetByIndex(index).(float64)
	return
}

// GetFloat attempts to retrieve a float value for the provided key
// If the provided key does not exist, or the value is not a float, the method
// returns the zero value and false
func (r *Record) GetFloat(key string) (value float64, ok bool) {
	v, ok := r.Get(key)
	if !ok {
		return 0, false
	}
	value, ok = v.(float64)
	if !ok {
		return 0, false
	}
	return
}

// GetBoolAtIndex retrieves the value for the record at the provided index
// asserting it as a bool, returning the value and boolean indicating
// whether the type was asserted correctly
func (r *Record) GetBoolAtIndex(index int) (value bool, ok bool) {
	value, ok = r.GetByIndex(index).(bool)
	return
}

// GetBool attempts to retrieve a rune value for the provided key
// If the provided key does not exist, or the value is not a rune, the method
// returns the zero value and false
func (r *Record) GetBool(key string) (value bool, ok bool) {
	v, ok := r.Get(key)
	if !ok {
		return false, false
	}
	value, ok = v.(bool)
	if !ok {
		return false, false
	}
	return
}
