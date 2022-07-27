package primatives

import (
	"fmt"

	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Circ struct {
	*pixel.Circle
	Primative
}

func NewCircle(radius float64, location Vector) Primative {
	return Circ{
		&pixel.Circle{
			Radius: radius,
			Center: location.toPixelVec(),
		},
		NewPrimative(NewCollider(NewShape(Circle))),
	}
}

func (c Circ) Move(v Vector) Primative {
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
	case Point:
		v := other.(Vector).toCollision2d()
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
	imd.Push(c.Center)
	imd.Circle(1, c.T())
	return imd
}
