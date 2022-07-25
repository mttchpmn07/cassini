package engine

import (
	"fmt"

	"github.com/faiface/pixel"
)

type Vector struct{ *pixel.Vec }
type Rect struct{ *pixel.Rect }
type Circle struct{ *pixel.Circle }
type Line struct{ *line }
type line struct {
	Start Vector
	End   Vector
}
type Polygon struct{ *poly }
type poly struct {
	Points []Vector
}

func VectorFromEvent(event Event) Vector {
	return event.Contents().(Vector)
}

func Vec(X, Y float64) Vector {
	v := pixel.V(X, Y)
	return Vector{&v}
}

func fromPixelVec(v pixel.Vec) Vector {
	return Vector{&v}
}

func (v *Vector) toPixelVec() pixel.Vec {
	return *v.Vec
}

func (v *Vector) String() string {
	return fmt.Sprintf("<%v, %v>", v.X, v.Y)
}

func NewRect(min Vector, max Vector) Rect {
	return Rect{&pixel.Rect{
		Min: min.toPixelVec(),
		Max: max.toPixelVec(),
	}}
}

func FromPixelRect(rect pixel.Rect) Rect {
	return Rect{&rect}
}

func NewCircle(radius float64, location Vector) Circle {
	return Circle{&pixel.Circle{
		Radius: radius,
		Center: location.toPixelVec(),
	}}
}

func NewLine(start Vector, end Vector) Line {
	return Line{&line{
		Start: start,
		End:   end,
	}}
}

func NewPolygon(points ...Vector) Polygon {
	var vectors []Vector
	vectors = append(vectors, points...)
	return Polygon{&poly{
		Points: vectors,
	}}
}

func NewPolygonFromLines(lines []line) Polygon {
	var vectors []Vector
	for _, l := range lines {
		vectors = append(vectors, l.Start)
	}
	return Polygon{&poly{
		Points: vectors,
	}}
}
