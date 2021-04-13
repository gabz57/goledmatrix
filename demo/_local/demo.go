package main

import (
	"github.com/gabz57/goledmatrix"
	"github.com/gabz57/goledmatrix/components/impl"
)

func main() {
	goledmatrix.RunMatrices(app)
}

func app() {
	goledmatrix.Run(func(config *goledmatrix.MatrixConfig) (goledmatrix.Matrix, error) {
		return goledmatrix.BuildMatrix(config)
	}, impl.Gameloop)
}
