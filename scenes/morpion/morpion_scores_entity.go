package morpion

import (
	"github.com/gabz57/goledmatrix/game"
	"time"
)

type MorpionScoresEntity struct {
	game.EntityBase
	controller    MorpionScoresControllerComponent
	scoresGraphic ScoresGraphicComponent
}

func NewMorpionScoresEntity() *MorpionScoresEntity {
	entity := MorpionScoresEntity{
		EntityBase: *game.NewEntityBase(),
	}
	entity.scoresGraphic = *newScoresGraphicComponent(&entity)
	entity.controller = *newMorpionScoresControllerComponent(&entity)
	return &entity
}

func (mse *MorpionScoresEntity) GetController() game.ControllerComponent {
	return &mse.controller
}

func (mse *MorpionScoresEntity) Enable() {
	mse.scoresGraphic.Enable()
}

func (mse *MorpionScoresEntity) Disable() {
	mse.scoresGraphic.Disable()
}

func (mse *MorpionScoresEntity) Update(engine *game.Engine, dt time.Duration) bool {
	// TODO: variate displayed info ? (~scroll)
	return false
}

func (mse *MorpionScoresEntity) FinalUpdate(dt time.Duration) bool {
	return false
}

func (mse *MorpionScoresEntity) showMenu(engine *game.Engine) error {
	// TODO: pause = true only when game state contains at least one cell AND not all
	showMainMenu(engine, true)
	return nil
}
