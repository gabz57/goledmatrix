package rpc

import (
	"encoding/gob"
	"github.com/gabz57/goledmatrix"
	"image/color"
	"net/rpc"
)

func init() {
	gob.Register(color.RGBA{})
}

type MatrixRpcClient struct {
	config  *goledmatrix.MatrixConfig
	network string
	addr    string
	client  *rpc.Client
}

func NewMatrixRpcClient(network, addr string, config *goledmatrix.MatrixConfig) (goledmatrix.Matrix, error) {
	rpcClient, err := rpc.DialHTTP(network, addr)
	if err != nil {
		return nil, err
	}

	return &MatrixRpcClient{
		config:  config,
		network: network,
		addr:    addr,
		client:  rpcClient,
	}, nil
}

func (m *MatrixRpcClient) Config() *goledmatrix.MatrixConfig {
	return m.config
}

func (m *MatrixRpcClient) Geometry() (width, height int) {
	var reply *GeometryReply
	err := m.client.Call("MatrixRPCServer.Geometry", &GeometryArgs{}, &reply)
	if err != nil {
		panic(err)
	}

	return reply.Width, reply.Height
}

func (m *MatrixRpcClient) RenderMethod(canvas *goledmatrix.Canvas) error {
	return m.Render(canvas)
}

func (m *MatrixRpcClient) Render(canvas *goledmatrix.Canvas) error {
	var reply *RenderReply
	return m.client.Call("MatrixRPCServer.Render", &RenderArgs{Colors: canvas.Leds()}, &reply)
}

func (m *MatrixRpcClient) Close() error {
	var reply *CloseReply
	return m.client.Call("MatrixRPCServer.Close", &CloseArgs{}, &reply)
}

func (m *MatrixRpcClient) MainThread(_ *goledmatrix.Canvas, done chan struct{}) {
	select {
	case <-done:
		break
	}
}

func (m *MatrixRpcClient) Send(_ interface{}) {
	// implemented only in emulator which handle data asynchronously
}

func (m *MatrixRpcClient) IsEmulator() bool {
	return false
}
