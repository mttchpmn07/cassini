package math

import (
	"fmt"
	"math"

	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
	"github.com/mttchpmn07/cassini/engine/events"
)

type Vector struct {
	X float64
	Y float64
}

func NewVector(x, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

func VectorFromEvent(event events.Event) Vector {
	v := event.Contents().(pixel.Vec)
	return NewVector(v.X, v.Y)
}

func VectorFromPixel(v pixel.Vec) Vector {
	return NewVector(v.X, v.Y)
}

func VectorFromCollision2d(v collision2d.Vector) Vector {
	return NewVector(v.X, v.Y)
}

func (v *Vector) String() string {
	return fmt.Sprintf("<%v, %v>", v.X, v.Y)
}

func (v *Vector) ToPixel() pixel.Vec {
	return pixel.V(v.X, v.Y)
}

func (v *Vector) ToCollision2d() collision2d.Vector {
	return collision2d.NewVector(v.X, v.Y)
}

func ZeroVector() Vector {
	return NewVector(0, 0)
}

func (vec Vector) Scale(factor float64) Vector {
	return NewVector(vec.X*factor, vec.Y*factor)
}

func (vec Vector) Add(other Vector) Vector {
	return NewVector(vec.X+other.X, vec.Y+other.Y)
}

func (vec Vector) Mid(other Vector) Vector {
	return vec.Add(other).Scale(.5)
}

func (vec Vector) Diff(other Vector) Vector {
	return NewVector(vec.X-other.X, vec.Y-other.Y)
}

func (vec Vector) Dist(other Vector) float64 {
	diff := vec.Diff(other)
	return math.Sqrt(diff.X*diff.X + diff.Y*diff.Y)
}
