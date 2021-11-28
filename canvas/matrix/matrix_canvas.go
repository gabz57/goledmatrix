package matrix

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"image"
	"image/color"
)

type matrixCanvas struct {
	w, h   int
	matrix Matrix
	leds   []color.Color
}

func NewCanvas(config *MatrixConfig, m Matrix) Canvas {
	w, h := config.Geometry()
	return &matrixCanvas{
		w:      w,
		h:      h,
		leds:   make([]color.Color, w*h),
		matrix: m,
	}
}

func NewSimpleCanvas(x, y int, leds *[]color.Color) Canvas {
	return &matrixCanvas{
		w:    x,
		h:    y,
		leds: *leds,
	}
}

// ColorModel returns the canvas' color model, always color.RGBAModel
func (c *matrixCanvas) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds return the topology of the Canvas
func (c *matrixCanvas) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.w, c.h)
}

// At returns the color of the pixel at (x, y) and SHOULD NOT be directly used by dev, only through image.Image interface
func (c *matrixCanvas) At(x, y int) color.Color {
	return c.leds[Position(x, y, c.w)]
}

func (c *matrixCanvas) GetLeds() *[]color.Color {
	return &c.leds
}

// Set set LED at position x,y to the provided 24-bit color value
func (c *matrixCanvas) Set(x, y int, ledColor color.Color) {
	//c.leds[c.position(x, y)] = color.RGBAModel.Convert(ledColor)
	point := image.Point{X: x, Y: y}
	//if x >= 0 && x < 0 && y >= 0 && position < len(c.leds) {
	if point.In(c.Bounds()) {
		c.leds[Position(x, y, c.w)] = color.RGBAModel.Convert(ledColor)
	}
}

func (c *matrixCanvas) SetPoint(point Point, ledColor color.Color) {
	if point.X >= 0 && point.Y >= 0 && Position(point.X, point.Y, c.w) < c.w*c.h {
		c.leds[Position(point.X, point.Y, c.w)] = ledColor
	}
}

// Clear set all the leds on the matrix to nil
func (c *matrixCanvas) Clear() {
	c.leds = make([]color.Color, c.w*c.h)
}

func (c *matrixCanvas) Render() error {
	return c.matrix.RenderMethod(c)
}

// Close clears the canvas and closes the matrix
func (c *matrixCanvas) Close() error {
	c.Clear()
	err := c.Render()
	if err != nil {
		return err
	}
	err = c.matrix.Close()
	if err != nil {
		return err
	}
	return err
}
