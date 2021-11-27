package game

import (
	"fmt"
	"github.com/gabz57/goledmatrix/canvas"
	"time"
)

const frameDurationInNanos = 33333333  // 30 FPS approximated in nanos
const updateDurationInNanos = 10000000 // 100 updates per second (to maintain physics & time), independent from FPS
const updateDuration = time.Duration(updateDurationInNanos)

type Engine struct {
	entities         map[EntityRef]Entity
	entityBuckets    []EntityBucket
	controllerEngine ControllerEngine
	animationEngine  AnimationEngine
	physicsEngine    PhysicsEngine
	collisionEngine  CollisionEngine
	graphicEngine    GraphicEngine
	gameDone         chan struct{}
}

type EntityBucket struct {
	Entities             map[EntityRef]Entity
	ControllerComponents []ControllerComponent
	AnimationComponents  []AnimationComponent
	PhysicsComponents    []PhysicsComponent
	CollisionComponents  []CollisionComponent
	GraphicComponents    []GraphicComponent
}

func NewGameEngine(canvas canvas.Canvas) *Engine {
	return &Engine{
		entityBuckets:    []EntityBucket{{}},
		controllerEngine: *NewControllerEngine(),
		animationEngine:  *NewAnimationEngine(),
		physicsEngine:    *NewPhysicsEngine(),
		collisionEngine:  *NewCollisionEngine(),
		graphicEngine:    *NewGraphicEngine(canvas),
		gameDone:         make(chan struct{}),
	}
}

func (e *Engine) LoadGame(game Game) {
	println("loading game into engine ...", game.Name())
	e.entities = map[EntityRef]Entity{}
	e.entityBuckets = game.InitializeBuckets()
	for _, bucket := range e.entityBuckets {
		for ref, entity := range bucket.Entities {
			e.entities[ref] = entity
		}
	}
	game.InitializeGame(e)
}

func (e *Engine) Run(matrixDone chan struct{}) {
	e.controllerEngine.Start()
	defer e.controllerEngine.Stop()

	e.runGameLoop(matrixDone)
}

func (e *Engine) Exit() {
	// Todo: shutdown gamepad, gameEngine, canvas, matrix
	go func() {
		e.gameDone <- struct{}{}
	}()
}

func (e *Engine) runGameLoop(done chan struct{}) {

	previous := time.Now()
	lag := time.Duration(0)
	fmt.Println("Starting Canvas Engine !")
	dirty := true
LOOP:
	for {
		select {
		case <-e.gameDone:
			fmt.Println("Game asked game loop BREAK")
			break LOOP
		case <-done:
			fmt.Println("Context asked for game loop BREAK")
			break LOOP
		default:
		}
		current := time.Now()
		elapsed := current.Sub(previous)
		previous = current
		lag += elapsed

		// using lag to catch up missing updates when UI renders to slow
		for lag >= updateDuration {
			e.controllerEngine.ConsumeGamepadEvents(updateDuration)
			e.controllerEngine.ProcessActions(e, updateDuration)
			for _, bucket := range e.entityBuckets {
				dirty = e.UpdateBucket(&bucket, updateDuration) || dirty
			}
			lag -= updateDuration
			if lag >= updateDuration {
				continue
			}
			select {
			case <-time.After(updateDurationInNanos - time.Now().Sub(current)):
			}
		}

		if dirty == true {
			dirty = false
			e.graphicEngine.RenderSceneAndSwapBuffers(&e.entityBuckets)
		}

		select {
		case <-time.After(frameDurationInNanos - time.Now().Sub(current)):
		}
	}
	fmt.Println("engine loop END")
}

func (e *Engine) UpdateBucket(bucket *EntityBucket, dt time.Duration) bool {
	dirty := false
	for _, entity := range bucket.Entities {
		dirty = entity.Update(e, dt) || dirty
	}
	e.animationEngine.CalculateIntermediatePoses(bucket, dt)
	e.physicsEngine.Simulate(bucket, dt)
	e.collisionEngine.DetectAndResolveCollisions(bucket, dt)
	e.animationEngine.FinalizePoseAndMatrixPalette(bucket)
	for _, entity := range bucket.Entities {
		dirty = entity.FinalUpdate(dt) || dirty
	}
	return dirty
}

func (e *Engine) GetEntity(ref EntityRef) Entity {
	return e.entities[ref]
}

func (e *Engine) SetActiveController(controllerComponent ControllerComponent) {
	e.controllerEngine.SetActiveController(controllerComponent)
}
