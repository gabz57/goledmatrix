package components

import (
	"fmt"
	. "github.com/gabz57/goledmatrix"
	"image/color"
	"math"
	"time"
)

var Origin = Point{}
var None = Origin

var defaultLayout = Layout{
	color:           color.White,
	backgroundColor: nil,
}

var ColorBlue = color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}
var ColorGreen = color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}
var ColorViolet = color.RGBA{R: 0xff, G: 0x00, B: 0xff, A: 0xff}
var ColorRed = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}

func NewGraphic(parent *Graphic, layout *Layout) Graphic {
	if layout == nil {
		if parent != nil {
			layout = parent.layout
		} else {
			layout = &defaultLayout
		}
	}
	return Graphic{
		layout: layout,
		parent: parent,
		offset: None,
	}
}

func NewLayout(fColor, bgColor color.Color) *Layout {
	return &Layout{
		color:           fColor,
		backgroundColor: bgColor,
	}
}

func Rotate(p, o Point, degrees float64) Point {
	rad := degToRad(degrees)
	fp := p.Floating()
	fo := o.Floating()
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return Point{
		X: int(cos*(fp.X-fo.X) - sin*(fp.Y-fo.Y) + fo.X),
		Y: int(sin*(fp.X-fo.X) + cos*(fp.Y-fo.Y) + fo.Y),
	}
}

func degToRad(x float64) float64 {
	return (x / 180) * math.Pi
}

func TimeToText(now time.Time) string {
	hour, min, sec := now.Clock()
	millis := now.Nanosecond() / 1000000
	return fmt.Sprintf("%02d", hour) + ":" + fmt.Sprintf("%02d", min) + ":" + fmt.Sprintf("%02d", sec) + "." + fmt.Sprintf("%03d", millis)
}
