package engine

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Window interface {
	GetWindow() *pixelgl.Window
}

type window struct {
	win *pixelgl.Window
}

func (w *window) GetWindow() *pixelgl.Window {
	return w.win
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
