package matrix

import (
	"fmt"
	"github.com/faiface/mainthread"
	. "github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/controller"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
)

const DefaultPixelPitch = 12
const windowTitle = "RGB led matrix emulator (shiny)"

type StopEvent struct{}

// UploadEvent signals that the shared pix slice should be uploaded to the
// screen.Texture via the screen.Buffer.
type UploadEvent struct {
	leds []color.Color
}

type MatrixEmulator struct {
	config                  *MatrixConfig
	PixelPitch              int
	Gutter                  int
	Width                   int
	Height                  int
	GutterColor             color.Color
	PixelPitchToGutterRatio int
	Margin                  int

	emulatorKeyboardChannel controller.KeyboardEventChannel

	s       screen.Screen
	w       screen.Window
	isReady bool
}

func NewEmulator(config *MatrixConfig) (Matrix, error) {
	log.Println("NewEmulator")
	w, h := config.Geometry()
	e := &MatrixEmulator{
		config:                  config,
		Width:                   w,
		Height:                  h,
		GutterColor:             color.Gray{Y: 20},
		PixelPitchToGutterRatio: 2,
		Margin:                  10,
		emulatorKeyboardChannel: make(controller.KeyboardEventChannel, 1000),
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

func (m *MatrixEmulator) RenderMethod(c Canvas) error {
	m.send(UploadEvent{
		leds: *c.GetLeds(),
	})
	return nil
}

// Render update the display with the data from the canvas content
func (m *MatrixEmulator) Render(canvas Canvas) error {
	if m.w != nil {
		matrixRectangle := m.matrixWithMarginsRect()
		buffer, err := m.s.NewBuffer(matrixRectangle.Max)
		if err != nil {
			return err
		}

		draw.Draw(buffer.RGBA(), buffer.Bounds(), image.NewUniform(color.Black), image.Point{}, draw.Src)
		// Fill matrix display rectangle with the gutter color.
		draw.Draw(buffer.RGBA(), matrixRectangle, image.NewUniform(m.GutterColor), image.Point{}, screen.Src)
		pixelPlusGutter := m.PixelPitch + m.Gutter
		for col := -1; col < m.Width; col++ {
			x := (col * pixelPlusGutter) + m.Margin + m.PixelPitch
			y := m.Margin
			draw.Draw(buffer.RGBA(), image.Rect(x, y, x+m.Gutter, (m.Height)*pixelPlusGutter+m.Margin+m.Gutter), image.NewUniform(color.Black), image.Point{}, screen.Src)
		}
		for row := -1; row < m.Height; row++ {
			x := m.Margin
			y := (row * pixelPlusGutter) + m.Margin + m.PixelPitch
			draw.Draw(buffer.RGBA(), image.Rect(x, y, (m.Width)*pixelPlusGutter+m.Margin+m.Gutter, y+m.Gutter), image.NewUniform(color.Black), image.Point{}, screen.Src)
		}

		var ledColor color.Color
		for x := 0; x < m.Width; x++ {
			for y := 0; y < m.Height; y++ {
				ledColor = canvas.At(x, y)
				if ledColor != nil {
					draw.Draw(buffer.RGBA(), m.ledRect(x, y), image.NewUniform(ledColor), image.Point{}, draw.Over)
				}
				ledColor = nil
			}
		}
		m.w.Upload(image.Point{}, buffer, matrixRectangle)
		m.w.Publish()
	}
	return nil
}

func (m *MatrixEmulator) Close() error {
	m.send(StopEvent{})
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

func (m *MatrixEmulator) MainThread(canvas Canvas, done chan struct{}) {
	mainthread.Call(func() {
		driver.Main(func(s screen.Screen) {
			log.Println("emulator.MainThread")
			var err error
			m.s = s
			// Calculate initial window size based on whatever our gutter/pixel pitch currently is.
			//dims := m.matrixWithMarginsRect()
			m.w, err = s.NewWindow(&screen.NewWindowOptions{
				Title:  windowTitle,
				Width:  786,
				Height: 786,
				//Width:  dims.Max.X,
				//Height: dims.Max.Y,
			})

			if err != nil {
				panic(err)
			}

			defer m.w.Release()
			var sz size.Event
		LOOP:
			for {
				//select {
				//case <-done:
				//	break LOOP
				//default:
				//}
				event := m.w.NextEvent()

				switch evn := event.(type) {
				case lifecycle.Event:
					// this doesn't happen on MacOS, Windows...
					if evn.To == lifecycle.StageDead {
						break LOOP
					}
				case paint.Event:
					m.isReady = true
				case StopEvent:
					break LOOP
				case UploadEvent:
					if m.isReady {
						max := canvas.Bounds().Max
						err = m.Render(NewSimpleCanvas(
							max.X,
							max.Y,
							&evn.leds,
						))
						if err != nil {
							panic(err)
						}
					}
				case size.Event:
					sz = evn
					m.updatePixelPitchForGutter(m.calculateGutterForViewableArea(sz.Size()))
					if evn.WidthPx == 0 && evn.HeightPx == 0 {
						log.Println("event : size.Event > closing window >> leaving UI loop")
						break LOOP
					}
				case key.Event:
					//format := "got %#v\n"
					//if _, ok := event.(fmt.Stringer); ok {
					//	format = "got %v\n"
					//}
					//fmt.Printf(format, event)
					keyboardEvent := convertKeyboardEvent(evn)
					if keyboardEvent != nil {
						m.emulatorKeyboardChannel <- keyboardEvent
					}
				case error:
					_, _ = fmt.Fprintln(os.Stderr, m)
				default:
				}
			}
			log.Println("screen loop END")
		})
	})
}

func convertKeyboardEvent(e key.Event) *controller.KeyboardEvent {
	var keyboardEvent *controller.KeyboardEvent
	//if e.Rune >= 0 {
	//	fmt.Sprintf("key.Event{%q (%v), %v, %v}", e.Rune, e.Code, e.Modifiers, e.Direction)
	//}
	//return fmt.Sprintf("key.Event{(%v), %v, %v}", e.Code, e.Modifiers, e.Direction)
	if e.Modifiers == key.ModControl && e.Code == key.CodeC {
		log.Println("Exiting...")
		os.Exit(0)
	}
	//r := event.Rune
	action := controller.PressKey
	if e.Direction == key.DirRelease {
		action = controller.ReleaseKey
	}
	if e.Code >= key.CodeA && e.Code <= key.Code0 || e.Code >= key.CodeHyphenMinus && e.Code <= key.CodeSlash {
		keyboardEvent = controller.NewKeyboardEvent(controller.KeyEventTypeChar, action, fmt.Sprint(e.Rune))
	} else if e.Code == key.CodeReturnEnter {
		keyboardEvent = controller.NewKeyboardEvent(controller.KeyEventTypeChar, action, "enter")
	} else if e.Code == key.CodeEscape {
		keyboardEvent = controller.NewKeyboardEvent(controller.KeyEventTypeChar, action, "esc")
	} else if e.Code == key.CodeDeleteBackspace {
		keyboardEvent = controller.NewKeyboardEvent(controller.KeyEventTypeChar, action, "deleteBackward")
		//} else if code == key.CodeTab {
		//	keyboardEvent =
	} else if e.Code == key.CodeSpacebar {
		keyboardEvent = controller.NewKeyboardEvent(controller.KeyEventTypeChar, action, " ")
	} else if e.Code == key.CodeRightArrow {
		keyboardEvent = controller.NewKeyboardEvent(controller.KeyEventTypeChar, action, "right")
	} else if e.Code == key.CodeLeftArrow {
		keyboardEvent = controller.NewKeyboardEvent(controller.KeyEventTypeChar, action, "left")
	} else if e.Code == key.CodeDownArrow {
		keyboardEvent = controller.NewKeyboardEvent(controller.KeyEventTypeChar, action, "down")
	} else if e.Code == key.CodeUpArrow {
		keyboardEvent = controller.NewKeyboardEvent(controller.KeyEventTypeChar, action, "up")
	} else {
		keyboardEvent = nil
	}
	return keyboardEvent
}

func (m *MatrixEmulator) send(event interface{}) {
	if m.w != nil {
		m.w.Send(event)
	}
}
