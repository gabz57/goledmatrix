package matrix

import (
	"encoding/gob"
	"fmt"
	. "github.com/gabz57/goledmatrix/canvas"
	"image/color"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func init() {
	gob.Register(color.RGBA{})
}

type MatrixRPCServer struct {
	m         *Matrix
	c         Canvas
	timestamp int64
}

type GeometryArgs struct{}
type GeometryReply struct{ Width, Height int }

func (m *MatrixRPCServer) Geometry(_ *GeometryArgs, reply *GeometryReply) error {
	fmt.Println("MatrixRPCServer.Geometry()")
	w, h := (*m.m).Geometry()
	reply.Width = w
	reply.Height = h
	return nil
}

type RpcPixel struct {
	X, Y int
	C    color.Color
}

type RenderArgs struct {
	Pixels []RpcPixel
	//Colors []color.Color
	Timestamp int64
}
type RenderReply struct{}

func (m *MatrixRPCServer) Render(args *RenderArgs, _ *RenderReply) error {
	defer m.c.Clear()
	for _, pixel := range args.Pixels {
		m.c.Set(pixel.X, pixel.Y, pixel.C)
	}
	return m.c.Render()
}

type CloseArgs struct{}
type CloseReply struct{}

func (m *MatrixRPCServer) Close(_ *CloseArgs, _ *CloseReply) error {
	fmt.Println("MatrixRPCServer.Close()")
	return (*m.m).Close()
}

func Serve(matrix *Matrix) func(c Canvas, done chan struct{}) {
	return func(c Canvas, done chan struct{}) {
		serve(matrix, c) // Blocking
		fmt.Println("RPC Server Stopped")
	}
}

func serve(m *Matrix, c Canvas) {
	server := MatrixRPCServer{m: m, c: c}
	rpc.Register(&server)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("Serving... @ " + l.Addr().String())
	http.Serve(l, nil)
}
