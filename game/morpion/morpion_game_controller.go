package morpion

import (
	"github.com/gabz57/goledmatrix/controller"
	"github.com/gabz57/goledmatrix/game/engine"
)

type MorpionGameControllerComponent struct {
	engine.ControllerComponentBase
	morpionGameEntity *MorpionGameEntity
}

func newMorpionGameControllerComponent(entity *MorpionGameEntity) *MorpionGameControllerComponent {
	return &MorpionGameControllerComponent{
		ControllerComponentBase: *engine.NewControllerComponentBase(&entity.Values),
		morpionGameEntity:       entity,
	}
}

func (mgcc *MorpionGameControllerComponent) ConsumeKeyboardEvents(events *[]controller.KeyboardEvent, projection controller.KeyboardProjection) {

	if mgcc.morpionGameEntity.GetValue(gameWinner).(PlayerRef) == NoPlayer &&
		!mgcc.morpionGameEntity.GetValue(exAequo).(bool) {
		for _, event := range *events {
			if event.Action == controller.PressKey {
				if event.Data == "up" {
					mgcc.AddAction(mgcc.morpionGameEntity.moveCursorUp)
				}
				if event.Data == "down" {
					mgcc.AddAction(mgcc.morpionGameEntity.moveCursorDown)
				}
				if event.Data == "left" {
					mgcc.AddAction(mgcc.morpionGameEntity.moveCursorLeft)
				}
				if event.Data == "right" {
					mgcc.AddAction(mgcc.morpionGameEntity.moveCursorRight)
				}
				if event.Data == "esc" {
					mgcc.AddAction(mgcc.morpionGameEntity.pauseGame)
				}
				if event.Data == " " || event.Data == "enter" {
					mgcc.AddAction(mgcc.morpionGameEntity.selectPosition)
				}
			}
		}
	} else {
		for _, event := range *events {
			if event.Action == controller.PressKey {
				mgcc.AddAction(mgcc.morpionGameEntity.startNewGame)
			}
		}
	}
}

func (mgcc *MorpionGameControllerComponent) ConsumeGamepadEvents(events *[]controller.GamepadEvent, projection controller.GamepadProjection) {
	if mgcc.morpionGameEntity.GetValue(gameWinner).(PlayerRef) == NoPlayer &&
		!mgcc.morpionGameEntity.GetValue(exAequo).(bool) {
		for _, event := range *events {
			if event.Action == controller.Press {
				if event.Name == controller.EventTypeDPadUp {
					mgcc.AddAction(mgcc.morpionGameEntity.moveCursorUp)
				}
				if event.Name == controller.EventTypeDPadDown {
					mgcc.AddAction(mgcc.morpionGameEntity.moveCursorDown)
				}
				if event.Name == controller.EventTypeDPadLeft {
					mgcc.AddAction(mgcc.morpionGameEntity.moveCursorLeft)
				}
				if event.Name == controller.EventTypeDPadRight {
					mgcc.AddAction(mgcc.morpionGameEntity.moveCursorRight)
				}
				if event.Name == controller.EventTypeOptions {
					mgcc.AddAction(mgcc.morpionGameEntity.pauseGame)
				}
				if event.Name == controller.EventTypeCross {
					mgcc.AddAction(mgcc.morpionGameEntity.selectPosition)
				}
			}
		}
	} else {
		for _, event := range *events {
			if event.Action == controller.Press {
				mgcc.AddAction(mgcc.morpionGameEntity.startNewGame)
			}
		}
	}
}
