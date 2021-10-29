package components

import "github.com/gabz57/goledmatrix/canvas"

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

func (cd *CompositeDrawable) Draw(canvas canvas.Canvas) error {
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
