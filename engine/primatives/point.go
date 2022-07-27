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

func (d Dot) Collides(other Collider) (bool, Collision) {
	switch other.Type() {
	case Point:
		return TestDotDot(other.(Dot), d)
	case Line:
		col, res := TestDotLine(d, other.(Lin))
		return col, res.Reverse()
	case Circle:
		col, res := TestDotCircle(d, other.(Circ))
		return col, res.Reverse()
	case Rectangle:
		col, res := TestDotRect(d, other.(Rect))
		return col, res.Reverse()
	case Polygon:
		col, res := TestDotPolygon(d, other.(Poly))
		return col, res.Reverse()
	}
	res := collision2d.NewResponse()
	res = res.NotColliding()
	return false, Collision{&res}
}

func (d Dot) Raster() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = d.C()
	imd.Push(d.Vector.ToPixel())
	imd.Circle(1, d.T())
	return imd
}
