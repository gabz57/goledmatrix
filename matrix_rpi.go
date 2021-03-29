package goledmatrix

/*
#cgo CFLAGS: -std=c99 -I${SRCDIR}/vendor/rpi-rgb-led-matrix/include -DSHOW_REFRESH_RATE
#cgo LDFLAGS: -lrgbmatrix -L${SRCDIR}/vendor/rpi-rgb-led-matrix/lib -lstdc++ -lm
#include <led-matrix-c.h>

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
	"unsafe"
)

func (c *MatrixConfig) toC() *C.struct_RGBLedMatrixOptions {
	o := &C.struct_RGBLedMatrixOptions{}
	o.rows = C.int(c.Rows)
	o.cols = C.int(c.Cols)
	o.chain_length = C.int(c.ChainLength)
	o.parallel = C.int(c.Parallel)
	o.pwm_bits = C.int(c.PWMBits)
	o.pwm_lsb_nanoseconds = C.int(c.PWMLSBNanoseconds)
	o.brightness = C.int(c.Brightness)
	o.scan_mode = C.int(c.ScanMode)
	o.hardware_mapping = C.CString(c.HardwareMapping)

	if c.ShowRefreshRate == true {
		C.set_show_refresh_rate(o, C.int(1))
	} else {
		C.set_show_refresh_rate(o, C.int(0))
	}

	if c.DisableHardwarePulsing == true {
		C.set_disable_hardware_pulsing(o, C.int(1))
	} else {
		C.set_disable_hardware_pulsing(o, C.int(0))
	}

	if c.InverseColors == true {
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
	//leds   []C.uint32_t
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
	c = &MatrixHardware{
		config: config,
		width:  w, height: h,
		matrix: m,
		buffer: b,
		//leds:   make([]C.uint32_t, w*h),
	}
	if m == nil {
		return nil, fmt.Errorf("unable to allocate memory")
	}

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

//
//// FIXME: DO NOT USE Set method anymore (Canvas concern)
//// Set set LED at position x,y to the provided 24-bit color value.
//func (m *MatrixHardware) Set(x, y int, color color.Color) {
//
//
//	var int32Color uint32 = C.uint32_t(colorToUint32(color))
//	C.led_canvas_set_pixel(
//		m.buffer,
//		x, y,
//		(int32Color >> 16) & 255,
//		(int32Color >> 8) & 255,
//		int32Color & 255)
//}

func (m *MatrixHardware) RenderMethod(canvas *Canvas) error {
	return m.Render(canvas)
}

// Render update the display with the data from the LED buffer
// NOTE: canvas is unused here as data are already set to buffer
func (m *MatrixHardware) Render(canvas *Canvas) error {
	C.led_matrix_swap(
		m.matrix,
		m.buffer,
		C.int(canvas.w), C.int(canvas.h),
		(*C.uint32_t)(unsafe.Pointer(&canvas.leds[0])),
	)

	canvas.leds = make([]C.uint32_t, canvas.w*canvas.h)

	m.buffer = C.led_matrix_swap_on_vsync(m.matrix, m.buffer)
	return nil
}

// Close finalizes the ws281x interface
func (m *MatrixHardware) Close() error {
	C.led_matrix_delete(m.matrix)
	return nil
}

//func colorToUint32(c color.Color) uint32 {
//	if c == nil {
//		return 0
//	}
//	// A color's RGBA method returns values in the range [0, 65535]
//	red, green, blue, _ := c.RGBA()
//	return (red>>8)<<16 | (green>>8)<<8 | blue>>8
//}
//
//
//func uint32ToColor(u C.uint32_t) color.Color {
//	return color.RGBA{
//		R: uint8(u>>16) & 255,
//		G: uint8(u>>8) & 255,
//		B: uint8(u>>0) & 255,
//	}
//}

func (m *MatrixHardware) MainThread(canvas *Canvas, done chan struct{}) {
	select {
	case <-done:
		break
	}
}

func (m *MatrixHardware) Send(event interface{}) {
	panic("implement me")
}

func (m *MatrixHardware) IsEmulator() bool {
	return false
}