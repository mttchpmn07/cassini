package engine

import (
	"fmt"

	"github.com/faiface/pixel"
)

/*
type Vector struct {
	X, Y float64
}
*/
//type Vector pixel.Vec
type Vector struct {
	pixel.Vec
}

func Vec(X, Y float64) Vector {
	return Vector{pixel.V(X, Y)}
}

func fromPixelVec(vec pixel.Vec) *Vector {
	return &Vector{vec}
	/*
		return &Vector{
			X: vec.X,
			Y: vec.Y,
		}
	*/
}

func (vec Vector) toPixelVec() pixel.Vec {
	return pixel.V(vec.X, vec.Y)
}

func (v *Vector) String() string {
	return fmt.Sprintf("<%v, %v>", v.X, v.Y)
}
