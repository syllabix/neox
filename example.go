package neox

import (
	"math/rand"
)

// CoolThing is the coolest thing that can be anything ever
type CoolThing interface{}

// Coolifyer does things to make a thing
// super cool
type Coolifyer interface {
	Coolify() CoolThing
}

// Exercise is a Coolifyer that takes a person and makes them a
// cool person
type Exercise struct {
	power int64
}

// Coolify is a great way to make an exercise cooler
func (e *Exercise) Coolify() CoolThing {
	power := rand.Int63n(e.power)
	return CoolThing(power * e.power)
}
