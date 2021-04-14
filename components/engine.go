package components

import (
	"fmt"
	"github.com/gabz57/goledmatrix"
	"github.com/gosuri/uilive"
	"github.com/paulbellamy/ratecounter"
	"time"
)

type Engine struct {
	canvas     *goledmatrix.Canvas
	components []*Component
}

func NewEngine(canvas *goledmatrix.Canvas, components []*Component) Engine {
	return Engine{
		canvas:     canvas,
		components: components,
	}
}

const FrameDurationInNanos = 33333333  // 30 FPS approximated in nanos
const UpdateDurationInNanos = 10000000 // 100 updates per second (to maintain physics & time), independent from FPS
const UpdateDuration = time.Duration(UpdateDurationInNanos)

func (e *Engine) Run(done chan struct{}) {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()
	updateCounter := ratecounter.NewRateCounter(1 * time.Second)
	renderCounter := ratecounter.NewRateCounter(1 * time.Second)

	previous := time.Now()
	lag := time.Duration(0)
	fmt.Println("Starting Canvas Engine !")
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

		// using lag to catch up missing updates when UI renders to slow
		for lag >= UpdateDuration {
			e.updateGame(UpdateDuration)
			updateCounter.Incr(1)
			lag -= UpdateDuration

			select {
			case <-time.After(UpdateDurationInNanos - time.Now().Sub(current)):
			}
		}

		e.render()
		renderCounter.Incr(1)

		select {
		case <-time.After(FrameDurationInNanos - time.Now().Sub(current)):
		}
	}
}

func (e *Engine) processInput() {
	// pipe user input events to be handled in respective component(s)
}

func (e *Engine) updateGame(elapsedBetweenUpdate time.Duration) {
	for _, component := range e.components {
		(*component).Update(elapsedBetweenUpdate)
	}
}

// Draw the components into the canvas and renders its content
func (e *Engine) render() error {
	(*e.canvas).Clear()
	for _, component := range e.components {
		err := (*component).Draw(*e.canvas)
		if err != nil {
			return err
		}
	}
	err := (*e.canvas).Render()
	return err
}
