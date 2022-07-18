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

func (mmcc *MainMenuControllerComponent) ConsumeKeyboardEvents(events *[]controller.KeyboardEvent, projection controller.KeyboardProjection) {
	for _, event := range *events {
		//actionTxt := "press"
		//if event.Action == controller.ReleaseKey {
		//	actionTxt = "release"
		//}
		//log.Println("Keyboard event", event.Data, " ", actionTxt)

		if event.Action == controller.PressKey {
			if event.Data == "up" {
				mmcc.AddAction(mmcc.mainMenuEntity.focusPreviousOption)
			}
			if event.Data == "down" {
				mmcc.AddAction(mmcc.mainMenuEntity.focusNextOption)
			}
			if event.Data == " " || event.Data == "enter" {
				mmcc.AddAction(mmcc.mainMenuEntity.selectCurrentOption)
			}
		}
	}
}

func (mmcc *MainMenuControllerComponent) ConsumeGamepadEvents(events *[]controller.GamepadEvent, projection controller.GamepadProjection) {
	for _, event := range *events {
		if event.Action == controller.Press {
			if event.Name == controller.EventTypeDPadUp {
				mmcc.AddAction(mmcc.mainMenuEntity.focusPreviousOption)
			}
			if event.Name == controller.EventTypeDPadDown {
				mmcc.AddAction(mmcc.mainMenuEntity.focusNextOption)
			}
			if event.Name == controller.EventTypeCross {
				mmcc.AddAction(mmcc.mainMenuEntity.selectCurrentOption)
			}
		}
	}
}
