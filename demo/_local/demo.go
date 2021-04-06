package main

import (
	"fmt"
	"github.com/gabz57/goledmatrix"
	"github.com/gabz57/goledmatrix/components"
	"time"
)

const (
	frameDuration = time.Second
)

func main() {
	goledmatrix.RunMatrices(app)
}

func app() {
	goledmatrix.Run(func(config *goledmatrix.MatrixConfig) (goledmatrix.Matrix, error) {
		return goledmatrix.BuildMatrix(config)
	}, Gameloop)
}

func Gameloop(c *goledmatrix.Canvas, done chan struct{}) {

	fmt.Println("wait for gameloop...")
	<-time.After(1000 * time.Millisecond)
	fmt.Println("gameloop !")

	var clock components.Component
	clock = NewClock(
		goledmatrix.Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		10,
	)
	world := components.NewWorld(c, []components.Component{clock})

	// TODO: call on exit :: done <- struct{}{}

	world.Run(done)
	//font := fonts.GetFont(fonts.Bdf7x13)
	//
	//var i = 0
	//var j = 0
	//
	//
	//
	//fpsTicker := time.NewTicker(frameDuration)
	////var tickerC <-chan time.Time
	//c.DrawLabel(10, 10, "Spring !", color.RGBA{R: 128, B: 128, A: 255}, font)
	//for _ = range fpsTicker.C {
	//
	//	c.Set(i, j, color.RGBA{R: 0x41, G: 0x69, B: 0xe1, A: 0xff})
	//	c.Render()
	//	i += 4
	//	if i%c.Bounds().Max.X == 0 {
	//		j += 4
	//	}
	//	if i == c.Bounds().Max.X && j == c.Bounds().Max.Y {
	//		break
	//	}
	//	i %= c.Bounds().Max.X
	//	if i == 0 {
	//		j %= c.Bounds().Max.Y
	//	}
	//}
	fmt.Println("Gameloop END")
	done <- struct{}{}
}
