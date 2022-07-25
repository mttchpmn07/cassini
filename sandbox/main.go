package main

import (
	"math"

	"github.com/mttchpmn07/cassini/engine"
)

type testLayer struct {
	engine.BaseLayer
	Name            string
	Sprite          *engine.DrawObject
	Shapes          []engine.Shape
	MouseCircle     engine.Circ
	DragLine        engine.Lin
	DragLineStarted bool
	SpriteLocs      []engine.Vector
}

func NewTestLyer() *testLayer {
	tl := &testLayer{
		BaseLayer:       engine.BaseLayer{},
		Name:            "Test Render Layer",
		Sprite:          nil,
		Shapes:          []engine.Shape{},
		MouseCircle:     engine.NewCircle(50, engine.NewVector(200, 200)),
		DragLine:        engine.NewLine(engine.NewVector(math.Inf(1), math.Inf(1)), engine.NewVector(math.Inf(1), math.Inf(1))),
		DragLineStarted: false,
		SpriteLocs:      []engine.Vector{},
	}
	tl.Shapes = append(tl.Shapes, engine.NewCircle(100, engine.NewVector(500, 500)))
	tl.Shapes = append(tl.Shapes, engine.NewPolygon([]engine.Vector{engine.NewVector(300, 300), engine.NewVector(200, 200), engine.NewVector(0, 250)}...))
	tl.Shapes = append(tl.Shapes, engine.NewLine(engine.NewVector(200, 700), engine.NewVector(700, 200)))
	tl.Shapes = append(tl.Shapes, engine.NewRectangle(engine.NewVector(700, 700), engine.NewVector(200, 200)))

	var err error
	tl.Sprite, err = engine.NewDrawObject("./celebrate.png", engine.NewVector(500, 500), 0, 0.5)
	if err != nil {
		panic(err)
	}

	return tl
}

func (l *testLayer) OnAttach() {}
func (l *testLayer) OnDetach() {}
func (l *testLayer) OnUpdate() {
	l.App.Ren.OpenBatch(l.Sprite.Spritesheet)
	for _, loc := range l.SpriteLocs {
		l.App.Ren.DrawSprite(l.Sprite.Moved(loc))
	}
	l.App.Ren.DrawShapes(l.Shapes)
	l.App.Ren.DrawShape(l.MouseCircle)
	l.App.Ren.DrawShape(l.DragLine)
	l.App.Ren.CloseBatch()
}

func (l *testLayer) OnEvent(event engine.Event) {
	mousePos := engine.VectorFromEvent(event)
	switch event.Key() {
	case "mouseMove":
		l.MouseCircle = engine.NewCircle(50, mousePos)
		if l.DragLineStarted {
			l.DragLine.End = mousePos
		}
	case "MOUSE_BUTTON_LEFT_Pressed":
		//l.Circles = append(l.Circles, l.MouseCircle)
		l.SpriteLocs = append(l.SpriteLocs, mousePos)
	case "MOUSE_BUTTON_RIGHT_JustPressed":
		if l.DragLineStarted {
			l.DragLine.End = mousePos
			l.DragLineStarted = false
		} else {
			l.DragLine.Start = mousePos
			l.DragLine.End = mousePos
			l.DragLineStarted = true
		}
	case "KEY_RIGHT_Pressed":
		for i := range l.Shapes {
			l.Shapes[i] = l.Shapes[i].Move(engine.NewVector(10, 0))
		}
	case "KEY_LEFT_Pressed":
		for i := range l.Shapes {
			l.Shapes[i] = l.Shapes[i].Move(engine.NewVector(-10, 0))
		}
	case "KEY_DOWN_Pressed":
		for i := range l.Shapes {
			l.Shapes[i] = l.Shapes[i].Move(engine.NewVector(0, -10))
		}
	case "KEY_UP_Pressed":
		for i := range l.Shapes {
			l.Shapes[i] = l.Shapes[i].Move(engine.NewVector(0, 10))
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

	//app.PushOverlay(engine.NewDemoLayer("Demo Overlay"))
	app.PushLayer(NewTestLyer())

	engine.Run()
}
