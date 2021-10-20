package objects

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/game"
	"time"
)

type (
	Bullet struct {
	}

	BulletPhysicsComponent struct {
		entity game.Entity
	}

	BulletCollisionComponent struct {
		entity game.Entity
	}

	BulletGraphicComponent struct {
		entity game.Entity
	}
)

func (b BulletPhysicsComponent) getEntity() game.Entity {
	return b.entity
}

func (b BulletPhysicsComponent) updatePhysics(elapsedBetweenUpdate time.Duration) {
	panic("implement me")
}

func (b BulletCollisionComponent) getEntity() game.Entity {
	return b.entity
}

func (b BulletCollisionComponent) updateCollision(elapsedBetweenUpdate time.Duration) {
	panic("implement me")
}

func (b BulletGraphicComponent) getEntity() game.Entity {
	return b.entity
}

func (b BulletGraphicComponent) UpdateGraphic(canvas canvas.Canvas) {
	panic("implement me")
}
