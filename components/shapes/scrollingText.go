package shapes

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/fonts"
	"image"
	"time"
)

type ScrollingText struct {
	visibleArea image.Rectangle
	mask        *canvas.VisibleMask
	text        *Text
	delay       time.Duration
	duration    time.Duration
	elapsed     time.Duration
	maxScroll   int
	masked      *components.Drawable

	//pauseStart bool
	//pauseEnd bool
}

func NewScrollingText(graphic *components.Graphic, c canvas.Canvas, txt string, f fonts.MatrixFont, position canvas.Point, visibleArea image.Rectangle, duration time.Duration) *ScrollingText {
	visible := visibleArea.Add(image.Point{
		X: position.X,
		Y: position.Y,
	}).Add(image.Point(graphic.ComputedOffset()))
	st := ScrollingText{
		visibleArea: visible,
		mask:        canvas.NewVisibleMask(c, visible),
		text:        NewText(graphic, position, txt, f),
		duration:    duration,
		delay:       2000 * time.Millisecond,
		elapsed:     0,
	}
	var drawableText components.Drawable = st.text
	var cMask canvas.Canvas = st.mask
	st.masked = components.Masked(&cMask, &drawableText)

	st.SetText(txt)
	return &st
}

func (st *ScrollingText) Draw(canvas canvas.Canvas) error {
	return (*st.masked).Draw(canvas)
}

func (st *ScrollingText) Update(elapsedBetweenUpdate time.Duration) bool {
	st.elapsed += elapsedBetweenUpdate
	if st.elapsed > (st.duration + 3*st.delay) {
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

	if st.mask.GetOffset().Eq(nextOffset) {
		return false
	} else {
		st.mask.SetOffset(nextOffset)
		return true
	}
}

func (st *ScrollingText) computeNextOffset() image.Point {
	var ratio float64 = 0
	if st.elapsed > st.delay {
		if st.elapsed < (st.delay + st.duration/2) {
			// advance
			ratio = float64(st.elapsed.Milliseconds()-st.delay.Milliseconds()) / float64(st.duration.Milliseconds()/2)
			if ratio > 1 {
				ratio = 1
			}
		} else if st.elapsed < (2*st.delay + st.duration/2) {
			// pause at the end
			ratio = 1
		} else {
			// backward
			ratio = float64(1) - float64(st.elapsed.Milliseconds()-2*st.delay.Milliseconds()-st.duration.Milliseconds()/2)/float64(st.duration.Milliseconds()/2)
			if ratio < 0 {
				ratio = 0
			}
		}
	}
	//println("ratio" + strconv.FormatFloat(ratio, 'f', 3, 64))
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
