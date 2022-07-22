package engine

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

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

type Picture pixel.Picture
type Rect *pixel.Rect

func NewRect(min Vector, max Vector) Rect {
	return &pixel.Rect{
		Min: min.toPixelVec(),
		Max: max.toPixelVec(),
	}
}

func FromPixelRect(rect pixel.Rect) Rect {
	return Rect(&rect)
}

type Circle *pixel.Circle

func NewCircle(radius float64, location Vector) Circle {
	return &pixel.Circle{
		Radius: radius,
		Center: location.toPixelVec(),
	}
}

type Line *line
type line struct {
	Start Vector
	End   Vector
}

func NewLine(start Vector, end Vector) Line {
	return &line{
		Start: start,
		End:   end,
	}
}

type Polygon *poly
type poly struct {
	Points []Vector
}

func NewPolygon(points ...Vector) Polygon {
	var vectors []Vector
	vectors = append(vectors, points...)
	return &poly{
		Points: vectors,
	}
}

type RenderSystem interface {
	DrawSprite(do *DrawObject)
	DrawCircle(c Circle)
	DrawQuad(rect Rect)
	DrawLine(line Line)
	DrawPoly(poly Polygon)
	OpenBatch(spritesheet Picture)
	CloseBatch()
	BatchRender()
}

type Renderer struct {
	platform *Platform
	batches  []*pixel.Batch
	curBatch int
}

func NewRenderer(platform *Platform) RenderSystem {
	return &Renderer{
		platform: platform,
		batches:  []*pixel.Batch{},
		curBatch: 0,
	}
}

func (ren *Renderer) OpenBatch(spritesheet Picture) {
	batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)
	ren.batches = append(ren.batches, batch)
}

func (ren *Renderer) CloseBatch() {
	ren.curBatch++
}

func (ren *Renderer) DrawSprite(do *DrawObject) {
	trans := pixel.IM.Scaled(pixel.ZV, do.Scale).Rotated(pixel.ZV, do.Angle)
	sprite := pixel.NewSprite(do.Spritesheet, *do.Frame)
	sprite.Draw(ren.batches[ren.curBatch], trans.Moved(do.Loc.toPixelVec()))
}

func (ren *Renderer) DrawCircle(c Circle) {
	imd := imdraw.New(nil)
	imd.Color = colornames.White
	imd.Push(c.Center)
	imd.Circle(c.Radius, 2)
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) DrawQuad(rect Rect) {
	imd := imdraw.New(nil)
	imd.Color = colornames.Blue
	imd.Push(rect.Min)
	imd.Push(pixel.V(rect.Min.X, rect.Max.Y))
	imd.Push(rect.Max)
	imd.Push(pixel.V(rect.Max.X, rect.Min.Y))
	imd.Polygon(0)
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) DrawLine(line Line) {
	imd := imdraw.New(nil)
	imd.Color = colornames.Green
	imd.Push(line.Start.toPixelVec())
	imd.Push(line.End.toPixelVec())
	imd.Polygon(2)
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) DrawPoly(poly Polygon) {
	imd := imdraw.New(nil)
	imd.Color = colornames.Red
	for _, p := range poly.Points {
		imd.Push(p.toPixelVec())
	}
	imd.Polygon(2)
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) BatchRender() {
	for _, batch := range ren.batches {
		batch.Draw(ren.platform)
		batch.Clear()
	}
	ren.curBatch = 0
}
