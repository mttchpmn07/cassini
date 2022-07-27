package primatives

import (
	"github.com/mttchpmn07/cassini/engine/graphics"
	m "github.com/mttchpmn07/cassini/engine/math"
)

type Primative interface {
	Collider
	graphics.Rasterable
	Move(v m.Vector) Primative
}

func (cp concretePrimative) Move(v m.Vector) Primative {
	return cp
}

type concretePrimative struct {
	Collider
	graphics.Rasterable
}

func NewPrimative(collider Collider) Primative {
	return &concretePrimative{
		collider,
		graphics.NewRasterable(),
	}
}
