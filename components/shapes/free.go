package shapes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
)

type Free struct {
	*Graphic
	pixels []Pixel
}

func NewFree(graphic *Graphic, pixels []Pixel) *Free {
	return &Free{
		Graphic: graphic,
		pixels:  pixels,
	}
}

func (f *Free) Draw(canvas Canvas) error {
	offset := f.ComputedOffset()
	for _, pixel := range f.pixels {
		canvas.Set(offset.X+pixel.X, offset.Y+pixel.Y, pixel.C)
	}
	return nil
}
