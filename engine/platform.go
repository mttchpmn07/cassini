package engine

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func NewPlatform(title string, width float64, height float64) (*pixelgl.Window, error) {
	cfg := pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, width, height),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return nil, err
	}
	return win, nil
}
