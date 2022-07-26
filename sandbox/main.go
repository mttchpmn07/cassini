package main

import (
	"fmt"
	"math"

	"github.com/mttchpmn07/cassini/engine"
	"golang.org/x/image/colornames"
)

type testLayer struct {
	engine.BaseLayer
	Name            string
	Sprite          *engine.DrawObject
	Shapes          []engine.Shape
	NoCollideShapes []engine.Shape
	MousePoint      engine.Shape
	MousePoly       engine.Shape
	DragLine        engine.Lin
	DragLineStarted bool
	SpriteLocs      []engine.Vector
	shapeSelector   int
}

func NewTestLyer() *testLayer {
	tl := &testLayer{
		BaseLayer:       engine.BaseLayer{},
		Name:            "Test Render Layer",
		Sprite:          nil,
		Shapes:          []engine.Shape{},
		MousePoint:      engine.NewVector(200, 200),
		MousePoly:       engine.NewPolygon([]engine.Vector{engine.NewVector(math.Inf(1), math.Inf(1)), engine.NewVector(math.Inf(1), math.Inf(1)), engine.NewVector(math.Inf(1), math.Inf(1))}...),
		DragLine:        engine.NewLine(engine.NewVector(math.Inf(1), math.Inf(1)), engine.NewVector(math.Inf(1), math.Inf(1))),
		DragLineStarted: false,
		SpriteLocs:      []engine.Vector{},
	}
	tl.Shapes = append(tl.Shapes, engine.NewCircle(100, engine.NewVector(600, 500)))
	tl.Shapes = append(tl.Shapes, engine.NewPolygon([]engine.Vector{engine.NewVector(0, 250), engine.NewVector(300, 300), engine.NewVector(200, 200)}...))
	tl.Shapes = append(tl.Shapes, engine.NewPolygon([]engine.Vector{engine.NewVector(0, 250), engine.NewVector(300, 300), engine.NewVector(200, 200)}...).Move(engine.NewVector(500, 300)))
	tl.Shapes = append(tl.Shapes, engine.NewPolygon([]engine.Vector{engine.NewVector(0, 250), engine.NewVector(300, 300), engine.NewVector(200, 200)}...).Move(engine.NewVector(0, 400)))
	tl.Shapes = append(tl.Shapes, engine.NewRectangle(engine.NewVector(0, 0), engine.NewVector(100, 100)))
	tl.Shapes = append(tl.Shapes, engine.NewRectangle(engine.NewVector(800, 800), engine.NewVector(700, 700)))
	tl.Shapes = append(tl.Shapes, engine.NewRectangle(engine.NewVector(0, 800), engine.NewVector(100, 700)))
	tl.Shapes = append(tl.Shapes, engine.NewLine(engine.NewVector(0, 0), engine.NewVector(1500, 1500)))
	tl.Shapes = append(tl.Shapes, engine.NewVector(600, 50))
	//tl.Shapes = append(tl.Shapes, engine.NewRectangle(engine.NewVector(400, 400), engine.NewVector(500, 500)))

	var err error
	tl.Sprite, err = engine.NewDrawObject("./celebrate.png", engine.NewVector(500, 500), 0, 0.5)
	if err != nil {
		panic(err)
	}

	return tl
}

func MouseShape(mousePos engine.Vector, shapeSelector int) engine.Shape {
	switch shapeSelector {
	case 0:
		return engine.NewVector(mousePos.X, mousePos.Y)
	case 1:
		return engine.NewPolygon([]engine.Vector{engine.NewVector(50, 0), engine.NewVector(50, 50), engine.NewVector(0, 50)}...).Move(mousePos)
	case 2:
		return engine.NewRectangle(engine.NewVector(0, 0), engine.NewVector(50, 50)).Move(mousePos)
	case 4:
		return engine.NewLine(engine.NewVector(10, 0), engine.NewVector(0, 10)).Move(mousePos)
	}
	return engine.NewCircle(5, mousePos)
}

func (l *testLayer) OnAttach() {}
func (l *testLayer) OnDetach() {}
func (l *testLayer) OnUpdate() {
	l.App.Ren.OpenBatch(l.Sprite.Spritesheet)
	for _, loc := range l.SpriteLocs {
		l.App.Ren.DrawSprite(l.Sprite.Moved(loc))
	}
	l.App.Ren.DrawShapes(l.Shapes)
	l.App.Ren.DrawShapes(l.NoCollideShapes)
	l.App.Ren.DrawShape(l.MousePoint)
	l.App.Ren.DrawShape(l.DragLine)
	l.App.Ren.DrawShape(l.MousePoly)
	l.App.Ren.CloseBatch()
}

func (l *testLayer) OnEvent(event engine.Event) {
	mousePos := engine.VectorFromEvent(event)
	mousePoint := engine.Shape(mousePos)
	switch event.Key() {
	case "mouseMove":
		l.MousePoly = MouseShape(mousePos, l.shapeSelector)
		l.MousePoint = mousePoint
		if l.DragLineStarted {
			l.DragLine.End = mousePos
		}
		l.App.Ren.SetColor(colornames.White)
		for _, s := range l.Shapes {
			//if _, col := engine.Collides(mousePoint, s); col {
			if _, col := engine.Collides(l.MousePoly, s); col {
				l.App.Ren.SetColor(colornames.Red)
				l.NoCollideShapes = append(l.NoCollideShapes, MouseShape(mousePos, l.shapeSelector))
				//l.NoCollideShapes = append(l.NoCollideShapes, engine.NewVector(mousePos.X, mousePos.Y))
			}
		}
	case "MOUSE_BUTTON_LEFT_Pressed":
		//l.Circles = append(l.Circles, l.MouseCircle)
		l.SpriteLocs = append(l.SpriteLocs, mousePos)
	case "MOUSE_BUTTON_RIGHT_JustPressed":
		engine.Log(fmt.Sprintf("%v", mousePos))
		if l.DragLineStarted {
			l.DragLine.End = mousePos
			l.DragLineStarted = false
		} else {
			l.Shapes = append(l.Shapes, engine.NewLine(l.DragLine.Start, l.DragLine.End))
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

	app.PushLayer(NewTestLyer())

	engine.Run()
}
