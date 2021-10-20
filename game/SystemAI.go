package game

import (
	"time"
)

type (
	AISystem struct {
		ais []AIComponent
	}

	AIComponent interface {
		getEntity() Entity
		updateAI(w *World, elapsedBetweenUpdate time.Duration)
	}
)

func NewAISystem() *AISystem {
	return &AISystem{}
}

func (s AISystem) Add(ai *AIComponent) {
	if ai != nil {
		s.ais = append(s.ais, *ai)
	}
}

func (s AISystem) Update(w *World, elapsedBetweenUpdate time.Duration) {
	for _, aiComponent := range s.ais {
		aiComponent.updateAI(w, elapsedBetweenUpdate)
	}
}
