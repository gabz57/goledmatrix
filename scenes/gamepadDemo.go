package scenes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"time"
)

type GamePadDemo struct {
	shape      *CompositeDrawable
	cross      *shapes.Cross
	controller GamepadDemoController
}

func NewGamePadDemoComponent(canvas Canvas) *GamePadDemo {
	c := GamePadDemo{shape: NewCompositeDrawable(NewGraphic(nil, NewLayout(ColorRed, ColorWhite)))}
	c.cross = shapes.NewCross(c.shape.Graphic, Point{
		X: canvas.Bounds().Max.X / 2,
		Y: canvas.Bounds().Max.Y / 2,
	}, 2)
	c.shape.AddDrawable(c.cross)
	c.controller = *NewGamepadDemoController(canvas.Bounds(), c.cross)
	return &c
}

func (f *GamePadDemo) Update(elapsedBetweenUpdate time.Duration) bool {
	return f.controller.Update(elapsedBetweenUpdate)
}

func (f *GamePadDemo) Draw(canvas Canvas) error {
	return f.shape.Draw(canvas)
}

func (f *GamePadDemo) Controller() SceneController {
	return &f.controller
}
