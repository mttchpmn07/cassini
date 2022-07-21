package engine

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Platform struct {
	*pixelgl.Window
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

func (p *Platform) UpdateWindow(dispatcher Publisher) {
	p.Update()
	p.updateKeys(dispatcher)

}

func (p *Platform) updateKeys(dispatcher Publisher) {
	if p.MouseInsideWindow() {
		mousePos := p.MousePosition()
		if mousePos != p.MousePreviousPosition() {
			dispatcher.Broadcast(NewEvent("mouseMove", fromPixelVec(mousePos)))
		}
	}
	for key, element := range KeyMap {
		if p.Pressed(pixelgl.Button(element)) {
			dispatcher.Broadcast(NewEvent(key+"_Pressed", nil))
		}
	}
	for key, element := range KeyMap {
		if p.JustPressed(pixelgl.Button(element)) {
			dispatcher.Broadcast(NewEvent(key+"_JustPressed", nil))
		}
	}
	for key, element := range KeyMap {
		if p.JustReleased(pixelgl.Button(element)) {
			dispatcher.Broadcast(NewEvent(key+"_JustReleased", nil))
		}
	}
}
