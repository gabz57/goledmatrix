package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"time"
)

type Images struct {
	shape *CompositeDrawable
	gif   *shapes.Img
}

func NewImages(imgPath string, position Point, targetSize Point) *Images {
	graphic := NewOffsetGraphic(nil, nil, position)
	images := Images{
		shape: NewCompositeDrawable(
			graphic),
		gif: buildGif(graphic, &imgPath, targetSize),
	}
	var drawableMario Drawable = images.gif
	images.shape.AddDrawable(&drawableMario)
	return &images
}

func (i *Images) Update(elapsedBetweenUpdate time.Duration) {
	i.gif.Update(elapsedBetweenUpdate)
}

func (i *Images) Draw(canvas Canvas) error {
	return i.shape.Draw(canvas)
}

func buildGif(graphic *Graphic, imgPath *string, size Point) *shapes.Img {
	return shapes.NewGif(
		NewGraphic(graphic, nil),
		imgPath,
		size,
	)
}
