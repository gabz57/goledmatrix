package components

import (
	"github.com/gabz57/goledmatrix/canvas"
)

type disablableDrawable struct {
	Drawable
	Enabled bool
}

type CompositeDrawable struct {
	Graphic   *Graphic
	Drawables []*disablableDrawable
}

func NewCompositeDrawable(graphic *Graphic) *CompositeDrawable {
	return &CompositeDrawable{
		Graphic:   graphic,
		Drawables: []*disablableDrawable{},
	}
}

func (cd *CompositeDrawable) AddDrawable(drawable Drawable) {
	cd.Drawables = append(cd.Drawables, &disablableDrawable{
		Drawable: drawable,
		Enabled:  true,
	})
}

func (cd *CompositeDrawable) AddDrawables(drawables ...Drawable) {
	for _, drawable := range drawables {
		cd.AddDrawable(drawable)
	}
}

func (cd *CompositeDrawable) RemoveDrawable(d Drawable) {
	for index, drawable := range cd.Drawables {
		if drawable.Drawable == d {
			cd.Drawables = append(cd.Drawables[:index], cd.Drawables[index+1:]...)
			return
		}
	}
}

func (cd *CompositeDrawable) EnableDrawable(drawable Drawable) {
	for _, disablable := range cd.Drawables {
		if disablable.Drawable == drawable {
			disablable.Enabled = true
			return
		}
	}
}
func (cd *CompositeDrawable) DisableDrawable(drawable Drawable) {
	for _, disablable := range cd.Drawables {
		if disablable.Drawable == drawable {
			disablable.Enabled = false
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
