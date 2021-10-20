package game

type (
	CollisionSystem struct {
		collisions []CollisionComponent
	}

	CollisionComponent interface {
		GetEntity() Entity
		UpdateCollision(world *World)
	}
)

func NewCollisionSystem() *CollisionSystem {
	return &CollisionSystem{}
}

func (s CollisionSystem) Add(collision *CollisionComponent) {
	if collision != nil {
		s.collisions = append(s.collisions, *collision)
	}
}

func (s CollisionSystem) Update(w *World) {
	// NOTE: try not to loop over all collisions components
	for _, collision := range s.collisions {
		collision.UpdateCollision(w)
	}
}
