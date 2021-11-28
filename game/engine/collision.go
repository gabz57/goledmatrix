package engine

import "time"

type (
	CollisionEngine struct{}

	CollisionComponent interface {
		UpdateCollision(dt time.Duration)
	}
)

// fix moves and produces fixed NextPositions
func (e CollisionEngine) DetectAndResolveCollisions(bucket *EntityBucket, dt time.Duration) {

}

func NewCollisionEngine() *CollisionEngine {
	return &CollisionEngine{}
}
