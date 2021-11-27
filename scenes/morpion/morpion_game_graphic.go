package morpion

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/fonts"
	"github.com/gabz57/goledmatrix/game"
	"image/color"
)

var gridTopLeft = canvas.Point{X: 10, Y: 4}
var cellSize = 33
var ringRadius = 9
var padding = 2
var contourSize = cellSize*3 + 4*padding

type GridGraphicComponent struct {
	game.GraphicComponentBase
	contour *shapes.Rectangle
}

func newGridGraphicComponent(entity *MorpionGameEntity) *GridGraphicComponent {
	graphic := components.NewGraphic(nil, components.NewLayout(components.ColorWhite, color.Transparent))
	graphicComponent := GridGraphicComponent{
		GraphicComponentBase: *game.NewGraphicComponentBase(&entity.Values, graphic, true, true),
	}
	graphicComponent.contour = shapes.NewRectangle(graphic, gridTopLeft, canvas.Point{X: contourSize, Y: contourSize}, false)
	graphicComponent.RegisterDrawable(graphicComponent.contour)

	return &graphicComponent
}

type TokenGraphicComponent struct {
	game.GraphicComponentBase
	contour *shapes.Rectangle
	cross1  *shapes.Line
	cross2  *shapes.Line
	cross3  *shapes.Line
	cross4  *shapes.Line
	cross5  *shapes.Line
	cross6  *shapes.Line
	circle  *shapes.Ring
}

var notSelectedColor = components.ColorWhite
var selectedColor = components.ColorRed
var winnerColor = components.ColorGreen

func newTokenGraphicComponent(entity *MorpionGameEntity, positionRef PositionRef) *TokenGraphicComponent {
	graphic := components.NewGraphic(nil, components.NewLayout(notSelectedColor, color.Transparent))
	graphicComponent := TokenGraphicComponent{
		GraphicComponentBase: *game.NewGraphicComponentBase(&entity.Values, graphic, true, true),
	}
	xCoef := int(positionRef) % 3
	yCoef := int(positionRef) / 3
	contourTopLeft := gridTopLeft.Add(canvas.Point{
		X: padding + xCoef*(cellSize+padding),
		Y: padding + yCoef*(cellSize+padding),
	})
	graphicComponent.contour = shapes.NewRectangle(graphic, contourTopLeft, canvas.Point{X: cellSize, Y: cellSize}, false)
	graphicComponent.RegisterDrawable(graphicComponent.contour)

	oGraphic := components.NewGraphic(graphic, components.NewLayout(PlayerOColor, PlayerOColor))
	graphicComponent.circle = shapes.NewRing(
		oGraphic,
		contourTopLeft.Add(canvas.Point{X: cellSize / 2, Y: cellSize / 2}),
		ringRadius+1,
		ringRadius+2,
		true,
	)
	graphicComponent.RegisterDrawable(graphicComponent.circle)

	xGraphic := components.NewGraphic(graphic, components.NewLayout(PlayerXColor, PlayerXColor))
	graphicComponent.cross1 = shapes.NewLine(xGraphic, contourTopLeft.Add(canvas.Point{X: 6, Y: 5}), contourTopLeft.Add(canvas.Point{X: 28, Y: 27}))
	graphicComponent.cross2 = shapes.NewLine(xGraphic, contourTopLeft.Add(canvas.Point{X: 5, Y: 5}), contourTopLeft.Add(canvas.Point{X: 28, Y: 28}))
	graphicComponent.cross3 = shapes.NewLine(xGraphic, contourTopLeft.Add(canvas.Point{X: 5, Y: 6}), contourTopLeft.Add(canvas.Point{X: 27, Y: 28}))
	graphicComponent.RegisterDrawable(graphicComponent.cross1)
	graphicComponent.RegisterDrawable(graphicComponent.cross2)
	graphicComponent.RegisterDrawable(graphicComponent.cross3)
	graphicComponent.cross4 = shapes.NewLine(xGraphic, contourTopLeft.Add(canvas.Point{X: 6, Y: 28}), contourTopLeft.Add(canvas.Point{X: 28, Y: 6}))
	graphicComponent.cross5 = shapes.NewLine(xGraphic, contourTopLeft.Add(canvas.Point{X: 5, Y: 28}), contourTopLeft.Add(canvas.Point{X: 28, Y: 5}))
	graphicComponent.cross6 = shapes.NewLine(xGraphic, contourTopLeft.Add(canvas.Point{X: 5, Y: 27}), contourTopLeft.Add(canvas.Point{X: 27, Y: 5}))
	graphicComponent.RegisterDrawable(graphicComponent.cross4)
	graphicComponent.RegisterDrawable(graphicComponent.cross5)
	graphicComponent.RegisterDrawable(graphicComponent.cross6)
	graphicComponent.setPlayer(NoPlayer)
	return &graphicComponent
}

