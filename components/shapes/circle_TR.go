package shapes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
)

type CircleTR struct {
	*Graphic
	center Point
	radius int
	fill   bool
	pixels []Pixel
}

func NewCircleTR(graphic *Graphic, center Point, radius int, fill bool) *CircleTR {
	c := CircleTR{
		Graphic: graphic,
		center:  center,
		radius:  radius,
		fill:    fill,
	}
	c.pixels = c.buildPixels(center, radius, fill)
	return &c
}

func (c *CircleTR) buildPixels(centerPoint Point, radius int, fill bool) []Pixel {
	var pixels []Pixel
	offset := c.ComputedOffset()
	center := offset.Add(centerPoint)
	if fill {
		c.fillCircleTR(&pixels, radius, center)
	}
	c.contourCircleTR(&pixels, radius, center)
	return pixels
}

func (c *CircleTR) fillCircleTR(pixels *[]Pixel, radius int, center Point) {
	bgColorColor := c.Layout().BackgroundColor()
	radiusSqr := radius * radius
	for x := 0; x <= radius; x++ {
		for y := 0; y <= radius; y++ {
			if x*x+y*y <= radiusSqr {
				*pixels = append(*pixels,
					Pixel{
						X: center.X + x,
						Y: center.Y - y,
						C: bgColorColor,
					},
				)
			}
		}
	}
}

func (c *CircleTR) contourCircleTR(pixels *[]Pixel, radius int, center Point) {
	fgColor := c.Layout().Color()
	var x = radius
	var y = 0
	var radiusError = 1 - x
	for y <= x {
		*pixels = append(*pixels,
			Pixel{
				X: center.X + x,
				Y: center.Y - y,
				C: fgColor,
			},
			Pixel{
				X: center.X + y,
				Y: center.Y - x,
				C: fgColor,
			},
		)
		y++
		if radiusError < 0 {
			radiusError += 2*y + 1
		} else {
			x--
			radiusError += 2 * (y - x + 1)
		}
	}
}

func (c *CircleTR) Draw(canvas Canvas) error {
	for _, pixel := range c.pixels {
		canvas.Set(pixel.X, pixel.Y, pixel.C)
	}
	return nil
}
