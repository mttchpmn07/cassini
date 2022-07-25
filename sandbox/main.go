package main

import (
	"github.com/mttchpmn07/cassini/engine"
)

type testLayer struct {
	engine.BaseLayer
	Name            string
	Sprite          *engine.DrawObject
	Circle          engine.Circle
	Circles         []engine.Circle
	MouseCircle     engine.Circle
	Rect            engine.Rect
	Line            engine.Line
	Poly            engine.Polygon
	DragLine        engine.Line
	DragLineStarted bool
	SpriteLocs      []engine.Vector
}

func NewTestLyer() *testLayer {
	tl := &testLayer{
		BaseLayer:       engine.BaseLayer{},
		Name:            "Test Render Layer",
		Sprite:          nil,
		Circle:          engine.NewCircle(100, engine.Vec(500, 500)),
		Circles:         []engine.Circle{},
		MouseCircle:     engine.NewCircle(50, engine.Vec(200, 200)),
		Rect:            engine.NewRect(engine.Vec(700, 700), engine.Vec(200, 200)),
		Line:            engine.NewLine(engine.Vec(200, 700), engine.Vec(700, 200)),
		Poly:            engine.NewPolygon([]engine.Vector{engine.Vec(300, 300), engine.Vec(200, 200), engine.Vec(0, 250)}...),
		DragLine:        engine.NewLine(engine.Vec(200, 700), engine.Vec(700, 200)),
		DragLineStarted: false,
		SpriteLocs:      []engine.Vector{},
	}
	var err error
	tl.Sprite, err = engine.NewDrawObject("./celebrate.png", engine.Vec(500, 500), 0, 0.5)
	if err != nil {
		panic(err)
	}

	return tl
}

func (l *testLayer) OnAttach() {}
func (l *testLayer) OnDetach() {}
func (l *testLayer) OnUpdate() {
	l.App.Ren.OpenBatch(l.Sprite.Spritesheet)
	l.App.Ren.DrawQuad(l.Rect)
	for _, loc := range l.SpriteLocs {
		l.App.Ren.DrawSprite(l.Sprite.Moved(loc))
	}
	//l.App.Ren.DrawSprite(l.Sprite)
	l.App.Ren.DrawCircle(l.Circle)
	for _, c := range l.Circles {
		l.App.Ren.DrawCircle(c)
	}
	l.App.Ren.DrawCircle(l.MouseCircle)
	l.App.Ren.DrawLine(l.Line)
	l.App.Ren.DrawLine(l.DragLine)
	l.App.Ren.DrawPoly(l.Poly)
	l.App.Ren.CloseBatch()
}

func (l *testLayer) OnEvent(event engine.Event) {
	mousePos := *event.Contents().(*engine.Vector)
	if event.Key() == "mouseMove" {
		l.MouseCircle = engine.NewCircle(50, mousePos)
		if l.DragLineStarted {
			l.DragLine.End = mousePos
		}
	}
	if event.Key() == "MOUSE_BUTTON_LEFT_Pressed" {
		//l.Circles = append(l.Circles, l.MouseCircle)
		l.SpriteLocs = append(l.SpriteLocs, mousePos)
	}
	if event.Key() == "MOUSE_BUTTON_RIGHT_JustPressed" {
		if l.DragLineStarted {
			l.DragLine.End = mousePos
			l.DragLineStarted = false
		} else {
			l.DragLine.Start = mousePos
			l.DragLine.End = mousePos
			l.DragLineStarted = true
		}
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
