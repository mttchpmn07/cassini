package engine

import "github.com/mttchpmn07/cassini/engine/events"

type BaseLayer struct {
	Name string
	App  *CassiniApp
}

func (bl *BaseLayer) SetApp(app *CassiniApp) {
	bl.App = app
}

type Layer interface {
	OnAttach()
	OnDetach()
	OnUpdate()
	OnEvent(event events.Event)
	SetApp(app *CassiniApp)
}

type DemoLayer struct {
	BaseLayer
}

func NewDemoLayer(name string) Layer {
	return &DemoLayer{
		BaseLayer: BaseLayer{
			Name: name,
		},
	}
}

func (l *DemoLayer) OnAttach() {}
func (l *DemoLayer) OnDetach() {}
func (l *DemoLayer) OnUpdate() {}
func (l *DemoLayer) OnEvent(event events.Event) {
	Log(event.Key())
}

type LayerStack interface {
	PushLayer(layer Layer)
	PushOverlay(overlay Layer)
	PopLayer(layer Layer)
	PopOverlay(overlay Layer)
	Get() []Layer
	PrintStack()
}

type layerStack struct {
	layers      []Layer
	insertLayer int
}

func NewLayerStack() LayerStack {
	return &layerStack{
		layers:      []Layer{},
		insertLayer: 0,
	}
}

func (ls *layerStack) find(layer Layer) int {
	for i, n := range ls.layers {
		if layer == n {
			return i
		}
	}
	return len(ls.layers)
}

func (ls *layerStack) PushLayer(layer Layer) {
	if ls.insertLayer == 0 {
		ls.layers = append([]Layer{layer}, ls.layers...)
	} else if ls.insertLayer == len(ls.layers) {
		ls.layers = append(ls.layers, layer)
	} else {
		ls.layers = append(ls.layers[:ls.insertLayer+1], ls.layers[ls.insertLayer:]...) // index < len(a)
		ls.layers[ls.insertLayer] = layer
	}
	ls.insertLayer += 1
}

func (ls *layerStack) PushOverlay(overlay Layer) {
	ls.layers = append(ls.layers, overlay)
}

func (ls *layerStack) PopLayer(layer Layer) {
	idx := ls.find(layer)
	ls.layers = append(ls.layers[:idx], ls.layers[idx+1:]...)
	ls.insertLayer -= 1

}

func (ls *layerStack) PopOverlay(overlay Layer) {
	idx := ls.find(overlay)
	ls.layers = append(ls.layers[:idx], ls.layers[idx+1:]...)
}

func (ls *layerStack) Get() []Layer {
	return ls.layers
}

func (ls *layerStack) PrintStack() {
	for _, l := range ls.layers {
		Print(l.(*DemoLayer).Name)
	}
	Print("")
}
