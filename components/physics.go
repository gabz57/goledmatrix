package components

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"image"
	"time"
)

const RIGHT float64 = 0
const BOTTOM_RIGHT float64 = 45
const BOTTOM float64 = 90
const BOTTOM_LEFT float64 = 135
const LEFT float64 = 180
const TOP_LEFT float64 = 225
const TOP float64 = 270
const TOP_RIGHT float64 = 315

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
	bounds          *image.Rectangle
}

func NewMovement(initialPosition, initialVelocity FloatingPoint, accelerations *[]Acceleration) *Movement {
	return &Movement{
		initialPosition: initialPosition,
		dXY:             FloatingPoint{},
		velocity:        initialVelocity,
		accelerations:   accelerations,
	}
}

func NewBoundedMovement(initialPosition, initialVelocity FloatingPoint, accelerations *[]Acceleration, bounds *image.Rectangle) *Movement {
	return &Movement{
		initialPosition: initialPosition,
		dXY:             FloatingPoint{},
		velocity:        initialVelocity,
		accelerations:   accelerations,
		bounds:          bounds,
	}
}

// TODO: consider max values for accelerations
// Return the next position to use
func (m *Movement) NextPosition(duration time.Duration) (FloatingPoint, FloatingPoint) {
	dtInSeconds := float64(duration.Nanoseconds()) / 1000000000
	if m.accelerations != nil {
		for _, acceleration := range *m.accelerations {
			m.velocity = m.velocity.Add(acceleration.NextVelocityDelta(dtInSeconds))
		}
	}
	dX := m.velocity.X * dtInSeconds
	dY := m.velocity.Y * dtInSeconds
	if m.bounds != nil {
		if int(m.initialPosition.X+m.dXY.X+dX) < m.bounds.Min.X || int(m.initialPosition.X+m.dXY.X+dX) >= m.bounds.Max.X {
			dX = 0
			m.velocity.X = 0
		}
		if int(m.initialPosition.Y+m.dXY.Y+dY) < m.bounds.Min.Y || int(m.initialPosition.Y+m.dXY.Y+dY) >= m.bounds.Max.Y {
			dY = 0
			m.velocity.Y = 0
		}
	}
	m.dXY = m.dXY.AddXY(dX, dY)

	return m.initialPosition.Add(m.dXY), m.velocity
}

func (m *Movement) SetVelocity(velocity FloatingPoint) {
	m.velocity = velocity
}

func (m *Movement) Velocity() FloatingPoint {
	return m.velocity
}
