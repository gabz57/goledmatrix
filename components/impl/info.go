package impl

import (
	"fmt"
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/fonts"
	"github.com/paulbellamy/ratecounter"
	"time"
)

type Info struct {
	shape         *CompositeDrawable
	now           time.Time
	lastFpsText   time.Time
	timeText      *shapes.Text
	fpsText       *shapes.Text
	updateText    *shapes.Text
	updateCounter ratecounter.RateCounter
	drawCounter   ratecounter.RateCounter
	location      *time.Location
}

var infoGraphic = NewGraphic(nil, nil)

func NewInfo(c Canvas) *Info {
	location, _ := time.LoadLocation("Europe/Paris")

	i := Info{
		now:           time.Now(),
		lastFpsText:   time.Now(),
		shape:         NewCompositeDrawable(infoGraphic),
		updateCounter: *ratecounter.NewRateCounter(1 * time.Second),
		drawCounter:   *ratecounter.NewRateCounter(1 * time.Second),
		location:      location,
	}

	i.timeText = i.buildTimeText()
	var drawableTimeText Drawable = i.timeText
	i.shape.AddDrawable(&drawableTimeText)

	i.fpsText = i.buildFPSText(c)
	var drawableFpsText Drawable = i.fpsText
	i.shape.AddDrawable(&drawableFpsText)

	i.updateText = i.buildUpdateText(c)
	var drawableUpdateText Drawable = i.updateText
	i.shape.AddDrawable(&drawableUpdateText)

	return &i
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
	i.updateText.SetText(i.updateTxt())
}

func (i *Info) Draw(canvas Canvas) error {
	defer i.drawCounter.Incr(1)
	return i.shape.Draw(canvas)
}

func (i *Info) fpsTxt() string {
	if i.updateCounter.Rate() != 0 && i.drawCounter.Rate() != 0 {
		return fmt.Sprintf("%3d", i.drawCounter.Rate()) + " FPS"
	}
	return "-"
}

func (i *Info) updateTxt() string {
	if i.updateCounter.Rate() != 0 && i.drawCounter.Rate() != 0 {
		return fmt.Sprintf("%3d", i.updateCounter.Rate()) + " upd/s"
	}
	return "-"
}

func (i *Info) buildTimeText() *shapes.Text {
	return shapes.NewText(
		NewGraphic(i.shape.Graphic, NewLayout(ColorGreen, nil)),
		Point{
			X: 1,
			Y: 0,
		},
		TimeToText(i.now),
		fonts.Bdf4x6,
	)
}
func (i *Info) buildUpdateText(c Canvas) *shapes.Text {
	return shapes.NewText(
		NewGraphic(i.shape.Graphic, NewLayout(ColorGreen, nil)),
		Point{
			X: c.Bounds().Max.X - 35,
			Y: c.Bounds().Max.Y - 7,
		},
		i.updateTxt(),
		fonts.Bdf4x6,
	)
}

func (i *Info) buildFPSText(c Canvas) *shapes.Text {
	return shapes.NewText(
		NewGraphic(i.shape.Graphic, NewLayout(ColorGreen, nil)),
		Point{
			X: 0,
			Y: c.Bounds().Max.Y - 7,
		},
		i.fpsTxt(),
		fonts.Bdf4x6,
	)
}
