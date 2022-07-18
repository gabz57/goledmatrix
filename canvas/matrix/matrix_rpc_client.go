package matrix

import (
	"encoding/gob"
	. "github.com/gabz57/goledmatrix/canvas"
	"image/color"
	"log"
	"net/rpc"
	"strconv"
	"time"
)

func init() {
	gob.Register(color.RGBA{})
}

type MatrixRpcClient struct {
	config            *MatrixConfig
	client            *rpc.Client
	lastFrameDuration time.Duration
	previous          []RpcPixel
}

func NewRpcClient(config *MatrixConfig) (Matrix, error) {
	rpcClient, err := rpc.DialHTTP("tcp", config.IpAddress+":8080")
	if err != nil {
		return nil, err
	}

	client := MatrixRpcClient{
		config:            config,
		client:            rpcClient,
		lastFrameDuration: 0 * time.Nanosecond,
	}

	if err != nil {
		return nil, err
	}
	if err = validateGeometry(config, &client); err != nil {
		return nil, err
	}
	return &client, nil
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

const frameDurationInNanos = 33333333 // 30 FPS approximated in nanos

func (m *MatrixRpcClient) Render(canvas Canvas) error {
	if m.lastFrameDuration < frameDurationInNanos {
		start := time.Now()
		err := m.doRender(canvas)
		duration := time.Now().Sub(start)
		log.Println("rpc render took " + strconv.FormatInt(duration.Milliseconds(), 10) + "ms")
		m.lastFrameDuration = duration
		return err

	} else {
		log.Println("rpc render skipped")

		m.lastFrameDuration = 0 * time.Nanosecond
		return nil
	}
}

func (m *MatrixRpcClient) doRender(canvas Canvas) error {
	var reply *RenderReply
	pixels := toRpcPixels(canvas)
	if m.sameAsPrevious(pixels) {
		println("skip same")
		m.previous = pixels
		return nil
	} else {
		m.previous = pixels
		err := m.client.Call("MatrixRPCServer.Render", &RenderArgs{
			Pixels:    pixels,
			Timestamp: time.Now().UnixNano(),
		}, &reply)
		return err
	}
}

func toRpcPixels(canvas Canvas) []RpcPixel {
	width := canvas.Bounds().Max.X
	pixels := make([]RpcPixel, canvas.Bounds().Max.X*canvas.Bounds().Max.Y)
	index := 0
	for i, c := range *canvas.GetLeds() {
		if c != nil && !isBlack(c) {
			pixels[index] = RpcPixel{
				X: i % width,
				Y: i / width,
				C: c,
			}
			index++
		}
	}
	return pixels[:index]
}

func isBlack(c color.Color) bool {
	br, bg, bb, _ := color.Black.RGBA()
	r, g, b, _ := c.RGBA()
	return r == br && g == bg && b == bb
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

func (m *MatrixRpcClient) sameAsPrevious(pixels []RpcPixel) bool {
	if m.previous == nil {
		return false
	}
	if len(pixels) != len(m.previous) {
		return false
	}
	for i, pixel := range pixels {
		rpcPixel := m.previous[i]
		br, bg, bb, _ := rpcPixel.C.RGBA()
		r, g, b, _ := rpcPixel.C.RGBA()
		if pixel.X != rpcPixel.X ||
			pixel.Y != rpcPixel.Y ||
			r != br ||
			g != bg ||
			b != bb {
			return false
		}
	}
	return true
}
