package game

import "time"

type (
	PhysicsEngine struct{}

	PhysicsComponent interface {
		UpdatePhysics(dt time.Duration)
	}
)

// Simulate apply physics (compute next positions & velocities)
func (pe PhysicsEngine) Simulate(bucket *EntityBucket, dt time.Duration) {
	//for _, physicComponent := range bucket.PhysicsComponents {
	//}
}

func NewPhysicsEngine() *PhysicsEngine {
	return &PhysicsEngine{}
}
