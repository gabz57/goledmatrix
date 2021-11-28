package morpion

import (
	"github.com/gabz57/goledmatrix/game/engine"
	"time"
)

type MorpionScoresEntity struct {
	engine.EntityBase
	controller    MorpionScoresControllerComponent
	scoresGraphic ScoresGraphicComponent
}

func NewMorpionScoresEntity() *MorpionScoresEntity {
	entity := MorpionScoresEntity{
		EntityBase: *engine.NewEntityBase(),
	}
	entity.scoresGraphic = *newScoresGraphicComponent(&entity)
	entity.controller = *newMorpionScoresControllerComponent(&entity)
	return &entity
}

func (mse *MorpionScoresEntity) GetController() engine.ControllerComponent {
	return &mse.controller
}

func (mse *MorpionScoresEntity) Enable() {
	mse.scoresGraphic.Enable()
}

func (mse *MorpionScoresEntity) Disable() {
	mse.scoresGraphic.Disable()
}

func (mse *MorpionScoresEntity) Update(engine *engine.Engine, dt time.Duration) bool {
	// TODO: variate displayed info ? (~scroll)
	return false
}

func (mse *MorpionScoresEntity) FinalUpdate(dt time.Duration) bool {
	return false
}

func (mse *MorpionScoresEntity) showMenu(engine *engine.Engine) error {
	// TODO: pause = true only when game state contains at least one cell AND not all
	showMainMenu(engine, true)
	return nil
}
