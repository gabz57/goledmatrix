package impl

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"image"
	"time"
)

type MovingDot struct {
	move                   *Movement
	dot                    *shapes.Dot
	dotAcceleration        *ConstantAcceleration
	bounds                 image.Rectangle
	elapsedSinceSceneStart time.Duration
}

func NewMovingDot(c Canvas, initialPosition Point, initialVelocity FloatingPoint, bounds image.Rectangle) Component {

	//var mask Canvas
	//mask = NewSingleColorMask(c, ColorRed)
	var accs []Acceleration
	acceleration := NewConstantAcceleration(
		0, // test value
		0, // test value
	)
	accs = append(accs, *acceleration)

	dot := MovingDot{
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
		//mask:   mask,
	}
	dot.dotAcceleration = acceleration
	return &dot

}

func (m *MovingDot) Update(elapsedBetweenUpdate time.Duration) {
	// changing erratically of direction by changing the velocity
	m.elapsedSinceSceneStart += elapsedBetweenUpdate
	if m.elapsedSinceSceneStart > 2*time.Second {
		m.move.SetVelocity(FloatingPoint{
			X: OneOrMinusOne() * Float64Between(0, 64),
			Y: OneOrMinusOne() * Float64Between(0, 64),
		})
		m.elapsedSinceSceneStart = 0
	}
	// advance the position by one step, make it bounce on bounds with exact values
	m.dot.SetPosition(m.applyNextPosition(m.move.NextPosition(elapsedBetweenUpdate)).Int())
}

func (m *MovingDot) applyNextPosition(nextPosition FloatingPoint, velocity FloatingPoint) FloatingPoint {
	var velocityCoefX, velocityCoefY float64 = 1, 1
	var accelCoefX, accelCoefY float64 = 1, 1
	if int(nextPosition.X) < 0 || int(nextPosition.X) >= m.bounds.Max.X {
		// moving to far to the LEFT or to the RIGHT, correcting overlaps
		if int(nextPosition.X) < 0 {
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
		//stop moving LEFT or RIGHT
		velocityCoefX = -1
	}
	if int(nextPosition.Y) < 0 || int(nextPosition.Y) >= m.bounds.Max.Y {
		// moving to far to the TOP or to the BOTTOM, correcting overlaps
		if int(nextPosition.Y) < 0 {
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
	var direction = DirectionToFloatingPoint(m.dotAcceleration.Direction())
	m.dotAcceleration.SetDirection(FloatingPointToDirection(FloatingPoint{
		X: direction.X * accelCoefX,
		Y: direction.Y * accelCoefY,
	}))
	return nextPosition
}

func (m *MovingDot) Draw(c Canvas) error {
	return (*m).dot.Draw(c)
	//return (*m).dot.Draw(m.mask)
}
