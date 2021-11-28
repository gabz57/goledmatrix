package engine

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

func (e GraphicEngine) RenderSceneAndSwapBuffers(buckets *[]EntityBucket) {
	e.canvas.Clear()
	for _, bucket := range *buckets {
		for _, graphicComponent := range bucket.GraphicComponents {
			if graphicComponent.IsEnabled() && graphicComponent.IsVisible() {
				err := graphicComponent.Draw(e.canvas)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	e.canvas.Render()
}

func NewGraphicComponentBase(values *EntityValues, graphic *components.Graphic, enabled, visible bool) *GraphicComponentBase {
	return &GraphicComponentBase{
		values:  values,
		enabled: enabled,
		visible: visible,
		shape:   *components.NewCompositeDrawable(graphic),
	}
}

func (b *GraphicComponentBase) Enable()         { b.enabled = true }
func (b *GraphicComponentBase) Disable()        { b.enabled = false }
func (b *GraphicComponentBase) Show()           { b.visible = true }
func (b *GraphicComponentBase) Hide()           { b.visible = false }
func (b *GraphicComponentBase) IsEnabled() bool { return b.enabled }
func (b *GraphicComponentBase) IsVisible() bool { return b.visible }

func (b *GraphicComponentBase) RegisterDrawable(drawable components.Drawable) {
	b.shape.AddDrawable(drawable)
}
func (b *GraphicComponentBase) Draw(canvas canvas.Canvas) error {
	return b.shape.Draw(canvas)
}

func (b *GraphicComponentBase) GetValue(ref ValueRef) interface{} {
	return b.values.Get(ref)
}

func (b *GraphicComponentBase) SetValue(ref ValueRef, value interface{}) {
	b.values.Set(ref, value)
}

func (b *GraphicComponentBase) EnableDrawable(drawable components.Drawable) {
	b.shape.EnableDrawable(drawable)
}

func (b *GraphicComponentBase) DisableDrawable(drawable components.Drawable) {
	b.shape.DisableDrawable(drawable)
}
