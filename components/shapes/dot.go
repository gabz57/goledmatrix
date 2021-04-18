package shapes

import (
	. "github.com/gabz57/goledmatrix/components"
	. "github.com/gabz57/goledmatrix/matrix"
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

func (d *Dot) Draw(canvas Canvas) error {
	p := d.position.Add(d.ComputedOffset())
	canvas.Set(p.X, p.Y, *d.Layout().Color())

	return nil
}

func (d *Dot) SetPosition(position Point) {
	d.position = position
}

func (d *Dot) GetPosition() *Point {
	return &d.position
}
