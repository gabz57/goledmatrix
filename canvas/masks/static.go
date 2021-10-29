package masks

import "image/color"

type StaticCanvasMask struct {
	colors []color.Color
}

//
//func (c *StaticCanvasMask) Set(x, y int, ledColor color.Color) {
//	position := Position(x, y, c.Bounds().Dx())
//	if x >= 0 && y >= 0 && position < len(*c.GetLeds()) {
//		(*c.GetLeds())[position] = c.colors[position]
//	}
//}
//
//// StaticCanvasMask specialization with a single color
//func NewSingleColorMask(canvas Canvas, maskColor color.Color) *StaticCanvasMask {
//	max := canvas.Bounds().Max
//	var mask = make([]color.Color, max.X*max.Y)
//	for col := 0; col < max.X; col++ {
//		for row := 0; row < max.Y; row++ {
//			mask[Position(col, row, canvas.Bounds().Dx())] = maskColor
//		}
//	}
//	return NewCanvasMask(canvas, mask)
//}
//
//func NewCanvasMask(canvas Canvas, colors []color.Color) *StaticCanvasMask {
//	return &StaticCanvasMask{
//		Mask:   canvas,
//		colors: colors,
//	}
//}
