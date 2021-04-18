package shapes

import (
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/fonts"
	. "github.com/gabz57/goledmatrix/matrix"
	"image/color"
	"time"
)

type TextPanel struct {
	panel        *Panel
	text         *Text
	value        string
	font         fonts.MatrixFont
	cornerRadius int
	graphic      *components.Graphic
	textGraphic  *components.Graphic
}

func NewTextPanel(text string, textColor color.Color, panelLayout *components.Layout, position, dimension Point, cornerRadius int, fill, border bool, font fonts.MatrixFont) *TextPanel {
	var textPanelGraphic = components.NewOffsetGraphic(nil, panelLayout, position)
	panel := TextPanel{
		graphic:      textPanelGraphic,
		panel:        NewPanel(textPanelGraphic, panelLayout, Point{}, dimension, cornerRadius, fill, border),
		value:        text,
		font:         font,
		cornerRadius: cornerRadius,
	}
	panel.textGraphic = components.NewGraphic(textPanelGraphic, components.NewLayout(textColor, nil))
	panel.SetText(text)
	return &panel
}

func (t *TextPanel) Update(elapsedBetweenUpdate time.Duration) {

}

func (t *TextPanel) Draw(canvas Canvas) error {
	err := t.panel.Draw(canvas)
	if err != nil {
		return err
	}
	return t.text.Draw(canvas)
}

func (t *TextPanel) SetText(text string) {
	// TODO: using Font, split 'text' into multiple 'dashed lines' to fill the panel, end with '...'
	metrics := fonts.GetFont(t.font).Metrics()
	height := int(metrics.Height / 64)
	var xAdjust, yAdjust int
	if height <= 7 {
		yAdjust = 1
	} else if height <= 9 {
		yAdjust = 2
	} else if height <= 11 {
		yAdjust = 3
	} else if height <= 12 {
		yAdjust = 5
	} else if height <= 13 {
		yAdjust = 4
	} else {
		xAdjust = 1
		yAdjust = 5
	}
	var textPosition = Point{
		X: -xAdjust,
		Y: -yAdjust,
	}
	if t.panel.cornerRadius > 0 {
		textPosition = textPosition.AddXY(
			t.cornerRadius/4,
			t.cornerRadius/4,
		)
	}
	if t.panel.border {
		textPosition = textPosition.AddXY(2, 2)
	}

	t.text = NewText(t.textGraphic, textPosition, text, t.font)
}
