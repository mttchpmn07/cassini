package primatives

import (
	"github.com/mttchpmn07/cassini/engine/graphics"
)

type Primative interface {
	Collider
	graphics.Rasterable
	Move(v Vector) Primative
}

func (cp concretePrimative) Move(v Vector) Primative {
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
