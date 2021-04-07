package main

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/fonts"
	"time"
)

type Clock struct {
	shape          CompositeDrawable
	now            time.Time
	center         Point
	radius         int
	text           *shapes.Text
	rotatingHour   *shapes.Line
	rotatingMinute *shapes.Line
	rotatingSecond *shapes.Dot
}

var clockGraphic = NewGraphic(nil, nil)

func NewClock(center Point, radius int) Clock {
	clock := Clock{
		now:    time.Now(),
		center: center,
		radius: radius,
		shape: CompositeDrawable{
			Graphic:   &clockGraphic,
			Drawables: []*Drawable{},
		},
	}
	clock.shape.AddDrawable(clock.buildBlabla(10, 10))
	clock.shape.AddDrawable(clock.buildBlabla(10, 20))
	clock.shape.AddDrawable(clock.buildBlabla(10, 30))
	clock.shape.AddDrawable(clock.buildBlabla(10, 40))
	clock.shape.AddDrawable(clock.buildBlabla(10, 50))
	clock.shape.AddDrawable(clock.buildBlabla(10, 60))
	clock.shape.AddDrawable(clock.buildBlabla(10, 70))
	clock.shape.AddDrawable(clock.buildBlabla(10, 80))
	//
	//clock.shape.AddDrawable(clock.buildBlabla(50, 10))
	//clock.shape.AddDrawable(clock.buildBlabla(50, 20))
	//clock.shape.AddDrawable(clock.buildBlabla(50, 30))
	//clock.shape.AddDrawable(clock.buildBlabla(50, 40))
	//clock.shape.AddDrawable(clock.buildBlabla(50, 50))
	//clock.shape.AddDrawable(clock.buildBlabla(50, 60))
	//clock.shape.AddDrawable(clock.buildBlabla(50, 70))
	//clock.shape.AddDrawable(clock.buildBlabla(50, 80))
	//
	//clock.shape.AddDrawable(clock.buildBlabla(90, 10))
	//clock.shape.AddDrawable(clock.buildBlabla(90, 20))
	//clock.shape.AddDrawable(clock.buildBlabla(90, 30))
	//clock.shape.AddDrawable(clock.buildBlabla(90, 40))
	//clock.shape.AddDrawable(clock.buildBlabla(90, 50))
	//clock.shape.AddDrawable(clock.buildBlabla(90, 60))
	//clock.shape.AddDrawable(clock.buildBlabla(90, 70))
	//clock.shape.AddDrawable(clock.buildBlabla(90, 80))
	//
	clock.shape.AddDrawable(clock.buildDrawableContourCircle())
	clock.shape.AddDrawable(clock.buildDrawableCenter())
	clock.shape.AddDrawables(clock.buildDrawableHours())
	clock.shape.AddDrawables(clock.buildDrawableMinutes())

	now := time.Now()
	hour, min, sec := now.Clock()
	clock.text = clock.buildText(now)
	var drawableText Drawable
	drawableText = clock.text
	clock.shape.AddDrawable(&drawableText)

	clock.rotatingHour = clock.buildRotatingHour(hour, min)
	var drawableHour Drawable
	drawableHour = clock.rotatingHour
	clock.shape.AddDrawable(&drawableHour)

	clock.rotatingMinute = clock.buildRotatingMinute(min, sec)
	var drawableMinute Drawable
	drawableMinute = clock.rotatingMinute
	clock.shape.AddDrawable(&drawableMinute)

	clock.rotatingSecond = clock.buildRotatingSecond(sec)
	var drawableSecond Drawable
	drawableSecond = clock.rotatingSecond
	clock.shape.AddDrawable(&drawableSecond)

	return clock
}

func (c Clock) Update() {
	now := time.Now()
	if now.Sub(c.now) > 10*time.Millisecond {
		c.now = now
		hour, min, sec := now.Clock()

		aDHour := angleDegreesHour(hour, min)
		c.rotatingHour.SetLine(
			c.center,
			c.hourLineEnd(aDHour),
		)

		c.rotatingMinute.SetLine(
			c.center,
			c.minuteLineEnd(angleDegreesMinute(min, sec)),
		)

		c.rotatingSecond.SetDot(
			c.secondDotPosition(angleDegreesSecond(sec, c.now)),
		)

		c.text.SetText(TimeToText(c.now))
	}
}

func angleDegreesHour(hour int, min int) float64 {
	return (float64(hour) + float64(min)/60) * 30
}

func angleDegreesMinute(min int, sec int) float64 {
	return (float64(min) + float64(sec)/60) * 6
}

func angleDegreesSecond(sec int, now time.Time) float64 {
	return (float64(sec) + (float64(now.Nanosecond()) / 1000000000)) * 6
}

