package masks

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"image"
	"image/color"
	"math"
)

type (
	ShadedColorCanvasMask struct {
		bounds image.Rectangle
		center Point
	}
)

// used in clock.go to color the hours and minutes
func NewShadedColorCanvasMask(bounds image.Rectangle) *ShadedColorCanvasMask {
	return &ShadedColorCanvasMask{
		bounds: bounds,
		center: Point{
			X: bounds.Max.X / 2,
			Y: bounds.Max.Y / 2,
		},
	}
}
func (m *ShadedColorCanvasMask) AdaptPixel() func(canvas Canvas, x int, y int, color color.Color) {
	return func(canvas Canvas, x int, y int, color color.Color) {
		if (image.Point{X: x, Y: y}).In(m.bounds) {
			canvas.Set(x, y, shadedAroundCenterColor(m.bounds.Max, m.center, x, y))
		}
	}
}

func shadedAroundCenterColor(max image.Point, center Point, x int, y int) color.Color {
	var saturation, value float64
	posY := y - center.Y
	posX := x - center.X
	hue := radToDeg(math.Atan2(float64(posY), float64(posX))) + 180 // angle
	saturation = 1
	value = math.Sqrt(float64(posY*posY)+float64(posX*posX)) / math.Sqrt(float64(max.X*max.X+max.Y*max.Y)/4)

	chroma := value * saturation
	hue1 := hue / 60
	x1 := chroma * (1 - math.Abs(math.Mod(hue1, 2)-1))

	var r1, g1, b1 float64
	if hue1 >= 0 && hue1 <= 1 {
		r1 = chroma
		g1 = x1
		b1 = 0
	} else if hue1 >= 1 && hue1 <= 2 {
		r1 = x1
		g1 = chroma
		b1 = 0
	} else if hue1 >= 2 && hue1 <= 3 {
		r1 = 0
		g1 = chroma
		b1 = x1
	} else if hue1 >= 3 && hue1 <= 4 {
		r1 = 0
		g1 = x1
		b1 = chroma
	} else if hue1 >= 4 && hue1 <= 5 {
		r1 = x1
		g1 = 0
		b1 = chroma
	} else if hue1 >= 5 && hue1 <= 6 {
		r1 = chroma
		g1 = 0
		b1 = x1
	}

	m := value - chroma
	r := r1 + m
	g := g1 + m
	b := b1 + m

	return color.RGBA{
		R: uint8(255 * r),
		G: uint8(255 * g),
		B: uint8(255 * b),
		A: uint8(255),
	}
}

func radToDeg(r float64) float64 {
	return (r * 180) / math.Pi
}
