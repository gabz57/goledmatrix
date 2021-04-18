package shapes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"image"
	"image/draw"
	"image/gif"
	"os"
	"time"
)

type Img struct {
	*Graphic
	images               []image.Image
	durations            []time.Duration
	dimensions           image.Rectangle
	activeImage          *image.Image
	activeDuration       *time.Duration
	elapsedSinceGifStart time.Duration
}

func (i *Img) Update(elapsedBetweenUpdate time.Duration) {
	if i.elapsedSinceGifStart > *i.activeDuration {
		i.runNextImage()
	}
	i.elapsedSinceGifStart += elapsedBetweenUpdate
}

func (i *Img) Draw(canvas Canvas) error {
	draw.Draw(
		canvas,
		image.Rectangle{
			Min: canvas.Bounds().Min.Add(
				image.Point(i.Graphic.ComputedOffset())),
			Max: canvas.Bounds().Max,
		},
		*i.activeImage,
		image.Point{},
		draw.Src,
	)
	return nil
}

func (i *Img) runNextImage() {
	for idx, img := range i.images {
		if img == *i.activeImage {
			i.activeImage = &i.images[(idx+1)%len(i.images)]
			i.activeDuration = &i.durations[(idx+1)%len(i.durations)]
			i.elapsedSinceGifStart = 0
			break
		}
	}
}

func NewGif(graphic *Graphic, path *string, targetSize Point) *Img {
	images, durations, dimensions := ReadGif(path, targetSize)
	if len(*durations) != len(*images) {
		panic("Images nb & durations nb differs in GIF : " + *path)
	}
	return &Img{
		Graphic:        graphic,
		images:         *images,
		durations:      *durations,
		dimensions:     *dimensions,
		activeImage:    &(*images)[0],
		activeDuration: &(*durations)[0],
	}
}

func NewGifFromFiles(graphic *Graphic, targetSize Point, imgDuration time.Duration, paths ...string) *Img {
	durations := make([]time.Duration, len(paths))
	images := make([]image.Image, len(paths))
	for i, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		g, err := gif.DecodeAll(f)
		if err != nil {
			panic(err)
		}
		images[i] = g.Image[0]
		durations[i] = imgDuration
	}
	return &Img{
		Graphic:   graphic,
		images:    images,
		durations: durations,
		dimensions: image.Rectangle{
			Min: image.Point{},
			Max: image.Point(targetSize),
		},
		activeImage:    &(images)[0],
		activeDuration: &(durations)[0],
	}
}