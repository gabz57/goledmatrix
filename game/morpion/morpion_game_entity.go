package morpion

import (
	"github.com/gabz57/goledmatrix/game/engine"
	"time"
)

type (
	PlayerRef int

	Grid struct {
		grid []PlayerRef
	}

	PositionRef int
)

const (
	NoPlayer PlayerRef = -1
	PlayerX  PlayerRef = 0
	PlayerO  PlayerRef = 1
)

func nextPlayer(ref PlayerRef) PlayerRef {
	return (ref + 1) % 2
}

const (
	out  PositionRef = -1000
	p1_1 PositionRef = 0
	p1_2 PositionRef = 1
	p1_3 PositionRef = 2
	p2_1 PositionRef = 3
	p2_2 PositionRef = 4
	p2_3 PositionRef = 5
	p3_1 PositionRef = 6
	p3_2 PositionRef = 7
	p3_3 PositionRef = 8
)
const (
	gameStateValue engine.ValueRef = iota
	gameWinner
	exAequo
	//justUpdatedGameState
	currentPlayerRefValue
	cursorPositionValue
)

type GameState []PlayerRef

func (s GameState) reset() {
	for i := range s {
		s[i] = NoPlayer
	}
}
func (s GameState) winner() (PlayerRef, []PositionRef) {
	winningLines := s.checkPlayer(PlayerX)
	if winningLines != nil {
		return PlayerX, winningLines
	}
	winningLines = s.checkPlayer(PlayerO)
	if winningLines != nil {
		return PlayerO, winningLines
	}
	return NoPlayer, nil
}

func (s GameState) checkPlayer(player PlayerRef) []PositionRef {
	lines := [][]PositionRef{{p1_1, p1_2, p1_3}, {p2_1, p2_2, p2_3}, {p3_1, p3_2, p3_3},
		{p1_1, p2_1, p3_1}, {p1_2, p2_2, p3_2}, {p1_3, p2_3, p3_3},
		{p1_1, p2_2, p3_3}, {p1_3, p2_2, p3_1},
	}
	for _, line := range lines {
		if s.check(line, player) {
			return line
		}
	}
	return nil
}

func (s GameState) check(positions []PositionRef, ref PlayerRef) bool {
	for _, position := range positions {
		if ref != s[position] {
			return false
		}
	}
	return true
}

func (s GameState) hasFreeCell() bool {
	for _, ref := range s {
		if ref == NoPlayer {
			return true
		}
	}
	return false
}

type MorpionGameEntity struct {
	engine.EntityBase
	controller    MorpionGameControllerComponent
	gridGraphic   GridGraphicComponent
	tokenGraphic1 TokenGraphicComponent
	tokenGraphic2 TokenGraphicComponent
	tokenGraphic3 TokenGraphicComponent
	tokenGraphic4 TokenGraphicComponent
	tokenGraphic5 TokenGraphicComponent
	tokenGraphic6 TokenGraphicComponent
	tokenGraphic7 TokenGraphicComponent
	tokenGraphic8 TokenGraphicComponent
	tokenGraphic9 TokenGraphicComponent
	textGraphic   TextGraphicComponent
}

var gameState = GameState{
	NoPlayer, NoPlayer, NoPlayer,
	NoPlayer, NoPlayer, NoPlayer,
	NoPlayer, NoPlayer, NoPlayer,
}
var initialPlayer = PlayerX

func NewMorpionGameEntity() *MorpionGameEntity {
	entity := MorpionGameEntity{
		EntityBase: *engine.NewEntityBase(),
	}
	entity.gridGraphic = *newGridGraphicComponent(&entity)
	entity.tokenGraphic1 = *newTokenGraphicComponent(&entity, p1_1)
	entity.tokenGraphic2 = *newTokenGraphicComponent(&entity, p1_2)
	entity.tokenGraphic3 = *newTokenGraphicComponent(&entity, p1_3)
	entity.tokenGraphic4 = *newTokenGraphicComponent(&entity, p2_1)
	entity.tokenGraphic5 = *newTokenGraphicComponent(&entity, p2_2)
	entity.tokenGraphic6 = *newTokenGraphicComponent(&entity, p2_3)
	entity.tokenGraphic7 = *newTokenGraphicComponent(&entity, p3_1)
	entity.tokenGraphic8 = *newTokenGraphicComponent(&entity, p3_2)
	entity.tokenGraphic9 = *newTokenGraphicComponent(&entity, p3_3)
	entity.textGraphic = *newTextGraphic(&entity)
	entity.controller = *newMorpionGameControllerComponent(&entity)

	return &entity
}

func (e *MorpionGameEntity) GetController() engine.ControllerComponent {
	return &e.controller
}

func (e *MorpionGameEntity) getTokenGraphic(ref PositionRef) *TokenGraphicComponent {
	switch ref {
	case p1_1:
		return &e.tokenGraphic1
	case p1_2:
		return &e.tokenGraphic2
	case p1_3:
		return &e.tokenGraphic3
	case p2_1:
		return &e.tokenGraphic4
	case p2_2:
		return &e.tokenGraphic5
	case p2_3:
		return &e.tokenGraphic6
	case p3_1:
		return &e.tokenGraphic7
	case p3_2:
		return &e.tokenGraphic8
	case p3_3:
		return &e.tokenGraphic9
	case out:
		return nil
	}
	return nil
}

func (e *MorpionGameEntity) Update(engine *engine.Engine, dt time.Duration) bool {
	return false
}

