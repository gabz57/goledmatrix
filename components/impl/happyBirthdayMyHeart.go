package impl

import (
	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/fonts"
	"time"
)

type HappyBirthday struct {
	joyeuxText       *shapes.Text
	anniversaireText *shapes.Text
	tatianaTextPanel *shapes.TextPanel
	movingHearts     *MovingHearts
	iText            *shapes.Text
	redHeart         *HeartRed
	uText            *shapes.Text
}

func NewHappyBirthday(canvas Canvas) *HappyBirthday {
	happyBirthdayGraphic := NewGraphic(nil, nil)
	textLayout := NewLayout(ColorViolet, nil)
	iLoveUTextLayout := NewLayout(ColorRed, nil)
	iLoveUOrigin := Point{
		X: 18,
		Y: 110,
	}
	c := HappyBirthday{
		joyeuxText: shapes.NewText(
			NewGraphic(happyBirthdayGraphic, textLayout),
			Point{
				X: 10,
				Y: 10,
			},
			"Joyeux",
			fonts.Bdf7x13,
		),
		anniversaireText: shapes.NewText(
			NewGraphic(happyBirthdayGraphic, textLayout),
			Point{
				X: 20,
				Y: 27,
			},
			"Anniversaire",
			fonts.Bdf7x13,
		),
		tatianaTextPanel: shapes.NewTextPanel(
			"Tatiana",
			ColorBlack,
			NewLayout(ColorRed, ColorWhite),
			Point{X: 40, Y: 65},
			Point{X: 55, Y: 15},
			7,
			true,
			true,
			fonts.Bdf7x13,
		),
		movingHearts: NewMovingHearts(canvas, Point{}, 10),
		iText: shapes.NewText(
			NewGraphic(happyBirthdayGraphic, iLoveUTextLayout),
			iLoveUOrigin,
			"with ",
			fonts.Bdf5x7,
		),
		redHeart: NewHeartRed(canvas,
			nil,
			iLoveUOrigin.AddXY(5*5-2, -1),
			time.Duration(Int64Between(2000, 5000))*time.Millisecond,
			0,
			true),
		uText: shapes.NewText(
			NewGraphic(happyBirthdayGraphic, iLoveUTextLayout),
			iLoveUOrigin.AddXY(5*5+1+HeartWidth-2*2, 0),
			" from Arnaud",
			fonts.Bdf5x7,
		),
	}
	return &c
}

func (c *HappyBirthday) Update(elapsedBetweenUpdate time.Duration) {
	c.tatianaTextPanel.Update(elapsedBetweenUpdate)
	c.movingHearts.Update(elapsedBetweenUpdate)
	c.redHeart.Update(elapsedBetweenUpdate)
}

func (c *HappyBirthday) Draw(canvas Canvas) error {
	err := c.joyeuxText.Draw(canvas)
	if err != nil {
		return err
	}
	err = c.anniversaireText.Draw(canvas)
	if err != nil {
		return err
	}
	err = c.tatianaTextPanel.Draw(canvas)
	if err != nil {
		return err
	}
	err = c.iText.Draw(canvas)
	if err != nil {
		return err
	}
	err = c.redHeart.Draw(canvas)
	if err != nil {
		return err
	}
	err = c.uText.Draw(canvas)
	if err != nil {
		return err
	}
	return c.movingHearts.Draw(canvas)
}
