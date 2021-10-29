package scenes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	"io/fs"

	"github.com/gabz57/goledmatrix/components/shapes"
	"image"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type MeteoIcons16 struct {
	//shape   *CompositeDrawable
	icons   []shapes.Img
	updated bool
}

func NewMeteoIcons16Component(canvas Canvas) *MeteoIcons16 {
	files1, err := ioutil.ReadDir("img/out/16/Box")
	if err != nil {
		log.Fatal(err)
	}
	icons1 := readIcons16(canvas, "img/out/16/Box/", files1, image.Point{})

	//files2, err := ioutil.ReadDir("img/out/16/Gaussian")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//icons2 := readIcons16(canvas, "img/out/16/Gaussian", files2, image.Point{Y: 4 * 17})

	files3, err := ioutil.ReadDir("img/out/16/Linear")
	if err != nil {
		log.Fatal(err)
	}
	icons3 := readIcons16(canvas, "img/out/16/Linear", files3, image.Point{Y: 4 * 17})

	var icons []shapes.Img
	for _, img := range icons1 {
		icons = append(icons, img)
	}
	//for _, img := range icons2 {
	//	icons = append(icons, img)
	//}
	for _, img := range icons3 {
		icons = append(icons, img)
	}

	meteoIcons := MeteoIcons16{
		icons: icons,
		//shape: NewCompositeDrawable(nil),
	}
	//for _, icon := range icons {
	//	var d Drawable = &icon
	//	meteoIcons.shape.AddDrawable(&d)
	//}
	return &meteoIcons
}

func readIcons16(canvas Canvas, dir string, files []fs.FileInfo, offset image.Point) []shapes.Img {
	var icons []shapes.Img
	var padding = 17
	var position = offset
	var length = 16
	var targetSize = Point{X: length, Y: length}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".png") {
			path := dir + "/" + file.Name()
			println(file.Name(), position.String())
			icons = append(icons, funcName(nil, position, path, targetSize))
		}
		position = position.Add(image.Point{X: padding})
		if !position.Add(image.Point{X: length, Y: length}).In(canvas.Bounds()) {
			position = image.Point{Y: position.Y + padding}
		}
	}
	return icons
}

func (mi *MeteoIcons16) Update(elapsedBetweenUpdate time.Duration) bool {
	var updated = false
	for _, icon := range mi.icons {
		updated = icon.Update(elapsedBetweenUpdate) || updated
	}
	return updated
}

func (mi *MeteoIcons16) Draw(canvas Canvas) error {
	for _, icon := range mi.icons {
		err := icon.Draw(canvas)
		if err != nil {
			return err
		}
	}
	return nil
	//return mi.shape.Draw(canvas)
}
