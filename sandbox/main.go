package main

import (
	"log"

	"github.com/mttchpmn07/cassini/engine"
)

type GUIOverlay struct {
	name string
}

func (g *GUIOverlay) OnAttach() {}
func (g *GUIOverlay) OnDetach() {}
func (g *GUIOverlay) OnUpdate() {}
func (g *GUIOverlay) OnEvent(event engine.Event) {
	if event.Key() == "mouseMove" {
		log.Printf("<%v, %v>\n", event.Contents().(engine.Vector).X, event.Contents().(engine.Vector).Y)
	} else {
		log.Printf("%v pressed\n", event.Key())
	}
}
func (g *GUIOverlay) Name() string {
	return g.name
}

func main() {
	appConfig := engine.AppConfig{
		Title:  "Cassini Test App",
		Width:  1024,
		Height: 768,
	}
	app := engine.NewCassiniApp(appConfig)
	guiOverlay := &GUIOverlay{
		name: "GUI",
	}
	app.PushOverlay(guiOverlay)

	engine.Start(app)
}
