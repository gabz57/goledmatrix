package fonts

import (
	"github.com/zachomedia/go-bdf"
	"golang.org/x/image/font"
	"io/ioutil"
	"log"
)

type MatrixFont struct {
	FileName string
}

var (
	Bdf4x6     = MatrixFont{"4x6.bdf"}
	Bdf5x7     = MatrixFont{"5x7.bdf"}
	Bdf5x8     = MatrixFont{"5x8.bdf"}
	Bdf6x9     = MatrixFont{"6x9.bdf"}
	Bdf6x10    = MatrixFont{"6x10.bdf"}
	Bdf6x12    = MatrixFont{"6x12.bdf"}
	Bdf6x13    = MatrixFont{"6x13.bdf"}
	Bdf6x13B   = MatrixFont{"6x13B.bdf"}
	Bdf6x13O   = MatrixFont{"6x13O.bdf"}
	Bdf7x13    = MatrixFont{"7x13.bdf"}
	Bdf7x13B   = MatrixFont{"7x13B.bdf"}
	Bdf7x13O   = MatrixFont{"7x13O.bdf"}
	Bdf7x14    = MatrixFont{"7x14.bdf"}
	Bdf7x14B   = MatrixFont{"7x14B.bdf"}
	Bdf8x13    = MatrixFont{"8x13.bdf"}
	Bdf8x13B   = MatrixFont{"8x13B.bdf"}
	Bdf8x13O   = MatrixFont{"8x13O.bdf"}
	Bdf9x15    = MatrixFont{"9x15.bdf"}
	Bdf9x15B   = MatrixFont{"9x15B.bdf"}
	Bdf9x18    = MatrixFont{"9x18.bdf"}
	Bdf9x18B   = MatrixFont{"9x18B.bdf"}
	Bdf10x20   = MatrixFont{"10x20.bdf"}
	BdfclR6x12 = MatrixFont{"clR6x12.bdf"}
	BdfhelvR12 = MatrixFont{"helvR12.bdf"}
)

var lazyBdfFonts = make(map[MatrixFont]*bdf.Font)

func loadFnt(matrixFont MatrixFont) (*bdf.Font, error) {
	log.Println("Loading Font " + matrixFont.FileName)

	fontBytes, err := ioutil.ReadFile("canvas/fonts/" + matrixFont.FileName)
	if err != nil {
		log.Println("Failed to read Font file " + matrixFont.FileName)
		return nil, err
	}
	fontPtr, err := bdf.Parse(fontBytes)
	if err != nil {
		log.Println("Failed to parse Font bytes " + matrixFont.FileName)
		return nil, err
	}
	return fontPtr, nil
}

func GetFont(f MatrixFont) font.Face {
	var err error
	if lazyBdfFonts[f] == nil {
		lazyBdfFonts[f], err = loadFnt(f)
		if err != nil {
			panic(err)
		}
	}
	return lazyBdfFonts[f].NewFace()
}
