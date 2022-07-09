package shapes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
)

type Rectangle struct {
	*Graphic
	min, max Point
	fill     bool
}

func NewRectangle(graphic *Graphic, position Point, dimensions Point, fill bool) *Rectangle {
	return &Rectangle{
		Graphic: graphic,
		min:     position,
		max:     position.Add(dimensions),
		fill:    fill,
	}
}

func (r *Rectangle) Draw(canvas Canvas) error {
	offset := r.ComputedOffset()
	start := Point{
		X: offset.X + r.min.X,
		Y: offset.Y + r.min.Y,
	}
	end := Point{
		X: offset.X + r.max.X,
		Y: offset.Y + r.max.Y,
	}

	color := r.Layout().Color()
	for y := start.Y; y <= end.Y; y++ {
		canvas.Set(start.X, y, color)
		canvas.Set(end.X, y, color)
	}
	for x := start.X + 1; x < end.X; x++ {
		canvas.Set(x, start.Y, color)
		canvas.Set(x, end.Y, color)
	}

	if r.fill {
		fillColor := r.Layout().BackgroundColor()
		for x := start.X + 1; x < end.X; x++ {
			for y := start.Y + 1; y < end.Y; y++ {
				canvas.Set(x, y, fillColor)
			}
		}
	}

	return nil
}

func (r *Rectangle) SetMin(min Point) {
	r.min = min
}

func (r *Rectangle) SetMax(max Point) {
	r.max = max
}
