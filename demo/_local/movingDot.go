package main

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"image"
	"time"
)

type MovingDot struct {
	pixelsPerSecond float64
	dot             shapes.Dot
	origin          Point
	dXY             FloatingPoint
	bounds          image.Rectangle
	mask            Canvas
}

func NewMovingDot(c Canvas, origin Point, pixelsPerSecond float64, bounds image.Rectangle) Component {
	var movingDotGraphic = NewGraphic(nil, nil)
	var mask Canvas
	mask = NewSingleColorMask(c, ColorRed)

	return &MovingDot{
		pixelsPerSecond: pixelsPerSecond,
		dot:             *shapes.NewDot(&movingDotGraphic, origin),
		origin:          origin,
		dXY:             FloatingPoint{},
		bounds:          bounds,
		mask:            mask,
	}
}

func (m *MovingDot) Update(elapsedBetweenUpdate time.Duration) {
	m.dXY = FloatingPoint{
		X: m.dXY.X + m.pixelsPerSecond*float64(elapsedBetweenUpdate.Nanoseconds())/1000000000,
		Y: m.dXY.Y + m.pixelsPerSecond*float64(elapsedBetweenUpdate.Nanoseconds())/1000000000,
	}
	m.dot.SetPosition(m.origin.AddXY(int(m.dXY.X)%m.bounds.Max.X, int(m.dXY.Y)%m.bounds.Max.Y))
}

func (m *MovingDot) Draw(c Canvas) error {
	return (*m).dot.Draw(m.mask)
}
