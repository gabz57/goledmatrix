package scenes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"time"
)

type Photo struct {
	dir      string
	shape    *CompositeDrawable
	image    *shapes.Img
	elapsed  time.Duration
	duration time.Duration
	graphic  *Graphic
}

func NewPhotoComponent(canvas Canvas) *Photo {
	var photoGraphic = NewGraphic(nil, NewLayout(ColorWhite, ColorBlack))

	var c = Photo{
		graphic:  photoGraphic,
		shape:    NewCompositeDrawable(photoGraphic),
		dir:      "img/Photos",
		duration: 5 * time.Second,
	}
	return &c
}

func (p *Photo) Update(elapsedBetweenUpdate time.Duration) bool {
	p.elapsed += elapsedBetweenUpdate
	if p.image != nil {
		p.image.Update(elapsedBetweenUpdate)
	}
	if p.image == nil || p.elapsed > p.duration {
		p.elapsed = 0
		file, err := RandomFile(p.dir, ImagesSuffixes)
		if err != nil {
			panic(err)
		}
		p.shape.RemoveDrawable(p.image)
		p.image = shapes.NewImg(p.graphic, file, Point{X: 128, Y: 128})
		p.shape.AddDrawable(p.image)
		return true
	}
	return false
}

func (p Photo) Draw(canvas Canvas) error {
	return p.shape.Draw(canvas)
}
