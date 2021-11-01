package components

import (
	"errors"
	"fmt"
	"github.com/anthonynsimon/bild/transform"
	. "github.com/gabz57/goledmatrix/canvas"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var randomSource = rand.NewSource(time.Now().UnixNano())
var Random = rand.New(randomSource)

var Origin = Point{}
var None = Origin

var ColorWhite color.Color = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
var ColorBlack color.Color = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
var ColorRed color.Color = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
var ColorGreen color.Color = color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}
var ColorBlue color.Color = color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}
var ColorViolet color.Color = color.RGBA{R: 0xff, G: 0x00, B: 0xff, A: 0xff}
var ColorYellow color.Color = color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}

func NewGraphic(parent *Graphic, layout *Layout) *Graphic {
	return NewOffsetGraphic(parent, layout, None)
}

func NewOffsetGraphic(parent *Graphic, layout *Layout, offset Point) *Graphic {
	if layout == nil {
		if parent != nil {
			layout = parent.layout
		} else {
			layout = DefaultLayout()
		}
	}
	return &Graphic{
		layout: layout,
		parent: parent,
		offset: offset,
	}
}

func DefaultLayout() *Layout {
	return &Layout{
		color:           color.White,
		backgroundColor: nil,
	}
}

func NewLayout(fColor, bgColor color.Color) *Layout {
	return &Layout{
		color:           fColor,
		backgroundColor: bgColor,
	}
}

func TimeToText(t time.Time) string {
	hour, min, sec := t.Clock()
	millis := t.Nanosecond() / 1000000
	return fmt.Sprintf("%02d", hour) + ":" + fmt.Sprintf("%02d", min) + ":" + fmt.Sprintf("%02d", sec) + "." + fmt.Sprintf("%03d", millis)
}

func Rotate(p, o Point, degrees float64) Point {
	rad := DegToRad(degrees)
	fp := p.Floating()
	fo := o.Floating()
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return Point{
		X: int(cos*(fp.X-fo.X) - sin*(fp.Y-fo.Y) + fo.X),
		Y: int(sin*(fp.X-fo.X) + cos*(fp.Y-fo.Y) + fo.Y),
	}
}

func RotateOrigin(fp FloatingPoint, degrees float64) FloatingPoint {
	rad := DegToRad(degrees)
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return FloatingPoint{
		X: cos*(fp.X) - sin*(fp.Y),
		Y: sin*(fp.X) + cos*(fp.Y),
	}
}

func DegToRad(x float64) float64 {
	return (x / 180) * math.Pi
}

func RadToDeg(x float64) float64 {
	return (x * 180) / math.Pi
}

// Angle (degrees) from cartesian coordinates
func FloatingPointToDirection(point FloatingPoint) float64 {
	return RadToDeg(math.Atan(point.Y / point.X))
}

// Cartesian coordinates from angle (degrees)
func DirectionToFloatingPoint(direction float64) FloatingPoint {
	rad := DegToRad(direction)
	return FloatingPoint{
		X: math.Cos(rad),
		Y: math.Sin(rad),
	}
}

func Int64Between(low, high int64) int64 {
	var gen int64
	for gen < low {
		gen = Random.Int63n(high)
	}
	return gen
}

func Float64Between(low, high float64) float64 {
	var gen float64
	for gen <= low {
		gen = Random.Float64() * high
	}
	return gen
}

func OneOrMinusOne() float64 {
	if Random.Intn(1) == 0 {
		return -1
	}
	return 1
}

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
	originFullSize := image.Rect
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
