package game

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/components"
)

type (
	GraphicEngine struct {
		canvas canvas.Canvas
	}
	GraphicComponent interface {
		Enable()
		Disable()
		Show()
		Hide()
		Draw(canvas canvas.Canvas) error
		IsEnabled() bool
		IsVisible() bool
	}
	GraphicComponentBase struct {
		values  *EntityValues
		enabled bool
		visible bool
		shape   components.CompositeDrawable
	}
)

func NewGraphicEngine(canvas canvas.Canvas) *GraphicEngine {
	return &GraphicEngine{
		canvas: canvas,
	}
}

func (ge GraphicEngine) RenderSceneAndSwapBuffers(buckets *[]EntityBucket) {
	ge.canvas.Clear()
	for _, bucket := range *buckets {
		for _, graphicComponent := range bucket.GraphicComponents {
			if graphicComponent.IsEnabled() && graphicComponent.IsVisible() {
				err := graphicComponent.Draw(ge.canvas)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	ge.canvas.Render()
}

func NewGraphicComponentBase(values *EntityValues, graphic *components.Graphic, enabled, visible bool) *GraphicComponentBase {
	return &GraphicComponentBase{
		values:  values,
		enabled: enabled,
		visible: visible,
		shape:   *components.NewCompositeDrawable(graphic),
	}
}

func (gcb *GraphicComponentBase) Enable()         { gcb.enabled = true }
func (gcb *GraphicComponentBase) Disable()        { gcb.enabled = false }
func (gcb *GraphicComponentBase) Show()           { gcb.visible = true }
func (gcb *GraphicComponentBase) Hide()           { gcb.visible = false }
func (gcb *GraphicComponentBase) IsEnabled() bool { return gcb.enabled }
func (gcb *GraphicComponentBase) IsVisible() bool { return gcb.visible }

func (gcb *GraphicComponentBase) RegisterDrawable(drawable components.Drawable) {
	gcb.shape.AddDrawable(drawable)
}
func (gcb *GraphicComponentBase) Draw(canvas canvas.Canvas) error {
	return gcb.shape.Draw(canvas)
}

func (gcb *GraphicComponentBase) GetValue(ref ValueRef) interface{} {
	return gcb.values.Get(ref)
}

func (gcb *GraphicComponentBase) SetValue(ref ValueRef, value interface{}) {
	gcb.values.Set(ref, value)
}

func (gcb *GraphicComponentBase) EnableDrawable(drawable components.Drawable) {
	for _, disablableDrawable := range gcb.shape.Drawables {
		if disablableDrawable.Drawable == drawable {
			disablableDrawable.Enabled = true
			return
		}
	}
}
func (gcb *GraphicComponentBase) DisableDrawable(drawable components.Drawable) {
	for _, disablableDrawable := range gcb.shape.Drawables {
		if disablableDrawable.Drawable == drawable {
			disablableDrawable.Enabled = false
			return
		}
	}
}
