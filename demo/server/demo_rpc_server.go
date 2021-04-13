package main

import (
	"github.com/gabz57/goledmatrix"
	"github.com/gabz57/goledmatrix/rpc"
)

func main() {
	goledmatrix.RunMatrices(rpcServer)
}

func rpcServer() {
	goledmatrix.Run(func(config *goledmatrix.MatrixConfig) (goledmatrix.Matrix, error) {
		return goledmatrix.BuildMatrix(config)
	}, rpc.Serve())
}
