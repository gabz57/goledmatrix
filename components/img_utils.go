package components

import (
	"errors"
	"github.com/anthonynsimon/bild/transform"
	. "github.com/gabz57/goledmatrix/canvas"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadGif(imgPath *string, targetSize Point) (*[]image.Image, *[]time.Duration, *image.Rectangle) {

	f, err := os.Open(*imgPath)
	if err != nil {
		panic(err)
	}

	g, err := gif.DecodeAll(f)
	if err != nil {
		panic(err)
	}

	durations := make([]time.Duration, len(g.Delay))
	images := make([]image.Image, len(g.Image))
	firstImage := g.Image[0]
	adjTargetSize, _ := computeTargetSize(firstImage, targetSize)
	images[0] = transform.Resize(firstImage, adjTargetSize.X, adjTargetSize.Y, transform.Linear)
	durations[0] = time.Millisecond * time.Duration(g.Delay[0]) * 10
	//writeBean(0, images)
	img := firstImage
	for i, nextImage := range g.Image {
		if i == 0 {
			continue
		}

		if firstImage.Rect != nextImage.Rect {
			// Replace only a specific part of the gif
			draw.Draw(img, img.Rect, nextImage, image.Point{}, draw.Over)
		} else {
			// Replace the whole image
			img = nextImage
		}
		images[i] = transform.Resize(img, adjTargetSize.X, adjTargetSize.Y, transform.Linear)
		//writeBean(i, images)
		durations[i] = time.Millisecond * time.Duration(g.Delay[i]) * 10
	}

	return &images, &durations, &image.Rectangle{
		Min: image.Point{},
		Max: image.Point(targetSize),
	}
}

func ReadPng(imgPath *string) *image.Image {
	f, err := os.Open(*imgPath)
	if err != nil {
		panic(err)
	}

	pngImg, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	return &pngImg
}

func ReadJpg(imgPath *string) *image.Image {
	f, err := os.Open(*imgPath)
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	return &img
}

func writeBean(i int, images []image.Image) {
	myfile, err := os.Create("img/bean/bean" + strconv.Itoa(i) + ".gif") // ... now lets save imag
	if err != nil {
		panic(err)
	}
	err = gif.Encode(myfile, images[i], nil)
	if err != nil {
		panic(err)
	}
}
func ReadBeanGif() (*[]image.Image, *[]time.Duration, *image.Rectangle) {
	durations := make([]time.Duration, 12)
	images := make([]image.Image, 12)
	for i := 0; i <= 11; i++ {
		f, err := os.Open("img/bean/bean" + strconv.Itoa(i) + ".gif")
		if err != nil {
			panic(err)
		}

		g, err := gif.DecodeAll(f)
		if err != nil {
			panic(err)
		}

		images[i] = g.Image[0]
		durations[i] = 100 * time.Millisecond
	}

	return &images, &durations, &image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: 128, Y: 128},
	}
}

// scale image to fit both target width and height
func computeTargetSize(image *image.Paletted, targetSize Point) (Point, float64) {
	ratioX := float64(targetSize.X) / float64(image.Rect.Max.X)
	ratioY := float64(targetSize.Y) / float64(image.Rect.Max.Y)
	ratio := ratioX
	if ratioX > ratioY {
		ratio = ratioY
	}
	adjTargetSize := Point{
		X: int(float64(image.Rect.Max.X) * ratio),
		Y: int(float64(image.Rect.Max.Y) * ratio),
	}
	return adjTargetSize, ratio
}

var ImagesSuffixes = []string{".jpg", ".png", ".gif"}

func RandomFile(dir string, suffixes []string) (*string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, errors.New("no files in " + dir)
	}
	var validImgIndex = -1
	var nbTests = 0
	// avoid infinite loop using size of directory
	for validImgIndex == -1 && nbTests < len(files) {
		var imgIndex = Random.Intn(len(files))
		name := files[imgIndex].Name()
		println("Testing :", name)
		if !files[imgIndex].IsDir() {
			for _, suffix := range suffixes {
				if strings.HasSuffix(name, suffix) {
					validImgIndex = imgIndex
					break
				}
			}
		}
		nbTests++
	}
	if validImgIndex == -1 {
		return nil, errors.New("no valid files found in " + dir)
	}
	randomFileName := dir + "/" + files[validImgIndex].Name()
	return &randomFileName, nil
}
