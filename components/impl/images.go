package impl

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"strings"
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
		img: buildImg(graphic, &imgPath, targetSize),
	}
	var drawableImg Drawable = images.img
	images.shape.AddDrawable(&drawableImg)
	return &images
}

func (i *Images) Update(elapsedBetweenUpdate time.Duration) bool {
	return i.img.Update(elapsedBetweenUpdate)
}

func (i *Images) Draw(canvas Canvas) error {
	return i.shape.Draw(canvas)
}

func buildImg(graphic *Graphic, imgPath *string, size Point) *shapes.Img {
	if strings.HasSuffix(*imgPath, ".gif") {
		return shapes.NewGif(
			NewGraphic(graphic, nil),
			imgPath,
			size,
		)
	}
	if strings.HasSuffix(*imgPath, ".png") {
		return shapes.NewPng(
			NewGraphic(graphic, nil),
			imgPath,
			size,
		)
	}
	return nil
	//panic(errors.New("Cannot read " + *imgPath))
}
