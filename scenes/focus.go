package scenes

import (
	"github.com/gabz57/goledmatrix/canvas"
	"time"
)

type Focus struct {
}

func NewFocusComponent(canvas canvas.Canvas) *Focus {
	var c = Focus{}
	return &c
}

func (f Focus) Update(elapsedBetweenUpdate time.Duration) bool {
	panic("implement me")
}

func (f Focus) Draw(canvas canvas.Canvas) error {
	panic("implement me")
}
