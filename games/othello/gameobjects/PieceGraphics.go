package gameobjects

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/components/shapes"
	"github.com/gabz57/goledmatrix/game"
	"time"
)

func pieceGraphics(entity game.Entity, player Player) game.GraphicComponent {
	var pieceColor = components.ColorRed
	if player == White {
		pieceColor = components.ColorWhite
	}
	circle := shapes.NewCircle(
		components.NewGraphic(nil, components.NewLayout(pieceColor, pieceColor)),
		canvas.Point{
			X: pieceRadius,
			Y: pieceRadius,
		},
		pieceRadius,
		true,
	)
	return &PieceGraphicComponent{
		entity: entity,
		drawablePiece: DrawablePiece{
			piece: *circle,
			// TODO: complete
		},
	}
}

type PieceGraphicComponent struct {
	entity        game.Entity
	drawablePiece DrawablePiece
}

func (pgc *PieceGraphicComponent) GetEntity() game.Entity {
	return pgc.entity
}

func (pgc *PieceGraphicComponent) UpdateGraphic(w *game.World) {
	object := w.GetEntityObject(pgc.entity)
	position := (*object).Position(w)
	// TODO ("implement me")
	pgc.drawablePiece.piece.Graphic.
}

func (pgc *PieceGraphicComponent) Render(canvas canvas.Canvas) {
	err := pgc.drawablePiece.Draw(canvas)
	if err != nil {
		return
	}
}

type DrawablePiece struct {
	piece shapes.Circle
}

func (p *DrawablePiece) Update(elapsedBetweenUpdate time.Duration) {
	panic("implement me")
}

func (p *DrawablePiece) Draw(canvas canvas.Canvas) error {
	// FIXME ?: too short implementation, missing parent relative position ?
	p.piece.Draw(canvas)
	return nil
}
