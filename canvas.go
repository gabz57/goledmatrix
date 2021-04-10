package goledmatrix

import (
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type RecursiveMutex struct {
	internalMutex    sync.Mutex
	currentGoRoutine int64  // keeps track of the current goroutine id
	lockCount        uint64 // lock count on the current goroutine
}

func goid() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return int64(id)
}

func (rm *RecursiveMutex) Lock() {
	// get the current goroutine id
	goRoutineID := goid()
	for {
		rm.internalMutex.Lock()
		if rm.currentGoRoutine == 0 {
			// no locks yet
			rm.currentGoRoutine = goRoutineID
			break
		} else if rm.currentGoRoutine == goRoutineID {
			// lock from the same go routine
			break
		} else {
			// lock from a different go routine, need to wait
			// until lock is released
			rm.internalMutex.Unlock()
			time.Sleep(time.Millisecond)
			continue
		}
	}
	// increase the lock count
	rm.lockCount++
	rm.internalMutex.Unlock()
}

func (rm *RecursiveMutex) Unlock() {
	rm.internalMutex.Lock()
	rm.lockCount--
	if rm.lockCount == 0 {
		rm.currentGoRoutine = 0
	}
	rm.internalMutex.Unlock()
}

// Canvas is a image.Image representation of a LED matrix, it implements
// image.Image interface and can be used with draw.Draw for example
type Canvas struct {
	w, h     int
	matrices []Matrix
	mutex    RecursiveMutex
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
	c.mutex.Lock()
	defer c.mutex.Unlock()
	//c.leds[c.position(x, y)] = color.RGBAModel.Convert(ledColor)
	if x >= 0 && y >= 0 && c.position(x, y) < c.w*c.h {
		c.leds[c.position(x, y)] = ledColor
	}
}

func (c *Canvas) SetPoint(point Point, ledColor color.Color) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if point.X >= 0 && point.Y >= 0 && c.position(point.X, point.Y) < c.w*c.h {
		c.leds[c.position(point.X, point.Y)] = ledColor
	}
}

func (c *Canvas) DrawLabel(x, y int, label string, ledColor color.Color, face font.Face) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
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
	//colorAt := tc.Canvas.At(x, y)
	//if colorAt == nil {
	//	colorAt = color.Black
	//}
	//return colorAt
}

func (c *Canvas) Render() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for _, m := range c.matrices {
		err := m.RenderMethod(c)
		if err != nil {
			return err
		}
	}
	return nil
}

// Clear set all the leds on the matrix with color.Black
func (c *Canvas) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.leds = nil
	c.leds = make([]color.Color, c.w*c.h)
	//draw.Draw(c, c.Bounds(), &image.Uniform{C: color.Black}, image.Point{}, draw.Src)
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

// TODO: fix design to avoid lock exposure
// added to avoid writes when drawing (emulator draw partial canvas)
func (c *Canvas) Lock() {
	c.mutex.Lock()
}

func (c *Canvas) Unlock() {
	c.mutex.Unlock()
}
