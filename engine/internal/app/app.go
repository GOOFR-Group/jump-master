package app

import (
	"fmt"
	"time"

	"github.com/goofr-group/game-engine/pkg/rendering"
	"github.com/goofr-group/go-math/rotation/matrix"
	"github.com/goofr-group/go-math/vector2"
	core "github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/domain"
	"github.com/goofr-group/jump-master/engine/internal/game"
	"github.com/goofr-group/jump-master/engine/internal/game/action"
)

// App defines the main application structure.
type App struct {
	gameEngine game.Engine // Represents the game engine being used.
	lastStep   time.Time   // Represents the time when the last step occurred.
}

// New creates a new application by initializing the game engine.
func New() *App {
	camera := rendering.NewCamera(1280, 720, 1, nil, nil)
	camera.Scale = vector2.Vector2{X: 1, Y: -1}

	return &App{
		gameEngine: game.NewEngine(camera),
	}
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

	// Perform the expected actions.
	if err := a.performActions(); err != nil {
		return domain.GameState{}, fmt.Errorf("failed to perform the game actions: %w", err)
	}

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

// performActions performs the expected actions based on the current state of the actions.
func (a *App) performActions() error {
	gameEngine := a.gameEngine.Engine()
	actionManager := a.gameEngine.ActionManager()

	// TODO: remove this, currently used for testing.
	if actionManager.ActionStarted(action.Space) {
		collider := core.NewCircleCollider(5, vector2.Zero())
		gameObject := core.Object{
			Active: true,
			Transform: core.Transform2D{
				Position: vector2.Vector2{
					X: 0,
					Y: 0,
				},
				Rotation: matrix.Identity(),
				Scale:    vector2.One(),
			},
			Collider: &collider,
			Renderer: &core.Renderer{
				Width:  10,
				Height: 10,
				Offset: vector2.Vector2{
					X: -5,
					Y: -5,
				},
				Layer: "default",
			},
		}

		if err := gameEngine.CreateGameObject(&gameObject, nil); err != nil {
			return fmt.Errorf("failed to create game object: %w", err)
		}
	}

	return nil
}
