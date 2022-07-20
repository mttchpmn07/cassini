package main

import (
	"github.com/mttchpmn07/cassini/engine"
)

func main() {
	appConfig := engine.AppConfig{
		Title:  "Cassini Test App",
		Width:  1024,
		Height: 768,
	}
	app := engine.NewCassiniApp(appConfig)
	layer := engine.NewDemoLayer("Demo Layer")
	app.PushOverlay(layer)

	engine.Start(app)
}
