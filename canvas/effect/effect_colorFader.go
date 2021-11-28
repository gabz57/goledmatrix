package effect

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"image/color"
)

type ColorFaderEffect struct {
	fade float64
}

func NewColorFaderMask() *ColorFaderEffect {
	return &ColorFaderEffect{
		fade: 0,
	}
}

func (m *ColorFaderEffect) SetFade(fade float64) {
	if fade >= 1 {
		m.fade = 1
	} else if fade <= 0 {
		m.fade = 0
	} else {
		m.fade = fade
	}
}

func (m *ColorFaderEffect) AdaptPixel() func(canvas Canvas, x int, y int, color color.Color) {
	return func(canvas Canvas, x int, y int, ledColor color.Color) {
		r, g, b, a := ledColor.RGBA()
		canvas.Set(x, y, color.RGBA{
			R: uint8((1 - m.fade) * float64(uint8(r))),
			G: uint8((1 - m.fade) * float64(uint8(g))),
			B: uint8((1 - m.fade) * float64(uint8(b))),
			A: uint8(a),
		})
	}
}
