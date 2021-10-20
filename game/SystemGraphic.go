package game

import (
	"github.com/gabz57/goledmatrix/canvas"
)

type (
	GraphicSystem struct {
		graphics []GraphicComponent
	}

	GraphicComponent interface {
		GetEntity() Entity
		UpdateGraphic(world *World)
		Render(c canvas.Canvas)
	}
)

func NewGraphicSystem() *GraphicSystem {
	return &GraphicSystem{}
}

func (s GraphicSystem) Add(graphics *GraphicComponent) {
	if graphics != nil {
		s.graphics = append(s.graphics, *graphics)
	}
}

func (s GraphicSystem) Update(w *World) {
	for _, graphicComponent := range s.graphics {
		graphicComponent.UpdateGraphic(w)
	}
}

func (s GraphicSystem) Render(canvas canvas.Canvas) {
	// maybe render by layer or something similar
	for _, graphicComponent := range s.graphics {
		// filter active components
		graphicComponent.Render(canvas)
	}
}
