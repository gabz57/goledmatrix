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
	clock := clockComponent(*c)
	info := infoComponent(*c)
	movingDot := movingDotComponent(*c)
	octoLogo := octoLogoComponent(*c)
	engine := components.NewEngine(c, []*components.Component{
		&octoLogo,
		&clock,
		&info,
		&movingDot,
	})
	engine.Run(done)
}

func infoComponent(c goledmatrix.Canvas) components.Component {
	return NewInfo(c)
}

func clockComponent(c goledmatrix.Canvas) components.Component {
	return NewClock(
		c,
		goledmatrix.Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		(c.Bounds().Max.X-10)/2,
	)
}

func movingDotComponent(c goledmatrix.Canvas) components.Component {
	return NewMovingDot(
		c,
		goledmatrix.Point{
			X: 0,
			Y: 0,
		},
		25,
		c.Bounds(),
	)
}

func octoLogoComponent(c goledmatrix.Canvas) components.Component {
	return NewOctoLogo(
		c,
		goledmatrix.Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		20,
	)
}
