package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/scenes"
	"math/rand"
	"time"
)

func Gameloop(c *Canvas, done chan struct{}) {
	infoCpnt := infoComponent(*c)
	sceneDuration := 12 * time.Second
	engine := NewEngine(c, []*Scene{
		//NewScene([]*Component{infoCpnt, octoLogoComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, octoLogoComponent(*c), clockComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, clockComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, beanComponent(*c)}, sceneDuration),
		//NewScene([]*Component{marioComponent(*c), infoCpnt}, sceneDuration),
		//NewScene([]*Component{infoCpnt, movingDotComponent(*c)}, sceneDuration),
		//NewScene([]*Component{
		//	infoCpnt,
		//	//bouncingDotComponent(*c, 10, 90),
		//	//bouncingDotComponent(*c, 20, 90),
		//	//bouncingDotComponent(*c, 30, 90),
		//	//bouncingDotComponent(*c, 40, 90),
		//	bouncingDotComponent(*c, 50, 90),
		//	//bouncingDotComponent(*c, 60, 90),
		//	//bouncingDotComponent(*c, 70, 90),
		//	//bouncingDotComponent(*c, 80, 90),
		//	//bouncingDotComponent(*c, 90, 90),
		//}, sceneDuration),
		//NewScene([]*Component{infoCpnt, heartComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, heartsComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, birthdayCakeComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, movingHeartsComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, happyBirthdayComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, meteoLocalComponent(*c)}, sceneDuration),
		NewScene([]*Component{infoCpnt, meteoForecastComponent(*c)}, sceneDuration),
		////NewScene([]*Component{infoCpnt, photoComponent(*c)}, sceneDuration),
		////NewScene([]*Component{infoCpnt, photoGalleryComponent(*c)}, sceneDuration),
		////NewScene([]*Component{infoCpnt, nextBirthdaysComponent(*c)}, sceneDuration),
		////NewScene([]*Component{infoCpnt, focusComponent(*c)}, sceneDuration),
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
		"img/mario.gif",
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

func bouncingDotComponent(c Canvas, x int, bottomAcceleration float64) *Component {
	var component Component = NewBouncingDot(
		c,
		Point{
			X: x,
			Y: 10,
		},
		FloatingPoint{
			X: 0,
			Y: 0,
		},
		bottomAcceleration,
		c.Bounds())
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

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////

func meteoLocalComponent(c Canvas) *Component {
	var component Component = scenes.NewMeteoCurrentComponent(c, "94016")
	return &component
}

func meteoForecastComponent(c Canvas) *Component {
	var component Component = scenes.NewMeteoForecastComponent(c, "94016")
	return &component
}

func photoComponent(c Canvas) *Component {
	var component Component = scenes.NewPhotoComponent(c)
	return &component
}

func photoGalleryComponent(c Canvas) *Component {
	var component Component = scenes.NewPhotoGalleryComponent(c)
	return &component
}

func nextBirthdaysComponent(c Canvas) *Component {
	var component Component = scenes.NewNextBirthdaysComponent(c)
	return &component
}

func focusComponent(c Canvas) *Component {
	var component Component = scenes.NewFocusComponent(c)
	return &component
}
