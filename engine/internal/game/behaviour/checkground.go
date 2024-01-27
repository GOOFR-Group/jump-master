package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/go-math/rotation/matrix"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/game/tag"
)

type CheckGroundBehaviour struct {
	object *game.Object

	isGrounded      bool
	currentGroundID int64
}

func NewCheckGroundBehaviour(object *game.Object) CheckGroundBehaviour {
	return CheckGroundBehaviour{
		object: object,
	}
}

func (b CheckGroundBehaviour) Enabled() bool {
	return true
}

func (b *CheckGroundBehaviour) Start(_ *engine.Engine) error {
	b.resetPosition()

	return nil
}

func (b *CheckGroundBehaviour) FixedUpdate(_ *engine.Engine) error {
	b.resetPosition()

	return nil
}

func (b *CheckGroundBehaviour) OnTriggerEnter(e *engine.Engine, otherID int64) error {
	otherObject := e.World().GetGameObjectByID(otherID)
	if otherObject == nil {
		return nil
	}

	if otherObject.Tag != tag.Platform {
		return nil
	}

	b.currentGroundID = otherID
	b.isGrounded = true

	return nil
}

func (b *CheckGroundBehaviour) OnTriggerExit(e *engine.Engine, otherID int64) error {
	otherObject := e.World().GetGameObjectByID(otherID)
	if otherObject == nil {
		return nil
	}

	if otherObject.Tag != tag.Platform {
		return nil
	}
	if b.currentGroundID != otherID {
		return nil
	}

	b.isGrounded = false

	return nil
}

func (b CheckGroundBehaviour) CurrentGroundID() int64 {
	return b.currentGroundID
}

func (b CheckGroundBehaviour) IsGrounded() bool {
	return b.isGrounded
}

func (b *CheckGroundBehaviour) resetPosition() {
	parent := b.object.Transform.Parent().GameObject()
	position := parent.Transform.Position
	if b.object.Collider != nil && parent.Collider != nil {
		bounds := b.object.Collider.Bounds()
		height := bounds.Max.Y - bounds.Min.Y

		position.Y = parent.Collider.Bounds().Min.Y - height*0.5
	}

	b.object.Transform.Position = position
	b.object.Transform.Rotation = matrix.Identity()
}
