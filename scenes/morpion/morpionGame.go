package morpion

import (
	. "github.com/gabz57/goledmatrix/canvas"
	. "github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"time"
)

type PlayerRef int

const (
	NoPlayer PlayerRef = iota
	PlayerX
	PlayerO
)

type Grid struct {
	grid []PlayerRef
}

type Morpion struct {
	shape *CompositeDrawable
	//cross      *shapes.Cross
	controller Controller
	gridCpt    Component
	selection  shapes.Rectangle

	morpionCell []Component

	morpionGrid      Grid
	currentPlayer    PlayerRef
	currentSelection int
	winner           PlayerRef
}

func NewMorpionComponent(canvas Canvas) *Morpion {
	c := Morpion{
		shape:            NewCompositeDrawable(NewGraphic(nil, NewLayout(ColorRed, ColorWhite))),
		morpionGrid:      *NewMorpionGrid(),
		currentPlayer:    NoPlayer,
		currentSelection: -1,
		winner:           NoPlayer,
	}
	//c.cross = shapes.NewCross(c.shape.Graphic, Point{
	//	X: canvas.Bounds().Max.X / 2,
	//	Y: canvas.Bounds().Max.Y / 2,
	//}, 2)
	//c.shape.AddDrawable(c.cross)
	//c.controller = *NewGamepadDemoController(canvas.Bounds(), c.cross)
	return &c
}

func NewMorpionGrid() *Grid {
	return &Grid{
		grid: []PlayerRef{
			NoPlayer, NoPlayer, NoPlayer,
			NoPlayer, NoPlayer, NoPlayer,
			NoPlayer, NoPlayer, NoPlayer,
		},
	}
}

func (f *Morpion) Update(elapsedBetweenUpdate time.Duration) bool {
	return f.controller.Update(elapsedBetweenUpdate)
}

func (f *Morpion) Draw(canvas Canvas) error {
	return f.shape.Draw(canvas)
}

func (f *Morpion) Controller() SceneController {
	return &f.controller
}
