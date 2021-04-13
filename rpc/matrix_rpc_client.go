package rpc

import (
	"encoding/gob"
	"fmt"
	"github.com/gabz57/goledmatrix"
	"image/color"
	"net/rpc"
	"strconv"
	"time"
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
	start := time.Now()
	err := m.client.Call("MatrixRPCServer.Render", &RenderArgs{
		Pixels:    toPixels(canvas),
		Timestamp: time.Now().UnixNano(),
	}, &reply)
	fmt.Println("rpc render took " + strconv.FormatInt(time.Now().Sub(start).Milliseconds(), 10) + "ms")
	return err
}

func (m *MatrixRpcClient) rpcRender(canvas *goledmatrix.Canvas, reply *RenderReply) error {
	start := time.Now()
	err := m.client.Call("MatrixRPCServer.Render", &RenderArgs{
		Pixels:    toPixels(canvas),
		Timestamp: time.Now().UnixNano(),
	}, &reply)
	fmt.Println("rpc render took" + strconv.FormatInt(time.Now().Sub(start).Milliseconds(), 10) + "ms")
	return err
}

func toPixels(canvas *goledmatrix.Canvas) (pixels []Pixel) {
	width := (*canvas).Bounds().Max.X
	for i, c := range *(*canvas).GetLeds() {
		if c != nil {
			pixels = append(pixels, Pixel{
				X: i % width,
				Y: i / width,
				C: c,
			})
		}
	}
	return pixels
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
