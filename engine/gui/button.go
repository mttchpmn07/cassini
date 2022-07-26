package gui

import (
	"github.com/faiface/pixel/imdraw"
	"github.com/mttchpmn07/cassini/engine/primatives"
)

type Button struct {
	Region primatives.Collider
	Visual []primatives.Collider
}

func NewButton(region primatives.Collider) *Button {
	button := &Button{
		Region: region,
	}
	return button
}

func (b *Button) ClickRegion()   {}
func (b *Button) ClickCallback() {}
func (b *Button) Raster() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	return imd
}
