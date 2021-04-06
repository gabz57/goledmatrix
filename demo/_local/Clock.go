package main

import (
	"fmt"
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
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
	hour, min, sec := time.Now().Clock()
	clock := Clock{
		now:    time.Now(),
		center: center,
		radius: radius,
		shape: CompositeDrawable{
			Graphic:   &clockGraphic,
			Drawables: []Drawable{},
		},
	}
	clock.shape.AddDrawable(clock.buildDrawableContourCircle())
	clock.shape.AddDrawable(clock.buildDrawableCenter())
	clock.shape.AddDrawables(clock.buildDrawableHours())
	clock.shape.AddDrawables(clock.buildDrawableMinutes())

	clock.text = clock.buildText(hour, min, sec)
	clock.rotatingHour = clock.buildRotatingHour(hour, min)
	clock.rotatingMinute = clock.buildRotatingMinute(min, sec)
	clock.rotatingSecond = clock.buildRotatingSecond(sec)

	return clock
}

func (c Clock) Update() {
	now := time.Now()
	if now.Sub(c.now) > 10*time.Millisecond {
		c.now = now
		hour, min, sec := now.Clock()

		angleDegreesHour := float64((hour + min/60) * 30)
		c.rotatingHour.SetLine(
			Rotate(Point{
				X: c.center.X,
				Y: 2,
			}, c.center, angleDegreesHour),
			Rotate(Point{
				X: c.center.X,
				Y: 5,
			}, c.center, angleDegreesHour),
		)

		angleDegreesMinute := (float64(min) + float64(sec/60)) * 6
		c.rotatingMinute.SetLine(
			c.center,
			Rotate(Point{
				X: c.center.X,
				Y: 10,
			}, c.center, angleDegreesMinute))

		angleDegreesSecond := (float64(sec) + (float64(c.now.Nanosecond()) / 1000000000)) * 6
		c.rotatingSecond.SetDot(
			Rotate(Point{
				X: c.center.X + c.radius - 2,
				Y: 0,
			}, c.center, angleDegreesSecond))
		c.text.SetText(timeToText(hour, min, sec))
	}
}

func timeToText(hour int, min int, sec int) string {
	return fmt.Sprintf("%02d", hour) + ":" + fmt.Sprintf("%02d", min) + ":" + fmt.Sprintf("%02d", sec)
}

func (c Clock) buildDrawableContourCircle() Drawable {
	graphic := NewGraphic(c.shape.Graphic, nil)
	return &*shapes.NewCircle(&graphic, c.center, c.radius, false)
}

func (c Clock) buildDrawableCenter() Drawable {
	graphic := NewGraphic(c.shape.Graphic, nil)
	return shapes.NewDot(&graphic, c.center)
}

func (c Clock) buildDrawableHours() []Drawable {
	drawables := make([]Drawable, 12)
	for hour := 0; hour < 12; hour++ {
		drawables[hour] = c.buildDrawableHour(c.center, c.radius, hour)
	}
	return drawables
}

func (c Clock) buildDrawableHour(center Point, radius int, hour int) Drawable {
	angleDegrees := float64(hour * 30)

	graphic := NewGraphic(c.shape.Graphic, nil)
	return shapes.NewLine(&graphic,
		Rotate(Point{
			X: center.X + radius - 5,
			Y: 0,
		}, center, angleDegrees),
		Rotate(Point{
			X: center.X + radius - 2,
			Y: 0,
		}, center, angleDegrees),
	)
}

func (c Clock) buildDrawableMinutes() []Drawable {
	drawables := make([]Drawable, 60)
	for minutes := 0; minutes < 60; minutes++ {
		drawables[minutes] = c.buildDrawableMinute(c.center, c.radius, minutes)
	}
	return drawables
}

func (c Clock) buildDrawableMinute(center Point, radius int, minute int) Drawable {
	angleDegrees := float64(minute * 6)
	graphic := NewGraphic(c.shape.Graphic, nil)
	return shapes.NewDot(&graphic,
		Rotate(Point{
			X: center.X + radius - 2,
			Y: 0,
		}, center, angleDegrees),
	)
}

func (c Clock) buildText(hour int, min int, sec int) *shapes.Text {
	graphic := NewGraphic(c.shape.Graphic, NewLayout(&ColorBlue, nil))
	return shapes.NewText(&graphic,
		Point{
			X: 0,
			Y: 0,
		},
		timeToText(hour, min, sec),
	)
}

// current hour (line)
func (c Clock) buildRotatingHour(hour, min int) *shapes.Line {
	angleDegrees := float64((hour + min/60) * 30)
	graphic := NewGraphic(c.shape.Graphic, NewLayout(&ColorBlue, nil))
	return shapes.NewLine(&graphic,
		Rotate(Point{
			X: c.center.X,
			Y: 2,
		}, c.center, angleDegrees),
		Rotate(Point{
			X: c.center.X,
			Y: 5,
		}, c.center, angleDegrees),
	)
}

// current minute (line)
func (c Clock) buildRotatingMinute(min, sec int) *shapes.Line {
	angleDegrees := (float64(min) + float64(sec/60)) * 6
	graphic := NewGraphic(c.shape.Graphic, NewLayout(ColorViolet, nil))
	return shapes.NewLine(&graphic,
		c.center,
		Rotate(Point{
			X: c.center.X,
			Y: 10,
		}, c.center, angleDegrees),
	)
}

// current second (dot)
func (c Clock) buildRotatingSecond(sec int) *shapes.Dot {
	angleDegrees := (float64(sec) + (float64(c.now.Nanosecond()) / 1000000000)) * 6
	//angleDegrees := (float64(sec)) * 6
	graphic := NewGraphic(c.shape.Graphic, NewLayout(ColorRed, nil))
	return shapes.NewDot(&graphic,
		Rotate(Point{
			X: c.center.X + c.radius - 2,
			Y: 0,
		}, c.center, angleDegrees),
	)
}

func (c Clock) Draw(canvas *Canvas) error {
	err := c.shape.Draw(canvas)
	if err != nil {
		return err
	}
	return c.DrawCurrentTime(canvas)
}

func (c Clock) DrawCurrentTime(canvas *Canvas) error {
	err := c.text.Draw(canvas)
	if err != nil {
		return err
	}
	err = c.rotatingHour.Draw(canvas)
	if err != nil {
		return err
	}
	err = c.rotatingMinute.Draw(canvas)
	if err != nil {
		return err
	}
	return c.rotatingSecond.Draw(canvas)
}
