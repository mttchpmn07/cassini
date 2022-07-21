package engine

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	*glfw.Window
}

func NewWindow(width int, height int, title string) (Window, error) {
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	return Window{
		window,
	}, nil
}
