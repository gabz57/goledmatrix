package components

import (
	"github.com/gabz57/goledmatrix/canvas"
)

type DisablableDrawable struct {
	Drawable
	Enabled bool
}

func NewDrawable(drawable Drawable) *DisablableDrawable {
	return &DisablableDrawable{
		Drawable: drawable,
		Enabled:  true,
	}
}

type CompositeDrawable struct {
	Graphic   *Graphic
	Drawables []*DisablableDrawable
}

func NewCompositeDrawable(graphic *Graphic) *CompositeDrawable {
	return &CompositeDrawable{
		Graphic:   graphic,
		Drawables: []*DisablableDrawable{},
	}
}

func (cd *CompositeDrawable) AddDrawable(drawable Drawable) {
	cd.Drawables = append(cd.Drawables, NewDrawable(drawable))
}

func (cd *CompositeDrawable) AddDrawables(drawables ...Drawable) {
	for _, drawable := range drawables {
		cd.Drawables = append(cd.Drawables, NewDrawable(drawable))
	}
}

func (cd *CompositeDrawable) RemoveDrawable(d Drawable) {
	for index, drawable := range cd.Drawables {
		if drawable == d {
			cd.Drawables = append(cd.Drawables[:index], cd.Drawables[index+1:]...)
			return
		}
	}
}

func (cd *CompositeDrawable) Draw(canvas canvas.Canvas) error {
	var err error
	for _, drawable := range cd.Drawables {
		if drawable.Enabled {
			err = drawable.Draw(canvas)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
