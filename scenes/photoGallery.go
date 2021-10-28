package scenes

import (
	"github.com/gabz57/goledmatrix/canvas"
	"time"
)

type PhotoGalleryData struct {
}

type PhotoGallery struct {
}

func NewPhotoGalleryComponent(canvas canvas.Canvas) *PhotoGallery {
	var c = PhotoGallery{}
	return &c
}

func (pg PhotoGallery) Update(elapsedBetweenUpdate time.Duration) bool {
	panic("implement me")
}

func (pg PhotoGallery) Draw(canvas canvas.Canvas) error {
	panic("implement me")
}
