package engine

import (
	"image"
	"os"

	"image/color"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Picture pixel.Picture
type DrawObject struct {
	Spritesheet Picture
	Frame       Rect
	Loc         Vector
	Angle       float64
	Scale       float64
}

func NewDrawObject(spritePath string, startLoc Vector, angle, scale float64) (*DrawObject, error) {
	pic, err := loadSpriteSheet(spritePath)
	if err != nil {
		return nil, err
	}
	return &DrawObject{
		Spritesheet: pic,
		Frame:       FromPixelRect(pic.Bounds()),
		Loc:         startLoc,
		Angle:       angle,
		Scale:       scale,
	}, nil
}

func (do DrawObject) Moved(loc Vector) *DrawObject {
	do.Loc = loc
	return &do
}

func loadSpriteSheet(path string) (Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	spritesheet := pixel.PictureDataFromImage(img)
	return spritesheet, nil
}

type drawProperties struct {
	color     color.Color
	thickness float64
}

type Renderer struct {
	platform   *Platform
	batches    []*pixel.Batch
	curBatch   int
	properties drawProperties
}

type RenderSystem interface {
	OpenBatch(spritesheet Picture)
	CloseBatch()
	BatchRender()
	SetColor(color color.Color)
	SetThickness(thickness float64)
	DrawSprite(do *DrawObject)
	DrawShape(shape Shape)
	DrawShapes(shapes []Shape)
}

func NewRenderSystem(platform *Platform) RenderSystem {
	return &Renderer{
		platform: platform,
		batches:  []*pixel.Batch{},
		curBatch: 0,
		properties: drawProperties{
			color:     colornames.White,
			thickness: 2,
		},
	}
}

func (ren *Renderer) OpenBatch(spritesheet Picture) {
	batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)
	ren.batches = append(ren.batches, batch)
}

func (ren *Renderer) CloseBatch() {
	ren.curBatch++
}

func (ren *Renderer) BatchRender() {
	for _, batch := range ren.batches {
		batch.Draw(ren.platform)
		batch.Clear()
	}
	ren.curBatch = 0
}

func (ren *Renderer) SetColor(color color.Color) {
	ren.properties.color = color
}

func (ren *Renderer) SetThickness(thickness float64) {
	ren.properties.thickness = thickness
}

func (ren *Renderer) DrawSprite(do *DrawObject) {
	trans := pixel.IM.Scaled(pixel.ZV, do.Scale).Rotated(pixel.ZV, do.Angle)
	sprite := pixel.NewSprite(do.Spritesheet, *do.Frame.Rect)
	sprite.Draw(ren.batches[ren.curBatch], trans.Moved(do.Loc.toPixelVec()))
}

func (ren *Renderer) DrawShape(shape Shape) {
	switch shape.Type() {
	case Point:
		ren.drawPoint(shape.(Vector))
	case Line:
		ren.drawLine(shape.(Lin))
	case Circle:
		ren.drawCircle(shape.(Circ))
	case Rectangle:
		ren.drawRectangle(shape.(Rect))
	case Polygon:
		ren.drawPolygon(shape.(Poly))
	default:
	}
}

func (ren *Renderer) DrawShapes(shapes []Shape) {
	for _, s := range shapes {
		ren.DrawShape(s)
	}
}

func (ren *Renderer) drawPoint(v Vector) {
	ren.drawCircle(NewCircle(1, v))
}

func (ren *Renderer) drawCircle(c Circ) {
	imd := imdraw.New(nil)
	imd.Color = ren.properties.color
	imd.Push(c.Center)
	imd.Circle(c.Radius, ren.properties.thickness)
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) drawRectangle(rect Rect) {
	imd := imdraw.New(nil)
	imd.Color = ren.properties.color
	imd.Push(rect.Min)
	imd.Push(pixel.V(rect.Min.X, rect.Max.Y))
	imd.Push(rect.Max)
	imd.Push(pixel.V(rect.Max.X, rect.Min.Y))
	imd.Polygon(ren.properties.thickness)
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) drawLine(line Lin) {
	imd := imdraw.New(nil)
	imd.Color = ren.properties.color
	imd.Push(line.Start.toPixelVec())
	imd.Push(line.End.toPixelVec())
	imd.Polygon(ren.properties.thickness)
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) drawPolygon(poly Poly) {
	imd := imdraw.New(nil)
	imd.Color = ren.properties.color
	for _, p := range poly.Points {
		imd.Push(p.toPixelVec())
	}
	imd.Polygon(ren.properties.thickness)
	imd.Draw(ren.batches[ren.curBatch])
}
