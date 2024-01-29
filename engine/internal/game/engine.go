package game

import (
	"github.com/goofr-group/game-engine/pkg/action"
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/game-engine/pkg/rendering"
	"github.com/goofr-group/game-engine/pkg/time"
	"github.com/goofr-group/physics-engine/pkg/collision/detector/naive"
	physics "github.com/goofr-group/physics-engine/pkg/engine"
	"github.com/goofr-group/physics-engine/pkg/integrator"
)

// Engine defines the physics engine, game engine and camera.
type Engine struct {
	physicsEngine *physics.Physics
	gameEngine    *engine.Engine
	camera        *rendering.Camera
	actionManager *action.Manager
}

// NewEngine initializes the physics and game engines and returns the Engine structure.
func NewEngine(camera rendering.Camera) Engine {
	physicsEngine := physics.NewEngine(&integrator.SymplecticEulerIntegrator{}, naive.MultipleDetector{})
	gameTime := time.NewTime()
	gameEngine := engine.NewEngine(physicsEngine, gameTime)
	actionManager := action.NewManager()

	return Engine{
		physicsEngine: physicsEngine,
		gameEngine:    gameEngine,
		camera:        &camera,
		actionManager: &actionManager,
	}
}

// Physics returns the physics engine being used to simulate the world.
func (e *Engine) Physics() *physics.Physics {
	return e.physicsEngine
}

// Engine returns the game engine being used to control the game objects.
func (e *Engine) Engine() *engine.Engine {
	return e.gameEngine
}

// Camera returns the game camera being used to render the game objects.
func (e *Engine) Camera() *rendering.Camera {
	return e.camera
}

// ActionManager returns the action manager being used to store the game inputs.
func (e *Engine) ActionManager() *action.Manager {
	return e.actionManager
}
