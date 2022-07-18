package matrix

import (
	"errors"
	"flag"
	"github.com/faiface/mainthread"
	. "github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/controller"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunMatrices(app func()) {
	mainthread.Run(app)
}

// Matrix is an interface that represent any RGB matrix, very useful for testing
type Matrix interface {
	Config() *MatrixConfig
	Geometry() (width, height int)
	// Render update the display with the data from the canvas content
	Render(canvas Canvas) error
	RenderMethod(canvas Canvas) error
	Close() error

	MainThread(canvas Canvas, done chan struct{})
	// extension method to delay Render in UI thread via UI custom event,
	send(event interface{})
}

type ScanMode int8

const (
	Progressive ScanMode = 0
	Interlaced  ScanMode = 1
)

// DefaultConfig default WS281x configuration
var DefaultConfig = MatrixConfig{
	Rows:                   32,
	Cols:                   32,
	ChainLength:            1,
	Parallel:               1,
	PWMBits:                11,
	PWMLSBNanoseconds:      130,
	Brightness:             100,
	ScanMode:               Progressive,
	DisableHardwarePulsing: false,
	ShowRefreshRate:        false,
	InverseColors:          false,
	HardwareMapping:        "",
	LedPixelMapper:         "",
	Emulator:               false,
	Client:                 false,
	Server:                 false,
}

// FLAGS
// NOTE: reading flags overwrites DefaultConfig values
var (
	rows                     = flag.Int("led-rows", 64, "number of rows supported")
	cols                     = flag.Int("led-cols", 64, "number of columns supported")
	parallel                 = flag.Int("led-parallel", 1, "number of daisy-chained panels")
	chain                    = flag.Int("led-chain", 4, "number of displays daisy-chained")
	brightness               = flag.Int("brightness", 100, "brightness (0-100)")
	hardware_mapping         = flag.String("led-gpio-mapping", "adafruit-hat-pwm", "Name of GPIO mapping used.")
	show_refresh             = flag.Bool("led-show-refresh", false, "Show refresh rate.")
	inverse_colors           = flag.Bool("led-inverse", false, "Switch if your matrix has inverse colors on.")
	disable_hardware_pulsing = flag.Bool("led-no-hardware-pulse", true, "Don't use hardware pin-pulse generation.")
	led_pixel_mapper         = flag.String("led-pixel-mapper", "U-mapper;Rotate:180", "Semicolon-separated list of pixel-mappers to arrange pixels.")
	//img                      = flag.String("image", "", "image path")
	//
	//rotate = flag.Int("rotate", 0, "rotate angle, 90, 180, 270")
)

const (
	MatrixEmulatorENV        = "MATRIX_EMULATOR"
	MatrixClientENV          = "MATRIX_CLIENT"
	MatrixServerENV          = "MATRIX_SERVER"
	MatrixServerIpAddressENV = "MATRIX_ADDRESS"
)

func ReadConfigFlags() (*MatrixConfig, error) {
	config := &DefaultConfig
	config.Rows = *rows
	config.Cols = *cols
	config.Parallel = *parallel
	config.ChainLength = *chain
	config.Brightness = *brightness
	config.PWMBits = 8
	config.HardwareMapping = *hardware_mapping
	config.ShowRefreshRate = *show_refresh
	config.InverseColors = *inverse_colors
	config.DisableHardwarePulsing = *disable_hardware_pulsing
	config.LedPixelMapper = *led_pixel_mapper
	config.Emulator = os.Getenv(MatrixEmulatorENV) == "1"
	config.Server = os.Getenv(MatrixServerENV) == "1"
	config.Client = os.Getenv(MatrixClientENV) == "1"
	config.IpAddress = os.Getenv(MatrixServerIpAddressENV)

	if config.Client && len(config.IpAddress) <= 7 {
		return nil, errors.New("client mode is enabled but server IP is missing, complete Environment variables with " + MatrixServerIpAddressENV + " value")
	}
	return config, nil
}

// MatrixConfig rgb-led-matrix configuration
type MatrixConfig struct {
	// Rows the number of rows supported by the display, so 32 or 16.
	Rows int
	// Cols the number of columns supported by the display, so 32 or 64 .
	Cols int
	// ChainLength is the number of displays daisy-chained together
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
	// Semicolon-separated list of pixel-mappers to arrange pixels.
	LedPixelMapper string
	// Using OpenGL emulator instead of driving hardware matrix
	Emulator bool
	// Driving remote matrix using GoRPC
	Client bool
	// Serving hardware matrix using GoRPC
	Server bool
	// Remote server address
	IpAddress string
}

func validateGeometry(config *MatrixConfig, remoteMatrix Matrix) error {
	width, height := config.Geometry()
	clientW, clientH := remoteMatrix.Geometry()
	if width != clientW {
		return errors.New("incorrect WIDTH detected between local (" + strconv.Itoa(width) + ") and received from remote (" + strconv.Itoa(clientW) + ")")
	}
	if height != clientH {
		return errors.New("incorrect HEIGHT detected between local (" + strconv.Itoa(height) + ") and received from remote (" + strconv.Itoa(clientH) + ")")
	}
	return nil
}

func (mc *MatrixConfig) Geometry() (width, height int) {
	var mapper string
	mapper = mc.LedPixelMapper
	if strings.Contains(mapper, "U-mapper") {
		return mc.Cols * mc.ChainLength / 2, mc.Rows * mc.Parallel * 2
	}
	return mc.Cols * mc.ChainLength, mc.Rows * mc.Parallel
}

func Run(gameloop func(_ Canvas, _ chan struct{}, _ *controller.KeyboardEventChannel)) {
	log.Println("Running...")
	config, err := ReadConfigFlags()
	if err != nil {
		log.Fatal(err)
	}
	matrix, err := BuildMatrix(config)
	if err != nil {
		log.Fatal(err)
	}
	canvas := NewCanvas(config, matrix)

	done := make(chan struct{})

	// Starting game loop on a separate routine
	go run(func(c Canvas, done chan struct{}, kbEventChannel *controller.KeyboardEventChannel) {
		if config.Server {
			// Only listen to matrix received by rpc (Ctrl+C to stop)
			RpcServe(matrix)(c, done)
		} else {
			// modify method interface to pass keyboard channel in gameloop
			gameloop(c, done, kbEventChannel)
		}
		err := canvas.Close()
		if err != nil {
			panic(err)
		}
	}, canvas, done, emulatorKeyboardChannel(matrix))

	log.Println("Starting LedMatrix")
	matrix.MainThread(canvas, done)
	log.Println("LedMatrix Stopped")
}

func emulatorKeyboardChannel(matrix Matrix) *controller.KeyboardEventChannel {
	var keyboardChannel *controller.KeyboardEventChannel
	emu, ok := matrix.(*MatrixEmulator)
	if ok {
		keyboardChannel = &emu.emulatorKeyboardChannel
	}
	return keyboardChannel
}

func run(gameloop func(c Canvas, done chan struct{}, keyboardChannel *controller.KeyboardEventChannel), canvas Canvas, done chan struct{}, keyboardChannel *controller.KeyboardEventChannel) {
	func() {
		// avoid drawing to early as emulator might not be ready, eventually fixed
		//<-time.After(10 * time.Millisecond)

		log.Println("Gameloop STARTED")
		gameloop(canvas, done, keyboardChannel)
		log.Println("Gameloop END")
		done <- struct{}{}
	}()
}
