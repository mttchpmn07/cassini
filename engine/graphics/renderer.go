package graphics

import (
	"image"
	"os"

	"image/color"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Rasterable interface {
	Raster() *imdraw.IMDraw
	Color(c color.Color)
	Thickness(t float64)
	C() color.Color
	T() float64
}

type RasterProperties struct {
	color     color.Color
	thickness float64
}

func NewRasterable() *RasterProperties {
	return &RasterProperties{
		color:     colornames.White,
		thickness: 2,
	}
}

func (rp *RasterProperties) Raster() *imdraw.IMDraw {
	return imdraw.New(nil)
}

func (rp *RasterProperties) Color(c color.Color) {
	rp.color = c
}

func (rp *RasterProperties) Thickness(t float64) {
	rp.thickness = t
}

func (rp *RasterProperties) C() color.Color {
	return rp.color
}

func (rp *RasterProperties) T() float64 {
	return rp.thickness
}

type Picture pixel.Picture
type DrawObject struct {
	Spritesheet Picture
	Frame       pixel.Rect
	Loc         pixel.Vec
	Angle       float64
	Scale       float64
}

func NewDrawObject(spritePath string, startLoc pixel.Vec, angle, scale float64) (*DrawObject, error) {
	pic, err := loadSpriteSheet(spritePath)
	if err != nil {
		return nil, err
	}
	return &DrawObject{
		Spritesheet: pic,
		Frame:       pic.Bounds(),
		Loc:         startLoc,
		Angle:       angle,
		Scale:       scale,
	}, nil
}

func (do DrawObject) Moved(x, y float64) *DrawObject {
	do.Loc = pixel.V(x, y)
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
	DrawShape(shape Rasterable)
	DrawShapes(shapes []Rasterable)
	Draw(pen Rasterable)
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
	sprite := pixel.NewSprite(do.Spritesheet, do.Frame)
	sprite.Draw(ren.batches[ren.curBatch], trans.Moved(do.Loc))
}

func (ren *Renderer) Draw(pen Rasterable) {
	imd := pen.Raster()
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) DrawShape(s Rasterable) {
	imd := s.Raster()
	imd.Draw(ren.batches[ren.curBatch])
}

func (ren *Renderer) DrawShapes(ss []Rasterable) {
	for _, s := range ss {
		ren.DrawShape(s)
	}
}
