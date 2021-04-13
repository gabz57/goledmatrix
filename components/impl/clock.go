package impl

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/fonts"
	"image/color"
	"time"
)

type Clock struct {
	shape             *CompositeDrawable
	now               time.Time
	center            Point
	radius            int
	text              *shapes.Text
	rotatingHour      *shapes.Line
	rotatingHourDot   *shapes.Dot
	rotatingMinute    *shapes.Line
	rotatingMinuteDot *shapes.Dot
	rotatingSecond    *shapes.Dot
	location          *time.Location
}

var clockGraphic = NewGraphic(nil, nil)

func NewClock(canvas Canvas, center Point, radius int) Component {
	location, _ := time.LoadLocation("Europe/Paris")
	var mask Canvas
	mask = NewShadedColorCanvasMask(canvas)

	c := Clock{
		now:      time.Now(),
		center:   center,
		radius:   radius,
		shape:    NewCompositeDrawable(clockGraphic),
		location: location,
	}

	c.shape.AddDrawable(c.buildStaticText(center.AddXY(-9, -radius/2), "Hello"))
	c.shape.AddDrawable(c.buildStaticText(center.AddXY(-7, radius/2-6), "OCTO"))

	c.shape.AddDrawable(Masked(mask, c.buildStaticContourCircle()))
	c.shape.AddDrawable(c.buildStaticCenter())
	c.shape.AddDrawables(c.buildStaticHours())
	c.shape.AddDrawables(c.buildStaticMinutes())

	now := time.Now()
	hour, min, sec := now.Clock()

	c.rotatingSecond = c.buildRotatingSecond(sec)
	var drawableSecond Drawable
	drawableSecond = c.rotatingSecond
	c.shape.AddDrawable(&drawableSecond)

	c.rotatingMinute = c.buildRotatingMinute(min, sec)
	var drawableMinute Drawable
	drawableMinute = c.rotatingMinute
	c.shape.AddDrawable(Masked(mask, &drawableMinute))

	c.rotatingMinuteDot = c.buildRotatingMinuteDot(min, sec)
	var drawableMinuteDot Drawable
	drawableMinuteDot = c.rotatingMinuteDot
	c.shape.AddDrawable(&drawableMinuteDot)

	c.rotatingHour = c.buildRotatingHour(hour, min)
	var drawableHour Drawable
	drawableHour = c.rotatingHour
	c.shape.AddDrawable(Masked(mask, &drawableHour))

	c.rotatingHourDot = c.buildRotatingHourDot(hour, min)
	var drawableHourDot Drawable
	drawableHourDot = c.rotatingHourDot
	c.shape.AddDrawable(&drawableHourDot)

	return &c
}

