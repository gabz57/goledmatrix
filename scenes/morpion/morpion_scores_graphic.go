package morpion

import (
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/game"
	"image/color"
)

type ScoresGraphicComponent struct {
	game.GraphicComponentBase
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
		GraphicComponentBase: *game.NewGraphicComponentBase(&entity.Values, graphic, true, true),
	}
}