func (c Clock) buildBlabla(x, y int) *Drawable {
	graphic := NewGraphic(c.shape.Graphic, nil)
	var circle Drawable
	circle = shapes.NewText(&graphic, Point{X: x, Y: y}, "bla bla", fonts.Bdf4x6)
	return &circle
}
func (c Clock) buildDrawableContourCircle() *Drawable {
	graphic := NewGraphic(c.shape.Graphic, nil)
	var circle Drawable
	circle = shapes.NewCircle(&graphic, c.center, c.radius, false)
	return &circle
}

func (c Clock) buildDrawableCenter() *Drawable {
	graphic := NewGraphic(c.shape.Graphic, nil)
	var dot Drawable
	dot = shapes.NewDot(&graphic, c.center)
	return &dot
}

func (c Clock) buildDrawableHours() []*Drawable {
	drawables := make([]*Drawable, 12)
	for hour := 0; hour < 12; hour++ {
		drawables[hour] = c.buildDrawableHour(hour)
	}
	return drawables
}

func (c Clock) buildDrawableHour(hour int) *Drawable {
	angleDegrees := angleDegreesHour(hour, 0)
	graphic := NewGraphic(c.shape.Graphic, nil)
	var line Drawable
	line = shapes.NewLine(&graphic,
		Rotate(Point{
			X: c.center.X,
			Y: c.center.Y - c.radius + 5,
		}, c.center, angleDegrees),
		Rotate(Point{
			X: c.center.X,
			Y: c.center.Y - c.radius + 2,
		}, c.center, angleDegrees),
	)
	return &line
}

func (c Clock) buildDrawableMinutes() []*Drawable {
	drawables := make([]*Drawable, 60)
	for minutes := 0; minutes < 60; minutes++ {
		drawables[minutes] = c.buildDrawableMinute(c.center, c.radius, minutes)
	}
	return drawables
}

func (c Clock) buildDrawableMinute(center Point, radius int, minute int) *Drawable {
	angleDegrees := angleDegreesMinute(minute, 0)
	graphic := NewGraphic(c.shape.Graphic, nil)
	var dot Drawable
	dot = shapes.NewDot(&graphic,
		Rotate(Point{
			X: center.X,
			Y: center.Y - radius + 2,
		}, center, angleDegrees),
	)
	return &dot
}

func (c Clock) buildText(now time.Time) *shapes.Text {
	graphic := NewGraphic(c.shape.Graphic, NewLayout(&ColorGreen, nil))
	return shapes.NewText(&graphic,
		Point{
			X: 0,
			Y: 0,
		},
		TimeToText(now),
		fonts.Bdf4x6,
	)
}

func (c Clock) buildRotatingHour(hour, min int) *shapes.Line {
	graphic := NewGraphic(c.shape.Graphic, NewLayout(&ColorBlue, nil))
	return shapes.NewLine(&graphic,
		c.center,
		c.hourLineEnd(angleDegreesHour(hour, min)),
	)
}

func (c Clock) buildRotatingMinute(min, sec int) *shapes.Line {
	graphic := NewGraphic(c.shape.Graphic, NewLayout(ColorViolet, nil))
	return shapes.NewLine(&graphic,
		c.center,
		c.minuteLineEnd(angleDegreesMinute(min, sec)),
	)
}

func (c Clock) buildRotatingSecond(sec int) *shapes.Dot {
	graphic := NewGraphic(c.shape.Graphic, NewLayout(ColorRed, nil))
	return shapes.NewDot(&graphic,
		c.secondDotPosition(angleDegreesSecond(sec, c.now)),
	)
}

func (c Clock) secondDotPosition(angleDegrees float64) Point {
	return Rotate(Point{
		X: c.center.X,
		Y: c.center.Y - c.radius + 2,
	}, c.center, angleDegrees)
}

func (c Clock) Draw(canvas *Canvas) error {
	return c.shape.Draw(canvas)
}

func (c Clock) hourLineStart(angleDegreesHour float64) Point {
	return Rotate(Point{
		X: c.center.X,
		Y: c.center.Y - c.radius + 2,
	}, c.center, angleDegreesHour)
}

func (c Clock) hourLineEnd(angleDegreesHour float64) Point {
	return Rotate(Point{
		X: c.center.X,
		Y: c.center.Y - c.radius + 5,
	}, c.center, angleDegreesHour)
}
func (c Clock) minuteLineEnd(angleDegrees float64) Point {
	return Rotate(Point{
		X: c.center.X,
		Y: c.center.Y - c.radius + 10,
	}, c.center, angleDegrees)
}
