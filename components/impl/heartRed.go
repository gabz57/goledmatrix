package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"time"
)

const (
	HeartRedWidth  = 13
	HeartRedHeight = 12
)

type HeartRed struct {
	shape        *CompositeDrawable
	heart        *shapes.Free
	mask         *ColorFaderCanvasMask
	fadeOut      bool
	fade         float64
	fadeDuration time.Duration
	step         float64
}

func NewHeartRed(canvas Canvas, parent *Graphic, initialPosition Point, fadeDuration time.Duration, initialFade float64, initialFadeOut bool) *HeartRed {
	var heartGraphic = NewOffsetGraphic(parent, nil, initialPosition)
	heart := HeartRed{
		shape:        NewCompositeDrawable(heartGraphic),
		fade:         initialFade,
		fadeOut:      initialFadeOut,
		fadeDuration: fadeDuration,
		mask:         NewColorFaderMask(canvas),
		heart:        shapes.NewFree(heartGraphic, redHeartPixels()),
	}

	var drawableHeartRed Drawable = heart.heart
	var cMask Canvas = heart.mask
	heart.shape.AddDrawable(Masked(&cMask, &drawableHeartRed))
	// hack to avoid flash on first rendering
	heart.Update(0)

	return &heart
}

func (h *HeartRed) Update(elapsedBetweenUpdate time.Duration) bool {
	h.step = float64(elapsedBetweenUpdate) / float64(h.fadeDuration)
	h.fade, h.fadeOut = nextFadeValue(true, h.fadeOut, h.fade, h.step)
	h.mask.SetFade(h.fade)
	return true
}

func (h *HeartRed) IsFaded() bool {
	// we cannot compare with 1 as fade is a float, we minor the value by 1 step
	return h.fadeOut && h.fade >= (1-h.step)
}

func (h HeartRed) Draw(canvas Canvas) error {
	return h.shape.Draw(canvas)
}

func (h *HeartRed) GetPosition() Point {
	return h.heart.Graphic.ComputedOffset()
}

func (h *HeartRed) SetPosition(point Point) {
	h.heart.Graphic.SetOffset(point)
}

