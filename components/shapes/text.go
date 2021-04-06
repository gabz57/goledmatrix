package shapes

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/fonts"
)

type Text struct {
	*Graphic
	position Point
	txt      string
}

func NewText(graphic *Graphic, position Point, txt string) *Text {
	return &Text{
		Graphic:  graphic,
		position: position,
		txt:      txt,
	}
}

func (t *Text) Draw(canvas *Canvas) error {
	position := t.position.Add(t.ComputedOffset())
	font := fonts.GetFont(fonts.Bdf7x13)
	canvas.DrawLabel(position.X, position.Y+font.Metrics().Height.Ceil(), t.txt, ColorBlue, font)
	return nil
}

func (t *Text) SetText(txt string) {
	t.txt = txt
}
