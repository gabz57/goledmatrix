package main

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"image/color"
	"time"
)

var darkBlue = color.RGBA{
	R: uint8(17),
	G: uint8(34),
	B: uint8(80),
	A: uint8(255),
}

var lightBlue = color.RGBA{
	R: uint8(73),
	G: uint8(163),
	B: uint8(190),
	A: uint8(255),
}

type OctoLogo struct {
	shape                           CompositeDrawable
	center                          Point
	radiusExt, radiusInt, radiusEye int
}

var octoLogoGraphic = NewGraphic(nil, nil)

func NewOctoLogo(canvas Canvas, center Point, radiusExt int) *OctoLogo {
	radiusInt := int(float64(radiusExt) * 2 / 3)
	logo := OctoLogo{
		shape: CompositeDrawable{
			Graphic:   &infoGraphic,
			Drawables: []*Drawable{},
		},
		center:    center,
		radiusExt: radiusExt,
		radiusInt: radiusInt,
		radiusEye: int(float64(radiusExt-radiusInt) * 1 / 2),
	}
	logo.shape.AddDrawable(logo.buildRing())
	logo.shape.AddDrawable(logo.buildEyeIn())
	logo.shape.AddDrawable(logo.buildEyeOut())

	return &logo
}
func (o *OctoLogo) Update(elapsedBetweenUpdate time.Duration) {
}

func (o *OctoLogo) Draw(canvas Canvas) error {
	return o.shape.Draw(canvas)
}

func (o *OctoLogo) buildRing() *Drawable {
	graphic := NewGraphic(o.shape.Graphic, NewLayout(darkBlue, nil))
	var ring Drawable
	ring = shapes.NewRing(&graphic, o.center, o.radiusExt, o.radiusInt, true)
	return &ring
}

func (o *OctoLogo) buildEyeIn() *Drawable {
	graphic := NewGraphic(o.shape.Graphic, NewLayout(lightBlue, nil))
	var eye Drawable
	eye = shapes.NewCircle(&graphic, o.center.AddXY(int(float64(o.radiusExt)/3), -int(float64(o.radiusExt)/3)), o.radiusEye, true)
	return &eye
}

func (o *OctoLogo) buildEyeOut() *Drawable {
	graphic := NewGraphic(o.shape.Graphic, NewLayout(lightBlue, nil))
	var eye Drawable
	eye = shapes.NewCircle(&graphic, o.center.AddXY(o.radiusExt-o.radiusEye, -(o.radiusExt-o.radiusEye)), o.radiusEye, true)
	return &eye
}
