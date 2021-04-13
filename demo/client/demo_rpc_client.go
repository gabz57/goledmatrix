package main

import (
	"github.com/gabz57/goledmatrix"
	"github.com/gabz57/goledmatrix/components/impl"
	"github.com/gabz57/goledmatrix/rpc"
	"os"
)

func main() {
	goledmatrix.RunMatrices(rpcClient)
}

func rpcClient() {
	goledmatrix.RunMany([]func(config *goledmatrix.MatrixConfig) (goledmatrix.Matrix, error){
		func(config *goledmatrix.MatrixConfig) (goledmatrix.Matrix, error) {
			return rpc.NewMatrixRpcClient("tcp", os.Getenv("SERVER_ADDRESS")+":8080", config)
		},
		// TODO: Add or Force EMULATOR UI (since we are running as RPC Client) ?
		//func() (goledmatrix.Matrix, error) {
		//	return goledmatrix.BuildMatrix()
		//},
	}, impl.Gameloop)
}