func (tgc *TokenGraphicComponent) setSelected(selected bool) {
	cellColor := notSelectedColor
	if selected {
		cellColor = selectedColor
	}
	tgc.contour.Graphic.Layout().SetColor(cellColor)
}

func (tgc *TokenGraphicComponent) setPlayer(ref PlayerRef) {
	switch ref {
	case NoPlayer:
		tgc.DisableDrawable(tgc.circle)
		tgc.DisableDrawable(tgc.cross1)
		tgc.DisableDrawable(tgc.cross2)
		tgc.DisableDrawable(tgc.cross3)
		tgc.DisableDrawable(tgc.cross4)
		tgc.DisableDrawable(tgc.cross5)
		tgc.DisableDrawable(tgc.cross6)
	case PlayerO:
		tgc.EnableDrawable(tgc.circle)
	case PlayerX:
		tgc.EnableDrawable(tgc.cross1)
		tgc.EnableDrawable(tgc.cross2)
		tgc.EnableDrawable(tgc.cross3)
		tgc.EnableDrawable(tgc.cross4)
		tgc.EnableDrawable(tgc.cross5)
		tgc.EnableDrawable(tgc.cross6)
	}
}

func (tgc *TokenGraphicComponent) setWinnerCell() {
	tgc.contour.Graphic.Layout().SetColor(winnerColor)
}

type TextGraphicComponent struct {
	game.GraphicComponentBase
	playerText *shapes.Text
	text       *shapes.Text
}

var NoPlayerColor = color.Transparent
var PlayerOColor = components.ColorRed
var PlayerXColor = components.ColorBlue

func (tgc *TextGraphicComponent) setNextPlayer(player PlayerRef) {
	tgc.setPlayer(player)
	tgc.text.SetText(" turn")
}

func (tgc *TextGraphicComponent) setWinner(player PlayerRef) {
	tgc.setPlayer(player)
	tgc.text.SetText(" Wins")
}

func (tgc *TextGraphicComponent) setPlayer(player PlayerRef) {
	switch player {
	case NoPlayer:
		tgc.playerText.SetColor(NoPlayerColor)
		tgc.playerText.SetText("")
	case PlayerO:
		tgc.playerText.SetColor(PlayerOColor)
		tgc.playerText.SetText("O")
	case PlayerX:
		tgc.playerText.SetColor(PlayerXColor)
		tgc.playerText.SetText("X")
	}
}

func (tgc *TextGraphicComponent) setExAequo() {
	tgc.playerText.SetColor(NoPlayerColor)
	tgc.playerText.SetText("")
	tgc.text.SetText("Ex Aequo !")
}

func newTextGraphic(entity *MorpionGameEntity) *TextGraphicComponent {
	graphic := components.NewGraphic(nil, components.NewLayout(components.ColorWhite, color.Transparent))
	component := TextGraphicComponent{
		GraphicComponentBase: *game.NewGraphicComponentBase(&entity.Values, graphic, true, true),
		playerText:           shapes.NewText(components.NewGraphic(graphic, components.NewLayout(components.ColorWhite, color.Transparent)), canvas.Point{X: 40, Y: 120}, "", fonts.Bdf5x8),
		text:                 shapes.NewText(graphic, canvas.Point{X: 50, Y: 120}, "", fonts.Bdf5x8),
	}
	component.RegisterDrawable(component.playerText)
	component.RegisterDrawable(component.text)
	return &component
}
