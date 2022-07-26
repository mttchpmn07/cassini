package primatives

import (
	"fmt"

	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/mttchpmn07/cassini/engine"
)

type Circ struct {
	*pixel.Circle
	engine.Shape
	engine.Rasterable
}

func NewCircle(radius float64, location Vector) Circ {
	return Circ{
		&pixel.Circle{
			Radius: radius,
			Center: location.toPixelVec(),
		},
		engine.NewShape(engine.Circle),
		engine.NewRasterable(),
	}
}

func (c Circ) Move(v Vector) Collider {
	c.Center = c.Center.Add(v.toPixelVec())
	return c
}

func (c Circ) String() string {
	return fmt.Sprintf("(%v, %v)", c.Center, c.Radius)
}

func (c Circ) toCollision2d() collision2d.Circle {
	return collision2d.NewCircle(collision2d.Vector(c.Center), c.Radius)
}

func (c Circ) Collides(other Collider) (Collision, bool) {
	cir := c.toCollision2d()
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := false
	switch other.Type() {
	case engine.Point:
		v := other.(Vector).toCollision2d()
		col = collision2d.PointInCircle(v, cir)
	case engine.Line:
		col = TestCircleLine(cir, other.(Lin))
	case engine.Circle:
		otherC := other.(Circ).toCollision2d()
		col, res = collision2d.TestCircleCircle(cir, otherC)
	case engine.Rectangle:
		rec := other.(Rect).toCollision2d()
		col, res = TestCirclePolygon(cir, rec)
	case engine.Polygon:
		p := other.(Poly).toCollision2d()
		col, res = TestCirclePolygon(cir, p)
	default:
	}
	return Collision{&res}, col
}

func (c Circ) Raster() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = c.C()
	imd.Push(c.Center)
	imd.Circle(1, c.T())
	return imd
}
