package primatives

import (
	"fmt"
	"math"

	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/mttchpmn07/cassini/engine/events"
)

type Vector struct {
	*pixel.Vec
	Primative
}

func PointFromEvent(event events.Event) Vector {
	return event.Contents().(Vector)
}

func VectorFromEvent(event events.Event) Vector {
	return event.Contents().(Vector)
}

func NewVector(X, Y float64) Vector {
	v := pixel.V(X, Y)
	return Vector{
		&v,
		NewPrimative(NewCollider(NewShape(Point))),
	}
}

func FromPixelVec(v pixel.Vec) Vector {
	return Vector{
		&v,
		NewPrimative(NewCollider(NewShape(Point))),
	}
}

func fromCollision2dVec(v collision2d.Vector) Vector {
	return NewVector(v.X, v.Y)
}

func (v *Vector) toPixelVec() pixel.Vec {
	return *v.Vec
}

func (v *Vector) String() string {
	return fmt.Sprintf("<%v, %v>", v.X, v.Y)
}

func ZeroVector() Vector {
	return NewVector(0, 0)
}

func (vec Vector) Add(other Vector) Vector {
	return NewVector(vec.X+other.X, vec.Y+other.Y)
}

func (vec Vector) Scale(factor float64) Vector {
	return NewVector(vec.X*factor, vec.Y*factor)
}

func (vec Vector) Mid(other Vector) Vector {
	return vec.Add(other).Scale(.5)
}

func (vec Vector) Diff(other Vector) Vector {
	return NewVector(vec.X-other.X, vec.Y-other.Y)
}

func (vec Vector) Dist(other Vector) float64 {
	diff := vec.Diff(other)
	return math.Sqrt(diff.X*diff.X + diff.Y*diff.Y)
}

func (vec Vector) Move(v Vector) Primative {
	vec.X = v.X
	vec.Y = v.Y
	return vec
}

func (vec Vector) toCollision2d() collision2d.Vector {
	return collision2d.Vector(vec.toPixelVec())
}

func (vec Vector) Collides(other Collider) (Collision, bool) {
	v := vec.toCollision2d()
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := false
	switch other.Type() {
	case Point:
		p := other.(Vector)
		c := collision2d.NewCircle(p.toCollision2d(), 1)
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

func (vec Vector) Raster() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = vec.C()
	imd.Push(vec.toPixelVec())
	imd.Circle(1, vec.T())
	return imd
}
