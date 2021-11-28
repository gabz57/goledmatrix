package effect

import (
	"github.com/gabz57/goledmatrix/canvas"
	"image/color"
	"time"
)

type FadeInOutDynamicEffect struct {
	fade          float64
	sceneDuration time.Duration
	elapsed       time.Duration
	fadeInEnd     time.Duration
	fadeOutStart  time.Duration
}

func (e *FadeInOutDynamicEffect) Update(elapsedBetweenUpdate time.Duration) bool {
	e.elapsed += elapsedBetweenUpdate
	if e.elapsed >= e.sceneDuration {
		e.elapsed = 0
	}
	return e.updateFadeValue()
}

func (e *FadeInOutDynamicEffect) AdaptPixel() func(canvas canvas.Canvas, x int, y int, color color.Color) {
	return func(canvas canvas.Canvas, x int, y int, colorPt color.Color) {
		r, g, b, a := colorPt.RGBA()
		canvas.Set(x, y, color.RGBA{
			R: uint8((1 - e.fade) * float64(uint8(r))),
			G: uint8((1 - e.fade) * float64(uint8(g))),
			B: uint8((1 - e.fade) * float64(uint8(b))),
			A: uint8(a),
		})
	}
}

func (e *FadeInOutDynamicEffect) updateFadeValue() bool {
	nextFade := e.computeNextFadeValue()
	if int(e.fade*100) == int(nextFade*100) {
		return false
	} else {
		e.fade = nextFade
		return true
	}
}

func (e *FadeInOutDynamicEffect) computeNextFadeValue() float64 {
	var nextFade = e.fade
	if e.elapsed < e.fadeInEnd {
		nextFade = float64(e.fadeInEnd.Nanoseconds()-e.elapsed.Nanoseconds()) / float64(e.fadeInEnd.Nanoseconds())
	} else if e.elapsed > e.fadeOutStart {
		nextFade = float64(e.elapsed.Nanoseconds()-e.fadeOutStart.Nanoseconds()) / float64((e.sceneDuration).Nanoseconds()-e.fadeOutStart.Nanoseconds())
	} else {
		nextFade = 0
	}
	return nextFade
}

func NewFadeInOutSceneEffect(sceneDuration time.Duration) *FadeInOutDynamicEffect {
	return &FadeInOutDynamicEffect{
		fade:          1,
		sceneDuration: sceneDuration,
		fadeInEnd:     time.Duration(int64(float64(sceneDuration.Nanoseconds()) * 0.25)),
		fadeOutStart:  time.Duration(int64(float64(sceneDuration.Nanoseconds()) * 0.75)),
	}
}
