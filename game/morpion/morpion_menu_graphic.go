package morpion

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/canvas/fonts"
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/game/engine"
	"image/color"
)

const optionYOffset = 12

var optionOffset = canvas.Point{Y: optionYOffset}
var noOffset = canvas.Point{}

type MainMenuGraphicComponent struct {
	engine.GraphicComponentBase
	gameTitle     *shapes.Text
	resumeText    *shapes.Text
	startText     *shapes.Text
	showText      *shapes.Text
	exitText      *shapes.Text
	optionPointer *shapes.Cross
}

func newMainMenuGraphicComponent(entity *MainMenuEntity) *MainMenuGraphicComponent {
	graphic := components.NewGraphic(nil, components.NewLayout(components.ColorWhite, color.Transparent))
	graphicComponent := MainMenuGraphicComponent{
		GraphicComponentBase: *engine.NewGraphicComponentBase(&entity.Values, graphic, true, true),
	}
	graphicComponent.gameTitle = newTitle(graphic, canvas.Point{X: 34, Y: 4}, "Morpion !")
	graphicComponent.resumeText = newTextOption(graphic, canvas.Point{X: 11, Y: 1*optionYOffset + 14}, "Resume")
	graphicComponent.startText = newTextOption(graphic, canvas.Point{X: 11, Y: 2*optionYOffset + 14}, "Start a new game")
	graphicComponent.showText = newTextOption(graphic, canvas.Point{X: 11, Y: 3*optionYOffset + 14}, "Show scores")
	graphicComponent.exitText = newTextOption(graphic, canvas.Point{X: 11, Y: 4*optionYOffset + 14}, "Exit")
	graphicComponent.optionPointer = shapes.NewCross(components.NewGraphic(graphic, components.NewLayout(components.ColorRed, components.ColorWhite)), canvas.Point{X: 5, Y: 30}, 3)
	graphicComponent.RegisterDrawable(graphicComponent.gameTitle)
	graphicComponent.RegisterDrawable(graphicComponent.resumeText)
	graphicComponent.RegisterDrawable(graphicComponent.startText)
	graphicComponent.RegisterDrawable(graphicComponent.showText)
	graphicComponent.RegisterDrawable(graphicComponent.exitText)
	graphicComponent.RegisterDrawable(graphicComponent.optionPointer)

	graphicComponent.DisableDrawable(graphicComponent.resumeText)

	return &graphicComponent
}

func newTitle(graphic *components.Graphic, position canvas.Point, txt string) *shapes.Text {
	return shapes.NewText(
		components.NewOffsetGraphic(graphic, nil, canvas.Point{}),
		position,
		txt,
		fonts.Bdf7x13,
	)
}
func newTextOption(graphic *components.Graphic, position canvas.Point, txt string) *shapes.Text {
	return shapes.NewText(
		components.NewOffsetGraphic(graphic, nil, canvas.Point{}),
		position,
		txt,
		fonts.Bdf5x8,
	)
}

func (mmgc *MainMenuGraphicComponent) refreshOptionsPosition() {
	menuOptionResumeVisible := mmgc.GetValue(isMenuOptionResumeVisible).(bool)
	if menuOptionResumeVisible {
		mmgc.GraphicComponentBase.EnableDrawable(mmgc.resumeText)
		mmgc.resumeText.SetOffset(optionOffset)
		mmgc.startText.SetOffset(optionOffset)
		mmgc.showText.SetOffset(optionOffset)
		mmgc.exitText.SetOffset(optionOffset)
	} else {
		mmgc.GraphicComponentBase.DisableDrawable(mmgc.resumeText)
		mmgc.resumeText.SetOffset(noOffset)
		mmgc.startText.SetOffset(noOffset)
		mmgc.showText.SetOffset(noOffset)
		mmgc.exitText.SetOffset(noOffset)
	}

	option := mmgc.GetValue(selectedMenuOption).(MenuOption)
	optionIndex := option
	if menuOptionResumeVisible {
		optionIndex = option
	} else {
		optionIndex = option - 1
	}
	mmgc.optionPointer.SetOffset(canvas.Point{X: 1, Y: 12 + int(optionIndex)*optionYOffset})
}
