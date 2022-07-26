package gui

import (
	"github.com/mttchpmn07/cassini/engine"
)

type GUI struct {
	engine.BaseLayer
}

func NewGUI(name string) engine.Layer {
	return &GUI{
		BaseLayer: engine.BaseLayer{
			Name: name,
		},
	}
}

func (g *GUI) OnAttach()                  {}
func (g *GUI) OnDetach()                  {}
func (g *GUI) OnUpdate()                  {}
func (g *GUI) OnEvent(event engine.Event) {}

type Clickable interface {
	ClickRegion()
	ClickCallback()
}
