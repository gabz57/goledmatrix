package shapes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"math"
)

type Line struct {
	*Graphic
	start, end Point
}

func NewLine(graphic *Graphic, start, end Point) *Line {
	return &Line{
		Graphic: graphic,
		start:   start,
		end:     end,
	}
}

func (l *Line) Draw(canvas Canvas) error {
	offset := l.ComputedOffset()

	start := offset.Add(l.start)
	end := offset.Add(l.end)

	dx := end.X - start.X
	dy := end.Y - start.Y

	var gradient int
	var x int
	var y int
	const shift = 0x10

	var posStart = start
	var posEnd = end
	if math.Abs(float64(dx)) > math.Abs(float64(dy)) {
		// x variation is bigger than y variation
		if posEnd.X < posStart.X {
			var tmp = posStart
			posStart = posEnd
			posEnd = tmp
		}
		gradient = (dy << shift) / dx
		y = 0x8000 + (posStart.Y << shift)
		for x = posStart.X; x <= posEnd.X; x++ {
			canvas.Set(x, y>>shift, l.Layout().Color())
			y += gradient
		}
	} else if dy != 0 {
		// y variation is bigger than x variation
		if posEnd.Y < posStart.Y {
			var tmp = posStart
			posStart = posEnd
			posEnd = tmp
		}
		gradient = (dx << shift) / dy
		x = 0x8000 + (posStart.X << shift)
		for y = posStart.Y; y <= posEnd.Y; y++ {
			canvas.Set(x>>shift, y, l.Layout().Color())
			x += gradient
		}
	} else {
		canvas.Set(posStart.X, posStart.Y, l.Layout().Color())
	}
	return nil
}

func (l *Line) SetLine(start Point, end Point) {
	l.start = start
	l.end = end
}
