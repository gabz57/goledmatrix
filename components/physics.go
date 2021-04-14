package components

import (
	"github.com/gabz57/goledmatrix"
	"time"
)

const EAST float64 = 0
const SOUTH float64 = 90
const WEST float64 = 180
const NORTH float64 = 270

type Acceleration interface {
	NextVelocity(dtInSeconds float64) (float64, float64) // dVx, dVy in Pixels per second
	Direction() float64                                  // Degrees
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

func (g ConstantAcceleration) NextVelocity(dtInSeconds float64) (float64, float64) {
	dV := RotateOrigin(goledmatrix.FloatingPoint{
		X: g.acceleration * dtInSeconds,
		Y: g.acceleration * dtInSeconds,
	}, g.direction)
	return dV.X, dV.Y
}

func (g ConstantAcceleration) Direction() float64 {
	return g.direction
}

var Gravity Acceleration = ConstantAcceleration{acceleration: 9.81, direction: SOUTH}

type Physics interface {
	// dV, dXY
	Advance(dt time.Duration) (goledmatrix.FloatingPoint, goledmatrix.FloatingPoint)
}

type Movement struct {
	dXY           goledmatrix.FloatingPoint // keep local floating dXY for accurate moves
	velocity      goledmatrix.FloatingPoint // in Pixel per second
	accelerations []Acceleration            // in Pixel per second2
}

func NewMovement(initialVelocity goledmatrix.FloatingPoint, accelerations []Acceleration) *Movement {
	return &Movement{
		dXY:           goledmatrix.FloatingPoint{},
		velocity:      initialVelocity,
		accelerations: accelerations,
	}
}

// TODO: consider max values for accelerations
// Return the next dXY to use
func (m *Movement) Advance(duration time.Duration) (goledmatrix.FloatingPoint, goledmatrix.FloatingPoint) {
	dtInSeconds := float64(duration.Nanoseconds()) / 1000000000
	for _, acceleration := range m.accelerations {
		dVX, dVY := acceleration.NextVelocity(dtInSeconds)
		m.velocity.AddXY(dVX*dtInSeconds, dVY*dtInSeconds)
	}
	m.dXY = m.dXY.AddXY(
		m.velocity.X*dtInSeconds,
		m.velocity.Y*dtInSeconds,
	)
	return m.velocity, m.dXY
}
