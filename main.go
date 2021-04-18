package main

import (
	"github.com/gabz57/goledmatrix/components/impl"
	"github.com/gabz57/goledmatrix/matrix"
)

func main() {
	matrix.RunMatrices(goLedApplication)
}

func goLedApplication() {
	matrix.Run(impl.Gameloop)
}