func redHeartPixels() (pixels []Pixel) {
	pixels = append(pixels,
		Pixel{X: 6, Y: 11, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 5, Y: 10, C: &ColorBlack},
		Pixel{X: 6, Y: 10, C: &ColorRed},
		Pixel{X: 7, Y: 10, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 4, Y: 9, C: &ColorBlack},
		Pixel{X: 5, Y: 9, C: &ColorRed},
		Pixel{X: 6, Y: 9, C: &ColorRed},
		Pixel{X: 7, Y: 9, C: &ColorRed},
		Pixel{X: 8, Y: 9, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 3, Y: 8, C: &ColorBlack},
		Pixel{X: 4, Y: 8, C: &ColorRed},
		Pixel{X: 5, Y: 8, C: &ColorRed},
		Pixel{X: 6, Y: 8, C: &ColorRed},
		Pixel{X: 7, Y: 8, C: &ColorRed},
		Pixel{X: 8, Y: 8, C: &ColorRed},
		Pixel{X: 9, Y: 8, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 2, Y: 7, C: &ColorBlack},
		Pixel{X: 3, Y: 7, C: &ColorRed},
		Pixel{X: 4, Y: 7, C: &ColorRed},
		Pixel{X: 5, Y: 7, C: &ColorRed},
		Pixel{X: 6, Y: 7, C: &ColorRed},
		Pixel{X: 7, Y: 7, C: &ColorRed},
		Pixel{X: 8, Y: 7, C: &ColorRed},
		Pixel{X: 9, Y: 7, C: &ColorRed},
		Pixel{X: 10, Y: 7, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 1, Y: 6, C: &ColorBlack},
		Pixel{X: 2, Y: 6, C: &ColorRed},
		Pixel{X: 3, Y: 6, C: &ColorRed},
		Pixel{X: 4, Y: 6, C: &ColorRed},
		Pixel{X: 5, Y: 6, C: &ColorRed},
		Pixel{X: 6, Y: 6, C: &ColorRed},
		Pixel{X: 7, Y: 6, C: &ColorRed},
		Pixel{X: 8, Y: 6, C: &ColorRed},
		Pixel{X: 9, Y: 6, C: &ColorRed},
		Pixel{X: 10, Y: 6, C: &ColorRed},
		Pixel{X: 11, Y: 6, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 0, Y: 5, C: &ColorBlack},
		Pixel{X: 1, Y: 5, C: &ColorRed},
		Pixel{X: 2, Y: 5, C: &ColorRed},
		Pixel{X: 3, Y: 5, C: &ColorRed},
		Pixel{X: 4, Y: 5, C: &ColorRed},
		Pixel{X: 5, Y: 5, C: &ColorRed},
		Pixel{X: 6, Y: 5, C: &ColorRed},
		Pixel{X: 7, Y: 5, C: &ColorRed},
		Pixel{X: 8, Y: 5, C: &ColorRed},
		Pixel{X: 9, Y: 5, C: &ColorRed},
		Pixel{X: 10, Y: 5, C: &ColorRed},
		Pixel{X: 11, Y: 5, C: &ColorRed},
		Pixel{X: 12, Y: 5, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 0, Y: 4, C: &ColorBlack},
		Pixel{X: 1, Y: 4, C: &ColorRed},
		Pixel{X: 2, Y: 4, C: &ColorWhite},
		Pixel{X: 3, Y: 4, C: &ColorRed},
		Pixel{X: 4, Y: 4, C: &ColorRed},
		Pixel{X: 5, Y: 4, C: &ColorRed},
		Pixel{X: 6, Y: 4, C: &ColorRed},
		Pixel{X: 7, Y: 4, C: &ColorRed},
		Pixel{X: 8, Y: 4, C: &ColorRed},
		Pixel{X: 9, Y: 4, C: &ColorRed},
		Pixel{X: 10, Y: 4, C: &ColorRed},
		Pixel{X: 11, Y: 4, C: &ColorRed},
		Pixel{X: 12, Y: 4, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 0, Y: 3, C: &ColorBlack},
		Pixel{X: 1, Y: 3, C: &ColorRed},
		Pixel{X: 2, Y: 3, C: &ColorWhite},
		Pixel{X: 3, Y: 3, C: &ColorRed},
		Pixel{X: 4, Y: 3, C: &ColorRed},
		Pixel{X: 5, Y: 3, C: &ColorRed},
		Pixel{X: 6, Y: 3, C: &ColorRed},
		Pixel{X: 7, Y: 3, C: &ColorRed},
		Pixel{X: 8, Y: 3, C: &ColorRed},
		Pixel{X: 9, Y: 3, C: &ColorRed},
		Pixel{X: 10, Y: 3, C: &ColorRed},
		Pixel{X: 11, Y: 3, C: &ColorRed},
		Pixel{X: 12, Y: 3, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 0, Y: 2, C: &ColorBlack},
		Pixel{X: 1, Y: 2, C: &ColorRed},
		Pixel{X: 2, Y: 2, C: &ColorRed},
		Pixel{X: 3, Y: 2, C: &ColorWhite},
		Pixel{X: 4, Y: 2, C: &ColorRed},
		Pixel{X: 5, Y: 2, C: &ColorRed},
		Pixel{X: 6, Y: 2, C: &ColorBlack},
		Pixel{X: 7, Y: 2, C: &ColorRed},
		Pixel{X: 8, Y: 2, C: &ColorRed},
		Pixel{X: 9, Y: 2, C: &ColorRed},
		Pixel{X: 10, Y: 2, C: &ColorRed},
		Pixel{X: 11, Y: 2, C: &ColorRed},
		Pixel{X: 12, Y: 2, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 1, Y: 1, C: &ColorBlack},
		Pixel{X: 2, Y: 1, C: &ColorRed},
		Pixel{X: 3, Y: 1, C: &ColorRed},
		Pixel{X: 4, Y: 1, C: &ColorRed},
		Pixel{X: 5, Y: 1, C: &ColorBlack},
		Pixel{X: 7, Y: 1, C: &ColorBlack},
		Pixel{X: 8, Y: 1, C: &ColorRed},
		Pixel{X: 9, Y: 1, C: &ColorRed},
		Pixel{X: 10, Y: 1, C: &ColorRed},
		Pixel{X: 11, Y: 1, C: &ColorBlack},
	)
	pixels = append(pixels,
		Pixel{X: 2, Y: 0, C: &ColorBlack},
		Pixel{X: 3, Y: 0, C: &ColorBlack},
		Pixel{X: 4, Y: 0, C: &ColorBlack},
		Pixel{X: 8, Y: 0, C: &ColorBlack},
		Pixel{X: 9, Y: 0, C: &ColorBlack},
		Pixel{X: 10, Y: 0, C: &ColorBlack},
	)
	return pixels
}
