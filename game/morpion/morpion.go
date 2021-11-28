package morpion

import (
	"github.com/gabz57/goledmatrix/game/engine"
)

const (
	morpionGameEntityRef engine.EntityRef = iota
	menuEntityRef
	scoresEntityRef
)

var rootRefs = []engine.EntityRef{menuEntityRef, morpionGameEntityRef, scoresEntityRef}

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

func (mg MorpionGame) InitializeBuckets() []engine.EntityBucket {
	// all in first bucket, no dependency between entities updates
	// (where computation depends on another Entity)
	return []engine.EntityBucket{
		{
			Entities: map[engine.EntityRef]engine.Entity{
				menuEntityRef:        &mg.mainMenu,
				morpionGameEntityRef: &mg.morpionGame,
				scoresEntityRef:      &mg.morpionScores,
			},
			ControllerComponents: []engine.ControllerComponent{
				&mg.mainMenu.controller,
				&mg.morpionGame.controller,
			},
			AnimationComponents: []engine.AnimationComponent{}, // empty
			PhysicsComponents:   []engine.PhysicsComponent{},   // empty
			CollisionComponents: []engine.CollisionComponent{}, // empty
			GraphicComponents:   mg.allGraphicComponents(),
		},
	}
}

func (mg *MorpionGame) allGraphicComponents() []engine.GraphicComponent {
	graphicComponents := []engine.GraphicComponent{
		&mg.mainMenu.menuGraphic,
		&mg.morpionScores.scoresGraphic,
		&mg.morpionGame.gridGraphic,
		&mg.morpionGame.textGraphic,
	}
	for _, tokenGraphicComponent := range mg.morpionGame.tokenGraphics() {
		var graphicComponent engine.GraphicComponent = tokenGraphicComponent
		graphicComponents = append(graphicComponents, graphicComponent)
	}
	return graphicComponents
}

func (mg MorpionGame) InitializeGame(engine *engine.Engine) {
	mg.mainMenu.initializeMenu()
	mg.morpionGame.initializeGame()
	showMainMenu(engine, false)
}

func showMainMenu(engine *engine.Engine, pauseGame bool) {
	// todo: something to shades game screen
	if pauseGame {
		menuEntity := engine.GetEntity(menuEntityRef).(*MainMenuEntity)
		menuEntity.SetValue(isMenuOptionResumeVisible, true)
		menuEntity.menuGraphic.refreshOptionsPosition()
	}
	useRootEntity(engine, menuEntityRef)
}
func resume(engine *engine.Engine) {
	// todo: something to remove shades game screen
	useRootEntity(engine, morpionGameEntityRef)
}
func startNewGame(engine *engine.Engine) {
	engine.GetEntity(morpionGameEntityRef).(*MorpionGameEntity).initializeGame()
	useRootEntity(engine, morpionGameEntityRef)
}
func showScores(engine *engine.Engine) {
	useRootEntity(engine, scoresEntityRef)
}
func exitGame(engine *engine.Engine) {
	// todo: close game resources ?
	engine.Exit()
}

func useRootEntity(e *engine.Engine, rootRef engine.EntityRef) {
	for _, ref := range rootRefs {
		if rootRef != ref {
			e.GetEntity(ref).(engine.Entity).Disable()
		}
	}
	entity := e.GetEntity(rootRef).(engine.Entity)
	e.SetActiveController(entity.GetController())
	entity.Enable()
}
