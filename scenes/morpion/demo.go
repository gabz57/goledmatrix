package morpion

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/game"
)

func Gameloop(c canvas.Canvas, done chan struct{}) {
	gameEngine := game.NewGameEngine(c)
	gameEngine.LoadGame(NewMorpionGame())
	gameEngine.Run(done)
}
