package gui

import (
	"github.com/faiface/pixel"
	"github.com/mttchpmn07/cassini/engine"
	"github.com/mttchpmn07/cassini/engine/events"
	p "github.com/mttchpmn07/cassini/engine/primatives"
	"golang.org/x/image/colornames"
)

type Button struct {
	p.Primative
}

func NewButton(region p.Primative) Clickable {
	button := &Button{
		region,
	}
	return button
}

func (b *Button) ClickCallback(d events.Publisher) {
	engine.Log("Button Clicked (in button)")
	d.Broadcast(events.NewEvent("button_click", pixel.V(0, 0)))
	b.Animate()
}

func (b *Button) Animate() {
	if b.C() == colornames.White {
		b.Color(colornames.Blue)
	} else {
		b.Color(colornames.White)
	}
}
