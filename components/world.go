package components

import (
	"fmt"
	"github.com/gabz57/goledmatrix"
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

const FrameDurationInNanos = 33333333  // 30 FPS approximated in nanos
const UpdateDurationInNanos = 10000000 // 30 FPS approximated in nanos
//const UpdateDurationInNanos = 33333333 // 30 FPS approximated in nanos
const UpdateDuration = time.Duration(UpdateDurationInNanos)

func (w World) Run(done chan struct{}) {
	fmt.Println("Run")
	previous := time.Now()
	lag := time.Duration(0)
LOOP:
	for {
		select {
		case <-done:
			break LOOP
		default:
		}
		current := time.Now()
		elapsed := current.Sub(previous)
		previous = current
		lag += elapsed

		w.processInput()
		//w.updateGame()
		//
		for lag >= UpdateDuration {
			//start := time.Now()
			w.updateGame()
			lag -= UpdateDuration
			//fmt.Println("w.update took " + strconv.FormatInt(time.Now().Sub(start).Milliseconds(), 10) + "ms")
		}
		//fmt.Println("updated after " + strconv.FormatInt(time.Now().Sub(current).Milliseconds(), 10) + "ms")

		w.render()

		elapsed = time.Now().Sub(current)

		if elapsed < FrameDurationInNanos {
			duration := FrameDurationInNanos - elapsed
			//fmt.Println(">sleep " + strconv.FormatInt(duration.Milliseconds(), 10) + "ms... (elapsed "+ strconv.FormatInt(elapsed.Milliseconds(), 10)+"ms)")
			time.Sleep(duration)
		}
	}
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
