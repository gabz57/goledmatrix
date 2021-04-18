package canvas

import (
	"golang.org/x/image/font"
	"image"
	"image/color"
)

// Canvas is a image.Image representation of a LED matrix, it implements
// image.Image interface and can be used with draw.Draw for example
type Canvas interface {
	image.Image
	Set(x, y int, ledColor color.Color)
	DrawLabel(x, y int, label string, ledColor color.Color, face font.Face)
	Render() error
	Clear()
	Close() error
	GetLeds() *[]color.Color
}

type Pixel struct {
	X, Y int
	C    *color.Color
}

type Point image.Point

type FloatingPoint struct {
	X, Y float64
}

func (p *Point) Floating() FloatingPoint {
	return FloatingPoint{
		X: float64(p.X),
		Y: float64(p.Y),
	}
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) AddXY(x, y int) Point {
	return Point{
		X: p.X + x,
		Y: p.Y + y,
	}
}

func (fp FloatingPoint) Add(other FloatingPoint) FloatingPoint {
	return FloatingPoint{
		X: fp.X + other.X,
		Y: fp.Y + other.Y,
	}
}

func (fp FloatingPoint) AddXY(x, y float64) FloatingPoint {
	return FloatingPoint{
		X: fp.X + x,
		Y: fp.Y + y,
	}
}

func (fp FloatingPoint) Int() Point {
	return Point{
		X: int(fp.X),
		Y: int(fp.Y),
	}
}

type TextCanvas struct {
	Canvas
}

func (tc *TextCanvas) At(x, y int) color.Color {
	return color.Black
}