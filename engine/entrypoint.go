package engine

import "github.com/faiface/pixel/pixelgl"

var GlobalApplication Application

func init() {
	Print("Cassini Game Engine veion 0")
}

func InitApp(config AppConfig) Application {
	GlobalApplication = NewCassiniApp(config)
	return GlobalApplication
}

func Run() {
	pixelgl.Run(run)
}

func run() {
	dispatcher := NewPublisher()
	dispatcher.Listen(GlobalApplication)
	platform, err := NewPlatform(
		GlobalApplication.GetConfig().Title,
		GlobalApplication.GetConfig().Width,
		GlobalApplication.GetConfig().Height,
	)
	if err != nil {
		panic(err)
	}
	GlobalApplication.Run(platform, dispatcher)
}
