package gameobjects

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/game"
)

const OthellierEntity game.Entity = 0

type Othellier struct {
	entity        game.Entity
	transforms    game.Transform
	graphics      game.GraphicComponent
	pieces        []Piece
	currentPlayer Player
	cursor        Point
}

func NewOthellier() *Othellier {
	var pieces []Piece
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			player := None
			if x == 3 && y == 3 {
				player = White
			}
			if x == 3 && y == 4 {
				player = Black
			}
			if x == 4 && y == 3 {
				player = Black
			}
			if x == 4 && y == 4 {
				player = White
			}
			pieces = append(pieces, *NewPiece(OthellierEntity, x, y, player))
		}
	}
	initialBlackPlayerDefaultPosition := Point{
		x: 2,
		y: 3,
	}
	return &Othellier{
		entity: OthellierEntity,
		transforms: *game.NewTransform(canvas.FloatingPoint{
			X: 0,
			Y: 0,
		}),
		pieces:        pieces,
		currentPlayer: Black,
		cursor:        initialBlackPlayerDefaultPosition,
		graphics:      othellierGraphics(OthellierEntity),
	}
}

func (o *Othellier) SetCursor(position Point) {
	o.cursor = position
}

func (o *Othellier) ApplyCurrentPosition() {
	// take new position
	var piece = o.GetPiece(o.cursor.x, o.cursor.y)
	piece.setPlayer(o.currentPlayer)

	o.processCaptureFromCursorPosition()

	// Game over ? (<=> grid complete or 1 untakable position for next player)
	var winner = o.findWinner()
	var gameOver = false
	if winner != None {
		gameOver = true
	}
	if !gameOver {
		o.currentPlayer = computeNextPlayer(o.currentPlayer)
	}
	// TODO: if Game Over, DISPLAY END with winner !!
	if gameOver {

	}
}

func (o *Othellier) processCaptureFromCursorPosition() {
	var nextPlayer = computeNextPlayer(o.currentPlayer)
	for _, adjacentsPosition := range adjacentPositions(o.cursor) {
		var dx, dy int
		dx = adjacentsPosition.x - o.cursor.x
		dy = adjacentsPosition.y - o.cursor.y

		firstPiece := o.GetPiece(adjacentsPosition.x, adjacentsPosition.y)
		if firstPiece.player == nextPlayer {
			// collect pieces & follow this direction (might be diag) until empty or current player piece
			var capturableOpponentPieces []*Piece
			capturableOpponentPieces = append(capturableOpponentPieces, firstPiece)
			nextPosition := Point{
				x: o.cursor.x,
				y: o.cursor.y,
			}
			for {
				nextPosition = Point{
					x: nextPosition.x + dx,
					y: nextPosition.y + dy,
				}
				nextPiece := o.GetPiece(adjacentsPosition.x, adjacentsPosition.y)
				if nextPiece.player == o.currentPlayer {
					// process capture
					for _, opponentPiece := range capturableOpponentPieces {
						opponentPiece.setPlayer(o.currentPlayer)
					}
					break
				} else if nextPiece.player == nextPlayer {
					// capturable
					capturableOpponentPieces = append(capturableOpponentPieces, nextPiece)
				} else if nextPiece.player == None {
					// can't use this direction
					break
				}
			}
		}
	}
}

func (o *Othellier) findWinner() Player {
	var winner = None
	var nbFreeSlots = 0
	var nbWhitePieces = 0
	var nbBlackPieces = 0
	for i := 0; i < 64; i++ {
		piecePlayer := o.pieces[i].player
		if piecePlayer == None {
			nbFreeSlots++
		}
		if piecePlayer == Black {
			nbBlackPieces++
		}
		if piecePlayer == White {
			nbWhitePieces++
		}
	}
	if nbFreeSlots == 0 || nbFreeSlots == 1 && o.untakeablePosition() {
		if nbBlackPieces > nbWhitePieces {
			winner = Black
		} else {
			winner = White
		}
	}
	return winner
}

func (o *Othellier) untakeablePosition() bool {
	var nextPlayer = computeNextPlayer(o.currentPlayer)
	var allAdjacentBelongToNextPlayer = true
	for _, adjacentsPosition := range adjacentPositions(o.lastFreePosition()) {
		allAdjacentBelongToNextPlayer = allAdjacentBelongToNextPlayer && o.GetPiece(adjacentsPosition.x, adjacentsPosition.y).player == nextPlayer
	}
	return allAdjacentBelongToNextPlayer
}

