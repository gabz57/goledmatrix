package main

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/fonts"
	"github.com/paulbellamy/ratecounter"
	"time"
)

type Info struct {
	shape         CompositeDrawable
	now           time.Time
	lastFpsText   time.Time
	timeText      *shapes.Text
	fpsText       *shapes.Text
	updateCounter ratecounter.RateCounter
	drawCounter   ratecounter.RateCounter
	location      *time.Location
}

var infoGraphic = NewGraphic(nil, nil)

func NewInfo(c Canvas) Component {
	location, _ := time.LoadLocation("Europe/Paris")

	info := Info{
		now:         time.Now(),
		lastFpsText: time.Now(),
		shape: CompositeDrawable{
			Graphic:   &infoGraphic,
			Drawables: []*Drawable{},
		},
		updateCounter: *ratecounter.NewRateCounter(1 * time.Second),
		drawCounter:   *ratecounter.NewRateCounter(1 * time.Second),
		location:      location,
	}

	info.timeText = info.buildTimeText()
	info.fpsText = info.buildFPSText(c)
	var drawableTimeText Drawable
	drawableTimeText = info.timeText
	info.shape.AddDrawable(&drawableTimeText)

	var drawableFpsText Drawable
	drawableFpsText = info.fpsText
	info.shape.AddDrawable(&drawableFpsText)

	return &info
}

func (i *Info) Update(elapsedBetweenUpdate time.Duration) {
	now := time.Now().In(i.location)

	defer i.updateCounter.Incr(1)
	if now.Sub(i.now).Milliseconds() < 10 {
		return
	}
	i.now = now
	i.timeText.SetText(TimeToText(i.now))
	if now.Sub(i.lastFpsText).Milliseconds() < 100 {
		return
	}
	i.lastFpsText = now
	i.fpsText.SetText(i.fpsTxt())
}

func (i *Info) Draw(canvas Canvas) error {
	defer i.drawCounter.Incr(1)
	return i.shape.Draw(canvas)
}

func (i *Info) fpsTxt() string {
	if i.updateCounter.Rate() != 0 && i.drawCounter.Rate() != 0 {
		return i.updateCounter.String() + " upd/s  " + i.drawCounter.String() + " FPS"
	}
	return "-"
}

func (i *Info) buildTimeText() *shapes.Text {
	graphic := NewGraphic(i.shape.Graphic, NewLayout(&ColorGreen, nil))
	return shapes.NewText(&graphic,
		Point{
			X: 0,
			Y: -1,
		},
		TimeToText(i.now),
		fonts.Bdf4x6,
	)
}

func (i *Info) buildFPSText(c Canvas) *shapes.Text {
	graphic := NewGraphic(i.shape.Graphic, NewLayout(&ColorGreen, nil))
	return shapes.NewText(&graphic,
		Point{
			X: 0,
			Y: c.Bounds().Max.Y - 6,
		},
		i.fpsTxt(),
		fonts.Bdf4x6,
	)
}
