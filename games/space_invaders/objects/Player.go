package objects

import (
	"github.com/gabz57/goledmatrix/canvas"
	"github.com/gabz57/goledmatrix/components"
	"github.com/gabz57/goledmatrix/game"
	"github.com/gabz57/goledmatrix/games/space_invaders/ui"
	"time"
)

type (
	Player struct {
		parent     game.Entity
		entity     game.Entity
		name       string
		transforms game.Transform
		input      game.InputComponent
		physics    game.PhysicsComponent
		collision  game.CollisionComponent
		graphics   game.GraphicComponent
		//ai        *game.AIComponent
	}

	//PlayerAIComponent struct {
	//	entity game.Entity
	//}

	PlayerInputComponent struct {
		entity game.Entity
	}

	PlayerPhysicsComponent struct {
		entity game.Entity
	}

	PlayerCollisionComponent struct {
		entity game.Entity
	}

	PlayerGraphicComponent struct {
		entity   game.Entity
		drawable components.Drawable //TODO: use real type ?
	}
)

func NewPlayer(entity game.Entity, name string) *Player {
	return &Player{
		entity: entity,
		name:   name,
		//ai:     playerAI(entity),
		input:     playerInput(entity),
		physics:   playerPhysics(entity),
		collision: playerCollision(entity),
		graphics:  playerGraphics(entity),
	}
}

func (p *Player) Register(world *game.World) {
	world.AddEntityObject(p)
}

func (p Player) Position(world *game.World) canvas.FloatingPoint {
	return p.transforms.Position(world, p.Parent())
}

func (p Player) LocalPosition() canvas.FloatingPoint {
	return p.transforms.LocalPosition()
}

func (p Player) Entity() game.Entity {
	return p.entity
}

func (p Player) Parent() game.Entity {
	return game.NoParent
}

func (p Player) Transform() *game.Transform {
	return &p.transforms
}

func (p Player) Input() *game.InputComponent {
	return &p.input
}

func (p Player) Ai() *game.AIComponent {
	return nil
	//return &p.ai
}

func (p Player) Physics() *game.PhysicsComponent {
	return &p.physics
}

func (p Player) Collision() *game.CollisionComponent {
	return &p.collision
}

func (p Player) Graphics() *game.GraphicComponent {
	return &p.graphics
}

func playerInput(entity game.Entity) game.InputComponent {
	return PlayerInputComponent{
		entity: entity,
	}
}

func (b PlayerInputComponent) GetEntity() game.Entity {
	return b.entity
}

func (b PlayerInputComponent) UpdateInputs(elapsedBetweenUpdate time.Duration) {
	// TODO ("implement me")
}

func playerPhysics(entity game.Entity) game.PhysicsComponent {
	return PlayerPhysicsComponent{
		entity: entity,
	}
}

func (b PlayerPhysicsComponent) GetEntity() game.Entity {
	return b.entity
}

func (b PlayerPhysicsComponent) UpdatePhysics(world *game.World, elapsedBetweenUpdate time.Duration) {
	// TODO ("implement me")
}

func playerCollision(entity game.Entity) game.CollisionComponent {
	return PlayerCollisionComponent{
		entity: entity,
	}
}

func (b PlayerCollisionComponent) GetEntity() game.Entity {
	return b.entity
}

func (b PlayerCollisionComponent) UpdateCollision(w *game.World) {
	// TODO ("implement me")
}

func (b PlayerCollisionComponent) onCollision(w *game.World, collidingWith game.CollisionComponent) {
	// do something, ie:
	// - bounce on collidingWith
	// - player explodes
	// - collidingWith explodes
}

func playerGraphics(entity game.Entity) game.GraphicComponent {
	return PlayerGraphicComponent{
		entity:   entity,
		drawable: ui.Spaceship{},
	}
}

func (b PlayerGraphicComponent) GetEntity() game.Entity {
	return b.entity
}

func (b PlayerGraphicComponent) UpdateGraphic(w *game.World) {
	// TODO ("implement me")
}

func (b PlayerGraphicComponent) Render(canvas canvas.Canvas) {
	err := b.drawable.Draw(canvas)
	if err != nil {
		return
	}
}
