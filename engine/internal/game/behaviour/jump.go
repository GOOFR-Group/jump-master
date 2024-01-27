package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/action"
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/go-math/mathf"
	"github.com/goofr-group/go-math/vector2"
	"github.com/goofr-group/physics-engine/pkg/game"

	input "github.com/goofr-group/jump-master/engine/internal/game/action"
)

// JumpOptions defines the structure of the jump options.
type JumpOptions struct {
	Impulse           float64 `json:"impulse"`           // Defines the base impulse of the jump.
	MaxImpulse        float64 `json:"maxImpulse"`        // Defines the maximum impulse of the jump.
	ImpulseMultiplier float64 `json:"impulseMultiplier"` // Defines the multiplier to apply in the base impulse each frame the jump action is performed.
	DiagonalAngle     float64 `json:"diagonalAngle"`     // Defines the angle in degrees to apply when jumping left or right.
}

// Jump defines the structure of the jump behaviour.
type Jump struct {
	object        *game.Object
	actionManager *action.Manager
	options       JumpOptions

	checkGround *CheckGround

	accumulatedImpulse float64 // Defines the current accumulated jump impulse.
	canJump            bool    // Defines if the object is able to jump.
}

// NewJump returns a new jump behaviour with the given options.
func NewJump(
	object *game.Object,
	actionManager *action.Manager,
	checkGround *CheckGround,
	options JumpOptions,
) Jump {
	return Jump{
		object:        object,
		actionManager: actionManager,
		checkGround:   checkGround,
		options:       options,
	}
}

func (b Jump) Enabled() bool {
	return true
}

func (b *Jump) FixedUpdate(_ *engine.Engine) error {
	// Check if the rigid body is accessible.
	if b.object == nil {
		return nil
	}
	if b.object.RigidBody == nil {
		return nil
	}

	// Check if the object can jump.
	if !b.canJump {
		return nil
	}

	// Apply the accumulated impulse.
	force := vector2.Up().Mul(b.accumulatedImpulse)
	b.object.RigidBody.AddVelocity(force)

	// Reset the accumulated impulse and jump flag.
	b.accumulatedImpulse = 0
	b.canJump = false

	return nil
}

// TODO: Jump to direction player is moving.

func (b *Jump) Update(_ *engine.Engine) error {
	// Check if the object is in contact with the ground.
	if !b.checkGround.IsGrounded() {
		b.accumulatedImpulse = 0
		return nil
	}

	// Check if the jump action is being performed.
	if b.actionManager.Action(input.Jump) {
		// Apply the impulse multiplier.
		b.accumulatedImpulse += b.options.Impulse * b.options.ImpulseMultiplier
		// Ensure that the accumulated impulse is not greater than the maximum defined.
		b.accumulatedImpulse = mathf.Min(b.accumulatedImpulse, b.options.MaxImpulse)
	}

	// Check if the jump action was released.
	if !b.actionManager.ActionEnded(input.Jump) {
		return nil
	}

	// The object is able to jump
	b.canJump = true

	return nil
}
