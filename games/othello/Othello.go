package othello

import (
	"github.com/gabz57/goledmatrix/game"
	"github.com/gabz57/goledmatrix/games/othello/gameobjects"
)

type OthelloGame struct {
	world game.World
}

func NewOthelloGame() *OthelloGame {
	world := game.NewWorld()
	world.RegisterEntityObject(gameobjects.NewOthellier())
	return &OthelloGame{
		world: *world,
	}
}

func (o *OthelloGame) GetWorld() *game.World {
	return &o.world
}
