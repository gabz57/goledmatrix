package impl

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"image"
	"time"
)

type MovingDot struct {
	Movement
	initialPosition Point
	dot             shapes.Dot
	bounds          image.Rectangle
	//mask   Canvas
}

func NewMovingDot(c Canvas, origin Point, initialVelocity FloatingPoint, bounds image.Rectangle) Component {
	var movingDotGraphic = NewGraphic(nil, nil)
	//var mask Canvas
	//mask = NewSingleColorMask(c, ColorRed)
	var accs []Acceleration
	accs = append(accs, *NewConstantAcceleration(
		0, // test value
		0, // test value
	))
	return &MovingDot{
		Movement: *NewMovement(
			FloatingPoint{
				X: initialVelocity.X,
				Y: initialVelocity.Y,
			},
			accs,
		),
		dot:             *shapes.NewDot(&movingDotGraphic, origin),
		bounds:          bounds,
		initialPosition: origin,
		//mask:   mask,
	}
}

func (m *MovingDot) Update(elapsedBetweenUpdate time.Duration) {
	// FIXME: also consider max getVelocity
	_, dXY := m.Advance(elapsedBetweenUpdate)
	//velocity, dXY := m.Advance(elapsedBetweenUpdate)
	//// note: maybe need to compare with absolute value
	//if velocity.X > m.maxVelocityX {
	//	velocity = m.maxVelocity
	//}
	m.dot.SetPosition(Point{
		X: (m.initialPosition.X + int(dXY.X)) % m.bounds.Max.X,
		Y: (m.initialPosition.Y + int(dXY.Y)) % m.bounds.Max.Y,
	})
}

func (m *MovingDot) Draw(c Canvas) error {
	return (*m).dot.Draw(c)
	//return (*m).dot.Draw(m.mask)
}
