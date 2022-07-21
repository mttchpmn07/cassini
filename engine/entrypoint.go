package engine

var GlobalApplication Application

func init() {
	Print("Cassini Game Engine veion 0")
}

func Run(app Application) {
	dispatcher := NewPublisher()
	GlobalApplication = app
	app.Run(nil, dispatcher)
}
