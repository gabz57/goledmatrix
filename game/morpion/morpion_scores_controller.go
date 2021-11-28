package morpion

import (
	"github.com/gabz57/goledmatrix/controller"
	"github.com/gabz57/goledmatrix/game/engine"
)

type MorpionScoresControllerComponent struct {
	engine.ControllerComponentBase
	morpionScoresEntity *MorpionScoresEntity
}

func newMorpionScoresControllerComponent(entity *MorpionScoresEntity) *MorpionScoresControllerComponent {
	return &MorpionScoresControllerComponent{
		ControllerComponentBase: *engine.NewControllerComponentBase(&entity.Values),
		morpionScoresEntity:     entity,
	}
}

func (mscc *MorpionScoresControllerComponent) ConsumeGamepadEvents(events *[]controller.GamepadEvent, projection controller.GamepadProjection) {
	for _, event := range *events {
		if event.Action == controller.Press {
			mscc.AddAction(mscc.morpionScoresEntity.showMenu)
		}
	}
}
