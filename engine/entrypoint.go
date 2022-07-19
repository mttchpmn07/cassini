package engine

var GlobalEvents Publisher
var GlobalApplication Application
var GlobalWindow Window

func init() {
	Print("Cassini Game Engine version 0")
	GlobalEvents = NewPublisher()
	/*
		LogInfo("Example Info")
		LogWarn("Example Warning")
		LogErr("Example Error")
		LogTrace("Example Trace")
	*/
}

func Start(app Application) {
	GlobalEvents.Listen(app)
	GlobalApplication = app
	runWindow()
}

func Run() {
	config := GlobalApplication.GetConfig()
	win, err := createWindow(
		config.Title,
		config.Width,
		config.Height)
	if err != nil {
		panic(err)
	}
	GlobalWindow = &window{
		win: win,
	}

	for !win.Closed() {
		GlobalWindow.UpdateKeys()
		GlobalApplication.Run()
		win.Update()
	}
}
