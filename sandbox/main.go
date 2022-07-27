package main

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/mttchpmn07/cassini/engine"
	"github.com/mttchpmn07/cassini/engine/events"
	"github.com/mttchpmn07/cassini/engine/graphics"
	m "github.com/mttchpmn07/cassini/engine/math"
	p "github.com/mttchpmn07/cassini/engine/primatives"
	"golang.org/x/image/colornames"
)

type testLayer struct {
	engine.BaseLayer
	Sprite          *graphics.DrawObject
	Shapes          []p.Primative
	NoCollideShapes []p.Primative
	MousePoint      p.Primative
	MousePoly       p.Primative
	DragLine        p.Primative
	DragLineStarted bool
	SpriteLocs      []m.Vector
	shapeSelector   int
}

func NewTestLyer() *testLayer {
	tl := &testLayer{
		BaseLayer: engine.BaseLayer{
			Name: "Test Render Layer",
		},
		Sprite:          nil,
		Shapes:          []p.Primative{},
		MousePoint:      p.NewDot(200, 200),
		MousePoly:       p.NewPolygon([]m.Vector{m.NewVector(math.Inf(1), math.Inf(1)), m.NewVector(math.Inf(1), math.Inf(1)), m.NewVector(math.Inf(1), math.Inf(1))}...),
		DragLine:        p.NewLine(m.NewVector(math.Inf(1), math.Inf(1)), m.NewVector(math.Inf(1), math.Inf(1))),
		DragLineStarted: false,
		SpriteLocs:      []m.Vector{},
	}
	tl.Shapes = append(tl.Shapes, p.NewCircle(100, m.NewVector(600, 500)))
	tl.Shapes = append(tl.Shapes, p.NewPolygon([]m.Vector{m.NewVector(0, 250), m.NewVector(300, 300), m.NewVector(200, 200)}...))
	tl.Shapes = append(tl.Shapes, p.NewPolygon([]m.Vector{m.NewVector(0, 250), m.NewVector(300, 300), m.NewVector(200, 200)}...).Move(m.NewVector(500, 300)))
	tl.Shapes = append(tl.Shapes, p.NewPolygon([]m.Vector{m.NewVector(0, 250), m.NewVector(300, 300), m.NewVector(200, 200)}...).Move(m.NewVector(0, 400)))
	tl.Shapes = append(tl.Shapes, p.NewRectangle(m.NewVector(0, 0), m.NewVector(100, 100)))
	tl.Shapes = append(tl.Shapes, p.NewRectangle(m.NewVector(800, 800), m.NewVector(700, 700)))
	tl.Shapes = append(tl.Shapes, p.NewRectangle(m.NewVector(0, 800), m.NewVector(100, 700)))
	tl.Shapes = append(tl.Shapes, p.NewLine(m.NewVector(0, 0), m.NewVector(1500, 1500)))
	tl.Shapes = append(tl.Shapes, p.NewDot(600, 50))

	var err error
	tl.Sprite, err = graphics.NewDrawObject("./celebrate.png", pixel.V(500, 500), 0, 0.5)
	if err != nil {
		panic(err)
	}

	return tl
}

func MouseShape(mousePos m.Vector, shapeSelector int) p.Primative {
	switch shapeSelector {
	case 0:
		return p.NewDot(mousePos.X, mousePos.Y)
	case 1:
		return p.NewPolygon([]m.Vector{m.NewVector(50, 0), m.NewVector(50, 50), m.NewVector(0, 50)}...).Move(mousePos)
	case 2:
		return p.NewRectangle(m.NewVector(0, 0), m.NewVector(50, 50)).Move(mousePos)
	case 4:
		return p.NewLine(m.NewVector(25, 0), m.NewVector(0, 25)).Move(mousePos)
	}
	return p.NewCircle(20, mousePos)
}

func (l *testLayer) OnAttach() {}
func (l *testLayer) OnDetach() {}
func (l *testLayer) OnUpdate() {
	l.App.Ren.OpenBatch(l.Sprite.Spritesheet)
	for _, loc := range l.SpriteLocs {
		l.App.Ren.DrawSprite(l.Sprite.Moved(loc.X, loc.Y))
	}
	for _, s := range l.Shapes {
		l.App.Ren.Draw(s)
	}
	for _, s := range l.NoCollideShapes {
		l.App.Ren.Draw(s)
	}
	l.App.Ren.DrawShape(l.MousePoint)
	l.App.Ren.DrawShape(l.DragLine)
	l.App.Ren.DrawShape(l.MousePoly)
	l.App.Ren.CloseBatch()
}

func (l *testLayer) OnEvent(event events.Event) {
	mousePos := m.VectorFromEvent(event)
	mousePoint := p.NewDot(mousePos.X, mousePos.Y)
	switch event.Key() {
	case "mouseMove":
		l.MousePoly = MouseShape(mousePos, l.shapeSelector)
		l.MousePoint = mousePoint
		if l.DragLineStarted {
			l.DragLine.(p.Lin).End = mousePos
		}
		l.App.Ren.SetColor(colornames.White)
		for _, s := range l.Shapes {
			if _, col := p.Collides(l.MousePoly, s); col {
				newShape := MouseShape(mousePos, l.shapeSelector)
				newShape.Color(colornames.Red)
				l.NoCollideShapes = append(l.NoCollideShapes, newShape)
			}
		}
	case "MOUSE_BUTTON_LEFT_Pressed":
		l.SpriteLocs = append(l.SpriteLocs, mousePos)
	case "MOUSE_BUTTON_RIGHT_JustPressed":
		engine.Log(fmt.Sprintf("%v", mousePos))
		if l.DragLineStarted {
			l.DragLine.(p.Lin).End = mousePos
			l.DragLineStarted = false
		} else {
			l.Shapes = append(l.Shapes, p.NewLine(l.DragLine.(p.Lin).Start, l.DragLine.(p.Lin).End))
			l.DragLine.(p.Lin).Start = mousePos
			l.DragLine.(p.Lin).End = mousePos
			l.DragLineStarted = true
		}
	case "KEY_RIGHT_Pressed":
		for i := range l.Shapes {
			l.Shapes[i] = l.Shapes[i].Move(m.NewVector(10, 0))
		}
	case "KEY_LEFT_Pressed":
		for i := range l.Shapes {
			l.Shapes[i] = l.Shapes[i].Move(m.NewVector(-10, 0))
		}
	case "KEY_DOWN_Pressed":
		for i := range l.Shapes {
			l.Shapes[i] = l.Shapes[i].Move(m.NewVector(0, -10))
		}
	case "KEY_UP_Pressed":
		for i := range l.Shapes {
			l.Shapes[i] = l.Shapes[i].Move(m.NewVector(0, 10))
		}
	case "KEY_SPACE_JustPressed":
		l.shapeSelector += 1
		if l.shapeSelector > 5 {
			l.shapeSelector = 0
		}
	default:
	}
}

func main() {
	config := engine.AppConfig{
		Title:  "Test App",
		Width:  1024,
		Height: 796,
	}
	app := engine.InitApp(config)
	//app.PushOverlay(engine.NewDemoLayer("DemoLayer"))

	app.PushLayer(NewTestLyer())

	engine.Run()
}
