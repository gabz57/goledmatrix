package main

import (
	"github.com/gabz57/goledmatrix"
	"github.com/gabz57/goledmatrix/rpc"
)

func main() {
	goledmatrix.RunMatrices(app)
}

func app() {
	var rpcDrivenMatrix *goledmatrix.Matrix
	goledmatrix.Run(func(config *goledmatrix.MatrixConfig) (goledmatrix.Matrix, error) {
		return rpcDriven(config, rpcDrivenMatrix)
	}, rpc.Serve(rpcDrivenMatrix))
}

func rpcDriven(config *goledmatrix.MatrixConfig, rpcDrivenMatrix *goledmatrix.Matrix) (goledmatrix.Matrix, error) {
	// can be Emulator or Hardware
	matrix, err := goledmatrix.BuildMatrix(config)
	rpcDrivenMatrix = &matrix
	return matrix, err
}
