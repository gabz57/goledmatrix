package components

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/controller"
	"image/color"
	"time"
)

type ControllableComponent interface {
	Component
	Controllable
}
type Controllable interface {
	Control(gamepad *controller.Gamepad)
}
type Component interface {
	// Controllable // watch for user/external events
	Updatable
	Drawable
}

type Updatable interface {
	// Update is called at its own rate to avoid depending on rendering time,
	// thus we can compute speed depending on this constant rate
	Update(elapsedBetweenUpdate time.Duration) bool
}

type Drawable interface {
	Draw(canvas Canvas) error
}

type Layout struct {
	color           color.Color
	backgroundColor color.Color
}

func (l *Layout) Color() *color.Color {
	return &l.color
}

func (l *Layout) BackgroundColor() *color.Color {
	if l.backgroundColor != nil {
		return &l.backgroundColor
	}
	return &l.color
}

type Graphic struct {
	layout *Layout
	parent *Graphic
	offset Point
	//enabled true
	//dirty   true // true for first rendering (static Graphics)
}

func (g *Graphic) Layout() *Layout {
	return g.layout
}

func (g *Graphic) ComputedOffset() Point {
	if g.parent != nil {
		return g.parent.ComputedOffset().Add(g.offset)
	}
	return g.offset
}

func (g *Graphic) SetOffset(offset Point) {
	g.offset = offset
}

type MaskedDrawable struct {
	drawable *Drawable
	mask     *Mask
}

func MaskDrawable(mask *Mask, drawable *Drawable) *Drawable {
	var maskedDrawable Drawable = &MaskedDrawable{
		drawable: drawable,
		mask:     mask,
	}
	return &maskedDrawable
}

// override Drawable.Draw method to perform indirection with mask
func (md MaskedDrawable) Draw(canvas Canvas) error {
	var c = wrapMask(md.mask, &canvas)
	return (*md.drawable).Draw(*c)
}

func wrapMask(mask *Mask, c *Canvas) *Canvas {
	var canvas Canvas = NewMaskAdapter(c, mask)
	return &canvas
}
