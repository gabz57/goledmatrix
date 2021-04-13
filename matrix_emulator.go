package goledmatrix

import (
	"fmt"
	"github.com/faiface/mainthread"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"image"
	"image/color"
	"os"
)

const DefaultPixelPitch = 12
const windowTitle = "RGB led matrix emulator (shiny)"

type MatrixEmulator struct {
	config                  *MatrixConfig
	PixelPitch              int
	Gutter                  int
	Width                   int
	Height                  int
	GutterColor             color.Color
	PixelPitchToGutterRatio int
	Margin                  int

	w       screen.Window
	s       screen.Screen
	isReady bool
}

func (m *MatrixEmulator) IsEmulator() bool {
	return true
}

func NewMatrixEmulator(config *MatrixConfig) (c Matrix, err error) {
	fmt.Println("NewEmulator")
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("error creating matrix: %v", r)
			}
		}
	}()
	w, h := config.Geometry()
	e := &MatrixEmulator{
		config:                  config,
		Width:                   w,
		Height:                  h,
		GutterColor:             color.Gray{Y: 20},
		PixelPitchToGutterRatio: 2,
		Margin:                  10,
	}
	e.updatePixelPitchForGutter(DefaultPixelPitch / e.PixelPitchToGutterRatio)
	return e, nil
}

/*
* Some formulas that allowed me to better understand the drawable area. I found that the math was
* easiest when put in terms of the Gutter width, hence the addition of PixelPitchToGutterRatio.
*
* PixelPitch = PixelPitchToGutterRatio * Gutter
* DisplayWidth = (PixelPitch * LEDColumns) + (Gutter * (LEDColumns - 1)) + (2 * Margin)
* Gutter = (DisplayWidth - (2 * Margin)) / (PixelPitchToGutterRatio * LEDColumns + LEDColumns - 1)
*
*  MMMMMMMMMMMMMMMM.....MMMM
*  MGGGGGGGGGGGGGGG.....GGGM
*  MGLGLGLGLGLGLGLG.....GLGM
*  MGGGGGGGGGGGGGGG.....GGGM
*  MGLGLGLGLGLGLGLG.....GLGM
*  MGGGGGGGGGGGGGGG.....GGGM
*  .........................
*  MGGGGGGGGGGGGGGG.....GGGM
*  MGLGLGLGLGLGLGLG.....GLGM
*  MGGGGGGGGGGGGGGG.....GGGM
*  MMMMMMMMMMMMMMMM.....MMMM
*
*  where:
*    M = Margin
*    G = Gutter
*    L = LED
 */
func (m *MatrixEmulator) updatePixelPitchForGutter(gutterWidth int) {
	m.PixelPitch = m.PixelPitchToGutterRatio * gutterWidth
	m.Gutter = gutterWidth
}

func (m *MatrixEmulator) Config() *MatrixConfig {
	return m.config
}

// Geometry returns the width and the height of the matrix
func (m *MatrixEmulator) Geometry() (width, height int) {
	return m.Width, m.Height
}

func (m *MatrixEmulator) RenderMethod(c *Canvas) error {
	m.Send(UploadEvent{
		leds: *(*c).GetLeds(),
	})
	return nil
}

// Render update the display with the data from the canvas content
func (m *MatrixEmulator) Render(canvas *Canvas) error {
	//start := time.Now()
	//var cnt = 0
	if m.w != nil {
		//var fillDuration = time.Duration(0)
		//var fillStart time.Time

		var ledColor color.Color
		for x := 0; x < m.Width; x++ {
			for y := 0; y < m.Height; y++ {
				ledColor = (*canvas).At(x, y)

				if ledColor != nil {
					//fillStart = time.Now()

					m.w.Fill(m.ledRect(x, y), ledColor, screen.Over)
					//fillDuration += time.Now().Sub(fillStart)
					//cnt++
				}
				ledColor = nil
			}
		}
		//fmt.Println("Render.m.w.fill " + strconv.Itoa(cnt) + " after " + strconv.FormatInt(time.Now().Sub(start).Milliseconds(), 10) + "ms")
		m.w.Publish()
	}
	return nil
}

func (m *MatrixEmulator) Close() error {
	return nil
}

// matrixWithMarginsRect Returns a Rectangle that describes entire emulated RGB Matrix, including margins.
func (m *MatrixEmulator) matrixWithMarginsRect() image.Rectangle {
	upperLeftLED := m.ledRect(0, 0)
	lowerRightLED := m.ledRect(m.Width-1, m.Height-1)
	return image.Rect(upperLeftLED.Min.X-m.Margin, upperLeftLED.Min.Y-m.Margin, lowerRightLED.Max.X+m.Margin, lowerRightLED.Max.Y+m.Margin)
}

