package shapes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
)

type Ring struct {
	*Graphic
	center               Point
	radiusExt, radiusInt int
	fill                 bool
	pixels               []Pixel
}

func NewRing(graphic *Graphic, center Point, radiusExt, radiusInt int, fill bool) *Ring {
	c := Ring{
		Graphic:   graphic,
		center:    center,
		radiusExt: radiusExt,
		radiusInt: radiusInt,
		fill:      fill,
	}
	c.pixels = c.buildPixels(center, radiusExt, radiusInt, fill)
	return &c
}

func (r *Ring) buildPixels(centerPoint Point, radiusExt, radiusInt int, fill bool) []Pixel {
	var pixels []Pixel
	offset := r.ComputedOffset()
	center := offset.Add(centerPoint)
	if fill {
		r.fillRing(&pixels, radiusExt, radiusInt, center)
	}
	r.contourRing(&pixels, radiusExt, radiusInt, center)
	return pixels
}

func (r *Ring) fillRing(pixels *[]Pixel, radiusExt int, radiusInt int, center Point) {
	bgColorColor := r.Layout().BackgroundColor()
	radiusExtSqr := radiusExt * radiusExt
	radiusIntSqr := radiusInt * radiusInt
	for x := 0; x <= radiusExt; x++ {
		for y := 0; y <= radiusExt; y++ {
			d := x*x + y*y
			if d >= radiusIntSqr && d <= radiusExtSqr {
				*pixels = append(*pixels,
					Pixel{
						X: center.X + x,
						Y: center.Y + y,
						C: bgColorColor,
					},
					Pixel{
						X: center.X + x,
						Y: center.Y - y,
						C: bgColorColor,
					},
					Pixel{
						X: center.X - x,
						Y: center.Y + y,
						C: bgColorColor,
					},
					Pixel{
						X: center.X - x,
						Y: center.Y - y,
						C: bgColorColor,
					},
				)
			}
		}
	}
}

func (r *Ring) contourRing(pixels *[]Pixel, radiusExt int, radiusInt int, center Point) {
	r.ring(pixels, radiusExt, center)
	r.ring(pixels, radiusInt, center)
}

func (r *Ring) ring(pixels *[]Pixel, radius int, center Point) {
	fgColor := r.Layout().Color()
	var x = radius
	var y = 0
	var radiusError = 1 - x
	for y <= x {
		*pixels = append(*pixels,
			Pixel{
				X: center.X + x,
				Y: center.Y + y,
				C: fgColor,
			},
			Pixel{
				X: center.X + x,
				Y: center.Y - y,
				C: fgColor,
			},
			Pixel{
				X: center.X - x,
				Y: center.Y + y,
				C: fgColor,
			},
			Pixel{
				X: center.X - x,
				Y: center.Y - y,
				C: fgColor,
			},
			Pixel{
				X: center.X + y,
				Y: center.Y + x,
				C: fgColor,
			},
			Pixel{
				X: center.X + y,
				Y: center.Y - x,
				C: fgColor,
			},
			Pixel{
				X: center.X - y,
				Y: center.Y + x,
				C: fgColor,
			},
			Pixel{
				X: center.X - y,
				Y: center.Y - x,
				C: fgColor,
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

func (r *Ring) Draw(canvas Canvas) error {
	for _, pixel := range r.pixels {
		canvas.Set(pixel.X, pixel.Y, *pixel.C)
	}
	return nil
}
