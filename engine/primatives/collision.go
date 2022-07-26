package primatives

import (
	"github.com/Tarliton/collision2d"
	"github.com/mttchpmn07/cassini/engine"
)

type Collider interface {
	engine.Shape
	Move(v Vector) Collider
	Collides(other Collider) (Collision, bool)
}

type Collision struct {
	*collision2d.Response
}

func Collides(s Collider, other Collider) (Collision, bool) {
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
