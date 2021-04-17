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
	acceleration := NewConstantAcceleration(
		10,
		TOP,
	)
	return &MovingHeart{
		move: NewMovement(
			initialPosition.Floating(),
			FloatingPoint{
				X: 0,
				Y: 0,
			},
			&[]Acceleration{acceleration},
		),
		heart: NewHeart(
			canvas,
			NewGraphic(nil, nil),
			initialPosition,
			fadeDuration,
			initialFade,
			initialFadeOut),
		heartAcceleration: acceleration,
	}
}

func (m *MovingHeart) Update(elapsedBetweenUpdate time.Duration) {
	m.heart.Update(elapsedBetweenUpdate)
	position, _ := m.move.NextPosition(elapsedBetweenUpdate)
	m.heart.SetPosition(position.Int())
}

func (m *MovingHeart) Draw(c Canvas) error {
	return m.heart.Draw(c)
}

func (m *MovingHeart) IsOut() bool {
	return m.heart.GetPosition().Y < -HeartHeight
}
