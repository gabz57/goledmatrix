package components

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/controller"
	"image/color"
	"time"
)

type SceneController interface {
	HandleGamepadEvent(event *controller.GamepadEvent, projection *controller.GamepadProjection)
}

type Scene struct {
	components []Component
	duration   time.Duration
	effects    []CanvasEffect
	controller SceneController
}

func NewScene(components []Component, duration time.Duration) *Scene {
	return &Scene{
		components: components,
		duration:   duration,
	}
}

//
//func NewControlledScene(components []Component, effects []CanvasEffect, controller SceneController) *Scene {
//	return &Scene{
//		components: components,
//		duration:   nil,
//		effects:    effects,
//		controller: controller,
//	}
//}

func (s *Scene) WithController(controller SceneController) *Scene {
	s.controller = controller
	return s
}

func (s *Scene) WithEffect(effect CanvasEffect) *Scene {
	s.effects = append(s.effects, effect)
	return s
}

func (s *Scene) WithEffects(effects []CanvasEffect) *Scene {
	s.effects = effects
	return s
}

func (s *Scene) Control(gamepad *controller.Gamepad) {
	if s.controller == nil {
		return
	}
	for {
		select {
		case event := <-(*(*gamepad).EventChannel()):
			s.controller.HandleGamepadEvent(event, (*gamepad).Projection())
		default:
			// avoid blocking select
			return
		}
	}
}

func (s *Scene) Update(elapsedBetweenUpdate time.Duration) bool {
	dirtyScene := false

	for _, effect := range s.effects {
		dirtyScene = effect.Update(elapsedBetweenUpdate) || dirtyScene
	}
	for _, component := range s.components {
		dirtyScene = component.Update(elapsedBetweenUpdate) || dirtyScene
	}

	return dirtyScene
}

func (s *Scene) Render(canvas canvas.Canvas) error {
	canvas.Clear()

	canvas = s.wrapWithEffects(canvas)

	for _, component := range s.components {
		err := component.Draw(canvas)
		if err != nil {
			return err
		}
	}

	return canvas.Render()
}

func (s *Scene) wrapWithEffects(canvas canvas.Canvas) canvas.Canvas {
	for _, effect := range s.effects {
		canvas = NewPixelCanvasAdapter(canvas, effect)
	}
	return canvas
}

type CanvasEffect interface {
	Update(elapsedBetweenUpdate time.Duration) bool
	AdaptPixel() func(canvas canvas.Canvas, x, y int, ledColor color.Color)
}

type PixelAdapter interface {
	Set(x, y int, ledColor color.Color)
}

type CanvasPixelAdapter struct {
	canvas.Canvas
	effect CanvasEffect
}

func (cpa CanvasPixelAdapter) Set(x, y int, ledColor color.Color) {
	cpa.effect.AdaptPixel()(cpa.Canvas, x, y, ledColor)
}

func NewPixelCanvasAdapter(canvas canvas.Canvas, effect CanvasEffect) *CanvasPixelAdapter {
	return &CanvasPixelAdapter{
		Canvas: canvas,
		effect: effect,
	}
}
