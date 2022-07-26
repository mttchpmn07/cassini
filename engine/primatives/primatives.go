package primatives

import "github.com/mttchpmn07/cassini/engine"

type Primative interface {
	engine.Shape
	Collider
	engine.Rasterable
}
