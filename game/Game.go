package game

import (
	"github.com/gabz57/goledmatrix/canvas"
	"time"
)

const NoParent = -1

type (
	Game interface {
		GetWorld() *World
		//GetActiveScene() *components.Scene
		//SetActiveScene(scene *components.Scene)
	}

	World struct {
		objects         map[Entity]*EntityObject
		inputSystem     InputSystem
		aiSystem        AISystem
		physicsSystem   PhysicsSystem
		collisionSystem CollisionSystem
		graphicSystem   GraphicSystem
	}

	Entity int

	Transform struct {
		position canvas.FloatingPoint
	}

	EntityObject interface {
		Register(world *World)
		Entity() Entity
		Parent() Entity
		Position(world *World) canvas.FloatingPoint
		LocalPosition() canvas.FloatingPoint
		Transform() *Transform
		Input() *InputComponent
		Ai() *AIComponent
		Physics() *PhysicsComponent
		Collision() *CollisionComponent
		Graphics() *GraphicComponent
	}
)

func NewWorld() *World {
	return &World{
		inputSystem:     *NewInputSystem(),
		aiSystem:        *NewAISystem(),
		physicsSystem:   *NewPhysicsSystem(),
		collisionSystem: *NewCollisionSystem(),
		graphicSystem:   *NewGraphicSystem(),
	}
}

func NewTransform(position canvas.FloatingPoint) *Transform {
	return &Transform{position: position}
}

func (t *Transform) LocalPosition() canvas.FloatingPoint {
	return t.position
}

func (t *Transform) Position(w *World, parent Entity) canvas.FloatingPoint {
	if parent != -1 {
		return (*w.GetEntityObject(parent)).Position(w).Add(t.position)
	}
	return t.position
}

func (w *World) RegisterEntityObject(object EntityObject) {
	object.Register(w)
}

func (w *World) AddEntityObject(object EntityObject) {
	w.objects[object.Entity()] = &object
	w.inputSystem.Add(object.Input())
	w.aiSystem.Add(object.Ai())
	physics := object.Physics()
	if physics != nil {
		w.physicsSystem.Add(physics)
	}
	w.collisionSystem.Add(object.Collision())
	w.graphicSystem.Add(object.Graphics())
}

func (w *World) Update(elapsedBetweenUpdate time.Duration) {
	w.inputSystem.Update(w)
	w.aiSystem.Update(w, elapsedBetweenUpdate)
	w.physicsSystem.Update(w, elapsedBetweenUpdate)
	w.collisionSystem.Update(w)
	w.graphicSystem.Update(w)
}

func (w *World) Render(canvas canvas.Canvas) {
	w.graphicSystem.Render(canvas)
}

func (w *World) GetEntityObject(entity Entity) *EntityObject {
	return w.objects[entity]
}

func Load() (*World, error) {
	world := World{}
	//NB_FILES := 3
	//for i := 0; i < NB_FILES; i++ {
	//	var entity Entity
	//	var object EntityObject
	//
	//	world.RegisterEntityObject(entity, object)
	//}
	return &world, nil
}

func Save(w *World) error {
	for entity, object := range w.objects {
		save(entity, object)
	}

	return nil
}

func save(entity Entity, object *EntityObject) {
	// write world state to file
}
