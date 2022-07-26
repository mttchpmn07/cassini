package engine

type shape int

const (
	Nothing shape = iota
	Point
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
}

type ConcreteShape struct {
	kind shape
}

func NewShape(kind shape) *ConcreteShape {
	return &ConcreteShape{
		kind: kind,
	}
}

func (cs *ConcreteShape) Type() shape {
	return cs.kind
}
