package engine

import (
	"fmt"
	"math"

	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
)

type shape int

const (
	Point shape = iota
	Line
	Circle
	Rectangle
	Polygon
)

func (s shape) String() string {
	return map[shape]string{
		Point:     "Point",
		Line:      "Line",
		Circle:    "Circle",
		Rectangle: "Rectangle",
		Polygon:   "Polygon",
	}[s]
}

type Shape interface {
	Type() shape
	Move(v Vector) Shape
	Collides(other Shape) (Collision, bool)
}

type Vector struct{ *pixel.Vec }
type Rect struct{ *pixel.Rect }
type Circ struct{ *pixel.Circle }
type Lin struct{ *line }
type line struct {
	Start Vector
	End   Vector
}
type Poly struct{ *poly }
type poly struct {
	Points []Vector
}

func PointFromEvent(event Event) Shape {
	return event.Contents().(Vector)
}

func VectorFromEvent(event Event) Vector {
	return event.Contents().(Vector)
}

func NewVector(X, Y float64) Vector {
	v := pixel.V(X, Y)
	return Vector{&v}
}

func fromPixelVec(v pixel.Vec) Vector {
	return Vector{&v}
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

func (v Vector) Type() shape {
	return Point
}

func (vec Vector) Move(v Vector) Shape {
	return v
}

func (vec Vector) toCollision2d() collision2d.Vector {
	return collision2d.Vector(vec.toPixelVec())
}

func (vec Vector) Collides(other Shape) (Collision, bool) {
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

func NewRectangle(min Vector, max Vector) Rect {
	return Rect{&pixel.Rect{
		Min: min.toPixelVec(),
		Max: max.toPixelVec(),
	}}
}

func (r Rect) Type() shape {
	return Rectangle
}

func (r Rect) Move(v Vector) Shape {
	return NewRectangle(fromPixelVec(r.Max).Add(v), fromPixelVec(r.Min).Add(v))
}

func (r Rect) Collides(other Shape) (Collision, bool) {
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

func (r Rect) toCollision2d() collision2d.Polygon {
	mid := fromPixelVec(r.Max).Mid(fromPixelVec(r.Min))
	diff := fromPixelVec(r.Max).Diff(fromPixelVec(r.Min))
	newPoly := collision2d.NewBox(mid.toCollision2d(), diff.X, diff.Y).ToPolygon()
	return newPoly.SetOffset(collision2d.NewVector(-diff.X/2, -diff.Y/2))
}

func FromPixelRect(rect pixel.Rect) Rect {
	return Rect{&rect}
}

func NewCircle(radius float64, location Vector) Circ {
	return Circ{&pixel.Circle{
		Radius: radius,
		Center: location.toPixelVec(),
	}}
}

func (c Circ) Type() shape {
	return Circle
}

func (c Circ) Move(v Vector) Shape {
	return NewCircle(c.Radius, fromPixelVec(c.Center).Add(v))
}

func (c Circ) String() string {
	return fmt.Sprintf("(%v, %v)", c.Center, c.Radius)
}

func (c Circ) Collides(other Shape) (Collision, bool) {
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

func (c Circ) toCollision2d() collision2d.Circle {
	return collision2d.NewCircle(collision2d.Vector(c.Center), c.Radius)
}

func NewLine(start Vector, end Vector) Lin {
	return Lin{&line{
		Start: start,
		End:   end,
	}}
}

func fromCollision2dEdges(edge1, edge2 collision2d.Vector) Lin {
	return NewLine(fromCollision2dVec(edge1), fromCollision2dVec(edge2))
}

func (l Lin) Type() shape {
	return Line
}

func (l Lin) Move(v Vector) Shape {
	return NewLine(l.Start.Add(v), l.End.Add(v))
}

func (l Lin) String() string {
	return fmt.Sprintf("(%v, %v) <-> (%v, %v)", l.Start.X, l.Start.Y, l.End.X, l.End.Y)
}

func (l Lin) Collides(other Shape) (Collision, bool) {
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

func NewPolygon(points ...Vector) Poly {
	var vectors []Vector
	vectors = append(vectors, points...)
	return Poly{&poly{
		Points: vectors,
	}}
}

func NewPolygonFromLines(lines []line) Poly {
	var vectors []Vector
	for _, l := range lines {
		vectors = append(vectors, l.Start)
	}
	return Poly{&poly{
		Points: vectors,
	}}
}

func (p Poly) Type() shape {
	return Polygon
}

func (p Poly) Move(v Vector) Shape {
	newPoints := []Vector{}
	for i := range p.Points {
		newPoints = append(newPoints, p.Points[i].Add(v))
	}
	return NewPolygon(newPoints...)
}

func (p Poly) Center() Vector {
	center := ZeroVector()
	for _, point := range p.Points {
		center = center.Add(point)
	}
	return center.Scale(1 / float64(len(p.Points)))
}

func (p Poly) Collides(other Shape) (Collision, bool) {
	pol := p.toCollision2d()
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := false
	switch other.Type() {
	case Point:
		v := other.(Vector).toCollision2d()
		col = collision2d.PointInPolygon(v, pol)
	case Line:
		l := other.(Lin)
		col, res = TestPolygonLine(pol, l)
	case Circle:
		c := other.(Circ).toCollision2d()
		col, res = collision2d.TestPolygonCircle(pol, c)
	case Rectangle:
		rec := other.(Rect).toCollision2d()
		col, res = collision2d.TestPolygonPolygon(pol, rec)
	case Polygon:
		otherP := other.(Poly).toCollision2d()
		col, res = collision2d.TestPolygonPolygon(pol, otherP)
	default:
	}
	return Collision{&res}, col
}

func (p Poly) toCollision2d() collision2d.Polygon {
	corners := []float64{}
	for _, point := range p.Points {
		corners = append(corners, point.X)
		corners = append(corners, point.Y)
	}
	return collision2d.NewPolygon(
		p.Center().toCollision2d(),
		p.Center().toCollision2d().Scale(-1),
		0,
		corners,
	)
}

type Collision struct {
	*collision2d.Response
}

func Collides(s Shape, other Shape) (Collision, bool) {
	res, col := s.Collides(other)
	return res, col
}

func TestCirclePolygon(circle collision2d.Circle, polygon collision2d.Polygon) (bool, collision2d.Response) {
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := collision2d.PointInPolygon(circle.Pos, polygon)
	if col {
		return col, res
	}
	nextIndex := 0
	points := polygon.Points
	for currentIndex := range points {
		nextIndex = currentIndex + 1
		if nextIndex == len(points) {
			nextIndex = 0
		}
		p1 := points[currentIndex].Add(polygon.Offset).Add(polygon.Pos)
		p2 := points[nextIndex].Add(polygon.Offset).Add(polygon.Pos)
		reconSide := fromCollision2dEdges(p1, p2)
		col = TestCircleLine(circle, reconSide)
		if col {
			return col, res
		}
	}
	return col, res
}

func TestCircleLine(circle collision2d.Circle, line Lin) bool {
	x1 := line.Start.X - circle.Pos.X
	y1 := line.Start.Y - circle.Pos.Y
	x2 := line.End.X - circle.Pos.X
	y2 := line.End.Y - circle.Pos.Y

	rSquared := circle.R * circle.R

	if x1*x1+y1*y1 <= rSquared {
		return true
	}
	if x2*x2+y2*y2 <= rSquared {
		return true
	}

	lenSquared := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)

	nx := y2 - y1
	ny := x1 - x2

	distSquared := (nx*x1 + ny*y1) * (nx*x1 + ny*y1)
	if distSquared > lenSquared*rSquared {
		return false
	}

	index := x1*(x1-x2) + y1*(y1-y2)
	if index < 0 {
		return false
	}
	if index > lenSquared {
		return false
	}
	return true
}

func TestLineLine(lin1, lin2 Lin) bool {
	x1, y1, x2, y2 := lin1.Start.X, lin1.Start.Y, lin1.End.X, lin1.End.Y
	x3, y3, x4, y4 := lin2.Start.X, lin2.Start.Y, lin2.End.X, lin2.End.Y

	uA := ((x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)) / ((y4-y3)*(x2-x1) - (x4-x3)*(y2-y1))
	uB := ((x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)) / ((y4-y3)*(x2-x1) - (x4-x3)*(y2-y1))

	if uA >= 0 && uA <= 1 && uB >= 0 && uB <= 1 {
		return true
	}
	return false
}

func TestPolygonLine(poly collision2d.Polygon, lin Lin) (bool, collision2d.Response) {
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := collision2d.PointInPolygon(lin.Start.toCollision2d(), poly) || collision2d.PointInPolygon(lin.End.toCollision2d(), poly)
	if col {
		return col, res
	}
	nextIndex := 0
	points := poly.Points
	for currentIndex := range points {
		nextIndex = currentIndex + 1
		if nextIndex == len(points) {
			nextIndex = 0
		}
		p1 := points[currentIndex].Add(poly.Offset).Add(poly.Pos)
		p2 := points[nextIndex].Add(poly.Offset).Add(poly.Pos)
		reconSide := fromCollision2dEdges(p1, p2)
		col = TestLineLine(lin, reconSide)
		if col {
			return col, res
		}
	}
	return col, res
}
