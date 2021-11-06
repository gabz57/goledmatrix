package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"time"
)

type Images struct {
	shape *CompositeDrawable
	img   *shapes.Img
}

func NewImages(imgPath string, position Point, targetSize Point) *Images {
	graphic := NewOffsetGraphic(nil, nil, position)
	images := Images{
		shape: NewCompositeDrawable(
			graphic),
		img: shapes.NewImg(graphic, &imgPath, targetSize),
	}
	images.shape.AddDrawable(images.img)
	return &images
}

func (i *Images) Update(elapsedBetweenUpdate time.Duration) bool {
	return i.img.Update(elapsedBetweenUpdate)
}

func (i *Images) Draw(canvas Canvas) error {
	return i.shape.Draw(canvas)
}
