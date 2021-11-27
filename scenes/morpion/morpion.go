package morpion

import (
	"github.com/gabz57/goledmatrix/game"
)

const (
	morpionGameEntityRef game.EntityRef = iota
	menuEntityRef
	scoresEntityRef
)

var rootRefs = []game.EntityRef{menuEntityRef, morpionGameEntityRef, scoresEntityRef}

type MorpionGame struct {
	mainMenu      MainMenuEntity
	morpionGame   MorpionGameEntity
	morpionScores MorpionScoresEntity
}

func NewMorpionGame() *MorpionGame {
	return &MorpionGame{
		mainMenu:      *NewMainMenuEntity(),
		morpionGame:   *NewMorpionGameEntity(),
		morpionScores: *NewMorpionScoresEntity(),
	}
}

func (mg MorpionGame) Name() string {
	return "Morpion"
}

func (mg MorpionGame) InitializeBuckets() []game.EntityBucket {
	// all in first bucket, no dependency between entities updates
	// (where computation depends on another Entity)
	return []game.EntityBucket{
		{
			Entities: map[game.EntityRef]game.Entity{
				menuEntityRef:        &mg.mainMenu,
				morpionGameEntityRef: &mg.morpionGame,
				scoresEntityRef:      &mg.morpionScores,
			},
			ControllerComponents: []game.ControllerComponent{
				&mg.mainMenu.controller,
				&mg.morpionGame.controller,
			},
			AnimationComponents: []game.AnimationComponent{}, // empty
			PhysicsComponents:   []game.PhysicsComponent{},   // empty
			CollisionComponents: []game.CollisionComponent{}, // empty
			GraphicComponents:   mg.allGraphicComponents(),
		},
	}
}

func (mg *MorpionGame) allGraphicComponents() []game.GraphicComponent {
	graphicComponents := []game.GraphicComponent{
		&mg.mainMenu.menuGraphic,
		&mg.morpionScores.scoresGraphic,
		&mg.morpionGame.gridGraphic,
		&mg.morpionGame.textGraphic,
	}
	for _, tokenGraphicComponent := range mg.morpionGame.tokenGraphics() {
		var graphicComponent game.GraphicComponent = tokenGraphicComponent
		graphicComponents = append(graphicComponents, graphicComponent)
	}
	return graphicComponents
}

func (mg MorpionGame) InitializeGame(engine *game.Engine) {
	mg.mainMenu.initializeMenu()
	mg.morpionGame.initializeGame()
	showMainMenu(engine, false)
}

func showMainMenu(engine *game.Engine, pauseGame bool) {
	// todo: something to shades game screen
	if pauseGame {
		menuEntity := engine.GetEntity(menuEntityRef).(*MainMenuEntity)
		menuEntity.SetValue(isMenuOptionResumeVisible, true)
		menuEntity.menuGraphic.refreshOptionsPosition()
	}
	useRootEntity(engine, menuEntityRef)
}
func resume(engine *game.Engine) {
	// todo: something to remove shades game screen
	useRootEntity(engine, morpionGameEntityRef)
}
func startNewGame(engine *game.Engine) {
	engine.GetEntity(morpionGameEntityRef).(*MorpionGameEntity).initializeGame()
	useRootEntity(engine, morpionGameEntityRef)
}
func showScores(engine *game.Engine) {
	useRootEntity(engine, scoresEntityRef)
}
func exitGame(engine *game.Engine) {
	// todo: close game resources ?
	engine.Exit()
}

func useRootEntity(engine *game.Engine, rootRef game.EntityRef) {
	for _, ref := range rootRefs {
		if rootRef != ref {
			engine.GetEntity(ref).(game.Entity).Disable()
		}
	}
	entity := engine.GetEntity(rootRef).(game.Entity)
	engine.SetActiveController(entity.GetController())
	entity.Enable()
}
