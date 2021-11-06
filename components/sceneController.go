package components

import (
	"github.com/gabz57/goledmatrix/controller"
)

type SceneController interface {
	HandleGamepadEvent(event *controller.GamepadEvent, projection *controller.GamepadProjection)
}
