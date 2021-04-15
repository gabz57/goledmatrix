package components

import (
	"fmt"
	. "github.com/gabz57/goledmatrix"
	"image/color"
	"math"
	"math/rand"
	"time"
)

var Origin = Point{}
var None = Origin

var ColorWhite color.Color = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
var ColorBlack color.Color = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
var ColorRed color.Color = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
var ColorGreen color.Color = color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}
var ColorBlue color.Color = color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}
var ColorViolet color.Color = color.RGBA{R: 0xff, G: 0x00, B: 0xff, A: 0xff}
var ColorYellow color.Color = color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}

func NewGraphic(parent *Graphic, layout *Layout) *Graphic {
	return NewOffsetGraphic(parent, layout, None)
}

func NewOffsetGraphic(parent *Graphic, layout *Layout, offset Point) *Graphic {
	if layout == nil {
		if parent != nil {
			layout = parent.layout
		} else {
			layout = DefaultLayout()
		}
	}
	return &Graphic{
		layout: layout,
		parent: parent,
		offset: offset,
	}
}

func DefaultLayout() *Layout {
	return &Layout{
		color:           color.White,
		backgroundColor: nil,
	}
}

func NewLayout(fColor, bgColor color.Color) *Layout {
	return &Layout{
		color:           fColor,
		backgroundColor: bgColor,
	}
}

func TimeToText(t time.Time) string {
	hour, min, sec := t.Clock()
	millis := t.Nanosecond() / 1000000
	return fmt.Sprintf("%02d", hour) + ":" + fmt.Sprintf("%02d", min) + ":" + fmt.Sprintf("%02d", sec) + "." + fmt.Sprintf("%03d", millis)
}

func Rotate(p, o Point, degrees float64) Point {
	rad := DegToRad(degrees)
	fp := p.Floating()
	fo := o.Floating()
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return Point{
		X: int(cos*(fp.X-fo.X) - sin*(fp.Y-fo.Y) + fo.X),
		Y: int(sin*(fp.X-fo.X) + cos*(fp.Y-fo.Y) + fo.Y),
	}
}

func RotateOrigin(fp FloatingPoint, degrees float64) FloatingPoint {
	rad := DegToRad(degrees)
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return FloatingPoint{
		X: cos*(fp.X) - sin*(fp.Y),
		Y: sin*(fp.X) + cos*(fp.Y),
	}
}

func DegToRad(x float64) float64 {
	return (x / 180) * math.Pi
}

func RadToDeg(x float64) float64 {
	return (x * 180) / math.Pi
}

// Angle (degrees) from cartesian coordinates
func FloatingPointToDirection(point FloatingPoint) float64 {
	return RadToDeg(math.Atan(point.Y / point.X))
}

// Cartesian coordinates from angle (degrees)
func DirectionToFloatingPoint(direction float64) FloatingPoint {
	rad := DegToRad(direction)
	return FloatingPoint{
		X: math.Cos(rad),
		Y: math.Sin(rad),
	}
}

func Int64Between(low, high int64) int64 {
	var gen int64
	for gen < low {
		gen = rand.Int63n(high)
	}
	return gen
}

func Float64Between(low, high float64) float64 {
	var gen float64
	for gen <= low {
		gen = rand.Float64() * high
	}
	return gen
}

func OneOrMinusOne() float64 {
	if rand.Intn(1) == 0 {
		return -1
	}
	return 1
}
