package gameobjects

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/game"
	"strconv"
)

type (
	Point struct {
		x, y int
	}
	Player int
	Piece  struct {
		parent     game.Entity
		entity     game.Entity
		transforms game.Transform
		graphics   game.GraphicComponent
		point      Point
		player     Player
	}
)

const (
	None Player = iota
	White
	Black
)

const margin = 2
const pieceRadius = 5
const pieceDiameter = 2 * pieceRadius
const spacing = 3

var xToCol = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

func (p *Point) getPositionLabel() string {
	return xToCol[p.x] + strconv.Itoa(p.y+1)
}

func NewPiece(parent game.Entity, x, y int, player Player) *Piece {
	entity := PieceEntity(parent, x, y)
	return &Piece{
		parent: parent,
		entity: entity,
		transforms: *game.NewTransform(canvas.FloatingPoint{
			// closest corner from origin for each piece
			X: float64(margin + x*(pieceDiameter+spacing)),
			Y: float64(margin + y*(pieceDiameter+spacing)),
		}),
		point:    Point{x, y},
		player:   player,
		graphics: pieceGraphics(entity, player),
	}
}

func PieceEntity(parent game.Entity, x int, y int) game.Entity {
	return game.Entity(int(parent) + 1 + x + 8*y)
}

func (p *Piece) Register(world *game.World) {
	world.AddEntityObject(p)
}

func (p *Piece) getPosition() Point {
	return p.point
}

func (p *Piece) setPlayer(player Player) {
	p.player = player
}

func (p *Piece) Entity() game.Entity {
	return p.entity
}

func (p *Piece) Parent() game.Entity {
	return p.parent
}

func (p *Piece) Position(world *game.World) canvas.FloatingPoint {
	return p.transforms.Position(world, p.Parent())
}

func (p *Piece) LocalPosition() canvas.FloatingPoint {
	return p.transforms.LocalPosition()
}

func (p *Piece) Transform() *game.Transform {
	return &p.transforms
}

func (p *Piece) Input() *game.InputComponent {
	return nil
}

func (p *Piece) Ai() *game.AIComponent {
	return nil
}

func (p *Piece) Physics() *game.PhysicsComponent {
	return nil
}

func (p *Piece) Collision() *game.CollisionComponent {
	return nil
}

func (p *Piece) Graphics() *game.GraphicComponent {
	return &p.graphics
}
