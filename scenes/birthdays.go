package scenes

import (
	"github.com/gabz57/goledmatrix/canvas"
	"time"
)

type NextBirthdays struct {
}

func NewNextBirthdaysComponent(canvas canvas.Canvas) *NextBirthdays {
	var c = NextBirthdays{}
	return &c
}

func (nb NextBirthdays) Update(elapsedBetweenUpdate time.Duration) bool {
	panic("implement me")
}

func (nb NextBirthdays) Draw(canvas canvas.Canvas) error {
	panic("implement me")
}
