package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/game-engine/pkg/rendering"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/game/tag"
)

// CameraController defines the structure of the camera controller behaviour.
type CameraController struct {
	camera *rendering.Camera

	playerObject *game.Object // Defines the player object.
}

// NewCameraController returns a new camera controller behaviour.
func NewCameraController(
	camera *rendering.Camera,
) CameraController {
	return CameraController{
		camera: camera,
	}
}

func (b CameraController) Enabled() bool {
	return true
}

func (b *CameraController) Start(e *engine.Engine) error {
	// Get the player object.
	b.playerObject = e.World().FindGameObjectWithTag(tag.Player)

	return nil
}

func (b *CameraController) Update(_ *engine.Engine) error {
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
	level := int((playerBounds.Min.Y + b.camera.PixelHeight*0.5) / b.camera.PixelHeight)
	b.camera.Position.Y = b.camera.PixelHeight * float64(level)

	return nil
}
