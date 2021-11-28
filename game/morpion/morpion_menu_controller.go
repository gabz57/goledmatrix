package morpion

import (
	"github.com/gabz57/goledmatrix/controller"
	"github.com/gabz57/goledmatrix/game/engine"
)

type MainMenuControllerComponent struct {
	engine.ControllerComponentBase
	mainMenuEntity *MainMenuEntity
}

func newMainMenuControllerComponent(entity *MainMenuEntity) *MainMenuControllerComponent {
	return &MainMenuControllerComponent{
		ControllerComponentBase: *engine.NewControllerComponentBase(&entity.Values),
		mainMenuEntity:          entity,
	}
}

func (mmcc *MainMenuControllerComponent) ConsumeGamepadEvents(events *[]controller.GamepadEvent, projection controller.GamepadProjection) {
	for _, event := range *events {
		if event.Name == controller.EventTypeDPadUp && event.Action == controller.Press {
			mmcc.AddAction(mmcc.mainMenuEntity.focusPreviousOption)
		}
		if event.Name == controller.EventTypeDPadDown && event.Action == controller.Press {
			mmcc.AddAction(mmcc.mainMenuEntity.focusNextOption)
		}
		if event.Name == controller.EventTypeCross && event.Action == controller.Press {
			mmcc.AddAction(mmcc.mainMenuEntity.selectCurrentOption)
		}
	}
}
