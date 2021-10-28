package shapes

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/fonts"
	"image/color"
	"time"
)

type TextPanel struct {
	panel        *Panel
	text         *Text
	font         fonts.MatrixFont
	cornerRadius int
	graphic      *Graphic
	textGraphic  *Graphic
	dirty        bool
}

func NewTextPanel(text string, textColor color.Color, panelLayout *Layout, position, dimension Point, cornerRadius int, fill, border bool, font fonts.MatrixFont) *TextPanel {
	var textPanelGraphic = NewOffsetGraphic(nil, panelLayout, position)
	panel := TextPanel{
		graphic:      textPanelGraphic,
		panel:        NewPanel(textPanelGraphic, panelLayout, Point{}, dimension, cornerRadius, fill, border),
		font:         font,
		cornerRadius: cornerRadius,
		dirty:        true,
	}
	panel.textGraphic = NewGraphic(textPanelGraphic, NewLayout(textColor, nil))
	panel.SetText(text)
	return &panel
}

func (t *TextPanel) Update(elapsedBetweenUpdate time.Duration) bool {
	return t.dirty
}

func (t *TextPanel) Draw(canvas Canvas) error {
	err := t.panel.Draw(canvas)
	if err != nil {
		return err
	}
	err = t.text.Draw(canvas)
	if err != nil {
		return err
	}
	t.dirty = false
	return err
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
	t.dirty = true
}
