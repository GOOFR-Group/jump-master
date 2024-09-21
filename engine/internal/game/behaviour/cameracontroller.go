package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/game-engine/pkg/rendering"
	"github.com/goofr-group/go-math/mathf"
	"github.com/goofr-group/go-math/vector2"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/game/tag"
)

// CameraController defines the structure of the camera controller behaviour.
type CameraController struct {
	camera *rendering.Camera

	playerObject    *game.Object    // Defines the player object.
	initialPosition vector2.Vector2 // Defines the camera initial position.

	previousPosition vector2.Vector2 // Defines the previous position of the camera.
	currentPosition  vector2.Vector2 // Defines the current position of the camera. It represents the target position of the current level.

	transition      float64 // Defines the current amount, in the range [0; 1], that has been transitioned from the previousPosition to the currentPosition.
	transitionSpeed float64 // Defines the speed of the animation transition.
}

// NewCameraController returns a new camera controller behaviour.
func NewCameraController(
	camera *rendering.Camera,
	cameraTransitionSpeed float64,
) CameraController {
	// Check that the camera transition speed is valid.
	if cameraTransitionSpeed <= 0 {
		// Invalid speed, defaults to 1.
		cameraTransitionSpeed = 1
	}

	return CameraController{
		camera: camera,

		transitionSpeed: cameraTransitionSpeed,
	}
}

func (b CameraController) Enabled() bool {
	return true
}

func (b *CameraController) Start(e *engine.Engine) error {
	// Get the player object.
	b.playerObject = e.World().FindGameObjectWithTag(tag.Player)

	// Save the camera initial position.
	b.initialPosition = b.camera.Position

	// Initialize the previous and current positions.
	b.previousPosition = b.initialPosition
	b.currentPosition = b.initialPosition

	return nil
}

func (b *CameraController) Update(e *engine.Engine) error {
	// Check if the player object is accessible.
	if b.playerObject == nil {
		return nil
	}
	if b.playerObject.Collider == nil {
		return nil
	}

	// Get the player bounds.
	playerBounds := b.playerObject.Collider.Bounds()

	// Compute the camera position based on the player minimum bound.
	level := int((playerBounds.Min.Y - b.initialPosition.Y + b.camera.PixelHeight*0.5) / b.camera.PixelHeight)

	// newPosition represents the new position of the camera considering the current level of the player.
	newPosition := vector2.Vector2{
		X: b.camera.Position.X,
		Y: b.camera.PixelHeight*float64(level) + b.initialPosition.Y,
	}

	// Check whether the camera position has changed.
	if newPosition != b.currentPosition {
		// Update the previous and current positions.
		b.previousPosition = b.camera.Position
		b.currentPosition = newPosition

		// Reset the animation transition.
		b.transition = 0
	}

	// If the transition has been completed, avoid unnecessary computation.
	if b.transition >= 1 {
		return nil
	}

	// Compute the new animation transition.
	b.transition = mathf.Clamp(b.transition+e.Time().DeltaTime*b.transitionSpeed, 0, 1)

	// Apply the current transition using an easing function.
	b.camera.Position = vector2.Lerp(b.previousPosition, b.currentPosition, easeOutSine(b.transition))

	return nil
}
