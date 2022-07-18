package main

import (
	"github.com/faiface/mainthread"
	"github.com/gabz57/goledmatrix/canvas/matrix"
	"github.com/gabz57/goledmatrix/game"
	"github.com/gabz57/goledmatrix/scenes"
)

func main() {
	mainthread.Run(goLedApplication)
}

// TODO: update or mix engines to avoid game/engine + scenes/engine

func goLedApplication() {
	playGame := false

	if playGame {
		// Using distinct files and objects to handle separately (game loop):
		// - control (keyboard, PSX DualShock4 via bluetooth)
		// - entities (~state, game logic)
		// - graphics (rendering on canvas)
		matrix.Run(game.Morpion)
	} else {
		// Direct canvas drawing using simpler Scene ([]Component) loop
		matrix.Run(scenes.Demos)
	}
}
