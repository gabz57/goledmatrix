package masks

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"image"
	"image/color"
)

type VisibleMask struct {
	visibleArea image.Rectangle
	offset      image.Point
}

func NewVisibleMask(visibleArea image.Rectangle) *VisibleMask {
	return &VisibleMask{
		visibleArea: visibleArea,
	}
}

func (c *VisibleMask) GetOffset() image.Point {
	return c.offset
}

func (c *VisibleMask) SetOffset(offset image.Point) {
	c.offset = offset
}

func (c *VisibleMask) AdaptPixel() func(canvas Canvas, x int, y int, color color.Color) {
	return func(canvas Canvas, x int, y int, color color.Color) {
		targetXY := image.Point{X: x, Y: y}.Add(c.offset)
		if targetXY.In(c.visibleArea) {
			canvas.Set(targetXY.X, targetXY.Y, color)
		}
	}
}
