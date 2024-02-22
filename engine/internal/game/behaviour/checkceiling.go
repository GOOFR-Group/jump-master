package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/go-math/rotation/matrix"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/game/tag"
)

// CheckCeiling defines the structure of the behaviour to check if the object is in contact with the ceiling.
type CheckCeiling struct {
	object *game.Object

	// ceilings defines the map of ceiling objects that the current object is in contact with. The map represents the
	// state of the contact by the ceiling object id.
	ceilings map[int64]bool
}

// NewCheckCeiling returns a new behaviour to check if the object is in contact with the ceiling.
func NewCheckCeiling(object *game.Object) CheckCeiling {
	return CheckCeiling{
		object: object,

		ceilings: make(map[int64]bool),
	}
}

func (b CheckCeiling) Enabled() bool {
	return true
}

func (b *CheckCeiling) Update(_ *engine.Engine) error {
	// Reset the position of the object.
	b.resetPosition()

	return nil
}

func (b *CheckCeiling) OnTriggerEnter(e *engine.Engine, otherID int64) error {
	// Get the colliding object.
	otherObject := e.World().GetGameObjectByID(otherID)
	if otherObject == nil {
		return nil
	}

	// Check if the colliding object contains the platform tag.
	// If not, there is not contact with the ceiling.
	if otherObject.Tag != tag.Platform {
		return nil
	}

	// Set the current ceiling as true since it is touching the object.
	b.ceilings[otherID] = true

	return nil
}

func (b *CheckCeiling) OnTriggerExit(e *engine.Engine, otherID int64) error {
	// Get the colliding object.
	otherObject := e.World().GetGameObjectByID(otherID)
	if otherObject == nil {
		return nil
	}

	// Check if the colliding object contains the platform tag.
	if otherObject.Tag != tag.Platform {
		return nil
	}

	// Set the current ceiling as false since it is not touching the object anymore.
	b.ceilings[otherID] = false

	return nil
}

// TouchingCeiling returns true if the current object is touching the ceiling. The ceiling is represented by any object
// with the Platform tag that is in contact with the head of this object.
func (b CheckCeiling) TouchingCeiling() bool {
	for _, touching := range b.ceilings {
		if touching {
			return true
		}
	}

	return false
}

// resetPosition resets the position of the object.
// Places the current object above its parent.
func (b *CheckCeiling) resetPosition() {
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
	position.Y = parentObject.Collider.Bounds().Max.Y + height*0.5

	// Update the position and rotation of the object.
	b.object.Transform.Position = position
	b.object.Transform.Rotation = matrix.Identity()
}
