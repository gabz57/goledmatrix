package fonts

import (
	"fmt"
	"github.com/gosuri/uilive"
	"github.com/zachomedia/go-bdf"
	"golang.org/x/image/font"
	"io/ioutil"
)

var writer *uilive.Writer

func init() {
	// TODO: lazily load the font when needed
	writer = uilive.New()
	writer.Start()
	defer writer.Stop()
	fmt.Println("Loading Fonts...")
	_, err := loadFonts()
	if err != nil {
		panic(err)
	}
}

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

var BdfFonts = make(map[MatrixFont]*bdf.Font)

func loadFonts() (map[MatrixFont]*bdf.Font, error) {
	var matrixFonts = []MatrixFont{
		Bdf4x6,
		Bdf5x7,
		Bdf5x8,
		Bdf6x9,
		Bdf6x10,
		Bdf6x12,
		Bdf6x13,
		Bdf6x13B,
		Bdf6x13O,
		Bdf7x13,
		Bdf7x13B,
		Bdf7x13O,
		Bdf7x14,
		Bdf7x14B,
		Bdf8x13,
		Bdf8x13B,
		Bdf8x13O,
		Bdf9x15,
		Bdf9x15B,
		Bdf9x18,
		Bdf9x18B,
		Bdf10x20,
		BdfclR6x12,
		BdfhelvR12,
	}

	for _, matrixFont := range matrixFonts {
		_, _ = fmt.Fprintf(writer.Newline(), "Loading Font "+matrixFont.FileName+"\n")

		fontBytes, err := ioutil.ReadFile("fonts/" + matrixFont.FileName)
		if err != nil {
			fmt.Println("Failed to read Font file " + matrixFont.FileName)
			return nil, err
		}
		fontPtr, err := bdf.Parse(fontBytes)
		if err != nil {
			fmt.Println("Failed to parse Font bytes " + matrixFont.FileName)
			return nil, err
		}
		BdfFonts[matrixFont] = fontPtr
	}
	_, _ = fmt.Fprintf(writer.Newline(), "Loaded all Fonts\n")
	return BdfFonts, nil
}

func GetFont(f MatrixFont) font.Face {
	// TODO: lazy init fonts
	return BdfFonts[f].NewFace()
}
