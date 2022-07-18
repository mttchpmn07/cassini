package engine

import (
	"fmt"
)

type Application interface {
	Run()
	OnEvent(event Event)
	PushLayer(layer Layer)
	PushOverlay(overlay Layer)
}

type cassiniApp struct {
	Log    LogLevel
	Layers LayerStack
}

func NewCassiniApp() Application {
	app := &cassiniApp{
		Log:    Err,
		Layers: NewLayerStack(),
	}

	return app
}

func (c cassiniApp) Run() {
	if c.Log == Trace {
		LogTrace("Enter: func (c CassiniApp) Run()")
	}

	Print("Hello from my cassini app!")

	layers := c.Layers.Get()
	for _, l := range layers {
		fmt.Println(l.(*layer).name)
		l.OnUpdate()
	}

	if c.Log == Trace {
		LogTrace("Exit: func (c CassiniApp) Run()")
	}
}

func (c cassiniApp) OnEvent(event Event) {
	if c.Log == Trace {
		LogTrace("Enter: func (c cassiniApp) OnEvent(event Event)")
	}

	layers := c.Layers.Get()
	for idx := len(layers) - 1; idx >= 0; idx-- {
		fmt.Println(layers[idx].(*layer).name)
		layers[idx].OnEvent(event)
		if event.Handled() {
			break
		}
	}

	if c.Log == Trace {
		LogTrace("Exit: func (c cassiniApp) OnEvent(event Event)")
	}
}

func (c cassiniApp) PushLayer(layer Layer) {
	c.Layers.PushLayer(layer)
}

func (c cassiniApp) PushOverlay(overlay Layer) {
	c.Layers.PushOverlay(overlay)
}
