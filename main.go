package main

import (
	"github.com/gabz57/goledmatrix/components/impl"
	"github.com/gabz57/goledmatrix/matrix"
)

func main() {
	matrix.RunMatrices(goLedApplication)
}

func goLedApplication() {
	matrix.Run(impl.Gameloop)
}

//
//import (
//	"github.com/anthonynsimon/bild/transform"
//	"image"
//	"image/png"
//	"io/ioutil"
//	"log"
//	"os"
//	"strings"
//
//	"github.com/srwiley/oksvg"
//	"github.com/srwiley/rasterx"
//)
//
//var filters = map[string]transform.ResampleFilter{
//	"Linear":   transform.Linear,
//	"Box":      transform.Box,
//	"Gaussian": transform.Gaussian,
//	//"Lanczos":           transform.Lanczos,
//	//"CatmullRom":        transform.CatmullRom,
//	//"MitchellNetravali": transform.MitchellNetravali,
//	//"NearestNeighbor":   transform.NearestNeighbor,
//}
//
//var svgSuffix = ".svg"
//
//func main() {
//	files, err := ioutil.ReadDir("img/in")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, file := range files {
//		name := file.Name()
//		if strings.HasSuffix(name, svgSuffix) {
//			convertFileToPng(name[:len(name)-len(svgSuffix)])
//		}
//	}
//}
//
//func convertFileToPng(imgName string) {
//	rgba := readSVG("img/in/" + imgName + ".svg")
//	mkdir("img/out/64")
//	write("img/out/64/"+imgName+".png", rgba)
//	for filterName, filter := range filters {
//		mkdir("img/out/16/" + filterName)
//		write("img/out/16/"+filterName+"/"+imgName+".png", transform.Resize(rgba, 16, 16, filter))
//	}
//}
//
//func readSVG(svgFileName string) *image.RGBA {
//
//	in, err := os.Open(svgFileName)
//	if err != nil {
//		panic(err)
//	}
//	defer in.Close()
//
//	w, h := 64, 64
//	icon, _ := oksvg.ReadIconStream(in)
//	icon.SetTarget(0, 0, float64(w), float64(h))
//	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
//	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)
//	return rgba
//}
//
//func mkdir(dir string) {
//	err := os.MkdirAll(dir, os.ModePerm)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func write(name string, img *image.RGBA) {
//	out12, err := os.Create(name)
//	defer out12.Close()
//
//	if err != nil {
//		panic(err)
//	}
//
//	err = png.Encode(out12, img)
//	if err != nil {
//		panic(err)
//	}
//}
