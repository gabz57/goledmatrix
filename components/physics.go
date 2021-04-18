package components

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"time"
)

const RIGHT float64 = 0
const BOTTOM float64 = 90
const LEFT float64 = 180
const TOP float64 = 270

type Acceleration interface {
	// Returns (dVx, dVy) in Pixels per second
	NextVelocityDelta(dtInSeconds float64) FloatingPoint
	// Direction in Degrees, reference is RIGHT
	Direction() float64
}

type ConstantAcceleration struct {
	acceleration float64 // Pixels per second^2
	direction    float64 // Degrees
}

func NewConstantAcceleration(acceleration, direction float64) *ConstantAcceleration {
	return &ConstantAcceleration{
		acceleration: acceleration,
		direction:    direction,
	}
}

func (a ConstantAcceleration) NextVelocityDelta(dtInSeconds float64) FloatingPoint {
	return RotateOrigin(FloatingPoint{
		X: a.acceleration * dtInSeconds,
		Y: 0,
	}, a.direction)
}

func (a ConstantAcceleration) Direction() float64 {
	return a.direction
}

func (a ConstantAcceleration) SetAcceleration(acceleration float64) {
	a.acceleration = acceleration
}

func (a ConstantAcceleration) SetDirection(direction float64) {
	a.direction = direction
}

var Gravity Acceleration = ConstantAcceleration{acceleration: 9.81, direction: BOTTOM}

type Physics interface {
	// Compute and return the next position and velocity
	NextPosition(dt time.Duration) (FloatingPoint, FloatingPoint)
}

type Movement struct {
	initialPosition FloatingPoint   // keep local floating dXY for accurate moves
	dXY             FloatingPoint   // keep local floating dXY for accurate moves
	velocity        FloatingPoint   // in Pixel per second
	accelerations   *[]Acceleration // in Pixel per second2
}

func NewMovement(initialPosition, initialVelocity FloatingPoint, accelerations *[]Acceleration) *Movement {
	return &Movement{
		initialPosition: initialPosition,
		dXY:             FloatingPoint{},
		velocity:        initialVelocity,
		accelerations:   accelerations,
	}
}

// TODO: consider max values for accelerations
// Return the next position to use
func (m *Movement) NextPosition(duration time.Duration) (FloatingPoint, FloatingPoint) {
	dtInSeconds := float64(duration.Nanoseconds()) / 1000000000
	for _, acceleration := range *m.accelerations {
		m.velocity = m.velocity.Add(acceleration.NextVelocityDelta(dtInSeconds))
	}
	m.dXY = m.dXY.AddXY(
		m.velocity.X*dtInSeconds,
		m.velocity.Y*dtInSeconds,
	)
	return m.initialPosition.Add(m.dXY), m.velocity
}

func (m *Movement) SetVelocity(velocity FloatingPoint) {
	m.velocity = velocity
}

func (m *Movement) Velocity() FloatingPoint {
	return m.velocity
}
