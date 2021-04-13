package goledmatrix

import (
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
)

// Canvas is a image.Image representation of a LED matrix, it implements
// image.Image interface and can be used with draw.Draw for example
type Canvas interface {
	draw.Image
	Set(x, y int, ledColor color.Color)
	DrawLabel(x, y int, label string, ledColor color.Color, face font.Face)
	Render() error
	Clear()
	Close() error
	position(x, y int) int
	GetLeds() *[]color.Color
}

type CanvasImpl struct {
	w, h   int
	matrix *Matrix
	leds   []color.Color
}

type Pixel struct {
	X, Y int
	C    *color.Color
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
func NewCanvas(config *MatrixConfig, m *Matrix) *Canvas {
	w, h := config.Geometry()
	var canvas Canvas
	canvas = &CanvasImpl{
		w:      w,
		h:      h,
		leds:   make([]color.Color, w*h),
		matrix: m,
	}
	return &canvas
}

func (c *CanvasImpl) register(matrix *Matrix) {
	c.matrix = matrix
	fmt.Println("Registered matrix !")
}

// ColorModel returns the canvas' color model, always color.RGBAModel
func (c *CanvasImpl) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds return the topology of the Canvas
func (c *CanvasImpl) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.w, c.h)
}

// At returns the color of the pixel at (x, y) and SHOULD NOT be directly used by dev, only through image.Image interface
func (c *CanvasImpl) At(x, y int) color.Color {
	return c.leds[c.position(x, y)]
}

// Set set LED at position x,y to the provided 24-bit color value
func (c *CanvasImpl) Set(x, y int, ledColor color.Color) {
	//c.leds[c.position(x, y)] = color.RGBAModel.Convert(ledColor)
	position := c.position(x, y)
	if x >= 0 && y >= 0 && position < len(c.leds) {
		c.leds[position] = color.RGBAModel.Convert(ledColor)
	}
}

func (c *CanvasImpl) SetPoint(point Point, ledColor color.Color) {
	if point.X >= 0 && point.Y >= 0 && c.position(point.X, point.Y) < c.w*c.h {
		c.leds[c.position(point.X, point.Y)] = ledColor
	}
}

func (c *CanvasImpl) DrawLabel(x, y int, label string, ledColor color.Color, face font.Face) {
	var canvas Canvas
	canvas = c
	d := &font.Drawer{
		Dst:  &textCanvas{canvas},
		Src:  image.NewUniform(ledColor),
		Face: face,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)},
	}
	d.DrawString(label)
}

type textCanvas struct {
	Canvas
}

func (tc *textCanvas) At(x, y int) color.Color {
	return color.Black
}

func (c *CanvasImpl) Render() error {
	var canvas Canvas
	canvas = c
	err := (*c.matrix).RenderMethod(&canvas)
	if err != nil {
		return err
	}
	return nil
}

// Clear set all the leds on the matrix to nil
func (c *CanvasImpl) Clear() {
	c.leds = make([]color.Color, c.w*c.h)
}

// Close clears the canvas and closes all the matrix
func (c *CanvasImpl) Close() error {
	c.Clear()
	err := c.Render()
	if err != nil {
		//return err
	}
	err = (*c.matrix).Close()
	if err != nil {
		return err
	}
	return err
}

func (c *CanvasImpl) position(x, y int) int {
	return x + (y * c.w)
}

func (c *CanvasImpl) GetLeds() *[]color.Color {
	return &c.leds
}
