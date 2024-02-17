package app

import (
	"fmt"
	"time"

	"github.com/goofr-group/game-engine/pkg/rendering"
	"github.com/goofr-group/go-math/rotation/matrix"
	"github.com/goofr-group/go-math/vector2"
	core "github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/config"
	"github.com/goofr-group/jump-master/engine/internal/domain"
	"github.com/goofr-group/jump-master/engine/internal/game"
	"github.com/goofr-group/jump-master/engine/internal/game/prefab"
)

// App defines the main application structure.
type App struct {
	gameEngine game.Engine // Represents the game engine being used.
	lastStep   time.Time   // Represents the time when the last step occurred.

	playerConfig config.Player // Represents the player configuration.
	worldConfig  config.World  // Represents the world configuration.
}

// New creates a new application by initializing the game engine.
func New(playerConfig config.Player, worldConfig config.World) *App {
	cameraConfig := worldConfig.Camera
	camera := rendering.NewCamera(cameraConfig.Width, cameraConfig.Height, cameraConfig.PPU, nil, nil)
	camera.Scale = vector2.Vector2{X: 1, Y: -1}

	return &App{
		gameEngine:   game.NewEngine(camera),
		playerConfig: playerConfig,
		worldConfig:  worldConfig,
	}
}

// StartGameWorld sets up the initial game world.
func (a *App) StartGameWorld() error {
	physicsEngine := a.gameEngine.Physics()
	gameEngine := a.gameEngine.Engine()

	physicsConfig := a.worldConfig.Physics

	// Set up physics configurations.
	a.lastStep = time.Now()
	gameEngine.SetFixedDeltaTime(physicsConfig.UpdateRate)
	physicsEngine.SetGravity(physicsConfig.Gravity)
	physicsEngine.CollisionSolvingIterations = 50

	// Create the camera controller object.
	err := prefab.NewCameraController(a.gameEngine)
	if err != nil {
		return fmt.Errorf("failed to create camera controller prefab: %w", err)
	}

	// Create the player object.
	err = prefab.NewPlayer(a.gameEngine, a.playerConfig)
	if err != nil {
		return fmt.Errorf("failed to create player prefab: %w", err)
	}

	// Create the grid objects (platforms and props).
	err = prefab.NewGridObjects(a.gameEngine, a.worldConfig.Grid)
	if err != nil {
		return fmt.Errorf("failed to create grid objects prefab: %w", err)
	}

	return nil
}

// GameStep performs an engine step and returns the current state of every game object in the world.
func (a *App) GameStep(actions map[string]bool) (domain.GameState, error) {
	// Update the state of the actions.
	for action, state := range actions {
		a.gameEngine.ActionManager().SetAction(action, state)
	}

	// Get the current time step.
	timeStep := time.Since(a.lastStep).Seconds()
	a.lastStep = time.Now()

	// Perform the actual game step.
	if err := a.gameEngine.Engine().Step(timeStep); err != nil {
		return domain.GameState{}, fmt.Errorf("failed to perform the game step: %w", err)
	}

	// Get the game camera.
	camera := a.gameEngine.Camera()

	// Get game objects to render.
	var gameObjects []core.Object
	for _, object := range a.gameEngine.Engine().GetState() {
		// Ignore objects that are not supposed to be visible.
		if !camera.IsVisible(object) {
			continue
		}

		// Convert the object's transform to screen space.
		camera.WorldToScreenTransform(&object.Transform)
		// Prevent the objects from rotating.
		object.Transform.Rotation = matrix.Identity()

		// Convert the object's collider to screen space.
		if object.Collider != nil {
			switch object.Collider.Type() {
			case core.CircleType:
				c := object.Collider.Collider().(*core.CircleCollider2D)

				c.Radius *= camera.PixelsPerUnit
				c.Pivot = c.Pivot.Mul(camera.PixelsPerUnit)

			case core.EdgeType:
				c := object.Collider.Collider().(*core.EdgeCollider2D)

				points := make([]vector2.Vector2, len(c.Points))
				for i := 0; i < len(c.Points); i++ {
					points[i] = c.Points[i].Mul(camera.PixelsPerUnit)
				}
				c.Points = points
				c.Offset = c.Offset.Mul(camera.PixelsPerUnit)

			case core.PolygonType:
				c := object.Collider.Collider().(*core.PolygonCollider2D)

				points := make([]vector2.Vector2, len(c.Points))
				for i := 0; i < len(c.Points); i++ {
					points[i] = c.Points[i].Mul(camera.PixelsPerUnit)
				}
				c.Points = points
				c.Offset = c.Offset.Mul(camera.PixelsPerUnit)
			}
		}

		// Convert the object's renderer to screen space.
		if object.Renderer != nil {
			object.Renderer.Height = object.Renderer.Height * camera.PixelsPerUnit
			object.Renderer.Width = object.Renderer.Width * camera.PixelsPerUnit
			object.Renderer.Offset = object.Renderer.Offset.Mul(camera.PixelsPerUnit)
		}

		gameObjects = append(gameObjects, object)
	}

	return domain.GameState{
		GameObjects: gameObjects,
		Camera:      *camera,
	}, nil
}
