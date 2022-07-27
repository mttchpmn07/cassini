package primatives

import (
	"fmt"

	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel/imdraw"
	m "github.com/mttchpmn07/cassini/engine/math"
)

type Dot struct {
	*m.Vector
	Primative
}

func NewDot(x, y float64) Dot {
	v := m.NewVector(x, y)
	return Dot{
		&v,
		NewPrimative(NewCollider(NewShape(Point))),
	}
}

func (d *Dot) String() string {
	return fmt.Sprintf("<%v, %v>", d.Vector.X, d.Vector.Y)
}

func (d Dot) Move(v m.Vector) Primative {
	d.Vector.X = v.X
	d.Vector.Y = v.Y
	return d
}

func (d Dot) Collides(other Collider) (Collision, bool) {
	v := d.Vector.ToCollision2d()
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := false
	switch other.Type() {
	case Point:
		p := other.(Dot).Vector.ToCollision2d()
		c := collision2d.NewCircle(p, 1)
		col = collision2d.PointInCircle(v, c)
	case Line:
		l := other.(Lin)
		c := collision2d.NewCircle(v, 1)
		col = TestCircleLine(c, l)
	case Circle:
		c := other.(Circ).toCollision2d()
		col = collision2d.PointInCircle(v, c)
	case Rectangle:
		r := other.(Rect).toCollision2d()
		col = collision2d.PointInPolygon(v, r)
	case Polygon:
		p := other.(Poly).toCollision2d()
		col = collision2d.PointInPolygon(v, p)
	}
	return Collision{&res}, col
}

func (d Dot) Raster() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = d.C()
	imd.Push(d.Vector.ToPixel())
	imd.Circle(1, d.T())
	return imd
}
