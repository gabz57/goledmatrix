package shapes

import (
	"fmt"
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/fonts"
	"golang.org/x/image/font"
	"strconv"
	"time"
)

type Text struct {
	*Graphic
	font     font.Face
	position Point
	txt      string
}

func NewText(graphic *Graphic, position Point, txt string, f fonts.MatrixFont) *Text {
	return &Text{
		Graphic:  graphic,
		position: position,
		txt:      txt,
		font:     fonts.GetFont(f),
	}
}

func (t *Text) Draw(canvas *Canvas) error {
	start := time.Now()
	position := t.position.Add(t.ComputedOffset())
	canvas.DrawLabel(position.X, position.Y+t.font.Metrics().Height.Ceil(), t.txt, *t.Layout().Color(), t.font)
	fmt.Println("Text: " + strconv.FormatInt(time.Now().Sub(start).Milliseconds(), 10) + " ms")
	return nil
}

func (t *Text) SetText(txt string) {
	t.txt = txt
}
