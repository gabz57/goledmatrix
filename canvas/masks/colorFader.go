package masks

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"image/color"
)

type ColorFaderCanvasMask struct {
	fade float64
}

func NewColorFaderMask() *ColorFaderCanvasMask {
	return &ColorFaderCanvasMask{
		fade: 0,
	}
}

func (c *ColorFaderCanvasMask) SetFade(fade float64) {
	if fade >= 1 {
		c.fade = 1
	} else if fade <= 0 {
		c.fade = 0
	} else {
		c.fade = fade
	}
}

func (c *ColorFaderCanvasMask) AdaptPixel() func(canvas *Canvas, x int, y int, color *color.Color) {
	return func(canvas *Canvas, x int, y int, ledColor *color.Color) {
		r, g, b, a := (*ledColor).RGBA()
		(*canvas).Set(x, y, color.RGBA{
			R: uint8((1 - c.fade) * float64(uint8(r))),
			G: uint8((1 - c.fade) * float64(uint8(g))),
			B: uint8((1 - c.fade) * float64(uint8(b))),
			A: uint8(a),
		})
	}
}
