package engine

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var Keys = map[string]int{
	"MouseButtonLeft":   0,
	"MouseButtonRight":  1,
	"MouseButtonMiddle": 2,
	"KeyW":              87,
	"KeyA":              65,
	"KeyS":              83,
	"KeyD":              68,
	"KeyEscape":         256,
	"KeyEnter":          257,
}

type Window interface {
	GetWindow() *pixelgl.Window
	UpdateKeys()
}

type window struct {
	win *pixelgl.Window
}

func (w *window) GetWindow() *pixelgl.Window {
	return w.win
}

func (w *window) UpdateKeys() {
	if w.win.MouseInsideWindow() {
		mousePos := w.win.MousePosition()
		if mousePos != w.win.MousePreviousPosition() {
			GlobalEvents.Broadcast(NewEvent("mouseMove", FromPixelVec(mousePos)))
		}
	}
	for key, element := range Keys {
		if w.win.Pressed(pixelgl.Button(element)) {
			GlobalEvents.Broadcast(NewEvent(key+"_Pressed", nil))
		}
	}
	for key, element := range Keys {
		if w.win.JustPressed(pixelgl.Button(element)) {
			GlobalEvents.Broadcast(NewEvent(key+"_JustPressed", nil))
		}
	}
	for key, element := range Keys {
		if w.win.JustReleased(pixelgl.Button(element)) {
			GlobalEvents.Broadcast(NewEvent(key+"_JustReleased", nil))
		}
	}
}

func runWindow() {
	pixelgl.Run(Run)
}

func createWindow(title string, width, height float64) (*pixelgl.Window, error) {
	cfg := pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	return pixelgl.NewWindow(cfg)
}