// ledRect Returns a Rectangle for the LED at col and row.
func (m *MatrixEmulator) ledRect(col int, row int) image.Rectangle {
	x := (col * (m.PixelPitch + m.Gutter)) + m.Margin
	y := (row * (m.PixelPitch + m.Gutter)) + m.Margin
	return image.Rect(x, y, x+m.PixelPitch, y+m.PixelPitch)
}

func (m *MatrixEmulator) drawBackground(sz size.Event) {
	// Fill entire background with BLACK.
	m.w.Fill(sz.Bounds(), color.Black, screen.Src)
	// Fill matrix display rectangle with the gutter color.
	m.w.Fill(m.matrixWithMarginsRect(), m.GutterColor, screen.Src)

	pixelPlusGutter := m.PixelPitch + m.Gutter
	for col := -1; col < m.Width; col++ {
		x := (col * pixelPlusGutter) + m.Margin + m.PixelPitch
		y := m.Margin
		m.w.Fill(image.Rect(x, y, x+m.Gutter, (m.Height)*pixelPlusGutter+m.Margin+m.Gutter), color.Black, screen.Src)
	}
	for row := -1; row < m.Height; row++ {
		x := m.Margin
		y := (row * pixelPlusGutter) + m.Margin + m.PixelPitch
		m.w.Fill(image.Rect(x, y, (m.Width)*pixelPlusGutter+m.Margin+m.Gutter, y+m.Gutter), color.Black, screen.Src)
	}
}

// calculateGutterForViewableArea As the name states, calculates the size of the gutter for a given viewable area.
// It's easier to understand the geometry of the matrix on screen when put in terms of the gutter,
// hence the shift toward calculating the gutter size.
func (m *MatrixEmulator) calculateGutterForViewableArea(size image.Point) int {
	maxGutterInX := (size.X - 2*m.Margin) / (m.PixelPitchToGutterRatio*m.Width + m.Width - 1)
	maxGutterInY := (size.Y - 2*m.Margin) / (m.PixelPitchToGutterRatio*m.Height + m.Height - 1)
	if maxGutterInX < maxGutterInY {
		return maxGutterInX
	}
	return maxGutterInY
}

func (m *MatrixEmulator) MainThread(canvas *Canvas, done chan struct{}) {
	mainthread.Call(func() {
		driver.Main(func(s screen.Screen) {
			fmt.Println("emulator.MainThread")
			var err error
			m.s = s
			// Calculate initial window size based on whatever our gutter/pixel pitch currently is.
			dims := m.matrixWithMarginsRect()
			m.w, err = s.NewWindow(&screen.NewWindowOptions{
				Title:  windowTitle,
				Width:  dims.Max.X,
				Height: dims.Max.Y,
			})

			if err != nil {
				panic(err)
			}

			defer m.w.Release()
			var sz size.Event
		LOOP:
			for {
				select {
				case <-done:
					break LOOP
				default:
				}
				event := m.w.NextEvent()
				//format := "got %#v\n"
				//if _, ok := event.(fmt.Stringer); ok {
				//	format = "got %v\n"
				//}
				//fmt.Printf(format, event)

				switch evn := event.(type) {
				case lifecycle.Event:
					// this doesn't happen on MacOS, Windows...
					if evn.To == lifecycle.StageDead {
						break LOOP
					}
				case paint.Event:
					fmt.Println("event : paint.Event")
					m.isReady = true
				case UploadEvent:
					if m.isReady {
						m.drawBackground(sz)
						max := (*canvas).Bounds().Max
						var canva Canvas
						canva = &CanvasImpl{
							w:    max.X,
							h:    max.Y,
							leds: evn.leds,
						}
						err = m.Render(&canva)
						if err != nil {
							panic(err)
						}
					}
				case size.Event:
					sz = evn
					m.updatePixelPitchForGutter(m.calculateGutterForViewableArea(sz.Size()))
					if evn.WidthPx == 0 && evn.HeightPx == 0 {
						fmt.Println("event : size.Event > closing window >> leaving UI loop")
						break LOOP
					}
				case error:
					//fmt.Println("event : error")
					fmt.Fprintln(os.Stderr, m)
					//default:
				}
			}
			fmt.Println("screen loop END")
		})
	})
}

func (m *MatrixEmulator) Send(event interface{}) {
	if m.w != nil {
		m.w.Send(event)
	}
}
