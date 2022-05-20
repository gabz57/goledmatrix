package scenes

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/canvas/effect"
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/controller"
	"time"
)

type SceneController interface {
	HandleGamepadEvent(event *controller.GamepadEvent, projection *controller.GamepadProjection)
	HandleKeyboardEvent(event *controller.KeyboardEvent, projection *controller.KeyboardProjection)
}

type Scene struct {
	components []components.Component
	duration   time.Duration
	effects    []effect.DynamicEffect
	controller SceneController
}

func NewScene(components []components.Component, duration time.Duration) *Scene {
	return &Scene{
		components: components,
		duration:   duration,
	}
}

func (s *Scene) WithController(controller SceneController) *Scene {
	s.controller = controller
	return s
}

func (s *Scene) WithEffect(effect effect.DynamicEffect) *Scene {
	s.effects = append(s.effects, effect)
	return s
}

func (s *Scene) WithEffects(effects []effect.DynamicEffect) *Scene {
	for _, dynamicEffect := range effects {
		_ = s.WithEffect(dynamicEffect)
	}
	return s
}

func (s *Scene) GamepadControl(gamepad *controller.Gamepad) {
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

func (s *Scene) KeyboardControl(keyboard *controller.Keyboard) {
	if s.controller == nil {
		return
	}
	for {
		select {
		case event := <-(*(*keyboard).EventChannel()):
			s.controller.HandleKeyboardEvent(event, (*keyboard).Projection())
		default:
			// avoid blocking select
			return
		}
	}
}

func (s *Scene) Update(elapsedBetweenUpdate time.Duration) bool {
	dirtyScene := false

	for _, dynamicEffect := range s.effects {
		dirtyScene = dynamicEffect.Update(elapsedBetweenUpdate) || dirtyScene
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
	for _, e := range s.effects {
		canvas = effect.NewAdapter(canvas, e)
	}
	return canvas
}
