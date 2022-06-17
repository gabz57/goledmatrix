package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/canvas/effect"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"image"
	"image/color"
	"math/rand"
	"time"
)

type FadingLine struct {
	shape        *CompositeDrawable
	length       int
	position     Point
	dots         []shapes.Dot
	fadeDuration time.Duration
	step         float64
	fade         float64
	fadeOut      bool
	mask         *effect.ColorFaderEffect
}

type ramp struct {
	left  int
	right int
}

func NewFadingLine(canvas Canvas, parent *Graphic, position Point, length int, fullLightLength int, fadeDuration time.Duration, initialFade float64, initialFadeOut bool, xColor color.Color) *FadingLine {
	var lineGraphic = NewGraphic(parent, nil)
	line := FadingLine{
		shape:        NewCompositeDrawable(lineGraphic),
		fade:         initialFade,
		fadeOut:      initialFadeOut,
		fadeDuration: fadeDuration,
		mask:         effect.NewColorFaderMask(),
		dots:         make([]shapes.Dot, length),
	}

	rampUp, middleLine, rampDown := prepareRamps(position, length, fullLightLength)
	var dots = make([]shapes.Dot, length)
	x := position.X
	for i := 0; i < length; i++ {
		rgba := fadeColor(x, middleLine, rampUp, rampDown, xColor)

		dots[i] = *shapes.NewDot(NewGraphic(lineGraphic, NewLayout(rgba, nil)), Point{X: x, Y: position.Y})
		x++
		line.shape.AddDrawable(MaskDrawable(line.mask, &dots[i]))
	}
	//line.shape.AddDrawable(MaskDrawable(line.mask, line.line))
	// hack to avoid flash on first rendering
	line.Update(0)

	return &line
}

func prepareRamps(position Point, length int, fullLightLength int) (ramp, ramp, ramp) {
	left := position.X
	var middle = left + (length / 2)
	var rampUp = ramp{
		left:  left,
		right: middle - (fullLightLength / 2),
	}

	var middleLine = ramp{
		left:  middle - (fullLightLength / 2),
		right: middle + (fullLightLength / 2),
	}

	var rampDown = ramp{
		left:  middle + (fullLightLength / 2),
		right: left + length,
	}
	return rampUp, middleLine, rampDown
}

func fadeColor(x int, middleLine ramp, rampUp ramp, rampDown ramp, xColor color.Color) color.RGBA {
	fade := float64(0)

	if x < middleLine.left {
		// rampUp 1..0
		max := rampUp.right - rampUp.left
		current := x - rampUp.left
		fade = float64(1) - float64(current)/float64(max)
	} else if x < rampDown.left {
		// middle 0
		fade = 0
	} else {
		// rampdown 0..1
		max := rampDown.right - rampDown.left
		current := x - rampDown.left
		fade = float64(current) / float64(max)
	}

	r, g, b, a := xColor.RGBA()
	rgba := color.RGBA{
		R: uint8((1 - fade) * float64(uint8(r))),
		G: uint8((1 - fade) * float64(uint8(g))),
		B: uint8((1 - fade) * float64(uint8(b))),
		A: uint8(a),
	}
	return rgba
}

func (h *FadingLine) Update(elapsedBetweenUpdate time.Duration) bool {
	h.step = float64(elapsedBetweenUpdate) / float64(h.fadeDuration)
	h.fade, h.fadeOut = nextFadeValue(true, h.fadeOut, h.fade, h.step)
	h.mask.SetFade(h.fade)
	return true
}

func (h *FadingLine) IsFaded() bool {
	// we cannot compare with 1 as fade is a float, we minor the value by 1 step
	return h.fadeOut && h.fade >= (1-h.step)
}

func (h FadingLine) Draw(canvas Canvas) error {
	return h.shape.Draw(canvas)
}

//
//func (h *FadingLine) GetPosition() Point {
//	return h.line.Graphic.ComputedOffset()
//}
//
//func (h *FadingLine) SetPosition(point Point) {
//	h.line.Graphic.SetOffset(point)
//}
type FadingLines struct {
	canvas     Canvas
	lines      []*FadingLine
	origin     Point
	maxX, maxY int
}

func NewFadingLines(canvas Canvas, origin Point, nbFadingLines int) *FadingLines {
	max := canvas.Bounds().Max
	h := FadingLines{
		canvas: canvas,
		origin: origin,
		maxX:   max.X - origin.X,
		maxY:   max.Y - origin.Y,
	}
	for i := 0; i < nbFadingLines; i++ {
		go h.addLine()
	}
	return &h
}

func (h *FadingLines) addLine() {
	//time.Sleep(time.Duration(components.Random.Int63n(2000)) * time.Millisecond)
	h.lines = append(h.lines, generateFadingLine(h.canvas, h.lines))
}

func (h *FadingLines) replaceLine(i int) {
	h.lines[i] = generateFadingLine(h.canvas, h.lines)
}

func generateFadingLine(c Canvas, currentLines []*FadingLine) *FadingLine {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x, y, length, fullLightLength := nextPositionAndLenghts(c.Bounds().Max, currentLines, r)

	return NewFadingLine(c,
		nil,
		Point{X: x, Y: y},
		length,
		fullLightLength,
		time.Duration(Random.Int63n(5000))*time.Millisecond+5000*time.Millisecond,
		1,
		false,
		color.RGBA{R: uint8(r.Intn(255)), G: uint8(r.Intn(255)), B: uint8(r.Intn(255)), A: 0xff},
	)
}

func nextPositionAndLenghts(bounds image.Point, lines []*FadingLine, r *rand.Rand) (int, int, int, int) {
	var x, y, length int

	freeFound := false
	for !freeFound {
		x = r.Intn(bounds.X)
		y = r.Intn(bounds.Y)
		length = 5
		if bounds.X-x > 5 {
			length = r.Intn(bounds.X - x)
		}
		freeFound = isFree(lines, x, y, length)
	}
	//println("x", x, "y", y, "length", length)
	fullLightLength := 0
	if float64(length)/5 >= 1 {
		fullLightLength = r.Intn(length / 5)
	}
	return x, y, length, fullLightLength
}

func isFree(lines []*FadingLine, xLeft int, y int, length int) bool {
	for _, line := range lines {
		if y == line.position.Y {
			lineLeft := line.position.X
			lineRight := lineLeft + line.length
			xRight := xLeft + length
			if lineLeft <= xLeft && xLeft <= lineRight ||
				lineLeft <= xRight && xRight <= lineRight ||
				xLeft <= lineLeft && xRight >= lineRight {
				return false
			}
		}
	}
	return true
}

func (h *FadingLines) Update(elapsedBetweenUpdate time.Duration) bool {
	dirty := false
	for i, line := range h.lines {
		dirty = line.Update(elapsedBetweenUpdate) || dirty
		if line.IsFaded() {
			h.replaceLine(i)
			dirty = true
		}
	}
	return dirty
}

func (h *FadingLines) Draw(canvas Canvas) error {
	for _, line := range h.lines {
		err := line.Draw(canvas)
		if err != nil {
			return err
		}
	}
	return nil
}
