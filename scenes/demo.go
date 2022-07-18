package scenes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/canvas/effect"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/impl"
	"github.com/gabz57/goledmatrix/controller"
	"time"
)

var infoCpnt Component

var sceneDuration = 30 * time.Second
var shortSceneDuration = 15 * time.Second
var effects = []effect.DynamicEffect{effect.NewFadeInOutSceneEffect(sceneDuration)}

func Demos(c Canvas, done chan struct{}, keyboardChannel *controller.KeyboardEventChannel) {
	infoCpnt = infoComponent(c, false)
	//galleryScene := photoGalleryScene(c, sceneDuration).WithEffects(effects)
	engine := NewEngine(c, []*Scene{
		//NewScene([]Component{infoCpnt, octoLogoComponent(c)}, sceneDuration),
		//NewScene([]Component{infoCpnt, octoLogoComponent(c), clockComponent(c)}, sceneDuration),
		fadingLinesScene(c, sceneDuration).WithEffects(effects),
		fadingDotsScene(c, sceneDuration, 100).WithEffects(effects),
		fadingDotsScene(c, sceneDuration, 1000), //.WithEffects(effects),
		NewScene([]Component{infoCpnt, dspComponent(c)}, sceneDuration),
		NewScene([]Component{infoCpnt, clockComponent(c)}, sceneDuration),
		//NewScene([]Component{ /*infoCpnt, */ clockComponent(c)}, sceneDuration).WithEffects(effects),
		//NewSceneWithEffect([]Component{infoCpnt, photoGalleryComponent(c)}, sceneDuration, []CanvasEffect{fadeEffect}),
		//NewScene([]Component{infoCpnt, beanComponent(c)}, sceneDuration),
		//NewScene([]Component{marioComponent(c), infoCpnt}, sceneDuration),
		//NewScene([]Component{infoCpnt, movingDotComponent(c)}, sceneDuration),
		//NewScene([]Component{
		//	infoCpnt,
		//	bouncingDotComponent(c, 10, 90),
		//	bouncingDotComponent(c, 20, 90),
		//	bouncingDotComponent(c, 30, 90),
		//	bouncingDotComponent(c, 40, 90),
		//	bouncingDotComponent(c, 50, 90),
		//	bouncingDotComponent(c, 60, 90),
		//	bouncingDotComponent(c, 70, 90),
		//	bouncingDotComponent(c, 80, 90),
		//	bouncingDotComponent(c, 90, 90),
		//}, sceneDuration),
		//NewScene([]Component{infoCpnt, heartComponent(c)}, sceneDuration),
		//NewScene([]Component{infoCpnt, heartsComponent(c)}, sceneDuration),
		//NewScene([]Component{infoCpnt, birthdayCakeComponent(c)}, sceneDuration),
		NewScene([]Component{infoCpnt, movingHeartsComponent(c)}, sceneDuration),
		//NewScene([]Component{infoCpnt, happyBirthdayComponent(c)}, sceneDuration),

		//gamepadDemoScene(c, effects),

		//meteoLocalScene(c, shortSceneDuration),
		//meteoLocalScene(c, shortSceneDuration).WithEffects(effects),
		//meteoForecastScene(c, shortSceneDuration, "94016"), // Cachan
		meteoForecastScene(c, shortSceneDuration, "94016").WithEffects(effects), // Cachan
		//galleryScene,
		//meteoForecastScene(c, shortSceneDuration, "57176"), // Diebling
		meteoForecastScene(c, shortSceneDuration, "57176").WithEffects(effects), // Diebling
		//galleryScene,
		//meteoForecastScene(c, shortSceneDuration, "75112").WithEffects(effects), // Paris 12 arr
		//galleryScene,
		//nextAnniversariesScene(c, shortSceneDuration),
		nextAnniversariesScene(c, shortSceneDuration).WithEffects(effects),

		//NewScene([]Component{meteoIconsComponent(c)}, sceneDuration),
		//NewScene([]Component{meteoIcons16Component(c)}, sceneDuration),
		//NewScene([]Component{ /*infoCpnt, */ photoComponent(c)}, sceneDuration),
		////NewScene([]Component{infoCpnt, focusComponent(c)}, sceneDuration),
		//NewScene([]Component{infoCpnt}, sceneDuration),
	})
	engine.Run(done)
}

