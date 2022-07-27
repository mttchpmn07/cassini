package primatives

import (
	"fmt"

	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel/imdraw"
)

type line struct {
	Start Vector
	End   Vector
}

type Lin struct {
	*line
	Primative
}

func NewLine(start Vector, end Vector) Lin {
	return Lin{
		&line{
			Start: start,
			End:   end,
		},
		NewPrimative(NewCollider(NewShape(Line))),
	}
}

func fromCollision2dEdges(edge1, edge2 collision2d.Vector) Lin {
	return NewLine(fromCollision2dVec(edge1), fromCollision2dVec(edge2))
}

func (l Lin) Move(v Vector) Primative {
	l.Start = l.Start.Add(v)
	l.End = l.End.Add(v)
	return l
}

func (l Lin) String() string {
	return fmt.Sprintf("(%v, %v) <-> (%v, %v)", l.Start.X, l.Start.Y, l.End.X, l.End.Y)
}

func (l Lin) Collides(other Collider) (Collision, bool) {
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := false
	switch other.Type() {
	case Point:
		v := other.(Vector).toCollision2d()
		c := collision2d.NewCircle(v, 1)
		col = TestCircleLine(c, l)
	case Line:
		l2 := other.(Lin)
		col = TestLineLine(l, l2)
	case Circle:
		c := other.(Circ).toCollision2d()
		col = TestCircleLine(c, l)
	case Rectangle:
		r := other.(Rect).toCollision2d()
		col, res = TestPolygonLine(r, l)
	case Polygon:
		p := other.(Poly).toCollision2d()
		col, res = TestPolygonLine(p, l)
	default:
	}
	return Collision{&res}, col
}

func (l Lin) Raster() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = l.C()
	imd.Push(l.Start.toPixelVec())
	imd.Push(l.End.toPixelVec())
	imd.Line(l.T())
	return imd
}
