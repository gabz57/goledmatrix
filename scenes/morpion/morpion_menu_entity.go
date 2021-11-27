package morpion

import (
	"github.com/gabz57/goledmatrix/game"
	"time"
)

const (
	selectedMenuOption game.ValueRef = iota
	isMenuOptionResumeVisible
)

type MenuOption int

const (
	menuOptionResume       MenuOption = 0
	menuOptionStartNewGame MenuOption = 1
	menuOptionShowScores   MenuOption = 2
	menuOptionExit         MenuOption = 3
)
const nbOptions = 4

type MainMenuEntity struct {
	game.EntityBase
	menuGraphic MainMenuGraphicComponent
	controller  MainMenuControllerComponent
}

func NewMainMenuEntity() *MainMenuEntity {
	entity := MainMenuEntity{
		EntityBase: *game.NewEntityBase(),
	}
	entity.menuGraphic = *newMainMenuGraphicComponent(&entity)
	entity.controller = *newMainMenuControllerComponent(&entity)
	return &entity
}

func (mme *MainMenuEntity) GetController() game.ControllerComponent {
	return &mme.controller
}

func (mme *MainMenuEntity) Update(engine *game.Engine, dt time.Duration) bool {
	return true
}

func (mme *MainMenuEntity) FinalUpdate(dt time.Duration) bool {
	return true
}

func (mme *MainMenuEntity) Enable() {
	mme.menuGraphic.Enable()
}

func (mme *MainMenuEntity) Disable() {
	mme.menuGraphic.Disable()
}

func (mme *MainMenuEntity) focusPreviousOption(engine *game.Engine) error {
	selectedMenuOptionValue := mme.GetValue(selectedMenuOption).(MenuOption)
	switch selectedMenuOptionValue {
	case menuOptionResume:
		// nothing
	case menuOptionStartNewGame:
		if mme.GetValue(isMenuOptionResumeVisible).(bool) {
			mme.SetValue(selectedMenuOption, menuOptionResume)
		}
	case menuOptionShowScores:
		mme.SetValue(selectedMenuOption, menuOptionStartNewGame)
	case menuOptionExit:
		mme.SetValue(selectedMenuOption, menuOptionShowScores)
	}
	mme.menuGraphic.refreshOptionsPosition()
	return nil
}

func (mme *MainMenuEntity) focusNextOption(engine *game.Engine) error {
	selectedMenuOptionValue := mme.GetValue(selectedMenuOption).(MenuOption)
	switch selectedMenuOptionValue {
	case menuOptionResume:
		mme.SetValue(selectedMenuOption, menuOptionStartNewGame)
	case menuOptionStartNewGame:
		mme.SetValue(selectedMenuOption, menuOptionShowScores)
	case menuOptionShowScores:
		mme.SetValue(selectedMenuOption, menuOptionExit)
	case menuOptionExit:
		// nothing
	}
	mme.menuGraphic.refreshOptionsPosition()
	return nil
}

func (mme *MainMenuEntity) selectCurrentOption(engine *game.Engine) error {
	switch mme.GetValue(selectedMenuOption).(MenuOption) {
	case menuOptionResume:
		resume(engine)
	case menuOptionStartNewGame:
		startNewGame(engine)
	case menuOptionShowScores:
		showScores(engine)
	case menuOptionExit:
		exitGame(engine)
	}
	return nil
}

func (mme *MainMenuEntity) initializeMenu() {
	mme.SetValue(selectedMenuOption, menuOptionStartNewGame)
	mme.SetValue(isMenuOptionResumeVisible, false)
	mme.menuGraphic.refreshOptionsPosition()
}
