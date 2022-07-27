package primatives

import (
	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	m "github.com/mttchpmn07/cassini/engine/math"
)

type quad struct {
	Min m.Vector
	Max m.Vector
}

type Rect struct {
	*quad
	Primative
}

func NewRectangle(min m.Vector, max m.Vector) Rect {
	return Rect{
		&quad{
			Min: min,
			Max: max,
		},
		NewPrimative(NewCollider(NewShape(Rectangle))),
	}
}

func (r Rect) Move(v m.Vector) Primative {
	r.Max = r.Max.Add(v)
	r.Min = r.Min.Add(v)
	return r
}

func (r Rect) toCollision2d() collision2d.Polygon {
	mid := r.Max.Mid(r.Min)
	diff := r.Max.Diff(r.Min)
	newPoly := collision2d.NewBox(mid.ToCollision2d(), diff.X, diff.Y).ToPolygon()
	return newPoly.SetOffset(collision2d.NewVector(-diff.X/2, -diff.Y/2))
}

func (r Rect) ToPolygon() Poly {
	points := []m.Vector{
		r.Min,
		m.NewVector(r.Min.X, r.Max.Y),
		r.Max,
		m.NewVector(r.Max.X, r.Min.Y),
	}
	return NewPolygon(points...)
}

func (r Rect) Collides(other Collider) (bool, Collision) {
	switch other.Type() {
	case Point:
		return TestDotRect(other.(Dot), r)
	case Line:
		col, res := TestRectLine(r, other.(Lin))
		return col, res.Reverse()
	case Circle:
		return TestCircleRect(other.(Circ), r)
	case Rectangle:
		return TestRectRect(other.(Rect), r)
	case Polygon:
		col, res := TestRectPolygon(r, other.(Poly))
		return col, res.Reverse()
	}
	res := collision2d.NewResponse()
	res = res.NotColliding()
	return false, Collision{&res}
}

func (r Rect) Raster() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = r.C()
	imd.Push(r.Min.ToPixel())
	imd.Push(pixel.V(r.Min.X, r.Max.Y))
	imd.Push(r.Max.ToPixel())
	imd.Push(pixel.V(r.Max.X, r.Min.Y))
	imd.Polygon(r.T())
	return imd
}
