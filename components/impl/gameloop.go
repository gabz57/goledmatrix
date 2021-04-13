package impl

import (
	"github.com/gabz57/goledmatrix"
	"github.com/gabz57/goledmatrix/components"
)

func Gameloop(c *goledmatrix.Canvas, done chan struct{}) {
	//clock := clockComponent(*c)
	//movingDot := movingDotComponent(*c)
	//octoLogo := octoLogoComponent(*c)
	heart := heartComponent(*c)
	info := infoComponent(*c)
	engine := components.NewEngine(c, []*components.Component{
		//&octoLogo,
		//&clock,
		//&movingDot,
		&heart,
		&info,
	})
	engine.Run(done)
}

func infoComponent(c goledmatrix.Canvas) components.Component {
	return NewInfo(c)
}

func heartComponent(c goledmatrix.Canvas) components.Component {
	return NewHeart(goledmatrix.Point{
		X: c.Bounds().Max.X / 5,
		Y: c.Bounds().Max.Y / 2,
	})
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
