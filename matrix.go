package goledmatrix

import (
	"errors"
	"flag"
	"fmt"
	"github.com/faiface/mainthread"
	"log"
	"os"
	"time"
)

func RunMatrices(app func()) {
	mainthread.Run(app)
}

const MatrixEmulatorENV = "MATRIX_EMULATOR"

// Matrix is an interface that represent any RGB matrix, very useful for testing
type Matrix interface {
	Config() *MatrixConfig
	Geometry() (width, height int)
	// Render update the display with the data from the canvas content
	Render(canvas *Canvas) error
	RenderMethod(canvas *Canvas) error
	Close() error

	MainThread(canvas *Canvas, done chan struct{})
	// extension method to delay Render in UI thread via UI custom event,
	Send(event interface{})
	IsEmulator() bool // can differ from configuration (todo: fix/enhance design ?)
}

type ScanMode int8

const (
	Progressive ScanMode = 0
	Interlaced  ScanMode = 1
)

// DefaultConfig default WS281x configuration
var DefaultConfig = MatrixConfig{
	Emulator:          false,
	Rows:              32,
	Cols:              32,
	ChainLength:       1,
	Parallel:          1,
	PWMBits:           11,
	PWMLSBNanoseconds: 130,
	Brightness:        100,
	ScanMode:          Progressive,
}

// FLAGS
var (
	rows                     = flag.Int("led-rows", 64, "number of rows supported")
	cols                     = flag.Int("led-cols", 64, "number of columns supported")
	parallel                 = flag.Int("led-parallel", 2, "number of daisy-chained panels")
	chain                    = flag.Int("led-chain", 2, "number of displays daisy-chained")
	brightness               = flag.Int("brightness", 100, "brightness (0-100)")
	hardware_mapping         = flag.String("led-gpio-mapping", "regular", "Name of GPIO mapping used.")
	show_refresh             = flag.Bool("led-show-refresh", false, "Show refresh rate.")
	inverse_colors           = flag.Bool("led-inverse", false, "Switch if your matrix has inverse colors on.")
	disable_hardware_pulsing = flag.Bool("led-no-hardware-pulse", false, "Don't use hardware pin-pulse generation.")
	//img                      = flag.String("image", "", "image path")
	//
	//rotate = flag.Int("rotate", 0, "rotate angle, 90, 180, 270")
)

func ReadConfigFlags() *MatrixConfig {
	config := &DefaultConfig
	config.Emulator = os.Getenv(MatrixEmulatorENV) == "1"
	config.Rows = *rows
	config.Cols = *cols
	config.Parallel = *parallel
	config.ChainLength = *chain
	config.Brightness = *brightness
	config.HardwareMapping = *hardware_mapping
	config.ShowRefreshRate = *show_refresh
	config.InverseColors = *inverse_colors
	config.DisableHardwarePulsing = *disable_hardware_pulsing
	return config
}

// MatrixConfig rgb-led-matrix configuration
type MatrixConfig struct {
	Emulator bool
	// Rows the number of rows supported by the display, so 32 or 16.
	Rows int
	// Cols the number of columns supported by the display, so 32 or 64 .
	Cols int
	// ChainLengthis the number of displays daisy-chained together
	// (output of one connected to input of next).
	ChainLength int
	// Parallel is the number of parallel chains connected to the Pi; in old Pis
	// with 26 GPIO pins, that is 1, in newer Pis with 40 interfaces pins, that
	// can also be 2 or 3. The effective number of pixels in vertical direction is
	// then thus rows * parallel.
	Parallel int
	// Set PWM bits used for output. Default is 11, but if you only deal with
	// limited comic-colors, 1 might be sufficient. Lower require less CPU and
	// increases refresh-rate.
	PWMBits int
	// Change the base time-unit for the on-time in the lowest significant bit in
	// nanoseconds.  Higher numbers provide better quality (more accurate color,
	// less ghosting), but have a negative impact on the frame rate.
	PWMLSBNanoseconds int // the DMA channel to use
	// Brightness is the initial brightness of the panel in percent. Valid range
	// is 1..100
	Brightness int
	// ScanMode progressive or interlaced
	ScanMode ScanMode // strip color layout
	// Disable the PWM hardware subsystem to create pulses. Typically, you don't
	// want to disable hardware pulsing, this is mostly for debugging and figuring
	// out if there is interference with the sound system.
	// This won't do anything if output enable is not connected to GPIO 18 in
	// non-standard wirings.
	DisableHardwarePulsing bool

	ShowRefreshRate bool
	InverseColors   bool

	// Name of GPIO mapping used
	HardwareMapping string
}

func (conf *MatrixConfig) Geometry() (width, height int) {
	return conf.Cols * conf.ChainLength, conf.Rows * conf.Parallel
}

// UploadEvent signals that the shared pix slice should be uploaded to the
// screen.Texture via the screen.Buffer.
type UploadEvent struct{}

func BuildMatrix(config *MatrixConfig) (m Matrix, err error) {
	if config.Emulator == true {
		return NewMatrixEmulator(config)
	} else {
		return NewMatrixEmulator(config)
		//return NewRGBLedMatrix(config)
	}
}

func Run(matrixCreator func(config *MatrixConfig) (Matrix, error), gameloop func(c *Canvas, done chan struct{})) {
	RunMany([]func(config *MatrixConfig) (Matrix, error){matrixCreator}, gameloop)
}

func RunMany(matrixCreators []func(config *MatrixConfig) (Matrix, error), gameloop func(c *Canvas, done chan struct{})) {
	fmt.Println("Running...")

	config := ReadConfigFlags()

	if len(matrixCreators) == 0 {
		log.Fatal("No matrix defined !")
	}

	canvas := NewCanvas(config)
	var w, h *int
	for _, matrixCreator := range matrixCreators {
		matrix, err := matrixCreator(config)

		if err != nil {
			log.Fatal(err)
		}
		width, height := matrix.Geometry()
		err = validateGeometry(width, height, w, h)
		if err != nil {
			log.Fatal(err)
		}
		canvas.register(matrix)
	}

	done := make(chan struct{})
	defer canvas.Close()

	// Starting game loop on a separate routine
	go func() {
		// avoid drawing to early as emulator might not be ready, eventually fixed
		<-time.After(1000 * time.Millisecond)

		fmt.Println("Starting Gameloop !")
		gameloop(canvas, done)
		fmt.Println("Gameloop END")
		done <- struct{}{}
	}()

	// run all matrices (UI is run on main thread)
	mainMatrix, otherMatrices := splitMatrices(&canvas.matrices)
	for _, otherMatrix := range otherMatrices {
		fmt.Println("go matrix.MainThread()")
		go (*otherMatrix).MainThread(canvas, done)
	}
	fmt.Println("matrix.MainThread()")
	(*mainMatrix).MainThread(canvas, done)
}

func splitMatrices(ms *[]Matrix) (mainMatrix *Matrix, others []*Matrix) {
	for _, matrix := range *ms {
		if matrix.IsEmulator() {
			mainMatrix = &matrix
		} else {
			others = append(others, &matrix)
		}
	}
	if mainMatrix == nil {
		if len(others) > 0 {
			mainMatrix = others[0]
		}
		if len(others) >= 2 {
			others = others[(1):]
		}
	}
	return mainMatrix, others
}

func validateGeometry(width, height int, w, h *int) error {
	if w == nil {
		w = &width
	} else {
		if *w != width {
			return errors.New("Incorrect WIDTH detected between matrices ")
		}
	}
	if h == nil {
		h = &height
	} else {
		if *h != height {
			return errors.New("Incorrect HEIGHT detected between matrices ")
		}
	}
	return nil
}
