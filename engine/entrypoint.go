package engine

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/mttchpmn07/cassini/engine/events"
	"github.com/mttchpmn07/cassini/engine/graphics"
)

var GlobalApplication Application

func init() {
	Print("Cassini Game Engine version 0")
}

func InitApp(config AppConfig) Application {
	GlobalApplication = NewCassiniApp(config)
	return GlobalApplication
}

func Run() {
	pixelgl.Run(run)
}

func run() {
	platform, err := graphics.NewPlatform(
		GlobalApplication.GetConfig().Title,
		GlobalApplication.GetConfig().Width,
		GlobalApplication.GetConfig().Height,
	)
	if err != nil {
		panic(err)
	}
	renderer := graphics.NewRenderSystem(platform)
	dispatcher := events.NewPublisher()
	dispatcher.Listen(GlobalApplication)
	GlobalApplication.Run(platform, renderer, dispatcher)
}
