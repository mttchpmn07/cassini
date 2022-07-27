package primatives

import (
	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Rect struct {
	*pixel.Rect
	Primative
}

func NewRectangle(min Vector, max Vector) Rect {
	return Rect{
		&pixel.Rect{
			Min: min.toPixelVec(),
			Max: max.toPixelVec(),
		},
		NewPrimative(NewCollider(NewShape(Rectangle))),
	}
}

func FromPixelRect(rect pixel.Rect) Rect {
	return Rect{
		&rect,
		NewPrimative(NewCollider(NewShape(Rectangle))),
	}
}

func (r Rect) Move(v Vector) Primative {
	r.Max = r.Max.Add(v.toPixelVec())
	r.Min = r.Min.Add(v.toPixelVec())
	return r
}

func (r Rect) toCollision2d() collision2d.Polygon {
	mid := FromPixelVec(r.Max).Mid(FromPixelVec(r.Min))
	diff := FromPixelVec(r.Max).Diff(FromPixelVec(r.Min))
	newPoly := collision2d.NewBox(mid.toCollision2d(), diff.X, diff.Y).ToPolygon()
	return newPoly.SetOffset(collision2d.NewVector(-diff.X/2, -diff.Y/2))
}

func (r Rect) Collides(other Collider) (Collision, bool) {
	rec := r.toCollision2d()
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := false
	switch other.Type() {
	case Point:
		v := other.(Vector).toCollision2d()
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
	imd.Push(r.Min)
	imd.Push(pixel.V(r.Min.X, r.Max.Y))
	imd.Push(r.Max)
	imd.Push(pixel.V(r.Max.X, r.Min.Y))
	imd.Polygon(r.T())
	return imd
}
