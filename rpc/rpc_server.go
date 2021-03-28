package rpc

import (
	"fmt"
	"github.com/gabz57/goledmatrix"
	"image/color"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type MatrixRPCServer struct {
	m goledmatrix.Matrix
	c *goledmatrix.Canvas
}

type GeometryArgs struct{}
type GeometryReply struct{ Width, Height int }

func (m *MatrixRPCServer) Geometry(_ *GeometryArgs, reply *GeometryReply) error {
	w, h := m.m.Geometry()
	reply.Width = w
	reply.Height = h
	return nil
}

type RenderArgs struct{ Colors []color.Color }
type RenderReply struct{}

func (m *MatrixRPCServer) Render(args *RenderArgs, _ *RenderReply) error {
	defer m.c.Clear()
	w, h := m.m.Geometry()

	var c color.Color
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c = args.Colors[x+y*w]
			if c != nil {
				m.c.Set(x, y, c)
			}
			c = nil
		}
	}
	return m.c.Render()
}

type CloseArgs struct{}
type CloseReply struct{}

func (m *MatrixRPCServer) Close(_ *CloseArgs, _ *CloseReply) error {
	return m.m.Close()
}

func Serve(serverMatrix *goledmatrix.Matrix) func(c *goledmatrix.Canvas, done chan struct{}) {
	return func(c *goledmatrix.Canvas, done chan struct{}) {
		serve(serverMatrix, c) // Blocking
		fmt.Println("RPC Server Stopped")
		done <- struct{}{}
	}
}

func serve(m *goledmatrix.Matrix, c *goledmatrix.Canvas) {
	rpc.Register(&MatrixRPCServer{*m, c})
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println(l)
	fmt.Println("Serving...")
	http.Serve(l, nil)
}
