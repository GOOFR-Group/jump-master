package app

import (
	"fmt"
	"time"

	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/game-engine/pkg/rendering"
	"github.com/goofr-group/go-math/rotation/matrix"
	"github.com/goofr-group/go-math/vector2"
	core "github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/config"
	"github.com/goofr-group/jump-master/engine/internal/domain"
	"github.com/goofr-group/jump-master/engine/internal/game"
	"github.com/goofr-group/jump-master/engine/internal/game/behaviour"
	"github.com/goofr-group/jump-master/engine/internal/game/tag"
)

// App defines the main application structure.
type App struct {
	gameEngine game.Engine // Represents the game engine being used.
	lastStep   time.Time   // Represents the time when the last step occurred.
}

// New creates a new application by initializing the game engine.
func New() *App {
	camera := rendering.NewCamera(1280, 720, 1, nil, nil) // TODO: this should be done in a config
	camera.Scale = vector2.Vector2{X: 1, Y: -1}

	return &App{
		gameEngine: game.NewEngine(camera),
	}
}

// StartGameWorld sets up the initial game world.
func (a *App) StartGameWorld() error {
	physicsEngine := a.gameEngine.Physics()
	gameEngine := a.gameEngine.Engine()
	actionManager := a.gameEngine.ActionManager()

	// TODO: test physics engine by always simulating at least a step

	// Set up physics configurations.
	gameEngine.SetFixedDeltaTime(1. / 75) // TODO: this should be done in a config
	physicsEngine.CollisionSolvingIterations = 50
	physicsEngine.SetGravity(vector2.Vector2{X: 0, Y: -9.8 * 100})

	// Load configurations.
	playerConfig, err := config.LoadPlayer()
	if err != nil {
		return fmt.Errorf("failed to load player config: %w", err)
	}

	colliderPlatform := core.NewBoxCollider(vector2.Vector2{X: 10000, Y: 100}, vector2.Vector2{X: -5000, Y: -50})
	gameObjectPlatform := core.Object{
		Active: true,
		Tag:    tag.Platform,
		Transform: core.Transform2D{
			Position: vector2.Vector2{
				X: 0,
				Y: -350,
			},
			Rotation: matrix.Identity(),
			Scale:    vector2.One(),
		},
		RigidBody: &core.RigidBody2D{
			BodyType:           core.BodyStatic,
			CollisionDetection: core.DiscreteDetection,
			Interpolation:      core.NoneInterpolation,
			AutoMass:           false,
			GravityScale:       0,
		},
		Collider: &colliderPlatform,
		Renderer: &core.Renderer{
			Width:  10000,
			Height: 100,
			Offset: vector2.Vector2{
				X: -5000,
				Y: -50,
			},
			Layer: "default",
		},
	}
	err = gameEngine.CreateGameObject(&gameObjectPlatform, nil)
	if err != nil {
		return fmt.Errorf("failed to create game object: %w", err)
	}

	colliderCheckGround := core.NewBoxCollider(vector2.Vector2{X: 100, Y: 1}, vector2.Vector2{X: -50, Y: -0.5})
	colliderCheckGround.IsTrigger = true
	gameObjectCheckGround := core.Object{
		Active: true,
		Transform: core.Transform2D{
			Position: vector2.Vector2{
				X: 0,
				Y: 0,
			},
			Rotation: matrix.Identity(),
			Scale:    vector2.One(),
		},
		RigidBody: &core.RigidBody2D{
			BodyType:           core.BodyKinematic,
			CollisionDetection: core.DiscreteDetection,
			Interpolation:      core.NoneInterpolation,
		},
		Collider: &colliderCheckGround,
		Renderer: &core.Renderer{
			Width:  100,
			Height: 1,
			Offset: vector2.Vector2{
				X: -50,
				Y: -0.5,
			},
			Layer: "helper",
		},
	}

	colliderPlayer := core.NewBoxCollider(vector2.Vector2{X: 96, Y: 96}, vector2.Vector2{X: -96 / 2, Y: -96 / 2})
	colliderPlayer.Material = core.Material{Elasticity: 0, Friction: 0.7}
	gameObjectPlayer := core.Object{
		Active: true,
		Tag:    tag.Player,
		Transform: core.Transform2D{
			Position: vector2.Vector2{
				X: 0,
				Y: 0,
			},
			Rotation: matrix.Identity(),
			Scale:    vector2.One(),
		},
		RigidBody: &core.RigidBody2D{
			BodyType:           core.BodyDynamic,
			CollisionDetection: core.ContinuousDetection,
			Interpolation:      core.Interpolate,
			Mass:               10,
			GravityScale:       1,
			Drag:               1,
		},
		Collider: &colliderPlayer,
		Renderer: &core.Renderer{
			Width:  96,
			Height: 96,
			Offset: vector2.Vector2{
				X: -96 / 2,
				Y: -96 / 2,
			},
			Layer: "default",
		},
	}

	checkGroundBehaviour := behaviour.NewCheckGround(&gameObjectCheckGround)
	animatorBehaviour := behaviour.NewAnimator(&gameObjectPlayer, playerConfig.Animations)
	movementBehaviour := behaviour.NewMovement(&gameObjectPlayer, actionManager, playerConfig.Movement, &checkGroundBehaviour, &animatorBehaviour)
	jumpBehaviour := behaviour.NewJump(&gameObjectPlayer, actionManager, &checkGroundBehaviour, &animatorBehaviour, playerConfig.Jump)

	err = gameEngine.CreateGameObject(&gameObjectPlayer, []engine.Behaviour{&movementBehaviour, &jumpBehaviour, &animatorBehaviour})
	if err != nil {
		return fmt.Errorf("failed to create game object: %w", err)
	}

	err = gameEngine.CreateGameObjectWithParent(&gameObjectCheckGround, &gameObjectPlayer.Transform, []engine.Behaviour{&checkGroundBehaviour})
	if err != nil {
		return fmt.Errorf("failed to create game object: %w", err)
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
