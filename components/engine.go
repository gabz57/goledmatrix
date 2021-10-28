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
	Scene struct {
		components []*Component
		duration   time.Duration
	}
)

func NewScene(component []*Component, duration time.Duration) *Scene {
	return &Scene{
		components: component,
		duration:   duration,
	}
}

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

func (s *Scene) Update(elapsedBetweenUpdate time.Duration) bool {
	dirtyScene := false
	for _, component := range s.components {
		dirtyScene = (*component).Update(elapsedBetweenUpdate) || dirtyScene
	}
	return dirtyScene
}

func (s *Scene) Render(c *canvas.Canvas) error {
	(*c).Clear()
	for _, component := range s.components {
		err := (*component).Draw(*c)
		if err != nil {
			return err
		}
	}
	return (*c).Render()
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
