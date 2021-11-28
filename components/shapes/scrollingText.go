package shapes

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/canvas/effect"
	"github.com/gabz57/goledmatrix/canvas/fonts"
	"github.com/gabz57/goledmatrix/components"
	"image"
	"time"
)

type ScrollingText struct {
	visibleArea image.Rectangle
	visibleMask effect.VisibleMask
	text        Text
	delay       time.Duration
	duration    time.Duration
	elapsed     time.Duration
	maxScroll   int
	masked      components.Drawable
}

func NewScrollingText(graphic *components.Graphic, c canvas.Canvas, txt string, f fonts.MatrixFont, position canvas.Point, visibleArea image.Rectangle, duration time.Duration) *ScrollingText {
	visible := visibleArea.Add(image.Point{
		X: position.X,
		Y: position.Y,
	}).Add(image.Point(graphic.ComputedOffset()))
	st := ScrollingText{
		visibleArea: visible,
		visibleMask: *effect.NewVisibleMask(visible),
		text:        *NewText(graphic, position, txt, f),
		duration:    duration,
		delay:       time.Duration(float64(duration.Nanoseconds()) * 0.15),
		elapsed:     0,
	}
	st.masked = components.MaskDrawable(&st.visibleMask, &st.text)
	st.SetText(txt)
	return &st
}

func (st *ScrollingText) Draw(canvas canvas.Canvas) error {
	return st.masked.Draw(canvas)
}

func (st *ScrollingText) Update(elapsedBetweenUpdate time.Duration) bool {
	st.elapsed += elapsedBetweenUpdate
	if st.elapsed >= st.duration {
		st.elapsed = 0
	}
	return st.updateOffset()
}

func (st *ScrollingText) updateOffset() bool {
	if st.maxScroll <= 0 {
		return false
	}
	// nextOffsetX: [0 to maxScroll]
	nextOffset := st.computeNextOffset()

	if st.visibleMask.GetOffset().Eq(nextOffset) {
		return false
	} else {
		st.visibleMask.SetOffset(nextOffset)
		return true
	}
}

func (st *ScrollingText) computeNextOffset() image.Point {
	var ratio float64 = 0
	if st.elapsed < st.delay {
		ratio = 0
	} else if st.elapsed < (st.duration/2 - st.delay) {
		// advance
		ratio = float64(st.elapsed.Milliseconds()-st.delay.Milliseconds()) / float64(st.duration.Milliseconds()/2-2*st.delay.Milliseconds())
		if ratio > 1 {
			ratio = 1
		}
	} else if st.elapsed > (st.delay + st.duration/2) {
		// backward
		ratio = float64(1) - float64(st.elapsed.Milliseconds()-(st.duration.Milliseconds()/2+st.delay.Milliseconds()))/float64(st.duration.Milliseconds()/2-2*st.delay.Milliseconds())
		if ratio < 0 {
			ratio = 0
		}
	} else {
		// pause at the end
		ratio = 1
	}
	//println("ratio " + strconv.FormatFloat(ratio, 'f', 2, 64), st.elapsed.String())
	nextOffset := image.Point{
		X: -int(float64(st.maxScroll) * ratio),
		Y: 0,
	}
	return nextOffset
}

func (st *ScrollingText) SetText(label string) {
	st.text.SetText(label)
	st.initMaxScroll()
}

func (st *ScrollingText) initMaxScroll() {
	maxScroll := st.text.Bounds().Dx() - st.visibleArea.Dx()
	if maxScroll > 0 {
		st.maxScroll = maxScroll
	}
}
