package engine

import (
	"fmt"

	"github.com/faiface/pixel"
)

type shape int

const (
	Point shape = iota
	Line
	Circle
	Rectangle
	Polygon
)

func (s shape) String() string {
	return map[shape]string{
		Point:     "Point",
		Line:      "Line",
		Circle:    "Circle",
		Rectangle: "Rectangle",
		Polygon:   "Polygon",
	}[s]
}

type Shape interface {
	Type() shape
}

type Vector struct{ *pixel.Vec }
type Rect struct{ *pixel.Rect }
type Circ struct{ *pixel.Circle }
type Lin struct{ *line }
type line struct {
	Start Vector
	End   Vector
}
type Poly struct{ *poly }
type poly struct {
	Points []Vector
}

func (v Vector) Type() shape {
	return Point
}

func (r Rect) Type() shape {
	return Rectangle
}

func (c Circ) Type() shape {
	return Circle
}

func (l Lin) Type() shape {
	return Line
}

func (p Poly) Type() shape {
	return Polygon
}

func VectorFromEvent(event Event) Vector {
	return event.Contents().(Vector)
}

func NewVector(X, Y float64) Vector {
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

func NewRectangle(min Vector, max Vector) Rect {
	return Rect{&pixel.Rect{
		Min: min.toPixelVec(),
		Max: max.toPixelVec(),
	}}
}

func FromPixelRect(rect pixel.Rect) Rect {
	return Rect{&rect}
}

func NewCircle(radius float64, location Vector) Circ {
	return Circ{&pixel.Circle{
		Radius: radius,
		Center: location.toPixelVec(),
	}}
}

func NewLine(start Vector, end Vector) Lin {
	return Lin{&line{
		Start: start,
		End:   end,
	}}
}

func NewPolygon(points ...Vector) Poly {
	var vectors []Vector
	vectors = append(vectors, points...)
	return Poly{&poly{
		Points: vectors,
	}}
}

func NewPolygonFromLines(lines []line) Poly {
	var vectors []Vector
	for _, l := range lines {
		vectors = append(vectors, l.Start)
	}
	return Poly{&poly{
		Points: vectors,
	}}
}
