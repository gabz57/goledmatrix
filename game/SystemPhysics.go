package game

import (
	"time"
)

type (
	PhysicsSystem struct {
		physics []PhysicsComponent
	}

	PhysicsComponent interface {
		GetEntity() Entity
		UpdatePhysics(world *World, elapsedBetweenUpdate time.Duration)
	}
)

func NewPhysicsSystem() *PhysicsSystem {
	return &PhysicsSystem{}
}

func (s PhysicsSystem) Add(physics *PhysicsComponent) {
	if physics != nil {
		s.physics = append(s.physics, *physics)
	}
}

// Update acceleration, velocity, position of object
// following updateDuration parameter
func (s PhysicsSystem) Update(world *World, updateDuration time.Duration) {
	for _, physicComponent := range s.physics {
		physicComponent.UpdatePhysics(world, updateDuration)
	}
}
