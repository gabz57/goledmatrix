package scenes

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/controller"
	"image"
	"time"
)

type GamepadDemoController struct {
	cross canvas.Positionable
	move  *components.Movement
}

func NewGamepadDemoController(bounds image.Rectangle, cross canvas.Positionable) *GamepadDemoController {
	return &GamepadDemoController{
		cross: cross,
		move: components.NewBoundedMovement(
			cross.GetPosition().Floating(),
			canvas.FloatingPoint{},
			nil,
			&bounds),
	}
}

func (c *GamepadDemoController) HandleGamepadEvent(event *controller.GamepadEvent, projection *controller.GamepadProjection) {
	if event.Name == "dpad" {
		dpadMove(c.move, projection.DPadDirection(), 10.)
	} else if event.Name == "left_stick" {
		stickMove(c.move, projection.LeftStick, 50.)
	} else if event.Name == "right_stick" {
		stickMove(c.move, projection.RightStick, 150.)
	}
}

func (c *GamepadDemoController) Update(elapsedBetweenUpdate time.Duration) bool {
	nextPosition, _ := c.move.NextPosition(elapsedBetweenUpdate)
	next := nextPosition.Int()
	current := c.cross.GetPosition()
	if current.X != next.X || current.Y != next.Y {
		c.cross.SetPosition(next)
		//f.applyNextPosition(nextPosition, nextVelocity).Int())
		return true
	}
	return false
}

func dpadMove(move *components.Movement, direction *float64, pixelPerSecond float64) {
	if direction != nil {
		directionFP := components.DirectionToFloatingPoint(*direction)
		move.SetVelocity(toVelocity(pixelPerSecond, directionFP))
	} else {
		move.SetVelocity(canvas.FloatingPoint{})
	}
}

func stickMove(move *components.Movement, stick controller.Stick, pixelPerSecond float64) {
	move.SetVelocity(toVelocity(pixelPerSecond, toDirection(stick)))
}

func toVelocity(pixelPerSecond float64, direction canvas.FloatingPoint) canvas.FloatingPoint {
	return canvas.FloatingPoint{
		X: pixelPerSecond * direction.X,
		Y: pixelPerSecond * direction.Y,
	}
}

func toDirection(stick controller.Stick) canvas.FloatingPoint {
	stickX := int(stick.X) - 128
	if stickX >= -10 && stickX < 10 {
		stickX = 0
	}
	stickY := int(stick.Y) - 128
	if stickY >= -10 && stickY < 10 {
		stickY = 0
	}
	return canvas.FloatingPoint{
		X: float64(stickX) / 255.0,
		Y: float64(stickY) / 255.0,
	}
}
