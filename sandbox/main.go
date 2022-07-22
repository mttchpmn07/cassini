package main

import (
	"github.com/mttchpmn07/cassini/engine"
)

type testLayer struct {
	engine.BaseLayer
	Name        string
	Sprite      *engine.DrawObject
	Circle      engine.Circle
	Circles     []engine.Circle
	MouseCircle engine.Circle
	Rect        engine.Rect
	Line        engine.Line
	Poly        engine.Polygon
}

func (l *testLayer) OnAttach() {}
func (l *testLayer) OnDetach() {}
func (l *testLayer) OnUpdate() {
	l.App.Ren.DrawQuad(l.Rect)
	l.App.Ren.DrawSprite(l.Sprite)
	l.App.Ren.DrawCircle(l.Circle)
	for _, c := range l.Circles {
		l.App.Ren.DrawCircle(c)
	}
	l.App.Ren.DrawCircle(l.MouseCircle)
	l.App.Ren.DrawLine(l.Line)
	l.App.Ren.DrawPoly(l.Poly)
}

func (l *testLayer) OnEvent(event engine.Event) {
	engine.Log(event.Key())
	if event.Key() == "mouseMove" {
		l.MouseCircle = engine.NewCircle(50, *event.Contents().(*engine.Vector))
	}
	if event.Key() == "MOUSE_BUTTON_LEFT_JustPressed" {
		l.Circles = append(l.Circles, l.MouseCircle)
	}
}

func main() {
	config := engine.AppConfig{
		Title:  "Test App",
		Width:  1024,
		Height: 796,
	}
	app := engine.InitApp(config)
	app.PushOverlay(engine.NewDemoLayer("Demo Overlay"))

	tl := &testLayer{
		Name:   "Test Render Layer",
		Sprite: nil,
	}
	pic, err := engine.LoadPicture("./celebrate.png")
	if err != nil {
		panic(err)
	}
	tl.Rect = engine.NewRect(engine.Vec(700, 700), engine.Vec(200, 200))
	tl.Circle = engine.NewCircle(100, engine.Vec(500, 500))
	tl.MouseCircle = engine.NewCircle(50, engine.Vec(200, 200))
	tl.Line = engine.NewLine(engine.Vec(200, 700), engine.Vec(700, 200))
	tl.Poly = engine.NewPolygon([]engine.Vector{
		engine.Vec(300, 300),
		engine.Vec(200, 200),
		engine.Vec(0, 250),
	}...)
	tl.Sprite = &engine.DrawObject{
		Spritesheet: pic,
		Frame:       engine.FromPixelRect(pic.Bounds()),
		Loc:         engine.Vec(500, 500),
		Angle:       0,
		Scale:       0.5,
	}
	app.PushLayer(tl)

	engine.Run()
}
