package impl

import (
	"github.com/gabz57/goledmatrix"
	"github.com/gabz57/goledmatrix/components"
	"math/rand"
	"time"
)

func Gameloop(c *goledmatrix.Canvas, done chan struct{}) {
	clock := clockComponent(*c)
	//movingDot := movingDotComponent(*c)
	octoLogo := octoLogoComponent(*c)
	//heart := heartComponent(*c)
	hearts := heartsComponent(*c)
	info := infoComponent(*c)
	engine := components.NewEngine(c, []*components.Component{
		&octoLogo,
		&clock,
		//&movingDot,
		//&heart,
		&hearts,
		&info,
	})
	engine.Run(done)
}

func infoComponent(c goledmatrix.Canvas) components.Component {
	return NewInfo(c)
}

func heartComponent(c goledmatrix.Canvas) components.Component {
	return NewHeart(
		c,
		goledmatrix.Point{
			X: c.Bounds().Max.X / 5,
			Y: c.Bounds().Max.Y / 2,
		},
		time.Duration(rand.Int63n(5000))*time.Millisecond,
		rand.Float64(),
		false)
}

func heartsComponent(c goledmatrix.Canvas) components.Component {
	return NewHearts(c, goledmatrix.Point{}, 10)
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
