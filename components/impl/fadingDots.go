package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/canvas/effect"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"image/color"
	"math/rand"
	"time"
)

type FadingDot struct {
	shape  *CompositeDrawable
	length int
	//dot          shapes.Dot
	fadeDuration time.Duration
	step         float64
	fade         float64
	fadeOut      bool
	mask         *effect.ColorFaderEffect
}

func NewFadingDot(parent *Graphic, position Point, fadeDuration time.Duration, initialFade float64, initialFadeOut bool, xColor color.Color) *FadingDot {
	var lineGraphic = NewGraphic(parent, nil)
	fadingDot := FadingDot{
		shape:        NewCompositeDrawable(lineGraphic),
		fade:         initialFade,
		fadeOut:      initialFadeOut,
		fadeDuration: fadeDuration,
		mask:         effect.NewColorFaderMask(),
	}
	dot := shapes.NewDot(NewGraphic(lineGraphic, NewLayout(xColor, nil)), position)
	fadingDot.shape.AddDrawable(MaskDrawable(fadingDot.mask, dot))
	fadingDot.Update(0)

	return &fadingDot
}

func (h *FadingDot) Update(elapsedBetweenUpdate time.Duration) bool {
	h.step = float64(elapsedBetweenUpdate) / float64(h.fadeDuration)
	h.fade, h.fadeOut = nextFadeValue(true, h.fadeOut, h.fade, h.step)
	h.mask.SetFade(h.fade)
	return true
}

func (h *FadingDot) IsFaded() bool {
	// we cannot compare with 1 as fade is a float, we minor the value by 1 step
	return h.fadeOut && h.fade >= (1-h.step)
}

func (h FadingDot) Draw(canvas Canvas) error {
	return h.shape.Draw(canvas)
}

type FadingDots struct {
	canvas     Canvas
	lines      []*FadingDot
	origin     Point
	maxX, maxY int
	color      color.Color
}

func NewFadingDots(canvas Canvas, origin Point, nbFadingDots int, optColor color.Color) *FadingDots {
	max := canvas.Bounds().Max
	h := FadingDots{
		canvas: canvas,
		origin: origin,
		maxX:   max.X - origin.X,
		maxY:   max.Y - origin.Y,
		color:  optColor,
	}
	for i := 0; i < nbFadingDots; i++ {
		go h.addDot()
	}
	return &h
}

func (h *FadingDots) addDot() {
	//time.Sleep(time.Duration(components.Random.Int63n(2000)) * time.Millisecond)
	h.lines = append(h.lines, h.generateFadingDot(h.canvas))
}

func (h *FadingDots) replaceDot(i int) {
	h.lines[i] = h.generateFadingDot(h.canvas)
}

func (h *FadingDots) generateFadingDot(c Canvas) *FadingDot {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := r.Intn(c.Bounds().Max.X)
	y := r.Intn(c.Bounds().Max.Y)
	//println("x", x, "y", y)
	var rgba = h.color
	if rgba == nil {
		rgba = randomColor(r)
	}
	return NewFadingDot(
		nil,
		Point{X: x, Y: y},
		time.Duration(Random.Int63n(5000))*time.Millisecond+5000*time.Millisecond,
		1,
		false,
		rgba,
	)
}

func randomColor(r *rand.Rand) color.RGBA {
	rgba := color.RGBA{R: uint8(r.Intn(255)), G: uint8(r.Intn(255)), B: uint8(r.Intn(255)), A: 0xff}
	return rgba
}

func (h *FadingDots) Update(elapsedBetweenUpdate time.Duration) bool {
	dirty := false
	for i, line := range h.lines {
		dirty = line.Update(elapsedBetweenUpdate) || dirty
		if line.IsFaded() {
			h.replaceDot(i)
			dirty = true
		}
	}
	return dirty
}

func (h *FadingDots) Draw(canvas Canvas) error {
	for _, line := range h.lines {
		err := line.Draw(canvas)
		if err != nil {
			return err
		}
	}
	return nil
}
