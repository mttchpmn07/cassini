package main

import (
	"github.com/mttchpmn07/cassini/engine"
)

/*
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)

	for !win.Closed() {
		win.Update()
	}
}
*/

func main() {
	app := engine.NewCassiniApp()
	layer1 := engine.NewLayer("layer1")
	layer2 := engine.NewLayer("layer2")
	overlay1 := engine.NewLayer("overlay1")
	overlay2 := engine.NewLayer("overlay2")
	app.PushLayer(layer1)
	app.PushLayer(layer2)
	app.PushOverlay(overlay1)
	app.PushOverlay(overlay2)

	//pixelgl.Run(engine.Run)
	engine.Start(app)
}
