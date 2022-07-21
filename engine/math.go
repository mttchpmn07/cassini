package engine

import (
	"fmt"

	"github.com/faiface/pixel"
)

type Vector struct {
	X, Y float64
}

func fromPixelVec(vec pixel.Vec) *Vector {
	return &Vector{
		X: vec.X,
		Y: vec.Y,
	}
}

func toPixelVec(vec Vector) pixel.Vec {
	return pixel.V(vec.X, vec.Y)
}

func (v *Vector) String() string {
	return fmt.Sprintf("<%v, %v>", v.X, v.Y)
}
