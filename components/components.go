package components

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"image/color"
	"time"
)

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

type CompositeDrawable struct {
	Graphic   *Graphic
	Drawables []*Drawable
}

func NewCompositeDrawable(graphic *Graphic) *CompositeDrawable {
	return &CompositeDrawable{
		Graphic:   graphic,
		Drawables: []*Drawable{},
	}
}

func (cd *CompositeDrawable) AddDrawable(drawables ...*Drawable) {
	for _, drawable := range drawables {
		cd.Drawables = append(cd.Drawables, drawable)
	}
}

func (cd *CompositeDrawable) Draw(canvas Canvas) error {
	var err error
	for _, drawable := range cd.Drawables {
		// FIXME: apply cd.Graphic.ComputedOffset()
		err = (*drawable).Draw(canvas)
		if err != nil {
			return err
		}
	}
	return nil
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

func Masked(mask *Canvas, drawable *Drawable) *Drawable {
	var d Drawable
	d = &MaskedDrawable{
		drawable: drawable,
		mask:     mask,
	}
	return &d
}

type MaskedDrawable struct {
	drawable *Drawable
	mask     *Canvas
}

// override Drawable.Draw method to perform indirection with mask
func (m MaskedDrawable) Draw(canvas Canvas) error {
	return (*m.drawable).Draw(*m.mask)
}
