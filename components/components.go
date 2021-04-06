package components

import (
	. "github.com/gabz57/goledmatrix"
	"image/color"
)

type Component interface {
	// processInput
	Updatable
	Drawable
}

type Drawable interface {
	Draw(canvas *Canvas) error
}

type Updatable interface {
	Update()
}

type CompositeDrawable struct {
	Graphic   *Graphic
	Drawables []Drawable
}

func (cd *CompositeDrawable) AddDrawable(drawable Drawable) {
	cd.Drawables = append(cd.Drawables, drawable)
}

func (cd *CompositeDrawable) AddDrawables(drawables []Drawable) {
	for _, drawable := range drawables {
		cd.Drawables = append(cd.Drawables, drawable)
	}
}

func (cd *CompositeDrawable) Draw(canvas *Canvas) error {
	var err error
	for _, drawable := range cd.Drawables {
		err = drawable.Draw(canvas)
		if err != nil {
			return err
		}
	}
	return nil
}

type Layout struct {
	color           color.Color
	backgroundColor color.Color
	//_transformers: PixelTransformer[] = [],
	//_reverseY = false,
	//_pixelEffect?: (_: Pixel) => Pixel[]
}

func (l *Layout) Color() color.Color {
	return l.color
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

func (g *Graphic) setParent(parent *Graphic) {
	g.parent = parent
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
