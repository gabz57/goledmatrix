package shapes

import (
	"fmt"
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"image/color"
	"strconv"
	"time"
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
	circle := Circle{
		Graphic: graphic,
		center:  center,
		radius:  radius,
		fill:    fill,
	}
	circle.pixels = buildPixels(&circle, center, radius, fill)
	return &circle
}

func buildPixels(circle *Circle, centerPoint Point, radius int, fill bool) []Pixel {
	var pixels []Pixel
	offset := circle.ComputedOffset()
	center := offset.Add(centerPoint)
	if fill {
		fillColor := circle.Layout().BackgroundColor()
		radiusSqr := radius * radius
		for x := 0; x <= radius; x++ {
			for y := 0; y <= radius; y++ {
				if x*x+y*y <= radiusSqr {
					pixels = append(pixels,
						Pixel{
							x: center.X + x,
							y: center.Y + y,
							c: fillColor,
						},
						Pixel{
							x: center.X + x,
							y: center.Y - y,
							c: fillColor,
						},
						Pixel{
							x: center.X - x,
							y: center.Y + y,
							c: fillColor,
						},
						Pixel{
							x: center.X - x,
							y: center.Y - y,
							c: fillColor,
						},
					)
				}
			}
		}
	}
	c := circle.Layout().Color()
	var x = radius
	var y = 0
	var radiusError = 1 - x
	for y <= x {
		pixels = append(pixels, Pixel{
			x: center.X + x,
			y: center.Y + y,
			c: c,
		},
			Pixel{
				x: center.X + x,
				y: center.Y - y,
				c: c,
			},
			Pixel{
				x: center.X - x,
				y: center.Y + y,
				c: c,
			},
			Pixel{
				x: center.X - x,
				y: center.Y - y,
				c: c,
			},
			Pixel{
				x: center.X + y,
				y: center.Y + x,
				c: c,
			},
			Pixel{
				x: center.X + y,
				y: center.Y - x,
				c: c,
			},
			Pixel{
				x: center.X - y,
				y: center.Y + x,
				c: c,
			},
			Pixel{
				x: center.X - y,
				y: center.Y - x,
				c: c,
			})
		y++
		if radiusError < 0 {
			radiusError += 2*y + 1
		} else {
			x--
			radiusError += 2 * (y - x + 1)
		}
	}
	return pixels
}

func (c *Circle) Draw(canvas *Canvas) error {
	start := time.Now()
	//offset := c.ComputedOffset()
	//center := offset.Add(c.center)
	//if c.fill {
	//	c.fillCircle(canvas, center)
	//}
	//c.circle(canvas, center)
	for _, pixel := range c.pixels {
		canvas.Set(pixel.x, pixel.y, *pixel.c)
	}
	fmt.Println("Circle: " + strconv.FormatInt(time.Now().Sub(start).Milliseconds(), 10) + " ms")
	return nil
}

func (c *Circle) fillCircle(canvas *Canvas, center Point) {
	fillColor := *c.Layout().BackgroundColor()
	radiusSqr := c.radius * c.radius
	for x := 0; x <= c.radius; x++ {
		for y := 0; y <= c.radius; y++ {
			if x*x+y*y <= radiusSqr {
				canvas.SetPoint(center.AddXY(x, y), fillColor)
				canvas.SetPoint(center.AddXY(x, -y), fillColor)
				canvas.SetPoint(center.AddXY(-x, y), fillColor)
				canvas.SetPoint(center.AddXY(-x, -y), fillColor)
			}
		}
	}
}

func (c *Circle) circle(canvas *Canvas, center Point) {
	fgColor := *c.Layout().Color()
	var x = c.radius
	var y = 0
	var radiusError = 1 - x
	for y <= x {
		canvas.SetPoint(center.AddXY(x, y), fgColor)
		canvas.SetPoint(center.AddXY(x, -y), fgColor)
		canvas.SetPoint(center.AddXY(-x, y), fgColor)
		canvas.SetPoint(center.AddXY(-x, -y), fgColor)
		canvas.SetPoint(center.AddXY(y, x), fgColor)
		canvas.SetPoint(center.AddXY(y, -x), fgColor)
		canvas.SetPoint(center.AddXY(-y, x), fgColor)
		canvas.SetPoint(center.AddXY(-y, -x), fgColor)
		y++
		if radiusError < 0 {
			radiusError += 2*y + 1
		} else {
			x--
			radiusError += 2 * (y - x + 1)
		}
	}
}
