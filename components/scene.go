package components

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/controller"
	"image/color"
	"time"
)

type Scene struct {
	components []Component
	duration   *time.Duration
	effects    []*CanvasEffect
	controller SceneController
}

func NewScene(components []Component, duration time.Duration) *Scene {
	return &Scene{
		components: components,
		duration:   &duration,
	}
}

func NewControlledScene(components []Component, effects []*CanvasEffect, controller SceneController) *Scene {
	return &Scene{
		components: components,
		duration:   nil,
		effects:    effects,
		controller: controller,
	}
}

func NewSceneWithEffect(components []Component, duration time.Duration, effects []*CanvasEffect) *Scene {
	return &Scene{
		components: components,
		duration:   &duration,
		effects:    effects,
	}
}

var nbMaxEventProcessed = 0

func (s *Scene) Control(gamepad *controller.Gamepad) {
	if s.controller == nil {
		return
	}
	//var nbEventProcessed = 0
	for {
		select {
		case event := <-(*(*gamepad).EventChannel()):
			s.controller.HandleGamepadEvent(event, (*gamepad).Projection())
			//if event != nil {
			//	nbEventProcessed++
			//}
		default:

			//if nbEventProcessed > nbMaxEventProcessed {
			//	nbMaxEventProcessed = nbEventProcessed
			//	println("nbMaxEventProcessed:", nbMaxEventProcessed)
			//}
			// avoid blocking select
			return
		}
	}
}

func (s *Scene) Update(elapsedBetweenUpdate time.Duration) bool {
	dirtyScene := false

	for _, effect := range s.effects {
		dirtyScene = (*effect).Update(elapsedBetweenUpdate) || dirtyScene
	}
	for _, component := range s.components {
		dirtyScene = component.Update(elapsedBetweenUpdate) || dirtyScene
	}

	return dirtyScene
}

func (s *Scene) Render(canvas *canvas.Canvas) error {
	(*canvas).Clear()

	c := s.applyEffects(canvas)

	for _, component := range s.components {
		err := component.Draw(*c)
		if err != nil {
			return err
		}
	}

	return (*c).Render()
}

func (s *Scene) applyEffects(canvas *canvas.Canvas) *canvas.Canvas {
	var c = canvas
	for _, effect := range s.effects {
		c = wrap(c, effect)
	}
	return c
}

func wrap(c *canvas.Canvas, effect *CanvasEffect) *canvas.Canvas {
	var canvasAdapter canvas.Canvas = NewPixelCanvasAdapter(c, &*effect)
	return &canvasAdapter
}

type CanvasEffect interface {
	Update(elapsedBetweenUpdate time.Duration) bool
	AdaptPixel() func(canvas *canvas.Canvas, x, y int, ledColor *color.Color)
}

type PixelAdapter interface {
	Set(x, y int, ledColor color.Color)
}

type CanvasPixelAdapter struct {
	canvas.Canvas
	effect *CanvasEffect
}

func (cpa CanvasPixelAdapter) Set(x, y int, ledColor color.Color) {
	(*cpa.effect).AdaptPixel()(&cpa.Canvas, x, y, &ledColor)
}

func NewPixelCanvasAdapter(canvas *canvas.Canvas, effect *CanvasEffect) *CanvasPixelAdapter {
	return &CanvasPixelAdapter{
		Canvas: *canvas,
		effect: effect,
	}
}
