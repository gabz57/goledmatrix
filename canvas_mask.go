package goledmatrix

import (
	"image/color"
	"math"
)

// Canvas wrapper to replace pixels color with colors from the mask
type CanvasMask struct {
	Canvas
	mask []color.Color
}

// Set set LED at position x,y to the provided 24-bit color value
func (c *CanvasMask) Set(x, y int, ledColor color.Color) {
	position := c.position(x, y)
	if x >= 0 && y >= 0 && position < len(*c.getLeds()) {
		(*c.getLeds())[position] = c.mask[position]
	}
}

func NewSingleColorMask(canvas Canvas, maskColor color.Color) *CanvasMask {
	max := canvas.Bounds().Max
	var mask = make([]color.Color, max.X*max.Y)
	for col := 0; col < max.X; col++ {
		for row := 0; row < max.Y; row++ {
			mask[canvas.position(col, row)] = maskColor
		}
	}
	return NewMask(canvas, mask)
}

type ShadedColorCanvasMask struct {
	Canvas
}

func NewShadedColorCanvasMask(canvas Canvas) *ShadedColorCanvasMask {
	return &ShadedColorCanvasMask{
		Canvas: canvas,
	}
}

// Set set LED at position x,y to the provided 24-bit color value
func (c *ShadedColorCanvasMask) Set(x, y int, ledColor color.Color) {
	position := c.position(x, y)
	center := Point{
		X: c.Bounds().Max.X / 2,
		Y: c.Bounds().Max.Y / 2,
	}

	if x >= 0 && y >= 0 && position < len(*c.getLeds()) {
		(*c.getLeds())[position] = shadedAroundCenterColor(center, x, y)
	}
}

func radToDeg(r float64) float64 {
	return (r * 180) / math.Pi
}

func shadedAroundCenterColor(center Point, x int, y int) color.Color {
	var saturation, value float64
	hue := radToDeg(math.Atan2(float64(y-center.Y), float64(x-center.X))) + 180 // angle
	saturation = 1
	value = 1

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
	//// Change r,g,b values from [0,1] to [0,255]
	//return [255*r,255*g,255*b];
}

func NewMask(canvas Canvas, mask []color.Color) *CanvasMask {
	return &CanvasMask{
		Canvas: canvas,
		mask:   mask,
	}
}
