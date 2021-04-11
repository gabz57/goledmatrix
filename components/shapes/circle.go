package shapes

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"image/color"
)

type Pixel struct {
	x, y int
	c    *color.Color
}

type Circle struct {
	*Graphic
	center Point
	radius int
	fill   bool
	pixels []Pixel
}

func NewCircle(graphic *Graphic, center Point, radius int, fill bool) *Circle {
	c := Circle{
		Graphic: graphic,
		center:  center,
		radius:  radius,
		fill:    fill,
	}
	c.pixels = c.buildPixels(center, radius, fill)
	return &c
}

func (c *Circle) buildPixels(centerPoint Point, radius int, fill bool) []Pixel {
	var pixels []Pixel
	offset := c.ComputedOffset()
	center := offset.Add(centerPoint)
	if fill {
		c.fillCircle(&pixels, radius, center)
	}
	c.contourCircle(&pixels, radius, center)
	return pixels
}

func (c *Circle) fillCircle(pixels *[]Pixel, radius int, center Point) {
	bgColorColor := c.Layout().BackgroundColor()
	radiusSqr := radius * radius
	for x := 0; x <= radius; x++ {
		for y := 0; y <= radius; y++ {
			if x*x+y*y <= radiusSqr {
				*pixels = append(*pixels,
					Pixel{
						x: center.X + x,
						y: center.Y + y,
						c: bgColorColor,
					},
					Pixel{
						x: center.X + x,
						y: center.Y - y,
						c: bgColorColor,
					},
					Pixel{
						x: center.X - x,
						y: center.Y + y,
						c: bgColorColor,
					},
					Pixel{
						x: center.X - x,
						y: center.Y - y,
						c: bgColorColor,
					},
				)
			}
		}
	}
}

func (c *Circle) contourCircle(pixels *[]Pixel, radius int, center Point) {
	fgColor := c.Layout().Color()
	var x = radius
	var y = 0
	var radiusError = 1 - x
	for y <= x {
		*pixels = append(*pixels,
			Pixel{
				x: center.X + x,
				y: center.Y + y,
				c: fgColor,
			},
			Pixel{
				x: center.X + x,
				y: center.Y - y,
				c: fgColor,
			},
			Pixel{
				x: center.X - x,
				y: center.Y + y,
				c: fgColor,
			},
			Pixel{
				x: center.X - x,
				y: center.Y - y,
				c: fgColor,
			},
			Pixel{
				x: center.X + y,
				y: center.Y + x,
				c: fgColor,
			},
			Pixel{
				x: center.X + y,
				y: center.Y - x,
				c: fgColor,
			},
			Pixel{
				x: center.X - y,
				y: center.Y + x,
				c: fgColor,
			},
			Pixel{
				x: center.X - y,
				y: center.Y - x,
				c: fgColor,
			})
		y++
		if radiusError < 0 {
			radiusError += 2*y + 1
		} else {
			x--
			radiusError += 2 * (y - x + 1)
		}
	}
}

func (c *Circle) Draw(canvas *Canvas) error {
	for _, pixel := range c.pixels {
		canvas.Set(pixel.x, pixel.y, *pixel.c)
	}
	return nil
}
