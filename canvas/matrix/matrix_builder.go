//go:build darwin
// +build darwin

package matrix

import "github.com/gabz57/goledmatrix/canvas"

func BuildMatrix(config *MatrixConfig) (Matrix, error) {
	if config.Client {
		if config.Emulator == true {
			return NewRpcClientWithEmulator(config)
		} else {
			return NewRpcClient(config)
		}
	} else {
		if config.Emulator == true {
			return NewEmulator(config)
		} else {
			// No hardware usable here :(
			return NewEmulator(config)
		}
	}
}

type RpcClientWithEmulator struct {
	rpcClient Matrix
	emulator  Matrix
	config    *MatrixConfig
}

func NewRpcClientWithEmulator(config *MatrixConfig) (Matrix, error) {
	client, err := NewRpcClient(config)
	if err != nil {
		return nil, err
	}
	emulator, err := NewEmulator(config)
	if err != nil {
		return nil, err
	}
	return &RpcClientWithEmulator{
		config:    config,
		rpcClient: client,
		emulator:  emulator,
	}, nil
}

func (r RpcClientWithEmulator) Config() *MatrixConfig {
	return r.config
}

func (r RpcClientWithEmulator) Geometry() (width, height int) {
	return r.emulator.Geometry()
}

func (r RpcClientWithEmulator) Render(c canvas.Canvas) error {
	go r.emulator.Render(c)
	return r.rpcClient.Render(c)
}

func (r RpcClientWithEmulator) RenderMethod(c canvas.Canvas) error {
	go r.emulator.RenderMethod(c)
	return r.rpcClient.RenderMethod(c)
}

func (r RpcClientWithEmulator) Close() error {
	err := r.rpcClient.Close()
	if err != nil {
		return err
	}
	return r.emulator.Close()
}

func (r RpcClientWithEmulator) MainThread(c canvas.Canvas, done chan struct{}) {
	//r.rpcClient.MainThread(c, done)
	r.emulator.MainThread(c, done)
}

func (r RpcClientWithEmulator) send(event interface{}) {
	r.rpcClient.send(event)
	r.emulator.send(event)
}
