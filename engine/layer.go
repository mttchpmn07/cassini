package engine

type Layer interface {
	OnAttach()
	OnDetach()
	OnUpdate()
	OnEvent(event Event)
}

type layer struct {
	name string
}

func NewLayer(name string) Layer {
	return &layer{
		name: name,
	}
}

func (l *layer) OnAttach()           {}
func (l *layer) OnDetach()           {}
func (l *layer) OnUpdate()           {}
func (l *layer) OnEvent(event Event) {}

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
		Print(l.(*layer).name)
	}
	Print("")
}
