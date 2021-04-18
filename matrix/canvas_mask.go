package matrix

import (
	"image"
	"image/color"
	"math"
)

type (
	// Canvas wrapper to replace pixels color with colors from the colors
	Mask interface {
		Canvas
		Set(x, y int, ledColor color.Color)
	}
	StaticCanvasMask struct {
		Mask
		colors []color.Color
	}
	ShadedColorCanvasMask struct {
		Mask
	}
	ColorFaderCanvasMask struct {
		Mask
		fade float64
	}
)

func (c *StaticCanvasMask) Set(x, y int, ledColor color.Color) {
	position := c.position(x, y)
	if x >= 0 && y >= 0 && position < len(*c.GetLeds()) {
		(*c.GetLeds())[position] = c.colors[position]
	}
}

// StaticCanvasMask specialization with a single color
func NewSingleColorMask(canvas Canvas, maskColor color.Color) *StaticCanvasMask {
	max := canvas.Bounds().Max
	var mask = make([]color.Color, max.X*max.Y)
	for col := 0; col < max.X; col++ {
		for row := 0; row < max.Y; row++ {
			mask[canvas.position(col, row)] = maskColor
		}
	}
	return NewCanvasMask(canvas, mask)
}

func NewCanvasMask(canvas Canvas, colors []color.Color) *StaticCanvasMask {
	return &StaticCanvasMask{
		Mask:   canvas,
		colors: colors,
	}
}

func NewShadedColorCanvasMask(canvas Canvas) *ShadedColorCanvasMask {
	return &ShadedColorCanvasMask{
		Mask: canvas,
	}
}

func (c *ShadedColorCanvasMask) Set(x, y int, ledColor color.Color) {
	position := c.position(x, y)
	if x >= 0 && y >= 0 && position < len(*c.GetLeds()) {
		center := Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		}
		(*c.GetLeds())[position] = shadedAroundCenterColor(c.Bounds().Max, center, x, y)
	}
}

func radToDeg(r float64) float64 {
	return (r * 180) / math.Pi
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

func NewColorFaderMask(canvas Canvas) *ColorFaderCanvasMask {
	return &ColorFaderCanvasMask{
		Mask: canvas,
		fade: 0,
	}
}

func (c *ColorFaderCanvasMask) SetFade(fade float64) {
	if fade >= 1 {
		c.fade = 1
	} else if fade <= 0 {
		c.fade = 0
	} else {
		c.fade = fade
	}
}

func (c *ColorFaderCanvasMask) Set(x, y int, ledColor color.Color) {
	r, g, b, a := ledColor.RGBA()
	c.Mask.Set(x, y, color.RGBA{
		R: uint8((1 - c.fade) * float64(uint8(r))),
		G: uint8((1 - c.fade) * float64(uint8(g))),
		B: uint8((1 - c.fade) * float64(uint8(b))),
		A: uint8(a),
	})
}
