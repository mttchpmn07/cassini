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

type RenderSystem interface {
	DrawSprite(do *DrawObject)
	DrawCircle(c Circle)
	DrawCircles(cs []Circle)
	DrawQuad(rect Rect)
	DrawQuads(rects []Rect)
	DrawLine(line Line)
	DrawLines(lines []Line)
	DrawPoly(poly Polygon)
	DrawPolys(polys []Polygon)
	OpenBatch(spritesheet Picture)
	CloseBatch()
	BatchRender()
	SetColor(color color.Color)
	SetThickness(thickness float64)
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

func NewRenderer(platform *Platform) RenderSystem {
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

func (ren *Renderer) DrawSprite(do *DrawObject) {
	trans := pixel.IM.Scaled(pixel.ZV, do.Scale).Rotated(pixel.ZV, do.Angle)
	sprite := pixel.NewSprite(do.Spritesheet, *do.Frame.Rect)
	sprite.Draw(ren.batches[ren.curBatch], trans.Moved(do.Loc.toPixelVec()))
}

func (ren *Renderer) DrawCircle(c Circle) {
	imd := imdraw.New(nil)
	imd.Color = ren.properties.color
	imd.Push(c.Center)
	imd.Circle(c.Radius, ren.properties.thickness)
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) DrawCircles(cs []Circle) {
	for _, c := range cs {
		ren.DrawCircle(c)
	}
}

func (ren *Renderer) DrawQuad(rect Rect) {
	imd := imdraw.New(nil)
	imd.Color = ren.properties.color
	imd.Push(rect.Min)
	imd.Push(pixel.V(rect.Min.X, rect.Max.Y))
	imd.Push(rect.Max)
	imd.Push(pixel.V(rect.Max.X, rect.Min.Y))
	imd.Polygon(ren.properties.thickness)
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) DrawQuads(rects []Rect) {
	for _, r := range rects {
		ren.DrawQuad(r)
	}
}

func (ren *Renderer) DrawLine(line Line) {
	imd := imdraw.New(nil)
	imd.Color = ren.properties.color
	imd.Push(line.Start.toPixelVec())
	imd.Push(line.End.toPixelVec())
	imd.Polygon(ren.properties.thickness)
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) DrawLines(lines []Line) {
	for _, l := range lines {
		ren.DrawLine(l)
	}
}

func (ren *Renderer) DrawPoly(poly Polygon) {
	imd := imdraw.New(nil)
	imd.Color = ren.properties.color
	for _, p := range poly.Points {
		imd.Push(p.toPixelVec())
	}
	imd.Polygon(ren.properties.thickness)
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) DrawPolys(polys []Polygon) {
	for _, p := range polys {
		ren.DrawPoly(p)
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
