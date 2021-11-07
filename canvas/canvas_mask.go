package canvas

import (
	"image/color"
)

type (
	Mask interface {
		AdaptPixel() func(canvas Canvas, x int, y int, color color.Color)
	}
	// Canvas wrapper to replace pixels color
	MaskAdapter struct {
		Canvas
		mask Mask
	}
)

func (a MaskAdapter) Set(x, y int, ledColor color.Color) {
	a.mask.AdaptPixel()(a.Canvas, x, y, ledColor)
}

func NewMaskAdapter(canvas Canvas, mask Mask) *MaskAdapter {
	return &MaskAdapter{
		Canvas: canvas,
		mask:   mask,
	}
}
