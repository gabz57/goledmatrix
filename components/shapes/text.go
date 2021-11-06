package shapes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/fonts"
	"github.com/gabz57/goledmatrix/matrix"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
)

type Text struct {
	*Graphic
	font     font.Face
	position Point
	txt      string
	pixels   []Pixel
	bounds   image.Rectangle
}

func NewText(graphic *Graphic, position Point, txt string, f fonts.MatrixFont) *Text {
	text := Text{
		Graphic:  graphic,
		position: position,
		txt:      txt,
		font:     fonts.GetFont(f),
	}
	text.computeBounds()
	text.computePixels()
	return &text
}

func (t *Text) Draw(canvas Canvas) error {
	//max := canvas.Bounds().Max
	for _, pixel := range t.pixels {
		//if pixel.X < max.X && pixel.Y < max.Y {
		canvas.Set(pixel.X, pixel.Y, pixel.C)
		//}
	}
	return nil
}

func (t *Text) SetText(txt string) {
	t.txt = txt
	t.computeBounds()
	t.computePixels()
}

func (t *Text) Bounds() image.Rectangle {
	return t.bounds
}

var colorPixel = color.White
var emptyPixel = color.Black

func (t *Text) computeBounds() {
	t.bounds = image.Rect(
		0,
		0,
		font.MeasureString(t.font, t.txt).Ceil(),
		t.font.Metrics().Ascent.Ceil()+t.font.Metrics().Descent.Ceil())
}

func (t *Text) computePixels() {
	dx := t.bounds.Dx()
	dy := t.bounds.Dy()

	var leds = make([]color.Color, dx*dy)
	var imgDst = matrix.NewSimpleCanvas(dx, dy, &leds)
	for rX := t.bounds.Min.X; rX < t.bounds.Max.X; rX++ {
		for rY := t.bounds.Min.Y; rY < t.bounds.Max.Y; rY++ {
			imgDst.Set(rX, rY, emptyPixel)
		}
	}
	d := font.Drawer{
		Dst:  imgDst,
		Src:  &image.Uniform{C: colorPixel},
		Face: t.font,
		Dot:  fixed.Point26_6{X: 0, Y: t.font.Metrics().Height - t.font.Metrics().Descent},
	}
	d.DrawString(t.txt)
	var pixels []Pixel = nil
	position := t.position.Add(t.ComputedOffset())
	for rX := 0; rX < dx; rX++ {
		for rY := 0; rY < dy; rY++ {
			at := imgDst.At(rX, rY)
			if at != emptyPixel {
				r, g, b, _ := at.RGBA()
				if r != 0 && g != 0 && b != 0 {
					pixels = append(pixels, Pixel{
						X: position.X + rX,
						Y: position.Y + rY,
						C: t.Layout().Color(),
					})
				} else {
					pixels = append(pixels, Pixel{
						X: position.X + rX,
						Y: position.Y + rY,
						C: t.Layout().BackgroundColor(),
					})
				}
			}
		}
	}
	t.pixels = pixels
}
