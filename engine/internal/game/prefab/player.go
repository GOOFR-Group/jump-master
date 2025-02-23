package prefab

import (
	"fmt"

	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/game-engine/pkg/rendering"
	"github.com/goofr-group/go-math/rotation/matrix"
	"github.com/goofr-group/go-math/vector2"
	core "github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/config"
	"github.com/goofr-group/jump-master/engine/internal/game"
	"github.com/goofr-group/jump-master/engine/internal/game/behaviour"
	"github.com/goofr-group/jump-master/engine/internal/game/tag"
)

// NewPlayer creates the player object and behaviours for the given configuration.
func NewPlayer(e game.Engine, config config.Player) error {
	gameEngine := e.Engine()
	actionManager := e.ActionManager()

	colliderSize := config.Object.ColliderSize
	colliderOffset := config.Object.ColliderOffset
	rendererSize := config.Object.RendererSize

	// Create the game object to check if the player is in contact with the ground.
	colliderCheckGround := core.NewBoxCollider(
		vector2.Vector2{X: colliderSize.X / 2, Y: 0.5},
		vector2.Vector2{X: -colliderSize.X / 4, Y: -0.25},
	)
	colliderCheckGround.IsTrigger = true
	gameObjectCheckGround := core.Object{
		Active: true,
		Transform: core.Transform2D{
			Rotation: matrix.Identity(),
			Scale:    vector2.One(),
		},
		RigidBody: &core.RigidBody2D{
			BodyType:           core.BodyKinematic,
			CollisionDetection: core.DiscreteDetection,
			Interpolation:      core.NoneInterpolation,
		},
		Collider: &colliderCheckGround,
	}

	// Create the game object to check if the player is in contact with the ceiling.
	colliderCheckCeiling := core.NewBoxCollider(
		vector2.Vector2{X: colliderSize.X / 2, Y: 0.5},
		vector2.Vector2{X: -colliderSize.X / 4, Y: -0.25},
	)
	colliderCheckCeiling.IsTrigger = true
	gameObjectCheckCeiling := core.Object{
		Active: true,
		Transform: core.Transform2D{
			Rotation: matrix.Identity(),
			Scale:    vector2.One(),
		},
		RigidBody: &core.RigidBody2D{
			BodyType:           core.BodyKinematic,
			CollisionDetection: core.DiscreteDetection,
			Interpolation:      core.NoneInterpolation,
		},
		Collider: &colliderCheckCeiling,
	}

	// Create the player game object.
	colliderPlayer := core.NewBoxCollider(colliderSize, colliderOffset)
	gameObjectPlayer := core.Object{
		Active: true,
		Tag:    tag.Player,
		Transform: core.Transform2D{
			Position: config.Object.Position,
			Rotation: matrix.Identity(),
			Scale:    vector2.One(),
		},
		RigidBody: &core.RigidBody2D{
			BodyType:           core.BodyDynamic,
			CollisionDetection: core.DiscreteDetection,
			Interpolation:      core.Interpolate,
			Mass:               config.Object.Mass,
			GravityScale:       1,
			Drag:               config.Object.Drag,
		},
		Collider: &colliderPlayer,
		Renderer: &core.Renderer{
			Width:  rendererSize.X,
			Height: rendererSize.Y,
			Offset: rendererSize.Div(-2),
			Layer:  rendering.DefaultRenderLayer,
		},
	}

	// Create the behaviours.
	checkGroundBehaviour := behaviour.NewCheckGround(&gameObjectCheckGround)
	checkCeilingBehaviour := behaviour.NewCheckCeiling(&gameObjectCheckCeiling)
	animatorBehaviour := behaviour.NewAnimator(&gameObjectPlayer, config.Animations)
	soundControllerBehaviour := behaviour.NewSoundController(&gameObjectPlayer)
	movementBehaviour := behaviour.NewMovement(&gameObjectPlayer, actionManager, config.Movement, &checkGroundBehaviour, &animatorBehaviour, &soundControllerBehaviour)
	jumpBehaviour := behaviour.NewJump(&gameObjectPlayer, actionManager, config.Jump, &checkGroundBehaviour, &animatorBehaviour, &soundControllerBehaviour)
	fallBehaviour := behaviour.NewFall(&gameObjectPlayer, config.Fall, &checkGroundBehaviour, &animatorBehaviour, &soundControllerBehaviour)
	knockBackBehaviour := behaviour.NewKnockBack(&gameObjectPlayer, config.KnockBack, &checkGroundBehaviour, &checkCeilingBehaviour, &jumpBehaviour, &animatorBehaviour, &soundControllerBehaviour)

	// Add the player game object to the game engine.
	err := gameEngine.CreateGameObject(&gameObjectPlayer, []engine.Behaviour{&movementBehaviour, &jumpBehaviour, &fallBehaviour, &knockBackBehaviour, &animatorBehaviour, &soundControllerBehaviour})
	if err != nil {
		return fmt.Errorf("failed to create player game object: %w", err)
	}

	// Add the check ground game object to the game engine.
	err = gameEngine.CreateGameObjectWithParent(&gameObjectCheckGround, &gameObjectPlayer.Transform, []engine.Behaviour{&checkGroundBehaviour})
	if err != nil {
		return fmt.Errorf("failed to create check ground game object: %w", err)
	}

	// Add the check ceiling game object to the game engine.
	err = gameEngine.CreateGameObjectWithParent(&gameObjectCheckCeiling, &gameObjectPlayer.Transform, []engine.Behaviour{&checkCeilingBehaviour})
	if err != nil {
		return fmt.Errorf("failed to create check ceiling game object: %w", err)
	}

	return nil
}
