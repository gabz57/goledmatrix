package matrix

import (
	"encoding/gob"
	"fmt"
	. "github.com/gabz57/goledmatrix/canvas"
	"image/color"
	"net/rpc"
	"strconv"
	"time"
)

func init() {
	gob.Register(color.RGBA{})
}

type MatrixRpcClient struct {
	config *MatrixConfig
	client *rpc.Client
}

func NewMatrixRpcClient(config *MatrixConfig) (Matrix, error) {
	rpcClient, err := rpc.DialHTTP("tcp", config.IpAddress+":8080")
	if err != nil {
		return nil, err
	}

	return &MatrixRpcClient{
		config: config,
		client: rpcClient,
	}, nil
}

func (m *MatrixRpcClient) Config() *MatrixConfig {
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

func (m *MatrixRpcClient) RenderMethod(canvas Canvas) error {
	return m.Render(canvas)
}

func (m *MatrixRpcClient) Render(canvas Canvas) error {
	var reply *RenderReply
	start := time.Now()
	err := m.client.Call("MatrixRPCServer.Render", &RenderArgs{
		Pixels:    toRpcPixels(canvas),
		Timestamp: time.Now().UnixNano(),
	}, &reply)
	fmt.Println("rpc render took " + strconv.FormatInt(time.Now().Sub(start).Milliseconds(), 10) + "ms")
	return err
}

func toRpcPixels(canvas Canvas) (pixels []RpcPixel) {
	width := canvas.Bounds().Max.X
	for i, c := range *canvas.GetLeds() {
		if c != nil {
			pixels = append(pixels, RpcPixel{
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

func (m *MatrixRpcClient) MainThread(_ Canvas, done chan struct{}) {
	select {
	case <-done:
		break
	}
}

func (m *MatrixRpcClient) send(_ interface{}) {
	// implemented only in emulator which handle data asynchronously
}
