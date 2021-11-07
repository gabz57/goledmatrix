package scenes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"
)

type PhotoGallery struct {
	dir          string
	images       []*shapes.Img
	elapsed      time.Duration
	graphic      *Graphic
	galleryIndex int
	fadeInRatio  *float64
	fadeOutRatio *float64
}

const gallerySize = 10
const fadeDuration = 1 * time.Second
const imgDuration = 3 * time.Second

func NewPhotoGalleryComponent(_ Canvas) *PhotoGallery {
	var photoGraphic = NewGraphic(nil, NewLayout(ColorWhite, ColorBlack))

	var pg = PhotoGallery{
		graphic: photoGraphic,
		dir:     "img/Photos",
		images:  make([]*shapes.Img, gallerySize),
	}
	pg.images[pg.galleryIndex] = pg.newGalleryImage()

	return &pg
}

func (p *PhotoGallery) Update(elapsedBetweenUpdate time.Duration) bool {
	p.elapsed += elapsedBetweenUpdate
	var updated = false
	for _, img := range p.images {
		if img != nil {
			updated = (*img).Update(elapsedBetweenUpdate) || updated
		}
	}

	if p.elapsed < fadeDuration {
		//p.updateFadeIn(p.galleryIndex)
		ratio := float64(p.elapsed.Nanoseconds()) / float64(fadeDuration.Nanoseconds())
		p.fadeInRatio = &ratio
		return true
	} else if p.elapsed > fadeDuration && p.elapsed < (imgDuration-fadeDuration) && p.fadeInRatio != nil {
		p.fadeInRatio = nil
		return true
	} else if p.elapsed > (imgDuration-fadeDuration) && p.elapsed < imgDuration {
		if p.imageCount() == gallerySize {
			//p.updateFadeOut(p.previousGalleryIndex())
			ratio := float64(1) - float64(p.elapsed.Nanoseconds()-(imgDuration.Nanoseconds()-fadeDuration.Nanoseconds()))/float64(fadeDuration.Nanoseconds())
			p.fadeOutRatio = &ratio
		}
		return true
	} else if p.elapsed > imgDuration {
		p.elapsed = 0
		p.fadeOutRatio = nil

		p.galleryIndex = p.nextGalleryIndex()
		if p.imageCount() == gallerySize {
			p.images[p.galleryIndex] = nil
		}
		p.images[p.galleryIndex] = p.newGalleryImage()
		return true
	}
	return updated
}

func (p *PhotoGallery) imageCount() int {
	var cnt = 0
	for _, img := range p.images {
		if img != nil {
			cnt++
		}
	}
	return cnt
}

func (p *PhotoGallery) previousGalleryIndex() int {
	return (p.galleryIndex + gallerySize - 1) % gallerySize
}

func (p *PhotoGallery) nextGalleryIndex() int {
	return (p.galleryIndex + 1) % gallerySize
}

func (p *PhotoGallery) buildImg(file *string) *shapes.Img {
	targetMaxSize := 96 - rand.Intn(48)
	maxSize := Point{X: targetMaxSize, Y: targetMaxSize}
	return shapes.NewLazyImg(
		NewOffsetGraphic(p.graphic, nil, Point{}),
		file,
		maxSize,
		randomRotateAndMove(randomRotateAndMoveValues(image.Rectangle{Max: image.Point(maxSize)})),
	)
}

func randomRotateAndMove(rotationAngle float64, randX, randY int) func(img *shapes.Img) {
	return func(img *shapes.Img) {
		img.Rotate(rotationAngle)
		img.Graphic.SetOffset(Point{
			X: randX,
			Y: randY,
		})
	}
}

func randomRotateAndMoveValues(bounds image.Rectangle) (float64, int, int) {
	rotationAngle := float64(rand.Intn(90) - 45)
	randX := rand.Intn(128 - bounds.Dx())
	randY := rand.Intn(128 - bounds.Dy())
	return rotationAngle, randX, randY
}

func (p PhotoGallery) Draw(canvas Canvas) error {
	rgba := image.NewRGBA(canvas.Bounds())

	nbImages := p.imageCount()
	for i := 0; i < nbImages; i++ {
		index := p.galleryIndex
		if nbImages > 1 {
			index = (i + p.galleryIndex + 1) % nbImages
		}
		var img = p.images[index]
		if img != nil {
			offset := (*img).Graphic.ComputedOffset()
			var fadeRatio *float64 = nil
			if i == 0 && p.fadeOutRatio != nil {
				fadeRatio = p.fadeOutRatio
			} else if i == nbImages-1 && p.fadeInRatio != nil {
				fadeRatio = p.fadeInRatio
			}
			err := drawOver(rgba, (*img).GetActiveImage(), offset, (*img).GetMask(), fadeRatio)
			if err != nil {
				return err
			}
		}
	}
	draw.Draw(canvas, canvas.Bounds(), rgba, image.Point{}, draw.Src)

	return nil
}

func drawOver(dest *image.RGBA, src *image.Image, offset Point, mask image.Image, fadeRatio *float64) error {
	draw.DrawMask(
		dest,
		dest.Bounds().Add(image.Point(offset)),
		*src,
		image.Point{},
		prepareMask(fadeRatio, mask),
		image.Point{},
		draw.Over,
	)
	return nil
}

func prepareMask(fadeRatio *float64, mask image.Image) image.Image {
	if fadeRatio != nil && mask != nil {
		return fade(mask, *fadeRatio)
	} else {
		return mask
	}
}

func fade(mask image.Image, ratio float64) *image.RGBA {
	var shadedMask = image.NewRGBA(mask.Bounds())
	max := shadedMask.Bounds().Max
	for x := 0; x < max.X; x++ {
		for y := 0; y < max.Y; y++ {
			at := mask.At(x, y)
			_, _, _, alpha := at.RGBA()
			var alphaAdapted float64 = 0
			if alpha > 0 {
				alphaAdapted = 255
			}
			shadedMask.Set(x, y, color.Alpha{A: uint8(alphaAdapted * ratio)})
		}
	}
	return shadedMask
}

func (p *PhotoGallery) newGalleryImage() *shapes.Img {
	file, err := RandomFile(p.dir, ImagesSuffixes)
	if err != nil {
		panic(err)
	}
	return p.buildImg(file)
}