func fadingLinesScene(c Canvas, sceneDuration time.Duration) *Scene {
	return NewScene([]Component{impl.NewFadingLines(c, Point{}, 15)}, sceneDuration)
}

func fadingDotsScene(c Canvas, sceneDuration time.Duration, nbDots int) *Scene {
	return NewScene([]Component{impl.NewFadingDots(c, Point{}, nbDots, nil)}, sceneDuration)
}

func nextAnniversariesScene(c Canvas, sceneDuration time.Duration) *Scene {
	return NewScene([]Component{infoCpnt, NewNextAnniversariesComponent(c)}, sceneDuration)
}

func photoGalleryScene(c Canvas, sceneDuration time.Duration) *Scene {
	return NewScene([]Component{infoCpnt, NewPhotoGalleryComponent(c)}, sceneDuration)
}

func meteoForecastScene(c Canvas, sceneDuration time.Duration, insee string) *Scene {
	return NewScene([]Component{infoCpnt, NewMeteoForecastComponent(c, insee)}, sceneDuration)
}

func meteoLocalScene(c Canvas, sceneDuration time.Duration) *Scene {
	return NewScene([]Component{infoCpnt, NewMeteoCurrentComponent(c, "94016")}, sceneDuration)
}

const maxDuration time.Duration = 1<<63 - 1

func gamepadDemoScene(c Canvas) *Scene {
	gamepadDemoComponent := NewGamePadDemoComponent(c)
	return NewScene([]Component{ /*infoCpnt, */ gamepadDemoComponent}, maxDuration).WithController(gamepadDemoComponent.Controller())
}

func infoComponent(c Canvas, enabled bool) Component {
	return impl.NewInfo(c, enabled)
}

func octoLogoComponent(c Canvas) Component {
	return impl.NewOctoLogo(
		c,
		Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		20,
	)
}

func clockComponent(c Canvas) Component {
	return impl.NewClock(
		c,
		Point{
			X: c.Bounds().Max.X / 2,
			Y: c.Bounds().Max.Y / 2,
		},
		(c.Bounds().Max.X-10)/2,
	)
}
func dspComponent(c Canvas) Component {
	return impl.NewDsp(c)
}

func marioComponent(c Canvas) Component {
	return impl.NewImages(
		"img/mario.gif",
		Point{
			X: (c.Bounds().Max.X - 32) / 2,
			Y: (c.Bounds().Max.Y - 32) / 2,
		}, Point{
			X: 32,
			Y: 32,
		})
}

func birthdayCakeComponent(c Canvas) Component {
	return impl.NewImages(
		"img/birthday-cake.gif",
		Point{
			X: (c.Bounds().Max.X - 100) / 2,
			Y: (c.Bounds().Max.Y - 58) / 2,
		},
		Point{
			X: 100,
			Y: 58,
		})
}

func beanComponent(c Canvas) Component {
	return impl.NewBeanNDot(c)
}

func movingDotComponent(c Canvas) Component {
	return impl.NewMovingDot(
		c,
		Point{
			X: Random.Intn(64),
			Y: Random.Intn(64),
		},
		FloatingPoint{
			X: Float64Between(32, 64),
			Y: Float64Between(32, 64),
		}, c.Bounds())
}

func bouncingDotComponent(c Canvas, x int, bottomAcceleration float64) Component {
	return impl.NewBouncingDot(
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
}

func heartComponent(c Canvas) Component {
	return impl.NewHeart(
		c,
		nil,
		Point{
			X: c.Bounds().Max.X / 5,
			Y: c.Bounds().Max.Y / 2,
		},
		time.Duration(Random.Int63n(5000))*time.Millisecond,
		Random.Float64(),
		false)
}

func heartsComponent(c Canvas) Component {
	return impl.NewHearts(c, Point{}, 15)
}

func movingHeartsComponent(c Canvas) Component {
	return impl.NewMovingHearts(c, Point{}, 30)
}

func happyBirthdayComponent(c Canvas) Component {
	return impl.NewHappyBirthday(c)
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////

//func meteoIcons16Component(c Canvas) Component {
//	return NewMeteoIcons16Component(c)
//}

//func meteoIconsComponent(c Canvas) Component {
//	return NewMeteoIconsComponent(c)
//}

func photoComponent(c Canvas) Component {
	return NewPhotoComponent(c)
}

func focusComponent(c Canvas) Component {
	return NewFocusComponent(c)
}
