package main

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/effect"
	"github.com/gabz57/goledmatrix/components/impl"
	"github.com/gabz57/goledmatrix/scenes"
	"time"
)

func Gameloop(c *Canvas, done chan struct{}) {
	//infoCpnt := infoComponent(*c)
	sceneDuration := 12 * time.Second
	var fadeEffect CanvasEffect = effect.NewFadeInOutSceneEffect(&sceneDuration)
	//photoGallery := photoGalleryComponent(*c)
	engine := NewEngine(c, []*Scene{
		//NewScene([]*Component{infoCpnt, octoLogoComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, octoLogoComponent(*c), clockComponent(*c)}, sceneDuration),
		//NewSceneWithEffect([]*Component{ /*infoCpnt, */ clockComponent(*c)}, sceneDuration, []*CanvasEffect{&fadeEffect}),
		//NewSceneWithEffect([]*Component{infoCpnt, photoGalleryComponent(*c)}, sceneDuration, []*CanvasEffect{&fadeEffect}),
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
		//NewSceneWithEffect([]*Component{ /*infoCpnt, */ meteoLocalComponent(*c)}, sceneDuration, []*CanvasEffect{&fadeEffect}),
		gamepadDemo(c, fadeEffect),
		//NewSceneWithEffect([]*Component{ /*infoCpnt, */ meteoForecastComponent(*c, "94016")}, sceneDuration, []*CanvasEffect{&fadeEffect}), // Cachan
		//NewSceneWithEffect([]*Component{ /*infoCpnt, */ photoGallery}, sceneDuration, []*CanvasEffect{&fadeEffect}),
		//NewSceneWithEffect([]*Component{ /*infoCpnt, */ meteoForecastComponent(*c, "57176")}, sceneDuration, []*CanvasEffect{&fadeEffect}), // Diebling
		//NewSceneWithEffect([]*Component{ /*infoCpnt, */ photoGallery}, sceneDuration, []*CanvasEffect{&fadeEffect}),
		//NewSceneWithEffect([]*Component{ /*infoCpnt, */ meteoForecastComponent(*c, "75112")}, sceneDuration, []*CanvasEffect{&fadeEffect}), // Paris 12 arr
		//NewSceneWithEffect([]*Component{ /*infoCpnt, */ photoGallery}, sceneDuration, []*CanvasEffect{&fadeEffect}),
		//NewScene([]*Component{meteoIconsComponent(*c)}, sceneDuration),
		//NewScene([]*Component{meteoIcons16Component(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt, photoComponent(*c)}, sceneDuration),
		//NewSceneWithEffect([]*Component{infoCpnt, nextAnniversariesComponent(*c)}, sceneDuration, []*CanvasEffect{&fadeEffect}),
		////NewScene([]*Component{infoCpnt, focusComponent(*c)}, sceneDuration),
		//NewScene([]*Component{infoCpnt}, sceneDuration),
	})
	engine.Run(done)
}

func gamepadDemo(c *Canvas, fadeEffect CanvasEffect) *Scene {
	gamepadDemoComponent := scenes.NewGamePadDemoComponent(*c)
	return NewControlledScene([]Component{ /*infoCpnt, */ gamepadDemoComponent}, []*CanvasEffect{&fadeEffect}, gamepadDemoComponent.Controller())
}

func infoComponent(c Canvas) *Component {
	var component Component = impl.NewInfo(c)
	return &component
}

func octoLogoComponent(c Canvas) *Component {
	var component Component = impl.NewOctoLogo(
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
	var component Component = impl.NewClock(
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
	var component Component = impl.NewImages(
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
	var component Component = impl.NewImages(
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
	var component Component = impl.NewBeanNDot(c)
	return &component
}

func movingDotComponent(c Canvas) *Component {
	var component Component = impl.NewMovingDot(
		c,
		Point{
			X: Random.Intn(64),
			Y: Random.Intn(64),
		},
		FloatingPoint{
			X: Float64Between(32, 64),
			Y: Float64Between(32, 64),
		}, c.Bounds())
	return &component
}

func bouncingDotComponent(c Canvas, x int, bottomAcceleration float64) *Component {
	var component Component = impl.NewBouncingDot(
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
	var component Component = impl.NewHeart(
		c,
		nil,
		Point{
			X: c.Bounds().Max.X / 5,
			Y: c.Bounds().Max.Y / 2,
		},
		time.Duration(Random.Int63n(5000))*time.Millisecond,
		Random.Float64(),
		false)
	return &component
}

func heartsComponent(c Canvas) *Component {
	var component Component = impl.NewHearts(c, Point{}, 10)
	return &component
}

func movingHeartsComponent(c Canvas) *Component {
	var component Component = impl.NewMovingHearts(c, Point{}, 10)
	return &component
}

func happyBirthdayComponent(c Canvas) *Component {
	var component Component = impl.NewHappyBirthday(c)
	return &component
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////

func gamePadDemo(c Canvas) *Component {
	var component Component = scenes.NewGamePadDemoComponent(c)
	return &component
}

func meteoLocalComponent(c Canvas) *Component {
	var component Component = scenes.NewMeteoCurrentComponent(c, "94016")
	return &component
}

func meteoForecastComponent(c Canvas, insee string) *Component {
	var component Component = scenes.NewMeteoForecastComponent(c, insee)
	return &component
}
func meteoIcons16Component(c Canvas) *Component {
	var component Component = scenes.NewMeteoIcons16Component(c)
	return &component
}

func meteoIconsComponent(c Canvas) *Component {
	var component Component = scenes.NewMeteoIconsComponent(c)
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

func nextAnniversariesComponent(c Canvas) *Component {
	var component Component = scenes.NewNextAnniversariesComponent(c)
	return &component
}

func focusComponent(c Canvas) *Component {
	var component Component = scenes.NewFocusComponent(c)
	return &component
}
