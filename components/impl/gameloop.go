package impl

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"math/rand"
	"time"
)

func Gameloop(c *Canvas, done chan struct{}) {
	octoLogo := octoLogoComponent(*c)
	//clock := clockComponent(*c)
	movingDot := movingDotComponent(*c)
	//heart := heartComponent(*c)
	//hearts := heartsComponent(*c)
	info := infoComponent(*c)
	engine := NewEngine(c, []*Component{
		&octoLogo,
		//&clock,
		&movingDot,
		//&heart,
		//&hearts,
		&info,
	})
	engine.Run(done)
}

func infoComponent(c Canvas) Component {
	return NewInfo(c)
}

func heartComponent(c Canvas) Component {
	return NewHeart(
		c,
		Point{
			X: c.Bounds().Max.X / 5,
			Y: c.Bounds().Max.Y / 2,
		},
		time.Duration(rand.Int63n(5000))*time.Millisecond,
		rand.Float64(),
		false)
}

func heartsComponent(c Canvas) Component {
	return NewHearts(c, Point{}, 10)
}

func clockComponent(c Canvas) Component {
	return NewClock(
		c,
		Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		(c.Bounds().Max.X-10)/2,
	)
}

func movingDotComponent(c Canvas) Component {
	return NewMovingDot(
		c,
		Point{
			X: 0,
			Y: 0,
		},
		FloatingPoint{
			X: 25,
			Y: 25,
		},
		c.Bounds(),
	)
}

func octoLogoComponent(c Canvas) Component {
	return NewOctoLogo(
		c,
		Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		20,
	)
}
