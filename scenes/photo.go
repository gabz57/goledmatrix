package scenes

import (
	"github.com/gabz57/goledmatrix/canvas"
	"time"
)

type Photo struct {
}

func NewPhotoComponent(canvas canvas.Canvas) *Photo {
	var c = Photo{}
	return &c
}

func (p Photo) Update(elapsedBetweenUpdate time.Duration) bool {
	panic("implement me")
}

func (p Photo) Draw(canvas canvas.Canvas) error {
	panic("implement me")
}
