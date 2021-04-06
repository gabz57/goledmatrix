package shapes

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
)

type Dot struct {
	*Graphic
	position Point
}

func NewDot(graphic *Graphic, position Point) *Dot {
	return &Dot{
		Graphic:  graphic,
		position: position,
	}
}

func (d *Dot) Draw(canvas *Canvas) error {
	position := d.position.Add(d.ComputedOffset())
	canvas.Set(position.X, position.Y, d.Layout().Color())

	return nil
}

func (d *Dot) SetDot(position Point) {
	d.position = position
}
