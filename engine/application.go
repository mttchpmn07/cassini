package engine

import "golang.org/x/image/colornames"

type Application interface {
	Run(platform *Platform, renderer RenderSystem, dispatcher Publisher)
	OnEvent(event Event)
	PushLayer(layer Layer)
	PushOverlay(overlay Layer)
	GetConfig() AppConfig
}

type AppConfig struct {
	Title         string
	Width, Height float64
}

type CassiniApp struct {
	Log    LogLevel
	Layers LayerStack
	Config AppConfig
	Plat   *Platform
	Ren    RenderSystem
	Dis    Publisher
}

func NewCassiniApp(config AppConfig) Application {
	app := &CassiniApp{
		Log:    Err,
		Layers: NewLayerStack(),
		Config: config,
		Plat:   nil,
		Ren:    nil,
		Dis:    nil,
	}

	return app
}

func (c *CassiniApp) GetConfig() AppConfig {
	return c.Config
}

func (c *CassiniApp) Run(platform *Platform, renderer RenderSystem, dispatcher Publisher) {
	if c.Log == Trace {
		LogTrace("Enter: func (c CassiniApp) Run()")
	}
	c.Plat = platform
	c.Ren = renderer
	c.Dis = dispatcher

	for !c.Plat.Closed() {
		c.Plat.Clear(colornames.Black)
		layers := c.Layers.Get()
		for _, l := range layers {
			l.OnUpdate()
		}
		c.Plat.UpdateWindow(dispatcher)
	}

	if c.Log == Trace {
		LogTrace("Exit: func (c CassiniApp) Run()")
	}
}

func (c *CassiniApp) OnEvent(event Event) {
	if c.Log == Trace {
		LogTrace("Enter: func (c cassiniApp) OnEvent(event Event)")
	}

	layers := c.Layers.Get()
	for idx := len(layers) - 1; idx >= 0; idx-- {
		layers[idx].OnEvent(event)
		if event.Handled() {
			break
		}
	}

	if c.Log == Trace {
		LogTrace("Exit: func (c cassiniApp) OnEvent(event Event)")
	}
}

func (c *CassiniApp) PushLayer(layer Layer) {
	layer.SetApp(c)
	c.Layers.PushLayer(layer)
}

func (c *CassiniApp) PushOverlay(overlay Layer) {
	overlay.SetApp(c)
	c.Layers.PushOverlay(overlay)
}
