package primatives

import (
	"github.com/Tarliton/collision2d"
)

type Collider interface {
	Shape
	Collides(other Collider) (bool, Collision)
}

type concreteCollider struct {
	Shape
}

func NewCollider(shape Shape) Collider {
	return &concreteCollider{
		shape,
	}
}

func (cc *concreteCollider) Collides(s Collider) (bool, Collision) {
	res := collision2d.NewResponse()
	res = res.NotColliding()
	return false, Collision{&res}
}

type Collision struct {
	*collision2d.Response
}

func (c Collision) Reverse() Collision {
	a := c.A
	aInB := c.AInB
	c.OverlapN = c.OverlapN.Reverse()
	c.OverlapV = c.OverlapV.Reverse()
	c.A = c.B
	c.B = a
	c.AInB = c.BInA
	c.BInA = aInB
	return c
}

func Collides(s Collider, other Collider) (bool, Collision) {
	return s.Collides(other)
}

func TestCircleCircle(circleA Circ, circleB Circ) (bool, Collision) {
	col, res := collision2d.TestCircleCircle(circleA.toCollision2d(), circleB.toCollision2d())
	return col, Collision{&res}
}

func TestCirclePolygon(circle Circ, polygon Poly) (bool, Collision) {
	col, res := TestDotPolygon(NewDot(circle.Center.X, circle.Center.Y), polygon)
	if col {
		return col, res
	}
	p := polygon.toCollision2d()
	nextIndex := 0
	points := p.Points
	for currentIndex := range points {
		nextIndex = currentIndex + 1
		if nextIndex == len(points) {
			nextIndex = 0
		}
		p1 := points[currentIndex].Add(p.Offset).Add(p.Pos)
		p2 := points[nextIndex].Add(p.Offset).Add(p.Pos)
		reconSide := fromCollision2dEdges(p1, p2)
		col, res = TestCircleLine(circle, reconSide)
		if col {
			return col, res
		}
	}
	return col, res
}

func TestCircleLine(circle Circ, line Lin) (bool, Collision) {
	res := collision2d.NewResponse()
	res = res.NotColliding()
	x1 := line.Start.X - circle.Center.X
	y1 := line.Start.Y - circle.Center.Y
	x2 := line.End.X - circle.Center.X
	y2 := line.End.Y - circle.Center.Y

	rSquared := circle.Radius * circle.Radius

	if x1*x1+y1*y1 <= rSquared {
		return true, Collision{&res}
	}
	if x2*x2+y2*y2 <= rSquared {
		return true, Collision{&res}
	}

	lenSquared := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)

	nx := y2 - y1
	ny := x1 - x2

	distSquared := (nx*x1 + ny*y1) * (nx*x1 + ny*y1)
	if distSquared > lenSquared*rSquared {
		return false, Collision{&res}
	}

	index := x1*(x1-x2) + y1*(y1-y2)
	if index < 0 {
		return false, Collision{&res}
	}
	if index > lenSquared {
		return false, Collision{&res}
	}
	return true, Collision{&res}
}

func TestLineLine(lin1, lin2 Lin) (bool, Collision) {
	res := collision2d.NewResponse()
	res = res.NotColliding()
	x1, y1, x2, y2 := lin1.Start.X, lin1.Start.Y, lin1.End.X, lin1.End.Y
	x3, y3, x4, y4 := lin2.Start.X, lin2.Start.Y, lin2.End.X, lin2.End.Y

	uA := ((x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)) / ((y4-y3)*(x2-x1) - (x4-x3)*(y2-y1))
	uB := ((x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)) / ((y4-y3)*(x2-x1) - (x4-x3)*(y2-y1))

	if uA >= 0 && uA <= 1 && uB >= 0 && uB <= 1 {
		return true, Collision{&res}
	}
	return false, Collision{&res}
}

func TestPolygonLine(poly Poly, lin Lin) (bool, Collision) {
	res := collision2d.NewResponse()
	res = res.NotColliding()
	p := poly.toCollision2d()
	col := collision2d.PointInPolygon(lin.Start.ToCollision2d(), p) || collision2d.PointInPolygon(lin.End.ToCollision2d(), p)
	if col {
		return col, Collision{&res}
	}
	nextIndex := 0
	points := p.Points
	for currentIndex := range points {
		nextIndex = currentIndex + 1
		if nextIndex == len(points) {
			nextIndex = 0
		}
		p1 := points[currentIndex].Add(p.Offset).Add(p.Pos)
		p2 := points[nextIndex].Add(p.Offset).Add(p.Pos)
		reconSide := fromCollision2dEdges(p1, p2)
		col, res := TestLineLine(lin, reconSide)
		if col {
			return col, res
		}
	}
	return false, Collision{&res}
}

func TestPolygonPolygon(polyA Poly, polyB Poly) (bool, Collision) {
	col, res := collision2d.TestPolygonPolygon(polyA.toCollision2d(), polyB.toCollision2d())
	return col, Collision{&res}
}

func TestDotLine(dot Dot, line Lin) (bool, Collision) {
	res := collision2d.NewResponse()
	res = res.NotColliding()
	c := NewCircle(1, *dot.Vector)
	return TestCircleLine(c.(Circ), line)
}

func TestRectLine(rect Rect, lin Lin) (bool, Collision) {
	return TestPolygonLine(rect.ToPolygon(), lin)
}

func TestDotCircle(dot Dot, circle Circ) (bool, Collision) {
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := collision2d.PointInCircle(dot.ToCollision2d(), circle.toCollision2d())
	return col, Collision{&res}
}

func TestDotDot(dotA Dot, dotB Dot) (bool, Collision) {
	c := NewCircle(1, *dotB.Vector)
	return TestDotCircle(dotA, c.(Circ))
}

func TestDotRect(dot Dot, rect Rect) (bool, Collision) {
	return TestDotPolygon(dot, rect.ToPolygon())
}

func TestDotPolygon(dot Dot, poly Poly) (bool, Collision) {
	res := collision2d.NewResponse()
	res = res.NotColliding()
	col := collision2d.PointInPolygon(dot.ToCollision2d(), poly.toCollision2d())
	return col, Collision{&res}
}

func TestCircleRect(circle Circ, rect Rect) (bool, Collision) {
	p := rect.ToPolygon()
	return TestCirclePolygon(circle, p)
}

func TestRectRect(rectA Rect, rectB Rect) (bool, Collision) {
	col, res := collision2d.TestPolygonPolygon(rectA.toCollision2d(), rectB.toCollision2d())
	return col, Collision{&res}
}

func TestRectPolygon(rect Rect, polygon Poly) (bool, Collision) {
	col, res := collision2d.TestPolygonPolygon(rect.toCollision2d(), polygon.toCollision2d())
	return col, Collision{&res}

}