func (e *MorpionGameEntity) FinalUpdate(dt time.Duration) bool {
	return false
}

func (e *MorpionGameEntity) Enable() {
	e.gridGraphic.Enable()
	for _, tokenGraphic := range e.tokenGraphics() {
		tokenGraphic.Enable()
	}
	e.textGraphic.Enable()
}

func (e *MorpionGameEntity) Disable() {
	e.gridGraphic.Disable()
	for _, tokenGraphic := range e.tokenGraphics() {
		tokenGraphic.Disable()
	}
	e.textGraphic.Disable()
}

func (e *MorpionGameEntity) tokenGraphics() []*TokenGraphicComponent {
	return []*TokenGraphicComponent{
		&e.tokenGraphic1,
		&e.tokenGraphic2,
		&e.tokenGraphic3,
		&e.tokenGraphic4,
		&e.tokenGraphic5,
		&e.tokenGraphic6,
		&e.tokenGraphic7,
		&e.tokenGraphic8,
		&e.tokenGraphic9,
	}
}

func (e *MorpionGameEntity) startNewGame(engine *engine.Engine) error {
	engine.GetEntity(menuEntityRef).(*MainMenuEntity).initializeMenu()
	e.initializeGame()
	return nil
}

func (e *MorpionGameEntity) moveCursorUp(engine *engine.Engine) error {
	ref := e.GetValue(cursorPositionValue).(PositionRef)
	if ref >= 3 {
		e.SetValue(cursorPositionValue, ref-3)
		e.updateCellSelection(ref, ref-3)
	}
	return nil
}
func (e *MorpionGameEntity) moveCursorDown(engine *engine.Engine) error {
	ref := e.GetValue(cursorPositionValue).(PositionRef)
	if ref <= 5 {
		e.SetValue(cursorPositionValue, ref+3)
		e.updateCellSelection(ref, ref+3)
	}
	return nil
}
func (e *MorpionGameEntity) moveCursorLeft(engine *engine.Engine) error {
	ref := e.GetValue(cursorPositionValue).(PositionRef)
	if ref >= 1 {
		e.SetValue(cursorPositionValue, ref-1)
		e.updateCellSelection(ref, ref-1)
	}
	return nil
}
func (e *MorpionGameEntity) moveCursorRight(engine *engine.Engine) error {
	ref := e.GetValue(cursorPositionValue).(PositionRef)
	if ref <= 7 {
		e.SetValue(cursorPositionValue, ref+1)
		e.updateCellSelection(ref, ref+1)
	}
	return nil
}

func (e *MorpionGameEntity) pauseGame(engine *engine.Engine) error {
	showMainMenu(engine, true)
	return nil
}

func (e *MorpionGameEntity) initializeGame() {
	gameState.reset()
	e.SetValue(gameStateValue, gameState)
	e.SetValue(currentPlayerRefValue, initialPlayer)
	e.SetValue(gameWinner, NoPlayer)
	e.SetValue(exAequo, false)
	for _, tokenGraphic := range e.tokenGraphics() {
		tokenGraphic.setPlayer(NoPlayer)
		tokenGraphic.setSelected(false)
	}
	e.SetValue(cursorPositionValue, p2_2)
	e.getTokenGraphic(p2_2).setSelected(true)
	e.textGraphic.setNextPlayer(initialPlayer)
}

func (e *MorpionGameEntity) selectPosition(engine *engine.Engine) error {
	gameState := e.GetValue(gameStateValue).(GameState)
	currentPositionRef := e.GetValue(cursorPositionValue).(PositionRef)
	currentPlayerRef := e.GetValue(currentPlayerRefValue).(PlayerRef)
	currentPositionPlayerRef := gameState[currentPositionRef]
	if currentPositionPlayerRef == NoPlayer {
		e.takePosition(gameState, currentPositionRef, currentPlayerRef)

		if winnerPlayerRef, winningCells := gameState.winner(); winnerPlayerRef != NoPlayer {
			if winnerPlayerRef == PlayerO {
				println("👑 player O")
			} else if winnerPlayerRef == PlayerX {
				println("👑 player X")
			}
			//e.SetValue(currentPlayerRefValue, NoPlayer)

			e.SetValue(gameWinner, winnerPlayerRef)
			e.textGraphic.setWinner(winnerPlayerRef)

			for _, winningCellRef := range winningCells {
				e.getTokenGraphic(winningCellRef).setWinnerCell()
			}
		} else if !gameState.hasFreeCell() {
			e.SetValue(exAequo, true)
			e.textGraphic.setExAequo()

		} else {
			nextPlayer := nextPlayer(currentPlayerRef)
			e.SetValue(currentPlayerRefValue, nextPlayer)
			e.textGraphic.setNextPlayer(nextPlayer)
		}
	} else {
		// position already occupied, do nothing
	}
	return nil
}

func (e *MorpionGameEntity) takePosition(gameState GameState, currentPositionRef PositionRef, currentPlayerRef PlayerRef) {
	gameState[currentPositionRef] = currentPlayerRef
	e.SetValue(gameStateValue, gameState)
	e.getTokenGraphic(currentPositionRef).setPlayer(currentPlayerRef)
}

func (e *MorpionGameEntity) updateCellSelection(oldRef, newRef PositionRef) {
	e.getTokenGraphic(oldRef).setSelected(false)
	e.getTokenGraphic(newRef).setSelected(true)
}
