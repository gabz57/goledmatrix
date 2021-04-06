package shapes

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
)

type Circle struct {
	*Graphic
	center Point
	radius int
	fill   bool
}

func NewCircle(graphic *Graphic, center Point, radius int, fill bool) *Circle {
	return &Circle{
		Graphic: graphic,
		center:  center,
		radius:  radius,
		fill:    fill,
	}
}

func (r *Circle) Draw(canvas *Canvas) error {
	offset := r.ComputedOffset()
	center := offset.Add(r.center)
	if r.fill {
		r.fillCircle(canvas, center)
	}
	r.circle(canvas, center)

	return nil
}

func (r *Circle) fillCircle(canvas *Canvas, center Point) {
	c := r.Layout().BackgroundColor()
	radiusSqr := r.radius * r.radius
	for x := 0; x <= r.radius; x++ {
		for y := 0; y <= r.radius; y++ {
			if x*x+y*y <= radiusSqr {
				canvas.SetPoint(center.AddXY(x, y), c)
				canvas.SetPoint(center.AddXY(x, -y), c)
				canvas.SetPoint(center.AddXY(-x, y), c)
				canvas.SetPoint(center.AddXY(-x, -y), c)
			}
		}
	}
}

func (r *Circle) circle(canvas *Canvas, center Point) {
	c := r.Layout().Color()
	var x = r.radius
	var y = 0
	var radiusError = 1 - x
	for y <= x {
		canvas.SetPoint(center.AddXY(x, y), c)
		canvas.SetPoint(center.AddXY(x, -y), c)
		canvas.SetPoint(center.AddXY(-x, y), c)
		canvas.SetPoint(center.AddXY(-x, -y), c)
		canvas.SetPoint(center.AddXY(y, x), c)
		canvas.SetPoint(center.AddXY(y, -x), c)
		canvas.SetPoint(center.AddXY(-y, x), c)
		canvas.SetPoint(center.AddXY(-y, -x), c)
		y++
		if radiusError < 0 {
			radiusError += 2*y + 1
		} else {
			x--
			radiusError += 2 * (y - x + 1)
		}
	}
}
