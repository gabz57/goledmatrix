package morpion

import (
	"github.com/gabz57/goledmatrix/controller"
	"github.com/gabz57/goledmatrix/game"
)

type MorpionGameControllerComponent struct {
	game.ControllerComponentBase
	morpionGameEntity *MorpionGameEntity
}

func newMorpionGameControllerComponent(entity *MorpionGameEntity) *MorpionGameControllerComponent {
	return &MorpionGameControllerComponent{
		ControllerComponentBase: *game.NewControllerComponentBase(&entity.Values),
		morpionGameEntity:       entity,
	}
}

func (mgcc *MorpionGameControllerComponent) ConsumeGamepadEvents(events *[]controller.GamepadEvent, projection controller.GamepadProjection) {
	if mgcc.morpionGameEntity.GetValue(gameWinner).(PlayerRef) == NoPlayer &&
		!mgcc.morpionGameEntity.GetValue(exAequo).(bool) {
		for _, event := range *events {
			if event.Name == controller.EventTypeDPadUp && event.Action == controller.Press {
				mgcc.AddAction(mgcc.morpionGameEntity.moveCursorUp)
			}
			if event.Name == controller.EventTypeDPadDown && event.Action == controller.Press {
				mgcc.AddAction(mgcc.morpionGameEntity.moveCursorDown)
			}
			if event.Name == controller.EventTypeDPadLeft && event.Action == controller.Press {
				mgcc.AddAction(mgcc.morpionGameEntity.moveCursorLeft)
			}
			if event.Name == controller.EventTypeDPadRight && event.Action == controller.Press {
				mgcc.AddAction(mgcc.morpionGameEntity.moveCursorRight)
			}
			if event.Name == controller.EventTypeOptions && event.Action == controller.Press {
				mgcc.AddAction(mgcc.morpionGameEntity.pauseGame)
			}
			if event.Name == controller.EventTypeCross && event.Action == controller.Press {
				mgcc.AddAction(mgcc.morpionGameEntity.selectPosition)
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
