package goledmatrix

import (
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
)

// Canvas is a image.Image representation of a LED matrix, it implements
// image.Image interface and can be used with draw.Draw for example
type Canvas struct {
	w, h     int
	matrices []Matrix
	leds     []color.Color
}

type Point struct {
	X, Y int
}

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
func NewCanvas(config *MatrixConfig) *Canvas {
	w, h := config.Geometry()
	c := Canvas{
		w:    w,
		h:    h,
		leds: make([]color.Color, w*h),
	}
	//draw.Draw(&c, c.Bounds(), &image.Uniform{C: color.Black}, image.Point{}, draw.Src)
	return &c
}

func (c *Canvas) register(matrix Matrix) {
	c.matrices = append(c.matrices, matrix)
	fmt.Println("Registered matrix !")
}

// ColorModel returns the canvas' color model, always color.RGBAModel
func (c *Canvas) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds return the topology of the Canvas
func (c *Canvas) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.w, c.h)
}

// At returns the color of the pixel at (x, y) and SHOULD NOT be directly used by dev, only through image.Image interface
func (c *Canvas) At(x, y int) color.Color {
	return c.leds[c.position(x, y)]
}

// Set set LED at position x,y to the provided 24-bit color value
func (c *Canvas) Set(x, y int, ledColor color.Color) {
	//c.leds[c.position(x, y)] = color.RGBAModel.Convert(ledColor)
	if x >= 0 && y >= 0 && c.position(x, y) < len(c.leds) {
		c.leds[c.position(x, y)] = ledColor
	}
}

func (c *Canvas) SetPoint(point Point, ledColor color.Color) {
	if point.X >= 0 && point.Y >= 0 && c.position(point.X, point.Y) < c.w*c.h {
		c.leds[c.position(point.X, point.Y)] = ledColor
	}
}

func (c *Canvas) DrawLabel(x, y int, label string, ledColor color.Color, face font.Face) {
	d := &font.Drawer{
		Dst:  &TextCanvas{c},
		Src:  image.NewUniform(ledColor),
		Face: face,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)},
	}
	d.DrawString(label)
}

type TextCanvas struct {
	*Canvas
}

func (tc *TextCanvas) At(x, y int) color.Color {
	return color.Black
}

func (c *Canvas) Render() error {
	for _, m := range c.matrices {
		err := m.RenderMethod(c)
		if err != nil {
			return err
		}
	}
	return nil
}

// Clear set all the leds on the matrix to nil
func (c *Canvas) Clear() {
	c.leds = make([]color.Color, c.w*c.h)
}

// Close clears the canvas and closes all the matrices
func (c *Canvas) Close() error {
	c.Clear()
	err := c.Render()
	if err != nil {
		//return err
	}
	for _, m := range c.matrices {
		err = m.Close()
		if err != nil {
			return err
		}
	}
	return err
}

func (c *Canvas) position(x, y int) int {
	return x + (y * c.w)
}

// NOTE: direct access (RPC Client) !
func (c *Canvas) Leds() []color.Color {
	return c.leds
}
