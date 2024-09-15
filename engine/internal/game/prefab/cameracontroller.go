package prefab

import (
	"fmt"

	"github.com/goofr-group/game-engine/pkg/engine"
	core "github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/config"
	"github.com/goofr-group/jump-master/engine/internal/game"
	"github.com/goofr-group/jump-master/engine/internal/game/behaviour"
)

// NewCameraController creates the camera controller object and its behaviour to control the camera position.
func NewCameraController(e game.Engine, config config.Camera) error {
	gameEngine := e.Engine()
	camera := e.Camera()

	// Create the camera controller game object.
	gameObject := core.Object{
		Active: true,
	}

	// Create the behaviour.
	cameraControllerBehaviour := behaviour.NewCameraController(camera, config.TransitionSpeed)

	// Add the camera controller game object to the game engine.
	err := gameEngine.CreateGameObject(&gameObject, []engine.Behaviour{&cameraControllerBehaviour})
	if err != nil {
		return fmt.Errorf("failed to create camera controller game object: %w", err)
	}

	return nil
}
