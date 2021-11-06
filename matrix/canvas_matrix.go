package matrix

import (
	"fmt"
	. "github.com/gabz57/goledmatrix/canvas"
	"image"
	"image/color"
)

type CanvasImpl struct {
	w, h   int
	matrix *Matrix
	leds   []color.Color
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

func NewSimpleCanvas(x, y int, leds *[]color.Color) Canvas {
	var canvas Canvas = &CanvasImpl{
		w:    x,
		h:    y,
		leds: *leds,
	}
	return canvas
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
	return c.leds[Position(x, y, c.w)]
}

func (c *CanvasImpl) GetLeds() *[]color.Color {
	return &c.leds
}

// Set set LED at position x,y to the provided 24-bit color value
func (c *CanvasImpl) Set(x, y int, ledColor color.Color) {
	//c.leds[c.position(x, y)] = color.RGBAModel.Convert(ledColor)
	point := image.Point{X: x, Y: y}
	//if x >= 0 && x < 0 && y >= 0 && position < len(c.leds) {
	if point.In(c.Bounds()) {
		c.leds[Position(x, y, c.w)] = color.RGBAModel.Convert(ledColor)
	}
}

func (c *CanvasImpl) SetPoint(point Point, ledColor color.Color) {
	if point.X >= 0 && point.Y >= 0 && Position(point.X, point.Y, c.w) < c.w*c.h {
		c.leds[Position(point.X, point.Y, c.w)] = ledColor
	}
}

// Clear set all the leds on the matrix to nil
func (c *CanvasImpl) Clear() {
	c.leds = make([]color.Color, c.w*c.h)
}

func (c *CanvasImpl) Render() error {
	var canvas Canvas
	canvas = c
	err := (*c.matrix).RenderMethod(canvas)
	if err != nil {
		return err
	}
	return nil
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
