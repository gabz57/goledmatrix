package main

import (
	"github.com/gabz57/goledmatrix"
	"github.com/gabz57/goledmatrix/components/impl"
)

func main() {
	goledmatrix.RunMatrices(goLedApplication)
}

func goLedApplication() {
	goledmatrix.Run(impl.Gameloop)
}
