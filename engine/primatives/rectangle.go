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

func (r Rect) Collides(other Collider) (Collision, bool) {
	rec := r.toCollision2d()
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := false
	switch other.Type() {
	case Point:
		v := other.(Dot).Vector.ToCollision2d()
		col = collision2d.PointInPolygon(v, rec)
	case Line:
		l := other.(Lin)
		col, res = TestPolygonLine(rec, l)
	case Circle:
		c := other.(Circ).toCollision2d()
		col, res = collision2d.TestCirclePolygon(c, rec)
	case Rectangle:
		otherR := other.(Rect).toCollision2d()
		col, res = collision2d.TestPolygonPolygon(rec, otherR)
	case Polygon:
		p := other.(Poly).toCollision2d()
		col, res = collision2d.TestPolygonPolygon(rec, p)
	default:
	}
	return Collision{&res}, col
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
