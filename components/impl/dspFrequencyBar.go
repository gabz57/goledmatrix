package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"image/color"
	"time"
)

type frequencyBar struct {
	shape           *CompositeDrawable
	current         *shapes.Line
	current2        *shapes.Line
	fill            *shapes.Rectangle
	nbLedsY         int
	prevNbLedsY     int
	initialPosition Point
	below           bool
	slowRate        float64
	prevNbSlowLedsY int
	nbSlowLedsY     int
}

func (f *frequencyBar) SetRate(rate float64) {
	if rate > f.slowRate {
		f.slowRate = rate
	}
	f.nbLedsY = rateToNbLeds(rate)
}

func rateToNbLeds(rate float64) int {
	return int(rate * float64(MAX_HEIGHT-1))
}

func (f *frequencyBar) Update(elapsedBetweenUpdate time.Duration) bool {
	updated := false
	heightTop := MAX_HEIGHT - 1 - f.nbLedsY
	fillWidth := BAR_WIDTH - 1
	if f.prevNbLedsY != f.nbLedsY {
		if f.below {
			f.fill.SetMin(f.initialPosition.Add(Point{X: 0, Y: heightTop}))
			f.fill.SetMax(f.initialPosition.Add(Point{X: 0, Y: heightTop}).Add(Point{X: fillWidth, Y: 2 * f.nbLedsY}))
		} else {
			f.fill.SetMin(f.initialPosition.Add(Point{X: 0, Y: heightTop}))
			f.fill.SetMax(f.initialPosition.Add(Point{X: 0, Y: heightTop}).Add(Point{X: fillWidth, Y: f.nbLedsY}))
		}

		f.prevNbLedsY = f.nbLedsY

		updated = true
	}
	f.slowRate = f.slowRate * 0.995
	f.nbSlowLedsY = rateToNbLeds(f.slowRate)
	if f.prevNbSlowLedsY != f.nbSlowLedsY {
		if f.below {
			f.current.SetLine(
				f.initialPosition.Add(Point{X: 0, Y: heightTop}),
				f.initialPosition.Add(Point{X: fillWidth, Y: heightTop}),
			)
			minY := MAX_HEIGHT - 1 + f.nbLedsY
			f.current2.SetLine(
				f.initialPosition.Add(Point{X: 0, Y: minY}),
				f.initialPosition.Add(Point{X: fillWidth, Y: minY}),
			)
		} else {
			f.current.SetLine(
				f.initialPosition.Add(Point{X: 0, Y: heightTop}),
				f.initialPosition.Add(Point{X: fillWidth, Y: heightTop}),
			)
		}

		f.prevNbSlowLedsY = f.nbSlowLedsY
		updated = true
	}

	return updated
}

func (f *frequencyBar) Draw(canvas Canvas) error {
	return f.shape.Draw(canvas)
}

func newFrequencyBar(graphic *Graphic, initialPosition Point, below bool, color color.Color) *frequencyBar {
	freqBar := frequencyBar{
		shape:           NewCompositeDrawable(graphic),
		initialPosition: initialPosition,
		prevNbLedsY:     -1,
		below:           below,
	}

	if freqBar.below {

		freqBar.fill = shapes.NewRectangle(NewGraphic(graphic,
			NewLayout(color, color)),
			initialPosition.Add(Point{X: 0, Y: MAX_HEIGHT - 1 - freqBar.nbLedsY}),
			Point{X: BAR_WIDTH - 1, Y: 2 * freqBar.nbLedsY},
			true)
		freqBar.shape.AddDrawable(freqBar.fill)
		freqBar.current = shapes.NewLine(
			NewGraphic(graphic, NewLayout(color, color)),
			initialPosition.Add(Point{X: 0, Y: MAX_HEIGHT - 1 - freqBar.nbLedsY}),
			initialPosition.Add(Point{X: BAR_WIDTH - 1, Y: MAX_HEIGHT - 1 - freqBar.nbLedsY}),
		)
		freqBar.shape.AddDrawable(freqBar.current)
		freqBar.current2 = shapes.NewLine(
			NewGraphic(graphic, NewLayout(color, color)),
			initialPosition.Add(Point{X: 0, Y: MAX_HEIGHT - 1 + freqBar.nbLedsY}),
			initialPosition.Add(Point{X: BAR_WIDTH - 1, Y: MAX_HEIGHT - 1 + freqBar.nbLedsY}),
		)
		freqBar.shape.AddDrawable(freqBar.current2)

	} else {

		freqBar.fill = shapes.NewRectangle(
			NewGraphic(graphic, NewLayout(color, color)),
			initialPosition.Add(Point{X: 0, Y: MAX_HEIGHT - 1 - freqBar.nbLedsY}),
			Point{X: BAR_WIDTH - 1, Y: freqBar.nbLedsY},
			true)
		freqBar.shape.AddDrawable(freqBar.fill)
		freqBar.current = shapes.NewLine(
			NewGraphic(graphic, NewLayout(color, color)),
			initialPosition.Add(Point{X: 0, Y: MAX_HEIGHT - 1 - freqBar.nbLedsY}),
			initialPosition.Add(Point{X: BAR_WIDTH - 1, Y: MAX_HEIGHT - 1 - freqBar.nbLedsY}),
		)
		freqBar.shape.AddDrawable(freqBar.current)
	}

	// hack to avoid flash on first rendering
	freqBar.Update(0)

	return &freqBar
}
