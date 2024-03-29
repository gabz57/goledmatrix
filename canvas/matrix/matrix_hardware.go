//go:build !darwin
// +build !darwin

package matrix

/*
#cgo CFLAGS: -std=c99 -I${SRCDIR}/../../vendor/rpi-rgb-led-matrix/include -DSHOW_REFRESH_RATE
#cgo LDFLAGS: -lrgbmatrix -L${SRCDIR}/../../vendor/rpi-rgb-led-matrix/lib -lstdc++ -lm
#include <led-matrix-c.h>

void led_matrix_swap(struct RGBLedMatrix *matrix, struct LedCanvas *offscreen_canvas,
                     int width, int height, const uint32_t pixels[]) {
  int i, x, y;
  uint32_t color;
  for (x = 0; x < width; ++x) {
    for (y = 0; y < height; ++y) {
      i = x + (y * width);
      color = pixels[i];

      led_canvas_set_pixel(offscreen_canvas, x, y,
        (color >> 16) & 255, (color >> 8) & 255, color & 255);
    }
  }

  offscreen_canvas = led_matrix_swap_on_vsync(matrix, offscreen_canvas);
}

void set_show_refresh_rate(struct RGBLedMatrixOptions *o, int show_refresh_rate) {
o->show_refresh_rate = show_refresh_rate != 0 ? 1 : 0;
}

void set_disable_hardware_pulsing(struct RGBLedMatrixOptions *o, int disable_hardware_pulsing) {
o->disable_hardware_pulsing = disable_hardware_pulsing != 0 ? 1 : 0;
}

void set_inverse_colors(struct RGBLedMatrixOptions *o, int inverse_colors) {
o->inverse_colors = inverse_colors != 0 ? 1 : 0;
}
*/
import "C"
import (
	"fmt"
	. "github.com/gabz57/goledmatrix/canvas"
	"github.com/gosuri/uilive"
	"image/color"
	"unsafe"
)

func (mc *MatrixConfig) toC() *C.struct_RGBLedMatrixOptions {
	o := &C.struct_RGBLedMatrixOptions{}
	o.rows = C.int(mc.Rows)
	o.cols = C.int(mc.Cols)
	o.chain_length = C.int(mc.ChainLength)
	o.parallel = C.int(mc.Parallel)
	o.pwm_bits = C.int(mc.PWMBits)
	o.pwm_lsb_nanoseconds = C.int(mc.PWMLSBNanoseconds)
	o.brightness = C.int(mc.Brightness)
	o.scan_mode = C.int(mc.ScanMode)
	o.hardware_mapping = C.CString(mc.HardwareMapping)
	o.pixel_mapper_config = C.CString(mc.LedPixelMapper)
	if mc.ShowRefreshRate == true {
		C.set_show_refresh_rate(o, C.int(1))
	} else {
		C.set_show_refresh_rate(o, C.int(0))
	}

	if mc.DisableHardwarePulsing == true {
		C.set_disable_hardware_pulsing(o, C.int(1))
	} else {
		C.set_disable_hardware_pulsing(o, C.int(0))
	}

	if mc.InverseColors == true {
		C.set_inverse_colors(o, C.int(1))
	} else {
		C.set_inverse_colors(o, C.int(0))
	}

	return o
}

// MatrixHardware matrix representation for ws281x
type MatrixHardware struct {
	config *MatrixConfig
	height int
	width  int
	matrix *C.struct_RGBLedMatrix
	buffer *C.struct_LedCanvas
	writer uilive.Writer
}

func (m *MatrixHardware) send(_ interface{}) {
	// implemented only in emulator which handle data asynchronously
}

// NewRGBLedMatrix returns a new matrix using the given size and config
func NewRGBLedMatrix(config *MatrixConfig) (c Matrix, err error) {
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
	m := C.led_matrix_create_from_options(config.toC(), nil, nil)
	b := C.led_matrix_create_offscreen_canvas(m)
	writer := uilive.New()
	c = &MatrixHardware{
		config: config,
		width:  w, height: h,
		matrix: m,
		buffer: b,
		writer: *writer,
	}
	if m == nil {
		return nil, fmt.Errorf("unable to allocate memory")
	}
	writer.Start()
	return c, nil
}

// Geometry returns the width and the height of the matrix
func (m *MatrixHardware) Config() *MatrixConfig {
	return m.config
}

// Geometry returns the width and the height of the matrix
func (m *MatrixHardware) Geometry() (width, height int) {
	return m.width, m.height
}

func (m *MatrixHardware) RenderMethod(canvas Canvas) error {
	return m.Render(canvas)
}

// Render update the display with the data from the LED buffer
func (m *MatrixHardware) Render(canvas Canvas) error {
	//start := time.Now()
	canvasSize := canvas.Bounds().Max
	leds := make([]C.uint32_t, canvasSize.X*canvasSize.Y)
	for i, led := range *canvas.GetLeds() {
		if led != nil {
			leds[i] = C.uint32_t(colorToUint32(led))
		}
	}
	//copyDuration := time.Now().Sub(start)

	//start = time.Now()
	C.led_matrix_swap(
		m.matrix,
		m.buffer,
		C.int(canvasSize.X), C.int(canvasSize.Y),
		(*C.uint32_t)(unsafe.Pointer(&leds[0])),
	)
	//swapDuration := time.Now().Sub(start)

	//_, _ = fmt.Fprintf(
	//m.writer.Newline(),
	//"copy: "+strconv.FormatInt(copyDuration.Milliseconds(), 10)+" ms - " +
	//	"swap: "+strconv.FormatInt(swapDuration.Milliseconds(), 10)+" ms\n")

	return nil
}

// Close finalizes the ws281x interface
func (m *MatrixHardware) Close() error {
	C.led_matrix_delete(m.matrix)
	return nil
}

func colorToUint32(c color.Color) uint32 {
	if c == nil {
		return 0
	}
	// A color's RGBA method returns values in the range [0, 65535]
	red, green, blue, _ := c.RGBA()
	return (red>>8)<<16 | (green>>8)<<8 | blue>>8
}

func (m *MatrixHardware) MainThread(_ Canvas, done chan struct{}) {
	select {
	case <-done:
		break
	}
}

func (m *MatrixHardware) Send(_ interface{}) {
}
