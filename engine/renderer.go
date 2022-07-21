package engine

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

//type Batch *pixel.Batch
type Picture pixel.Picture
type Rect *pixel.Rect
type Circle *pixel.Circle
type Line *line
type line struct {
	Start Vector
	End   Vector
}
type Polygon *poly
type poly struct {
	Points []Vector
}

func NewRect(min Vector, max Vector) Rect {
	return &pixel.Rect{
		Min: toPixelVec(min),
		Max: toPixelVec(max),
	}
}

func NewCircle(radius float64, location Vector) Circle {
	return &pixel.Circle{
		Radius: radius,
		Center: toPixelVec(location),
	}
}

func NewLine(start Vector, end Vector) Line {
	return &line{
		Start: start,
		End:   end,
	}
}

func NewPolygon(points ...Vector) Polygon {
	var vectors []Vector
	vectors = append(vectors, points...)
	return &poly{
		Points: vectors,
	}
}

func FromPixelRect(rect pixel.Rect) Rect {
	return Rect(&rect)
}

func LoadPicture(path string) (Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

type RenderSystem interface {
	DrawSprite(do *DrawObject)
	DrawCircle(c Circle)
	DrawQuad(rect Rect)
	DrawLine(line Line)
}

type Renderer struct {
	platform *Platform
}

func NewRenderer(platform *Platform) RenderSystem {
	return &Renderer{
		platform: platform,
	}
}

type DrawObject struct {
	//Batch       Batch
	Spritesheet Picture
	Frame       Rect
	Loc         Vector
	Angle       float64
	Scale       float64
}

func (ren *Renderer) DrawSprite(do *DrawObject) {
	trans := pixel.IM.Scaled(pixel.ZV, do.Scale).Rotated(pixel.ZV, do.Angle)
	sprite := pixel.NewSprite(do.Spritesheet, *do.Frame)
	sprite.Draw(ren.platform, trans.Moved(toPixelVec(do.Loc)))
}

/*
	poly := imdraw.New(nil)
	for _, p := range shape.Points {
		poly.Push(pixelPoint(p.Add(shape.Pos).Sub(shape.Offset)))
	}
	poly.Polygon(2).DrawCircle()
*/

func (ren *Renderer) DrawCircle(c Circle) {
	imd := imdraw.New(nil)
	imd.Color = colornames.White
	imd.Push(c.Center)
	imd.Circle(c.Radius, 2)
	imd.Draw(ren.platform)
}

func (ren *Renderer) DrawQuad(rect Rect) {
	imd := imdraw.New(nil)
	imd.Color = colornames.Blue
	imd.Push(rect.Min)
	imd.Push(pixel.V(rect.Min.X, rect.Max.Y))
	imd.Push(rect.Max)
	imd.Push(pixel.V(rect.Max.X, rect.Min.Y))
	imd.Polygon(0)
	imd.Draw(ren.platform)
}

func (ren *Renderer) DrawLine(line Line) {
	imd := imdraw.New(nil)
	imd.Color = colornames.Green
	imd.Push(toPixelVec(line.Start))
	imd.Push(toPixelVec(line.End))
	imd.Polygon(2)
	imd.Draw(ren.platform)
}

func (ren *Renderer) DrawPoly(poly Polygon) {

}

// Update renders each entity to its batch and draws the batches based on the z value from CLocation
//sudo code of what is happening here
//
//	layers = map layer => []drawObjs
//	batches = set batch
//	for entity e
//		build drawObj(location, rotation, batch, frame)
//		layers[layer].append(drawObj)
//	for layer l in sorted layers
//		for drawObj do in layers[l]
//			do.Render
//		for batch b in batches
//			b.Draw(win)
//			b.Clear()
/*
func (br *SBatchRenderer) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)

	layers := map[int][]*drawObject{}
	var exists = struct{}{}
	batches := map[*pixel.Batch]struct{}{}
	for _, e := range br.controlEntities {
		sp, err := components.GetCProperties(e)
		if err != nil {
			return err
		}
		if !sp.Active {
			continue
		}
		an, err := components.GetCAnimation(e)
		if err != nil {
			return err
		}
		curFrame := an.GetCurrentFrame()
		ba, err := components.GetCBatchAsset(e)
		if err != nil {
			return err
		}
		loc, err := components.GetCLocation(e)
		if err != nil {
			return err
		}
		do := &drawObject{
			Batch:       ba.Batch,
			Spritesheet: &ba.Spritesheet,
			Frame:       &curFrame,
			Loc:         &loc.Loc,
			Angle:       sp.Angle,
			Scale:       sp.Scale,
		}
		if _, OK := layers[loc.Z]; !OK {
			layers[loc.Z] = []*drawObject{}
		}
		layers[loc.Z] = append(layers[loc.Z], do)
		if _, c := batches[ba.Batch]; !c {
			batches[ba.Batch] = exists
		}
	}
	keys := make([]int, 0, len(layers))
	for k := range layers {
		keys = append(keys, k)
	}
	sKeys := sort.IntSlice(keys)
	sort.Sort(sKeys)
	for _, k := range sKeys {
		layer := layers[k]
		for _, do := range layer {
			do.render()
		}
		for b := range batches {
			b.Draw(win)
			b.Clear()
		}
	}
	return nil
}
*/
