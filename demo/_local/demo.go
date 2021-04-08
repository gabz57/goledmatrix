package main

import (
	"github.com/gabz57/goledmatrix"
	"github.com/gabz57/goledmatrix/components"
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
	engine := components.NewEngine(c, []components.Component{
		clockComponent(c),
		infoComponent(c),
	})
	engine.Run(done)
}

func infoComponent(c *goledmatrix.Canvas) components.Component {
	return NewInfo(c)
}

func clockComponent(c *goledmatrix.Canvas) components.Component {
	return NewClock(
		goledmatrix.Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		(c.Bounds().Max.X-10)/2,
	)
}
