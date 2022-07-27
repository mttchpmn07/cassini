package primatives

import (
	"fmt"

	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel/imdraw"
	m "github.com/mttchpmn07/cassini/engine/math"
)

type circ struct {
	Radius float64
	Center m.Vector
}

type Circ struct {
	*circ
	Primative
}

func NewCircle(radius float64, location m.Vector) Primative {
	return Circ{
		&circ{
			Radius: radius,
			Center: location,
		},
		NewPrimative(NewCollider(NewShape(Circle))),
	}
}

func (c Circ) Move(v m.Vector) Primative {
	c.Center = c.Center.Add(v)
	return c
}

func (c Circ) String() string {
	return fmt.Sprintf("Circle(%v, %v)", c.Center, c.Radius)
}

func (c Circ) toCollision2d() collision2d.Circle {
	return collision2d.NewCircle(c.Center.ToCollision2d(), c.Radius)
}

func (c Circ) Collides(other Collider) (Collision, bool) {
	cir := c.toCollision2d()
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := false
	switch other.Type() {
	case Point:
		v := other.(Dot).Vector.ToCollision2d()
		col = collision2d.PointInCircle(v, cir)
	case Line:
		col = TestCircleLine(cir, other.(Lin))
	case Circle:
		otherC := other.(Circ).toCollision2d()
		col, res = collision2d.TestCircleCircle(cir, otherC)
	case Rectangle:
		rec := other.(Rect).toCollision2d()
		col, res = TestCirclePolygon(cir, rec)
	case Polygon:
		p := other.(Poly).toCollision2d()
		col, res = TestCirclePolygon(cir, p)
	default:
	}
	return Collision{&res}, col
}

func (c Circ) Raster() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = c.C()
	imd.Push(c.Center.ToPixel())
	imd.Circle(c.Radius, c.T())
	return imd
}
