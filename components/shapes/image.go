package shapes

import (
	"errors"
	"github.com/anthonynsimon/bild/transform"
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
	"strings"
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
	updated              bool // used with .png & .jpg for 1st rendering
	mask                 image.Image
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
		//max := min.Add(i.dimensions.Max)
		draw.DrawMask(
			canvas,
			image.Rectangle{
				Min: min,
				Max: canvas.Bounds().Max,
			},
			*i.activeImage,
			image.Point{},
			i.mask,
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

func (i *Img) GetActiveImage() *image.Image {
	return i.activeImage
}

func (i *Img) Rotate(angle float64) {
	imgBounds := i.dimensions
	for index, img := range *i.Images() {
		(*i.Images())[index] = transform.Rotate(img, angle, &transform.RotationOptions{
			ResizeBounds: true,
			Pivot:        nil,
		})
	}
	rgbaMask := image.NewRGBA(imgBounds)
	for x := 0; x < imgBounds.Max.X; x++ {
		for y := 0; y < imgBounds.Max.Y; y++ {
			rgbaMask.Set(x, y, color.Alpha{
				A: uint8(255),
			})
			//rgbaMask.Set(x, y, color.Opaque)
		}
	}
	rgbaMask = transform.Rotate(rgbaMask, angle, &transform.RotationOptions{
		ResizeBounds: true,
		Pivot:        nil,
	})
	//imgMask := image.NewAlpha(imgBounds)
	//	for x := 0; x < imgBounds.Max.X; x++ {
	//		for y := 0; y < imgBounds.Max.Y; y++ {
	//			imgMask.SetAlpha(color.Opaque.RGBA())
	//		}
	//	}
	//	var imgMaskI image.Image = imgMask
	//	imgMaskI = transform.Rotate(imgMaskI, angle, &transform.RotationOptions{
	//		ResizeBounds: true,
	//		Pivot:        nil,
	//	})

	i.mask = rgbaMask
}

func (i *Img) GetMask() *image.Image {
	return &i.mask
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
		if !images[i].Bounds().Max.In(image.Rectangle{Max: image.Point(targetSize)}) {
			images[i] = transform.Resize(images[i], targetSize.X, targetSize.Y, transform.Linear)
		}
	}
	return &Img{
		Graphic:   graphic,
		images:    images,
		durations: []time.Duration{},
		dimensions: image.Rectangle{
			Max: image.Point(targetSize),
		},
		activeImage: &(images)[0],
	}
}

func NewImg(graphic *Graphic, path *string, targetSize Point) *Img {
	var img *Img
	if strings.HasSuffix(*path, ".png") {
		return asImg(graphic, *ReadPng(path), targetSize)
	} else if strings.HasSuffix(*path, ".jpg") {
		return asImg(graphic, *ReadJpg(path), targetSize)
	} else if strings.HasSuffix(*path, ".gif") {
		img = NewGif(graphic, path, targetSize)
		return img
	} else {
		panic(errors.New("Image format not supported : " + *path))
	}
	return img
}

func asImg(graphic *Graphic, img image.Image, targetSize Point) *Img {
	rectangle := image.Rectangle{Max: image.Point(targetSize)}
	if !img.Bounds().Max.In(rectangle) {
		size := computeTargetSize(img, targetSize)
		img = transform.Resize(img, size.X, size.Y, transform.Linear)
	}
	var images = []image.Image{img}
	return &Img{
		Graphic:   graphic,
		images:    images,
		durations: []time.Duration{},
		dimensions: image.Rectangle{
			Max: img.Bounds().Max,
		},
		activeImage: &(images)[0],
	}
}

func computeTargetSize(image image.Image, targetSize Point) Point {
	originFullSize := image.Bounds()
	originalMaxSize := originFullSize.Max
	ratioX := float64(targetSize.X) / float64(originalMaxSize.X)
	ratioY := float64(targetSize.Y) / float64(originalMaxSize.Y)
	ratio := ratioX
	if ratioX > ratioY {
		ratio = ratioY
	}
	adjTargetSize := Point{
		X: int(float64(originalMaxSize.X) * ratio),
		Y: int(float64(originalMaxSize.Y) * ratio),
	}
	return adjTargetSize
}