func adjacentPositions(position Point) []Point {
	var arr []Point
	for i := position.x - 1; i < position.x+1; i++ {
		for j := position.y - 1; j < position.x+1; i++ {
			if i < 0 || j < 0 || i >= 8 || j >= 8 {
				continue
			}
			if i == position.x && j == position.y {
				continue
			}
			arr = append(arr, Point{x: i, y: j})
		}
	}
	return arr
}

func (o *Othellier) lastFreePosition() Point {
	var lastPiece Piece
	for _, piece := range o.pieces {
		if piece.player == None {
			lastPiece = piece
			break
		}
	}
	return lastPiece.point
}

func (o *Othellier) GetPiece(x int, y int) *Piece {
	return &o.pieces[x+8*y]
}

func computeNextPlayer(player Player) Player {
	if player == Black {
		return White
	}
	return Black
}

func (o *Othellier) SetTurn(player Player) {
	o.currentPlayer = player
}

func (o *Othellier) Register(world *game.World) {
	world.AddEntityObject(o)
	for _, piece := range o.pieces {
		world.RegisterEntityObject(&piece)
	}
}

func (o *Othellier) Entity() game.Entity {
	return o.entity
}

func (o *Othellier) Parent() game.Entity {
	return game.NoParent
}

func (o *Othellier) Position(world *game.World) canvas.FloatingPoint {
	return canvas.FloatingPoint{}
}

func (o *Othellier) LocalPosition() canvas.FloatingPoint {
	return canvas.FloatingPoint{}
}

func (o *Othellier) Transform() *game.Transform {
	panic("implement me")
}

func (o *Othellier) Input() *game.InputComponent {
	panic("implement me")
}

func (o *Othellier) Ai() *game.AIComponent {
	panic("implement me")
}

func (o *Othellier) Physics() *game.PhysicsComponent {
	panic("implement me")
}

func (o *Othellier) Collision() *game.CollisionComponent {
	panic("implement me")
}

func (o *Othellier) Graphics() *game.GraphicComponent {
	return &o.graphics
}

func othellierGraphics(entity game.Entity) game.GraphicComponent {

	return &OthellierGraphicsComponent{
		entity:             entity,
		drawableBackground: *drawableBackground(),
		drawableGrid:       drawableGrid(),
		//drawablePiece: DrawablePiece{
		//	piece: *circle,
		//},
	}
}

func drawableBackground() *shapes.Rectangle {
	var backgroundColor = components.ColorGreen
	backgroundGraphic := components.NewGraphic(nil, components.NewLayout(backgroundColor, backgroundColor))
	return shapes.NewRectangle(
		backgroundGraphic,
		canvas.Point{
			X: 0,
			Y: 0,
		},
		canvas.Point{
			X: 2*margin + 8*(pieceDiameter+spacing),
			Y: 2*margin + 8*(pieceDiameter+spacing),
		},
		true,
	)
}

func drawableGrid() []shapes.Line {
	var separatorColor = components.ColorGreen
	components.NewGraphic(nil, components.NewLayout(separatorColor, separatorColor))
	separatorGraphic := components.NewGraphic(nil, components.NewLayout(separatorColor, separatorColor))
	var lines []shapes.Line
	for i := 1; i < 8; i++ {
		// vertical
		lines = append(lines, *shapes.NewLine(
			separatorGraphic,
			canvas.Point{
				X: 0,
				Y: 0,
			},
			canvas.Point{
				X: margin + i*(pieceDiameter+spacing), //FIXME
				Y: 0,                                  //FIXME
			},
		),
			*shapes.NewLine(
				separatorGraphic,
				canvas.Point{
					X: 0,
					Y: 0,
				},
				canvas.Point{
					X: 0, //FIXME
					Y: 0, //FIXME
				},
			))
	}
	return lines
}

type OthellierGraphicsComponent struct {
	entity             game.Entity
	drawableBackground shapes.Rectangle
	drawableGrid       []shapes.Line
}

func (o OthellierGraphicsComponent) GetEntity() game.Entity {
	panic("implement me")
}

func (o OthellierGraphicsComponent) UpdateGraphic(world *game.World) {
	panic("implement me")
}

func (o OthellierGraphicsComponent) Render(canvas canvas.Canvas) {
	err := o.drawableBackground.Draw(canvas)
	if err != nil {
		return
	}
	for _, line := range o.drawableGrid {
		err = line.Draw(canvas)
		if err != nil {
			return
		}
	}
}
