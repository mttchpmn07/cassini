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

func (c Circ) Collides(other Collider) (bool, Collision) {
	switch other.Type() {
	case Point:
		return TestDotCircle(other.(Dot), c)
	case Line:
		col, res := TestCircleLine(c, other.(Lin))
		return col, res.Reverse()
	case Circle:
		return TestCircleCircle(other.(Circ), c)
	case Rectangle:
		col, res := TestCircleRect(c, other.(Rect))
		return col, res.Reverse()
	case Polygon:
		col, res := TestCirclePolygon(c, other.(Poly))
		return col, res.Reverse()
	}
	res := collision2d.NewResponse()
	res = res.NotColliding()
	return false, Collision{&res}
}

func (c Circ) Raster() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = c.C()
	imd.Push(c.Center.ToPixel())
	imd.Circle(c.Radius, c.T())
	return imd
}
