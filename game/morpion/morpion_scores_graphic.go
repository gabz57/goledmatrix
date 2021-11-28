package morpion

import (
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/game/engine"
	"image/color"
)

type ScoresGraphicComponent struct {
	engine.GraphicComponentBase
}

//func (t ScoresGraphicComponent) Draw(canvas canvas.Canvas) error {
//	// TODO: paint:
//	//  - menu + option
//	//  - selection
//	return nil
//}

func newScoresGraphicComponent(entity *MorpionScoresEntity) *ScoresGraphicComponent {
	graphic := components.NewGraphic(nil, components.NewLayout(components.ColorWhite, color.Transparent))
	return &ScoresGraphicComponent{
		GraphicComponentBase: *engine.NewGraphicComponentBase(&entity.Values, graphic, true, true),
	}
}
