package components

import (
	"fmt"
	"github.com/gabz57/goledmatrix"
	"github.com/gosuri/uilive"
	"github.com/paulbellamy/ratecounter"
	"strconv"
	"time"
)

type World struct {
	canvas     *goledmatrix.Canvas
	components []Component
}

func NewWorld(canvas *goledmatrix.Canvas, components []Component) World {
	return World{
		canvas:     canvas,
		components: components,
	}
}

const FrameDurationInNanos = 33333333 // 30 FPS approximated in nanos
//const FrameDurationInNanos  = 16666666
const UpdateDurationInNanos = 20000000

//const UpdateDurationInNanos = 5000000
const UpdateDuration = time.Duration(UpdateDurationInNanos)

func (w World) Run(done chan struct{}) {
	writer := uilive.New()
	// start listening for updates and render
	writer.Start()
	updateCounter := ratecounter.NewRateCounter(1 * time.Second)
	renderCounter := ratecounter.NewRateCounter(1 * time.Second)

	fmt.Println("Run")
	previous := time.Now()
	lag := time.Duration(0)
LOOP:
	for {
		_, _ = fmt.Fprintf(writer.Newline(), "Updates: "+strconv.FormatInt(updateCounter.Rate(), 10)+" updates/sec - RenderRate: "+strconv.FormatInt(renderCounter.Rate(), 10)+" FPS\n")

		select {
		case <-done:
			break LOOP
			//TODO: case e:= <-ui.event;// w.processInput()
		default:
		}
		current := time.Now()
		elapsed := current.Sub(previous)
		previous = current
		lag += elapsed

		for lag >= UpdateDuration {
			w.updateGame()
			updateCounter.Incr(1)
			lag -= UpdateDuration

			elapsed = time.Now().Sub(current)

			if elapsed < UpdateDurationInNanos {
				duration := UpdateDurationInNanos - elapsed
				time.Sleep(duration)
			}
		}

		w.render()
		renderCounter.Incr(1)

		elapsed = time.Now().Sub(current)

		if elapsed < FrameDurationInNanos {
			duration := FrameDurationInNanos - elapsed
			time.Sleep(duration)
		}
	}
	writer.Stop() // flush and stop rendering to CLI
}

func (w World) processInput() {
	//fmt.Println("processInput")
	// pipe user input events to be handled in respective component(s)
}

func (w World) updateGame() {
	//fmt.Println("updateGame")
	// update one or more components state
	for _, component := range w.components {
		component.Update()
	}
}

// Draw the world into a canvas and render its content
func (w World) render() error {
	//start := time.Now()

	//fmt.Println("render (frameProgress=" + strconv.FormatInt(frameProgress, 10) + ")")
	w.canvas.Clear()
	for _, component := range w.components {
		err := component.Draw(w.canvas)
		if err != nil {
			return err
		}
	}
	err := w.canvas.Render()
	//fmt.Println("w.render took " + strconv.FormatInt(time.Now().Sub(start).Milliseconds(), 10) + "ms")
	return err
}
