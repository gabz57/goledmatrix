package effect

import (
	"github.com/gabz57/goledmatrix/canvas"
	"image/color"
	"time"
)

type FadeInOutSceneEffect struct {
	fade          float64
	sceneDuration *time.Duration
	elapsed       time.Duration
	fadeInEnd     time.Duration
	fadeOutStart  time.Duration
}

func (fe *FadeInOutSceneEffect) Update(elapsedBetweenUpdate time.Duration) bool {
	fe.elapsed += elapsedBetweenUpdate
	if fe.elapsed >= *fe.sceneDuration {
		fe.elapsed = 0
	}
	return fe.updateFadeValue()
}

func (fe *FadeInOutSceneEffect) AdaptPixel() func(canvas canvas.Canvas, x int, y int, color *color.Color) {
	return func(canvas canvas.Canvas, x int, y int, colorPt *color.Color) {
		r, g, b, a := (*colorPt).RGBA()
		canvas.Set(x, y, color.RGBA{
			R: uint8((1 - fe.fade) * float64(uint8(r))),
			G: uint8((1 - fe.fade) * float64(uint8(g))),
			B: uint8((1 - fe.fade) * float64(uint8(b))),
			A: uint8(a),
		})
	}
}

func (fe *FadeInOutSceneEffect) updateFadeValue() bool {
	nextFade := fe.computeNextFadeValue()
	if int(fe.fade*100) == int(nextFade*100) {
		return false
	} else {
		fe.fade = nextFade
		return true
	}
}

func (fe *FadeInOutSceneEffect) computeNextFadeValue() float64 {
	var nextFade = fe.fade
	if fe.elapsed < fe.fadeInEnd {
		nextFade = float64(fe.fadeInEnd.Nanoseconds()-fe.elapsed.Nanoseconds()) / float64(fe.fadeInEnd.Nanoseconds())
	} else if fe.elapsed > fe.fadeOutStart {
		nextFade = float64(fe.elapsed.Nanoseconds()-fe.fadeOutStart.Nanoseconds()) / float64((*fe.sceneDuration).Nanoseconds()-fe.fadeOutStart.Nanoseconds())
	} else {
		nextFade = 0
	}
	return nextFade
}

func NewFadeInOutSceneEffect(sceneDuration *time.Duration) *FadeInOutSceneEffect {
	return &FadeInOutSceneEffect{
		fade:          1,
		sceneDuration: sceneDuration,
		fadeInEnd:     time.Duration(int64(float64(sceneDuration.Nanoseconds()) * 0.25)),
		fadeOutStart:  time.Duration(int64(float64(sceneDuration.Nanoseconds()) * 0.75)),
	}
}
