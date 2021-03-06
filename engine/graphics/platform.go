package graphics

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"github.com/mttchpmn07/cassini/engine/events"
)

type Platform struct {
	*pixelgl.Window
}

type Button struct {
	*pixelgl.Button
}

func ButtonFromInt(i int) Button {
	b := pixelgl.Button(i)
	return Button{&b}
}

func NewPlatform(title string, width float64, height float64) (*Platform, error) {
	cfg := pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, width, height),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return nil, err
	}
	return &Platform{win}, nil
}

func (p *Platform) UpdateWindow(dispatcher events.Publisher) {
	p.Update()
	p.updateKeys(dispatcher)

}

func (p *Platform) updateKeys(dispatcher events.Publisher) {
	mousePos := pixel.V(math.Inf(1), math.Inf(1))
	if p.MouseInsideWindow() {
		mousePos = p.MousePosition()
		if mousePos != p.MousePreviousPosition() {
			dispatcher.Broadcast(events.NewEvent("mouseMove", mousePos))
		}
	}
	for key, element := range KeyMap {
		if p.Press(ButtonFromInt(element)) {
			dispatcher.Broadcast(events.NewEvent(key+"_Pressed", mousePos))
		}
	}
	for key, element := range KeyMap {
		if p.Tap(ButtonFromInt(element)) {
			dispatcher.Broadcast(events.NewEvent(key+"_JustPressed", mousePos))
		}
	}
	for key, element := range KeyMap {
		if p.Release(ButtonFromInt(element)) {
			dispatcher.Broadcast(events.NewEvent(key+"_JustReleased", mousePos))
		}
	}
}

/*
func (p *Platform) MousePos() Vector {
	return fromPixelVec(p.MousePosition())
}

func (p *Platform) MousePrevPos() Vector {
	return fromPixelVec(p.MousePreviousPosition())
}
*/

func (p *Platform) Press(button Button) bool {
	return p.Pressed(*button.Button)
}

func (p *Platform) Tap(button Button) bool {
	return p.JustPressed(*button.Button)
}

func (p *Platform) Release(button Button) bool {
	return p.JustReleased(*button.Button)
}
