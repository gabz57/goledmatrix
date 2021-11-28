package effect

import (
	"github.com/gabz57/goledmatrix/canvas"
	"image/color"
	"time"
)

type (
	Effect interface {
		AdaptPixel() func(canvas canvas.Canvas, x int, y int, color color.Color)
	}
	DynamicEffect interface {
		Effect
		Update(elapsedBetweenUpdate time.Duration) bool
	}
	// Adapter is a Canvas wrapper to replace pixels with Effect implementation
	Adapter struct {
		canvas.Canvas
		effect Effect
	}
)

func (a Adapter) Set(x, y int, ledColor color.Color) {
	a.effect.AdaptPixel()(a.Canvas, x, y, ledColor)
}

func NewAdapter(canvas canvas.Canvas, effect Effect) *Adapter {
	return &Adapter{
		Canvas: canvas,
		effect: effect,
	}
}
