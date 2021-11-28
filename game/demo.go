package game

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/game/engine"
	"github.com/gabz57/goledmatrix/game/morpion"
)

func Morpion(c canvas.Canvas, done chan struct{}) {
	gameEngine := engine.NewGameEngine(c)
	gameEngine.LoadGame(morpion.NewMorpionGame())
	gameEngine.Run(done)
}
