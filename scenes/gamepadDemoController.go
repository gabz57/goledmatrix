package scenes

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/controller"
	"image"
	"time"
)

type GamepadDemoController struct {
	positionable canvas.Positionable
	move         *components.Movement
}

func NewGamepadDemoController(bounds image.Rectangle, positionable canvas.Positionable) *GamepadDemoController {
	return &GamepadDemoController{
		positionable: positionable,
		move: components.NewBoundedMovement(
			positionable.GetPosition().Floating(),
			canvas.FloatingPoint{},
			nil,
			&bounds),
	}
}

func (c *GamepadDemoController) HandleKeyboardEvent(event *controller.KeyboardEvent, projection *controller.KeyboardProjection) {

}

func (c *GamepadDemoController) HandleGamepadEvent(event *controller.GamepadEvent, projection *controller.GamepadProjection) {
	if event.IsDPad() {
		dpadMove(c.move, projection.DPadDirection(), 10.)
	} else if event.Name == controller.EventTypeLeftStick {
		stickMove(c.move, projection.LeftStick, 150.)
	} else if event.Name == controller.EventTypeRightStick {
		stickMove(c.move, projection.RightStick, 300.)
	}
}

func (c *GamepadDemoController) Update(elapsedBetweenUpdate time.Duration) bool {
	nextPosition, _ := c.move.NextPosition(elapsedBetweenUpdate)
	next := nextPosition.Int()
	current := c.positionable.GetPosition()
	if current.X != next.X || current.Y != next.Y {
		c.positionable.SetPosition(next)
		return true
	}
	return false
}

func dpadMove(move *components.Movement, direction *float64, pixelPerSecond float64) {
	if direction != nil {
		directionFP := components.DirectionToFloatingPoint(*direction)
		directionFPCorrected := canvas.FloatingPoint{}
		if directionFP.X > 0.1 {
			directionFPCorrected.X = 1
		} else if directionFP.X < -0.1 {
			directionFPCorrected.X = -1
		}
		if directionFP.Y > 0.1 {
			directionFPCorrected.Y = 1
		} else if directionFP.Y < -0.1 {
			directionFPCorrected.Y = -1
		}
		move.SetVelocity(toVelocity(pixelPerSecond, directionFPCorrected))
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
	// remove noise around (0,0)
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
