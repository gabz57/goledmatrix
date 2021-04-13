package impl

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"image/color"
	"time"
)

type Heart struct {
	Graphic *Graphic
	pixels  []Pixel
	origin  Point
}

func NewHeart(origin Point) *Heart {
	var heartGraphic = NewGraphic(nil, nil)
	heart := Heart{
		Graphic: &heartGraphic,
		pixels:  heartPixels(),
		origin:  origin,
	}
	return &heart
}

var pink color.Color = color.RGBA{R: 218, G: 50, B: 166, A: 0xff}
var pinkShadow1 color.Color = color.RGBA{R: 203, G: 47, B: 157, A: 0xff}
var pinkShadow2 color.Color = color.RGBA{R: 179, G: 40, B: 137, A: 0xff}
var pinkDark1 color.Color = color.RGBA{R: 130, G: 30, B: 150, A: 0xff}
var pinkDark2 color.Color = color.RGBA{R: 104, G: 25, B: 147, A: 0xff}
var pinkDark3 color.Color = color.RGBA{R: 86, G: 22, B: 145, A: 0xff}

func heartPixels() (pixels []Pixel) {
	pixels = append(pixels,
		Pixel{X: 6, Y: 11, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 5, Y: 10, C: &ColorBlack},
		Pixel{X: 6, Y: 10, C: &pinkShadow1},
		Pixel{X: 7, Y: 10, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 4, Y: 9, C: &ColorBlack},
		Pixel{X: 5, Y: 9, C: &pink},
		Pixel{X: 6, Y: 9, C: &pink},
		Pixel{X: 7, Y: 9, C: &pinkShadow1},
		Pixel{X: 8, Y: 9, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 3, Y: 8, C: &ColorBlack},
		Pixel{X: 4, Y: 8, C: &pink},
		Pixel{X: 5, Y: 8, C: &pink},
		Pixel{X: 6, Y: 8, C: &pink},
		Pixel{X: 7, Y: 8, C: &pink},
		Pixel{X: 8, Y: 8, C: &pinkShadow1},
		Pixel{X: 9, Y: 8, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 2, Y: 7, C: &ColorBlack},
		Pixel{X: 3, Y: 7, C: &pink},
		Pixel{X: 4, Y: 7, C: &pink},
		Pixel{X: 5, Y: 7, C: &pink},
		Pixel{X: 6, Y: 7, C: &pink},
		Pixel{X: 7, Y: 7, C: &pink},
		Pixel{X: 8, Y: 7, C: &pink},
		Pixel{X: 9, Y: 7, C: &pinkShadow1},
		Pixel{X: 10, Y: 7, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 1, Y: 6, C: &ColorBlack},
		Pixel{X: 2, Y: 6, C: &pinkDark1},
		Pixel{X: 3, Y: 6, C: &pink},
		Pixel{X: 4, Y: 6, C: &pink},
		Pixel{X: 5, Y: 6, C: &pink},
		Pixel{X: 6, Y: 6, C: &pink},
		Pixel{X: 7, Y: 6, C: &pink},
		Pixel{X: 8, Y: 6, C: &pink},
		Pixel{X: 9, Y: 6, C: &pinkShadow1},
		Pixel{X: 10, Y: 6, C: &pinkShadow2},
		Pixel{X: 11, Y: 6, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 0, Y: 5, C: &ColorBlack},
		Pixel{X: 1, Y: 5, C: &pinkDark1},
		Pixel{X: 2, Y: 5, C: &pinkDark1},
		Pixel{X: 3, Y: 5, C: &pinkDark1},
		Pixel{X: 4, Y: 5, C: &pink},
		Pixel{X: 5, Y: 5, C: &pink},
		Pixel{X: 6, Y: 5, C: &pink},
		Pixel{X: 7, Y: 5, C: &pink},
		Pixel{X: 8, Y: 5, C: &pink},
		Pixel{X: 9, Y: 5, C: &pink},
		Pixel{X: 10, Y: 5, C: &pinkShadow1},
		Pixel{X: 11, Y: 5, C: &pinkShadow2},
		Pixel{X: 12, Y: 5, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 0, Y: 4, C: &ColorBlack},
		Pixel{X: 1, Y: 4, C: &pinkDark2},
		Pixel{X: 2, Y: 4, C: &ColorWhite},
		Pixel{X: 3, Y: 4, C: &pinkDark1},
		Pixel{X: 4, Y: 4, C: &pink},
		Pixel{X: 5, Y: 4, C: &pink},
		Pixel{X: 6, Y: 4, C: &pink},
		Pixel{X: 7, Y: 4, C: &pink},
		Pixel{X: 8, Y: 4, C: &pink},
		Pixel{X: 9, Y: 4, C: &pink},
		Pixel{X: 10, Y: 4, C: &pinkShadow1},
		Pixel{X: 11, Y: 4, C: &pinkShadow2},
		Pixel{X: 12, Y: 4, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 0, Y: 3, C: &ColorBlack},
		Pixel{X: 1, Y: 3, C: &pinkDark3},
		Pixel{X: 2, Y: 3, C: &ColorWhite},
		Pixel{X: 3, Y: 3, C: &pinkDark2},
		Pixel{X: 4, Y: 3, C: &pinkDark1},
		Pixel{X: 5, Y: 3, C: &pink},
		Pixel{X: 6, Y: 3, C: &pink},
		Pixel{X: 7, Y: 3, C: &pink},
		Pixel{X: 8, Y: 3, C: &pink},
		Pixel{X: 9, Y: 3, C: &pink},
		Pixel{X: 10, Y: 3, C: &pinkShadow1},
		Pixel{X: 11, Y: 3, C: &pinkShadow2},
		Pixel{X: 12, Y: 3, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 0, Y: 2, C: &ColorBlack},
		Pixel{X: 1, Y: 2, C: &pinkDark3},
		Pixel{X: 2, Y: 2, C: &pinkDark3},
		Pixel{X: 3, Y: 2, C: &ColorWhite},
		Pixel{X: 4, Y: 2, C: &pinkDark2},
		Pixel{X: 5, Y: 2, C: &pinkDark2},
		Pixel{X: 6, Y: 2, C: &ColorBlack},
		Pixel{X: 7, Y: 2, C: &pink},
		Pixel{X: 8, Y: 2, C: &pink},
		Pixel{X: 9, Y: 2, C: &pinkShadow1},
		Pixel{X: 10, Y: 2, C: &pinkShadow1},
		Pixel{X: 11, Y: 2, C: &pinkShadow2},
		Pixel{X: 12, Y: 2, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 1, Y: 1, C: &ColorBlack},
		Pixel{X: 2, Y: 1, C: &pinkDark3},
		Pixel{X: 3, Y: 1, C: &pinkDark3},
		Pixel{X: 4, Y: 1, C: &pinkDark2},
		Pixel{X: 5, Y: 1, C: &ColorBlack},
		Pixel{X: 7, Y: 1, C: &ColorBlack},
		Pixel{X: 8, Y: 1, C: &pinkShadow1},
		Pixel{X: 9, Y: 1, C: &pinkShadow1},
		Pixel{X: 10, Y: 1, C: &pinkShadow2},
		Pixel{X: 11, Y: 1, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 2, Y: 0, C: &ColorBlack},
		Pixel{X: 3, Y: 0, C: &ColorBlack},
		Pixel{X: 4, Y: 0, C: &ColorBlack},
		Pixel{X: 8, Y: 0, C: &ColorBlack},
		Pixel{X: 9, Y: 0, C: &ColorBlack},
		Pixel{X: 10, Y: 0, C: &ColorBlack},
	)
	return pixels
}

func (h Heart) Update(elapsedBetweenUpdate time.Duration) {
}

func (h Heart) Draw(canvas Canvas) error {
	for _, pixel := range h.pixels {
		canvas.Set(h.origin.X+pixel.X, h.origin.Y+pixel.Y, *pixel.C)
	}
	return nil
}
