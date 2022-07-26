package primatives

import (
	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/mttchpmn07/cassini/engine"
)

type Rect struct {
	*pixel.Rect
	engine.Shape
	engine.Rasterable
}

func NewRectangle(min Vector, max Vector) Rect {
	return Rect{
		&pixel.Rect{
			Min: min.toPixelVec(),
			Max: max.toPixelVec(),
		},
		engine.NewShape(engine.Rectangle),
		engine.NewRasterable(),
	}
}

func FromPixelRect(rect pixel.Rect) Rect {
	return Rect{
		&rect,
		engine.NewShape(engine.Rectangle),
		engine.NewRasterable(),
	}
}

func (r Rect) Move(v Vector) Collider {
	r.Max = r.Max.Add(v.toPixelVec())
	r.Min = r.Min.Add(v.toPixelVec())
	return r
}

func (r Rect) toCollision2d() collision2d.Polygon {
	mid := fromPixelVec(r.Max).Mid(fromPixelVec(r.Min))
	diff := fromPixelVec(r.Max).Diff(fromPixelVec(r.Min))
	newPoly := collision2d.NewBox(mid.toCollision2d(), diff.X, diff.Y).ToPolygon()
	return newPoly.SetOffset(collision2d.NewVector(-diff.X/2, -diff.Y/2))
}

func (r Rect) Collides(other Collider) (Collision, bool) {
	rec := r.toCollision2d()
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := false
	switch other.Type() {
	case engine.Point:
		v := other.(Vector).toCollision2d()
		col = collision2d.PointInPolygon(v, rec)
	case engine.Line:
		l := other.(Lin)
		col, res = TestPolygonLine(rec, l)
	case engine.Circle:
		c := other.(Circ).toCollision2d()
		col, res = collision2d.TestCirclePolygon(c, rec)
	case engine.Rectangle:
		otherR := other.(Rect).toCollision2d()
		col, res = collision2d.TestPolygonPolygon(rec, otherR)
	case engine.Polygon:
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
