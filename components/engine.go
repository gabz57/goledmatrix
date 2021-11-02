package components

import (
	"fmt"
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gosuri/uilive"
	"github.com/paulbellamy/ratecounter"
	"time"
)

type (
	Engine struct {
		canvas                 *canvas.Canvas
		scenes                 []*Scene
		activeScene            *Scene
		elapsedSinceSceneStart time.Duration
	}
)

func NewEngine(canvas *canvas.Canvas, scenes []*Scene) Engine {
	return Engine{
		canvas:      canvas,
		scenes:      scenes,
		activeScene: scenes[0],
	}
}

const FrameDurationInNanos = 33333333  // 30 FPS approximated in nanos
const UpdateDurationInNanos = 10000000 // 100 updates per second (to maintain physics & time), independent from FPS
const UpdateDuration = time.Duration(UpdateDurationInNanos)

func (e *Engine) Run(done chan struct{}) {
	//joystickAdaptor := joystick.NewAdaptor()
	//stick := joystick.NewDriver(joystickAdaptor, "controller/dualshock4.json",
	//)
	//
	//work := func() {
	//	// buttons
	//	stick.On(joystick.SquarePress, func(data interface{}) {
	//		fmt.Println("square_press")
	//	})
	//	stick.On(joystick.SquareRelease, func(data interface{}) {
	//		fmt.Println("square_release")
	//	})
	//	stick.On(joystick.TrianglePress, func(data interface{}) {
	//		fmt.Println("triangle_press")
	//	})
	//	stick.On(joystick.TriangleRelease, func(data interface{}) {
	//		fmt.Println("triangle_release")
	//	})
	//	stick.On(joystick.CirclePress, func(data interface{}) {
	//		fmt.Println("circle_press")
	//	})
	//	stick.On(joystick.CircleRelease, func(data interface{}) {
	//		fmt.Println("circle_release")
	//	})
	//	stick.On(joystick.XPress, func(data interface{}) {
	//		fmt.Println("x_press")
	//	})
	//	stick.On(joystick.XRelease, func(data interface{}) {
	//		fmt.Println("x_release")
	//	})
	//	stick.On(joystick.StartPress, func(data interface{}) {
	//		fmt.Println("start_press")
	//	})
	//	stick.On(joystick.StartRelease, func(data interface{}) {
	//		fmt.Println("start_release")
	//	})
	//	stick.On(joystick.SelectPress, func(data interface{}) {
	//		fmt.Println("select_press")
	//	})
	//	stick.On(joystick.SelectRelease, func(data interface{}) {
	//		fmt.Println("select_release")
	//	})
	//	stick.On(joystick.ShareRelease, func(data interface{}) {
	//		fmt.Println("share_press")
	//	})
	//	stick.On(joystick.ShareRelease, func(data interface{}) {
	//		fmt.Println("share_release")
	//	})
	//	stick.On(joystick.OptionsPress, func(data interface{}) {
	//		fmt.Println("options_press")
	//	})
	//	stick.On(joystick.OptionsRelease, func(data interface{}) {
	//		fmt.Println("options_release")
	//	})
	//	stick.On(joystick.HomePress, func(data interface{}) {
	//		fmt.Println("home_press")
	//	})
	//	stick.On(joystick.HomeRelease, func(data interface{}) {
	//		fmt.Println("home_release")
	//	})
	//	stick.On(joystick.UpPress, func(data interface{}) {
	//		fmt.Println("up_press")
	//	})
	//	stick.On(joystick.UpRelease, func(data interface{}) {
	//		fmt.Println("up_release")
	//	})
	//	stick.On(joystick.DownPress, func(data interface{}) {
	//		fmt.Println("down_press")
	//	})
	//	stick.On(joystick.DownRelease, func(data interface{}) {
	//		fmt.Println("down_release")
	//	})
	//	stick.On(joystick.LeftPress, func(data interface{}) {
	//		fmt.Println("left_press")
	//	})
	//	stick.On(joystick.LeftRelease, func(data interface{}) {
	//		fmt.Println("left_release")
	//	})
	//	stick.On(joystick.RightPress, func(data interface{}) {
	//		fmt.Println("right_press")
	//	})
	//	stick.On(joystick.RightRelease, func(data interface{}) {
	//		fmt.Println("right_release")
	//	})
	//	//// joysticks
	//	//stick.On(joystick.LeftX, func(data interface{}) {
	//	//	fmt.Println("left_x", data)
	//	//})
	//	//stick.On(joystick.LeftY, func(data interface{}) {
	//	//	fmt.Println("left_y", data)
	//	//})
	//	//stick.On(joystick.RightX, func(data interface{}) {
	//	//	fmt.Println("right_x", data)
	//	//})
	//	//stick.On(joystick.RightY, func(data interface{}) {
	//	//	fmt.Println("right_y", data)
	//	//})
	//	stick.On(joystick.L2, func(data interface{}) {
	//		fmt.Println("l2", data)
	//	})
	//	stick.On(joystick.R2, func(data interface{}) {
	//		fmt.Println("r2", data)
	//	})
	//
	//	// triggers
	//	stick.On(joystick.R1Press, func(data interface{}) {
	//		fmt.Println("R1Press", data)
	//	})
	//	stick.On(joystick.R2Press, func(data interface{}) {
	//		fmt.Println("R2Press", data)
	//	})
	//	stick.On(joystick.L1Press, func(data interface{}) {
	//		fmt.Println("L1Press", data)
	//	})
	//	stick.On(joystick.L2Press, func(data interface{}) {
	//		fmt.Println("L2Press", data)
	//	})
	//}
	//
	//robot := gobot.NewRobot("joystickBot",
	//	[]gobot.Connection{joystickAdaptor},
	//	[]gobot.Device{stick},
	//	work,
	//)
	//
	//robot.Start()
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()
	updateCounter := ratecounter.NewRateCounter(1 * time.Second)
	renderCounter := ratecounter.NewRateCounter(1 * time.Second)

	previous := time.Now()
	lag := time.Duration(0)
	fmt.Println("Starting Canvas Engine !")
	dirty := true

LOOP:
	for {
		// print and updates to console without writing on a new line each time
		//_, _ = fmt.Fprintf(writer.Newline(), "Updates: "+strconv.FormatInt(updateCounter.Rate(), 10)+" update/sec - RenderRate: "+strconv.FormatInt(renderCounter.Rate(), 10)+" FPS\n")

		select {
		case <-done:
			break LOOP
			//TODO: plug user event
			//case e:= <-ui.event;// e.processInput()
		default:
		}
		current := time.Now()
		elapsed := current.Sub(previous)
		previous = current
		lag += elapsed

		if e.elapsedSinceSceneStart > e.activeScene.duration {
			e.runNextScene()
		}
		// using lag to catch up missing updates when UI renders to slow
		for lag >= UpdateDuration {
			dirty = e.updateGame(UpdateDuration) || dirty
			updateCounter.Incr(1)
			lag -= UpdateDuration
			e.elapsedSinceSceneStart += UpdateDuration
			if lag >= UpdateDuration {
				continue
			}
			select {
			case <-time.After(UpdateDurationInNanos - time.Now().Sub(current)):
			}
		}

		if dirty == true {
			dirty = false
			e.render()
			renderCounter.Incr(1)
		}

		select {
		case <-time.After(FrameDurationInNanos - time.Now().Sub(current)):
		}
	}
}

func (e *Engine) processInput() {
	// pipe user input events to be handled in respective component(s)
}

func (e *Engine) updateGame(elapsedBetweenUpdate time.Duration) bool {
	return e.activeScene.Update(elapsedBetweenUpdate)
}

// Draw the components into the canvas and renders its content
func (e *Engine) render() error {
	return e.activeScene.Render(e.canvas)
}

func (e *Engine) runNextScene() {
	for i, scene := range e.scenes {
		if scene == e.activeScene {
			e.activeScene = e.scenes[(i+1)%len(e.scenes)]
			e.elapsedSinceSceneStart = 0
			break
		}
	}
}
