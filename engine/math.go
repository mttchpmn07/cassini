package engine

import "github.com/faiface/pixel"

type Vector struct {
	X, Y float64
}

func FromPixelVec(vec pixel.Vec) Vector {
	return Vector{
		X: vec.X,
		Y: vec.Y,
	}
}