func (c *Clock) Update(elapsedBetweenUpdate time.Duration) {
	now := time.Now().In(c.location)
	if now.Sub(c.now).Milliseconds() < 10 {
		// skip
		return
	}
	c.now = now
	hour, min, sec := c.now.Clock()

	aDHour := angleDegreesHour(hour, min)
	c.rotatingHour.SetLine(
		c.center,
		c.hourLineEnd(aDHour),
	)
	c.rotatingHourDot.SetPosition(
		c.hourLineEnd(aDHour),
	)

	c.rotatingMinute.SetLine(
		c.center,
		c.minuteLineEnd(angleDegreesMinute(min, sec)),
	)
	c.rotatingMinuteDot.SetPosition(
		c.minuteLineEnd(angleDegreesMinute(min, sec)),
	)

	c.rotatingSecond.SetPosition(
		c.secondDotPosition(angleDegreesSecond(sec, c.now)),
	)
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

func (c *Clock) buildStaticText(position Point, txt string) *Drawable {
	graphic := NewGraphic(c.shape.Graphic, nil)
	var circle Drawable
	circle = shapes.NewText(&graphic, position, txt, fonts.Bdf4x6)
	return &circle
}

func (c *Clock) buildStaticContourCircle() *Drawable {
	graphic := NewGraphic(c.shape.Graphic, nil)
	var circle Drawable
	circle = shapes.NewCircle(&graphic, c.center, c.radius, false)
	return &circle
}

func (c *Clock) buildStaticCenter() *Drawable {
	graphic := NewGraphic(c.shape.Graphic, nil)
	var dot Drawable
	dot = shapes.NewDot(&graphic, c.center)
	return &dot
}

func (c *Clock) buildStaticHours() []*Drawable {
	drawables := make([]*Drawable, 12)
	for hour := 0; hour < 12; hour++ {
		drawables[hour] = c.buildStaticHour(hour)
	}
	return drawables
}

func (c *Clock) buildStaticHour(hour int) *Drawable {
	graphic := NewGraphic(c.shape.Graphic, nil)
	var line Drawable
	line = shapes.NewLine(&graphic,
		Rotate(Point{
			X: c.center.X,
			Y: c.center.Y - c.radius + 5,
		}, c.center, angleDegreesHour(hour, 0)),
		Rotate(Point{
			X: c.center.X,
			Y: c.center.Y - c.radius + 2,
		}, c.center, angleDegreesHour(hour, 0)),
	)
	return &line
}

func (c *Clock) buildStaticMinutes() []*Drawable {
	drawables := make([]*Drawable, 60)
	for minutes := 0; minutes < 60; minutes++ {
		drawables[minutes] = c.buildStaticMinute(c.center, c.radius, minutes)
	}
	return drawables
}

func (c *Clock) buildStaticMinute(center Point, radius int, minute int) *Drawable {
	graphic := NewGraphic(c.shape.Graphic, nil)
	var dot Drawable
	dot = shapes.NewDot(&graphic,
		Rotate(Point{
			X: center.X,
			Y: center.Y - radius + 2,
		}, center, angleDegreesMinute(minute, 0)),
	)
	return &dot
}

func (c *Clock) buildRotatingHour(hour, min int) *shapes.Line {
	graphic := NewGraphic(c.shape.Graphic, NewLayout(ColorBlue, nil))
	return shapes.NewLine(&graphic,
		c.center,
		c.hourLineEnd(angleDegreesHour(hour, min)),
	)
}

func (c *Clock) buildRotatingHourDot(hour, min int) *shapes.Dot {
	graphic := NewGraphic(c.shape.Graphic, NewLayout(ColorRed, nil))
	return shapes.NewDot(&graphic,
		c.hourLineEnd(angleDegreesHour(hour, min)),
	)
}

func (c *Clock) buildRotatingMinute(min, sec int) *shapes.Line {
	graphic := NewGraphic(c.shape.Graphic, NewLayout(color.White, nil))
	return shapes.NewLine(&graphic,
		c.center,
		c.minuteLineEnd(angleDegreesMinute(min, sec)),
	)
}

func (c *Clock) buildRotatingMinuteDot(min, sec int) *shapes.Dot {
	graphic := NewGraphic(c.shape.Graphic, NewLayout(ColorRed, nil))
	return shapes.NewDot(&graphic,
		c.minuteLineEnd(angleDegreesMinute(min, sec)),
	)
}

func (c *Clock) buildRotatingSecond(sec int) *shapes.Dot {
	graphic := NewGraphic(c.shape.Graphic, NewLayout(ColorRed, nil))
	return shapes.NewDot(&graphic,
		c.secondDotPosition(angleDegreesSecond(sec, c.now)),
	)
}

func (c *Clock) secondDotPosition(angleDegrees float64) Point {
	return Rotate(Point{
		X: c.center.X,
		Y: c.center.Y - c.radius + 2,
	}, c.center, angleDegrees)
}

func (c *Clock) Draw(canvas Canvas) error {
	return c.shape.Draw(canvas)
}

func (c *Clock) hourLineEnd(angleDegreesHour float64) Point {
	return Rotate(Point{
		X: c.center.X,
		Y: c.center.Y - c.radius/2,
	}, c.center, angleDegreesHour)
}

func (c *Clock) minuteLineEnd(angleDegrees float64) Point {
	return Rotate(Point{
		X: c.center.X,
		Y: c.center.Y - c.radius + 10,
	}, c.center, angleDegrees)
}
