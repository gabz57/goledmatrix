package scenes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"io/fs"

	"github.com/gabz57/goledmatrix/components/shapes"
	"image"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type MeteoIcons struct {
	//shape   *CompositeDrawable
	icons   []shapes.Img
	updated bool
}

func NewMeteoIconsComponent(canvas Canvas) *MeteoIcons {
	files1, err := ioutil.ReadDir("img/out/Box")
	if err != nil {
		log.Fatal(err)
	}
	icons1 := readIcons(canvas, "img/out/Box/", files1, image.Point{})

	files2, err := ioutil.ReadDir("img/out/Gaussian")
	if err != nil {
		log.Fatal(err)
	}
	icons2 := readIcons(canvas, "img/out/Gaussian", files2, image.Point{Y: 3 * 13})

	files3, err := ioutil.ReadDir("img/out/Linear")
	if err != nil {
		log.Fatal(err)
	}
	icons3 := readIcons(canvas, "img/out/Linear", files3, image.Point{Y: 6 * 13})

	var icons []shapes.Img
	for _, img := range icons1 {
		icons = append(icons, img)
	}
	for _, img := range icons2 {
		icons = append(icons, img)
	}
	for _, img := range icons3 {
		icons = append(icons, img)
	}

	meteoIcons := MeteoIcons{
		icons: icons,
		//shape: NewCompositeDrawable(nil),
	}
	//for _, icon := range icons {
	//	var d Drawable = &icon
	//	meteoIcons.shape.AddDrawable(&d)
	//}
	return &meteoIcons
}

func readIcons(canvas Canvas, dir string, files []fs.FileInfo, offset image.Point) []shapes.Img {
	var icons []shapes.Img
	var padding = 13
	var position = offset
	var targetSize = Point{X: 12, Y: 12}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".png") {
			path := dir + "/" + file.Name()
			println(file.Name(), position.String())
			icons = append(icons, funcName(nil, position, path, targetSize))
		}
		position = position.Add(image.Point{X: padding})
		if !position.Add(image.Point{X: 12, Y: 12}).In(canvas.Bounds()) {
			position = image.Point{Y: position.Y + padding}
		}
	}
	return icons
}

func funcName(parent *Graphic, position image.Point, path string, targetSize Point) shapes.Img {
	return *shapes.NewPng(
		NewOffsetGraphic(parent, nil, Point{X: position.X, Y: position.Y}),
		&path,
		targetSize,
	)
}

func (mi *MeteoIcons) Update(elapsedBetweenUpdate time.Duration) bool {
	var updated = false
	for _, icon := range mi.icons {
		updated = icon.Update(elapsedBetweenUpdate) || updated
	}
	return updated
}

func (mi *MeteoIcons) Draw(canvas Canvas) error {
	for _, icon := range mi.icons {
		err := icon.Draw(canvas)
		if err != nil {
			return err
		}
	}
	return nil
	//return mi.shape.Draw(canvas)
}
