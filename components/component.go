package components

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/canvas/effect"
	"image/color"
	"time"
)

type Drawable interface {
	Draw(canvas Canvas) error
}

type Updatable interface {
	// Update is called at its own rate to avoid depending on rendering time,
	// thus we can compute speed depending on this constant rate
	Update(elapsedBetweenUpdate time.Duration) bool
}

type Component interface {
	Updatable
	Drawable
}

type Layout struct {
	color           color.Color
	backgroundColor color.Color
}

func (l *Layout) Color() color.Color {
	return l.color
}
func (l *Layout) SetColor(color color.Color) {
	l.color = color
}

func (l *Layout) BackgroundColor() color.Color {
	if l.backgroundColor != nil {
		return l.backgroundColor
	}
	return l.color
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
	drawable Drawable
	effect   effect.Effect
}

func MaskDrawable(mask effect.Effect, drawable Drawable) Drawable {
	return &MaskedDrawable{
		drawable: drawable,
		effect:   mask,
	}
}

// override Drawable.Draw method to perform indirection with effect
func (md MaskedDrawable) Draw(canvas Canvas) error {
	return md.drawable.Draw(effect.NewAdapter(canvas, md.effect))
}
