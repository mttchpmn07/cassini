package gui

import (
	"github.com/mttchpmn07/cassini/engine"
	"github.com/mttchpmn07/cassini/engine/events"
	m "github.com/mttchpmn07/cassini/engine/math"
	p "github.com/mttchpmn07/cassini/engine/primatives"
	"golang.org/x/image/colornames"
)

type GUI struct {
	engine.BaseLayer
	elements []Clickable
}

func NewGUI(name string) engine.Layer {
	firstButton := NewButton(p.NewRectangle(m.NewVector(100, 100), m.NewVector(200, 200)))
	firstButton.Color(colornames.Blue)
	firstButton.Thickness(0)
	gui := &GUI{
		BaseLayer: engine.BaseLayer{
			Name: name,
		},
		elements: []Clickable{
			firstButton,
		},
	}
	return gui
}

func (g *GUI) OnAttach() {}
func (g *GUI) OnDetach() {}
func (g *GUI) OnUpdate() {
	//engine.LogTrace("Enter OnUpdate")
	g.App.Ren.OpenBatch(nil)
	for _, b := range g.elements {
		//engine.Log(fmt.Sprintf("%v", g.App))
		if g.App != nil {
			g.App.Ren.Draw(b)
		}
	}
	g.App.Ren.CloseBatch()
	//engine.LogTrace("Exit OnUpdate")
}
func (g *GUI) OnEvent(event events.Event) {
	if event.Handled() {
		return
	}
	mousePos := m.VectorFromEvent(event)
	mousePoint := p.NewDot(mousePos.X, mousePos.Y)
	switch event.Key() {
	case "MOUSE_BUTTON_LEFT_JustPressed":
		for _, b := range g.elements {
			if click, _ := b.Collides(mousePoint); click {
				b.ClickCallback(g.App.Dis)
				event.Handle()
			}
		}
	}
}
func (g *GUI) Broadcast(event events.Event) {
	g.App.Broadcast(event)
}

type Clickable interface {
	p.Primative
	ClickCallback(d events.Publisher)
}
