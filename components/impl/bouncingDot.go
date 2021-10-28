package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"image"
	"time"
)

type BouncingDot struct {
	move *Movement
	dot  *shapes.Dot
	//dotAcceleration        *ConstantAcceleration
	bounds                 image.Rectangle
	elapsedSinceSceneStart time.Duration
}

var groundReaction Acceleration = NewConstantAcceleration(9.81, TOP)
var noReaction Acceleration = NewConstantAcceleration(0, 0)

func NewBouncingDot(c Canvas, initialPosition Point, initialVelocity FloatingPoint, initialBottomAcceleration float64, bounds image.Rectangle) *BouncingDot {
	var accs []Acceleration
	//acceleration := NewConstantAcceleration(0, 0)
	//accs = append(accs, *acceleration)
	accs = append(accs, NewConstantAcceleration(initialBottomAcceleration, BOTTOM))

	dot := BouncingDot{
		move: NewMovement(
			initialPosition.Floating(),
			FloatingPoint{
				X: initialVelocity.X,
				Y: initialVelocity.Y,
			},
			&accs,
		),
		dot: shapes.NewDot(
			NewGraphic(nil, nil),
			initialPosition,
		),
		bounds: bounds,
	}
	//dot.dotAcceleration = acceleration
	return &dot

}

func (m *BouncingDot) Update(elapsedBetweenUpdate time.Duration) bool {
	// changing erratically of direction by changing the velocity
	m.elapsedSinceSceneStart += elapsedBetweenUpdate

	// advance the position by one step, make it bounce on bounds with exact values
	m.dot.SetPosition(
		m.applyNextPosition(
			m.move.NextPosition(elapsedBetweenUpdate)).Int())
	return true
}

// FIXME: losing some POWER when bouncing, while we expect to bounce infinitely (float approximation after time integration)
func (m *BouncingDot) applyNextPosition(nextPosition FloatingPoint, velocity FloatingPoint) FloatingPoint {
	var velocityCoefX, velocityCoefY float64 = 1, 1
	//var accelCoefX, accelCoefY float64 = 1, 1
	if int(nextPosition.X) < m.bounds.Min.X || int(nextPosition.X) >= m.bounds.Max.X {
		// moving to far to the LEFT or to the RIGHT, correcting overlaps
		if int(nextPosition.X) < m.bounds.Min.X {
			nextPosition = FloatingPoint{
				X: -nextPosition.X,
				Y: nextPosition.Y,
			}
		} else {
			nextPosition = FloatingPoint{
				X: 2*float64(m.bounds.Max.X) - nextPosition.X,
				Y: nextPosition.Y,
			}
		}
		// reverse X velocity
		velocityCoefX = -1
	}
	if int(nextPosition.Y) < m.bounds.Min.Y || int(nextPosition.Y) >= m.bounds.Max.Y {
		// moving to far to the TOP or to the BOTTOM, correcting overlaps
		if int(nextPosition.Y) < m.bounds.Min.Y {
			nextPosition = FloatingPoint{
				X: nextPosition.X,
				Y: -nextPosition.Y,
			}
		} else {
			nextPosition = FloatingPoint{
				X: nextPosition.X,
				Y: 2*float64(m.bounds.Max.Y) - nextPosition.Y,
			}
		}
		// reverse Y velocity
		velocityCoefY = -1
	}
	m.move.SetVelocity(FloatingPoint{
		X: velocity.X * velocityCoefX,
		Y: velocity.Y * velocityCoefY,
	})
	//var direction = DirectionToFloatingPoint(m.dotAcceleration.Direction())
	//m.dotAcceleration.SetDirection(FloatingPointToDirection(FloatingPoint{
	//	X: direction.X * accelCoefX,
	//	Y: direction.Y * accelCoefY,
	//}))
	return nextPosition
}

func (m *BouncingDot) Draw(c Canvas) error {
	return m.dot.Draw(c)
}
