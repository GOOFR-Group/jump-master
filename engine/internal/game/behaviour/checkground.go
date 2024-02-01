package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/go-math/rotation/matrix"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/game/tag"
)

// CheckGround defines the structure of the behaviour to check if the object is in contact with the ground.
type CheckGround struct {
	object *game.Object

	isGrounded      bool  // Defines if the object is in contact with the ground.
	currentGroundID int64 // Defines the id of the object that the current object is in contact with.
}

// NewCheckGround returns a new behaviour to check if the object is in contact with the ground.
func NewCheckGround(object *game.Object) CheckGround {
	return CheckGround{
		object: object,

		isGrounded:      false,
		currentGroundID: -1,
	}
}

func (b CheckGround) Enabled() bool {
	return true
}

func (b *CheckGround) Start(_ *engine.Engine) error {
	// Reset the position of the object.
	b.resetPosition()

	return nil
}

func (b *CheckGround) OnTriggerStay(e *engine.Engine, otherID int64) error {
	// Get the colliding object.
	otherObject := e.World().GetGameObjectByID(otherID)
	if otherObject == nil {
		return nil
	}

	// Check if the colliding object contains the platform tag.
	// If not, there is not contact with the ground.
	if otherObject.Tag != tag.Platform {
		return nil
	}

	// Set the flag as true since the current object is touching the ground.
	// Also update the id of the current ground object being touched.
	b.isGrounded = true
	b.currentGroundID = otherID

	return nil
}

func (b *CheckGround) OnTriggerExit(e *engine.Engine, otherID int64) error {
	// Check if the collision being exited is related to the current ground object in contact.
	if b.currentGroundID != otherID {
		return nil
	}

	// Get the colliding object.
	otherObject := e.World().GetGameObjectByID(otherID)
	if otherObject == nil {
		return nil
	}

	// Check if the colliding object contains the platform tag.
	if otherObject.Tag != tag.Platform {
		return nil
	}

	// Set the flag as false since the current object is not touching the ground anymore.
	b.isGrounded = false

	return nil
}

// IsGrounded returns true if the current object is touching the ground. The ground is represented by an object with the
// Platform tag.
func (b CheckGround) IsGrounded() bool {
	return b.isGrounded
}

// resetPosition resets the position of the object.
// Places the current object below its parent.
func (b *CheckGround) resetPosition() {
	// Get the parent object.
	parent := b.object.Transform.Parent()
	if parent == nil {
		return
	}

	parentObject := parent.GameObject()
	if parentObject == nil {
		return
	}

	// Check if any of the objects are misconfigured.
	if b.object.Collider == nil {
		return
	}
	if parentObject.Collider == nil {
		return
	}

	// Get the height of the current object based on its collider.
	bounds := b.object.Collider.Bounds()
	height := bounds.Max.Y - bounds.Min.Y

	// Compute the position to place the current object below its parent.
	position := parentObject.Transform.Position
	position.Y = parentObject.Collider.Bounds().Min.Y - height*0.5

	// Update the position and rotation of the object.
	b.object.Transform.Position = position
	b.object.Transform.Rotation = matrix.Identity()
}
