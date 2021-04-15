package impl

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"math/rand"
	"time"
)

func Gameloop(c *Canvas, done chan struct{}) {
	infoCpnt := infoComponent(*c)
	sceneDuration := 12 * time.Second
	engine := NewEngine(c, []*Scene{
		NewScene([]*Component{infoCpnt, octoLogoComponent(*c)}, sceneDuration),
		NewScene([]*Component{infoCpnt, octoLogoComponent(*c), clockComponent(*c)}, sceneDuration),
		NewScene([]*Component{infoCpnt, movingDotComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, heartComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, heartsComponent(*c)}, sceneDuration),
		NewScene([]*Component{infoCpnt, movingHeartsComponent(*c)}, sceneDuration),
		NewScene([]*Component{infoCpnt, happyBirthdayComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt}, sceneDuration),
	})
	engine.Run(done)
}

func infoComponent(c Canvas) *Component {
	var component Component
	component = NewInfo(c)
	return &component
}

func heartComponent(c Canvas) *Component {
	var component Component
	component = NewHeart(
		c,
		nil,
		Point{
			X: c.Bounds().Max.X / 5,
			Y: c.Bounds().Max.Y / 2,
		},
		time.Duration(rand.Int63n(5000))*time.Millisecond,
		rand.Float64(),
		false)
	return &component
}

func heartsComponent(c Canvas) *Component {
	var component Component
	component = NewHearts(c, Point{}, 10)
	return &component
}

func happyBirthdayComponent(c Canvas) *Component {
	var component Component
	component = NewHappyBirthday(c)
	return &component
}

func movingHeartsComponent(c Canvas) *Component {
	var component Component
	component = NewMovingHearts(c, Point{}, 10)
	return &component
}

func clockComponent(c Canvas) *Component {
	var component Component
	component = NewClock(
		c,
		Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		(c.Bounds().Max.X-10)/2,
	)
	return &component
}

func movingDotComponent(c Canvas) *Component {
	var component Component
	component = NewMovingDot(
		c,
		Point{
			X: rand.Intn(64),
			Y: rand.Intn(64),
		},
		FloatingPoint{
			X: Float64Between(32, 64),
			Y: Float64Between(32, 64),
		},
		c.Bounds(),
	)
	return &component
}

func octoLogoComponent(c Canvas) *Component {
	var component Component
	component = NewOctoLogo(
		c,
		Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		20,
	)
	return &component
}
