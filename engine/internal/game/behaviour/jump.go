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
}

// Jump defines the structure of the jump behaviour.
type Jump struct {
	object        *game.Object
	actionManager *action.Manager
	options       JumpOptions

	accumulatedImpulse float64
	canJump            bool
}

// NewJump returns a new jump behaviour with the given options.
func NewJump(
	object *game.Object,
	actionManager *action.Manager,
	options JumpOptions,
) Jump {
	return Jump{
		object:        object,
		actionManager: actionManager,
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

	// Check if the player should jump.
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

func (b *Jump) Update(_ *engine.Engine) error {
	// Check if the jump action is being performed.
	if b.actionManager.Action(input.Jump) {
		// Apply the impulse multiplier.
		b.accumulatedImpulse += b.options.Impulse * b.options.ImpulseMultiplier
		// Ensure that the accumulated impulse is not greater than the maximum defined.
		b.accumulatedImpulse = mathf.Min(b.accumulatedImpulse, b.options.MaxImpulse)
	}

	// Check if the jump action was released.
	if b.actionManager.ActionEnded(input.Jump) {
		b.canJump = true
	}

	return nil
}
