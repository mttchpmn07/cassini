package engine

type Application interface {
	Run(window interface{}, dispatcher Publisher)
	OnEvent(event Event)
	PushLayer(layer Layer)
	PushOverlay(overlay Layer)
	GetConfig() AppConfig
}

type AppConfig struct {
	Title         string
	Width, Height float64
}

type cassiniApp struct {
	Log        LogLevel
	Layers     LayerStack
	Config     AppConfig
	window     interface{}
	dispatcher Publisher
}

func NewCassiniApp(config AppConfig) Application {
	app := &cassiniApp{
		Log:        Err,
		Layers:     NewLayerStack(),
		Config:     config,
		window:     nil,
		dispatcher: nil,
	}

	return app
}

func (c *cassiniApp) GetConfig() AppConfig {
	return c.Config
}

func (c *cassiniApp) Run(window interface{}, dispatcher Publisher) {
	if c.Log == Trace {
		LogTrace("Enter: func (c CassiniApp) Run()")
	}
	c.window = window
	c.dispatcher = dispatcher

	//Print("Hello from my cassini app!")

	layers := c.Layers.Get()
	for _, l := range layers {
		//fmt.Println(l.Name())
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
