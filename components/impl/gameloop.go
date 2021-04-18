package impl

import (
	. "github.com/gabz57/goledmatrix/components"
	. "github.com/gabz57/goledmatrix/matrix"
	"math/rand"
	"time"
)

func Gameloop(c *Canvas, done chan struct{}) {
	infoCpnt := infoComponent(*c)
	sceneDuration := 12 * time.Second
	engine := NewEngine(c, []*Scene{
		//NewScene([]*Component{infoCpnt, octoLogoComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, octoLogoComponent(*c), clockComponent(*c)}, sceneDuration),
		NewScene([]*Component{infoCpnt, beanComponent(*c)}, sceneDuration),
		//NewScene([]*Component{marioComponent(*c), infoCpnt}, sceneDuration),
		//NewScene([]*Component{infoCpnt, movingDotComponent(*c)}, sceneDuration),
		////NewScene([]*Component{infoCpnt, heartComponent(*c)}, sceneDuration),
		////NewScene([]*Component{infoCpnt, heartsComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, movingHeartsComponent(*c)}, sceneDuration),
		////NewScene([]*Component{infoCpnt, birthdayCakeComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, happyBirthdayComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt}, sceneDuration),
	})
	engine.Run(done)
}

func infoComponent(c Canvas) *Component {
	var component Component = NewInfo(c)
	return &component
}

func octoLogoComponent(c Canvas) *Component {
	var component Component = NewOctoLogo(
		c,
		Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		20,
	)
	return &component
}

func clockComponent(c Canvas) *Component {
	var component Component = NewClock(
		c,
		Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		(c.Bounds().Max.X-10)/2,
	)
	return &component
}

func marioComponent(c Canvas) *Component {
	var component Component = NewImages(
		"img/gif.gif",
		Point{
			X: (c.Bounds().Max.X - 32) / 2,
			Y: (c.Bounds().Max.Y - 32) / 2,
		}, Point{
			X: 32,
			Y: 32,
		})
	return &component
}

func birthdayCakeComponent(c Canvas) *Component {
	var component Component = NewImages(
		"img/birthday-cake.gif",
		Point{
			X: (c.Bounds().Max.X - 100) / 2,
			Y: (c.Bounds().Max.Y - 58) / 2,
		},
		Point{
			X: 100,
			Y: 58,
		})
	return &component
}
func beanComponent(c Canvas) *Component {
	var component Component = NewBeanNDot(c)
	return &component
}

func movingDotComponent(c Canvas) *Component {
	var component Component = NewMovingDot(
		c,
		Point{
			X: rand.Intn(64),
			Y: rand.Intn(64),
		},
		FloatingPoint{
			X: Float64Between(32, 64),
			Y: Float64Between(32, 64),
		}, c.Bounds())
	return &component
}

func heartComponent(c Canvas) *Component {
	var component Component = NewHeart(
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
	var component Component = NewHearts(c, Point{}, 10)
	return &component
}

func movingHeartsComponent(c Canvas) *Component {
	var component Component = NewMovingHearts(c, Point{}, 10)
	return &component
}

func happyBirthdayComponent(c Canvas) *Component {
	var component Component = NewHappyBirthday(c)
	return &component
}
