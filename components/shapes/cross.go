package shapes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
)

type Cross struct {
	*Graphic
	position Point
	length   int
}

func NewCross(graphic *Graphic, position Point, length int) *Cross {
	return &Cross{
		Graphic:  graphic,
		position: position,
		length:   length,
	}
}

func (c *Cross) Draw(canvas Canvas) error {
	p := c.position.Add(c.ComputedOffset())
	for x := p.X - c.length; x <= p.X+c.length; x++ {
		canvas.Set(x, p.Y, c.Layout().BackgroundColor())
	}
	for y := p.Y - c.length; y <= p.Y+c.length; y++ {
		canvas.Set(p.X, y, c.Layout().BackgroundColor())
	}
	canvas.Set(p.X, p.Y, c.Layout().Color())
	return nil
}

func (c *Cross) SetPosition(position Point) {
	c.position = position
}

func (c *Cross) GetPosition() *Point {
	return &c.position
}
