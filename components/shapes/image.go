package shapes

import (
	"github.com/anthonynsimon/bild/transform"
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
	updated              bool // used with .png
}

func (i *Img) Update(elapsedBetweenUpdate time.Duration) bool {
	if i.activeDuration != nil {
		if i.elapsedSinceGifStart > *i.activeDuration {
			i.runNextImage()
			i.elapsedSinceGifStart += elapsedBetweenUpdate
			return true
		} else {
			i.elapsedSinceGifStart += elapsedBetweenUpdate
			return false
		}
	} else {
		if !i.updated {
			i.updated = true
			return true
		}
	}
	return false
}

func (i *Img) Draw(canvas Canvas) error {
	if i.activeImage != nil {
		min := image.Point(i.Graphic.ComputedOffset())
		max := min.Add(i.dimensions.Max)
		draw.Draw(
			canvas,
			image.Rectangle{
				Min: min,
				Max: max,
			},
			*i.activeImage,
			image.Point{},
			draw.Src,
		)
	}
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

func (i *Img) SetActiveImage(activeImage *image.Image) {
	i.activeImage = activeImage
	if i.activeDuration != nil {
		i.elapsedSinceGifStart = 0
		for index, anImg := range i.images {
			if &anImg == activeImage {
				i.activeDuration = &i.durations[index]
				break
			}
		}
	}
	i.updated = true
}

func (i *Img) Images() *[]image.Image {
	return &i.images
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

func NewPngFromPaths(graphic *Graphic, targetSize Point, paths ...string) *Img {
	images := make([]image.Image, len(paths))
	for i, path := range paths {
		images[i] = *ReadPng(&path)
	}
	return &Img{
		Graphic:   graphic,
		images:    images,
		durations: []time.Duration{},
		dimensions: image.Rectangle{
			Min: image.Point{},
			Max: image.Point(targetSize),
		},
		activeImage: &(images)[0],
	}
}

func NewPng(graphic *Graphic, path *string, targetSize Point) *Img {
	var png = *ReadPng(path)
	if !png.Bounds().Max.Eq(image.Point(targetSize)) {
		png = transform.Crop(png, image.Rectangle{
			Max: image.Point{X: targetSize.X, Y: targetSize.Y},
		})
	}
	var images = []image.Image{png}
	return &Img{
		Graphic:   graphic,
		images:    images,
		durations: []time.Duration{},
		dimensions: image.Rectangle{
			Min: image.Point{},
			Max: image.Point(targetSize),
		},
		activeImage: &(images)[0],
	}
}
