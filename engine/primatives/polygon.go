package primatives

import (
	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel/imdraw"
)

type poly struct {
	Points []Vector
}

type Poly struct {
	*poly
	Primative
}

func NewPolygon(points ...Vector) Poly {
	var vectors []Vector
	vectors = append(vectors, points...)
	return Poly{
		&poly{
			Points: vectors,
		},
		NewPrimative(NewCollider(NewShape(Polygon))),
	}
}

func NewPolygonFromLines(lines []line) Poly {
	var vectors []Vector
	for _, l := range lines {
		vectors = append(vectors, l.Start)
	}
	return Poly{
		&poly{
			Points: vectors,
		},
		NewPrimative(NewCollider(NewShape(Polygon))),
	}
}

func (p Poly) Move(v Vector) Primative {
	newPoints := []Vector{}
	for i := range p.Points {
		newPoints = append(newPoints, p.Points[i].Add(v))
	}
	p.Points = newPoints
	return p
}

func (p Poly) Center() Vector {
	center := ZeroVector()
	for _, point := range p.Points {
		center = center.Add(point)
	}
	return center.Scale(1 / float64(len(p.Points)))
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

func (p Poly) Collides(other Collider) (Collision, bool) {
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

func (p Poly) Raster() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = p.C()
	for _, p := range p.Points {
		imd.Push(p.toPixelVec())
	}
	imd.Polygon(p.T())
	return imd
}
