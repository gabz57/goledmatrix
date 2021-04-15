package impl

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"time"
)

type MovingHeart struct {
	move              *Movement
	heart             *Heart
	heartAcceleration *ConstantAcceleration
}

func NewMovingHeart(canvas Canvas, initialPosition Point, fadeDuration time.Duration, initialFade float64, initialFadeOut bool) *MovingHeart {
	var accs []Acceleration
	acceleration := NewConstantAcceleration(
		10,
		TOP,
	)
	accs = append(accs, *acceleration)

	heart := MovingHeart{
		move: NewMovement(
			initialPosition.Floating(),
			FloatingPoint{
				X: 0,
				Y: 0,
			},
			&accs,
		),
		heart: NewHeart(
			canvas,
			NewGraphic(nil, nil),
			initialPosition,
			fadeDuration,
			initialFade,
			initialFadeOut),
	}
	heart.heartAcceleration = acceleration
	return &heart

}

func (m *MovingHeart) Update(elapsedBetweenUpdate time.Duration) {
	m.heart.Update(elapsedBetweenUpdate)
	position, _ := m.move.NextPosition(elapsedBetweenUpdate)
	m.heart.SetPosition(position.Int())
}

func (m *MovingHeart) Draw(c Canvas) error {
	return (*m).heart.Draw(c)
}

func (m *MovingHeart) IsOut() bool {
	return m.heart.GetPosition().Y < -HeartHeight
}
