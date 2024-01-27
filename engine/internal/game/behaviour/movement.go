package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/action"
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/go-math/vector2"
	"github.com/goofr-group/physics-engine/pkg/game"

	input "github.com/goofr-group/jump-master/engine/internal/game/action"
)

// MovementOptions defines the structure of the movement options.
type MovementOptions struct {
	Speed float64 `json:"speed"` // Defines the movement speed.
}

// Movement defines the structure of the movement behaviour.
type Movement struct {
	object        *game.Object
	actionManager *action.Manager
	options       MovementOptions

	checkGround *CheckGround
}

// NewMovement returns a new movement behaviour with the given options.
func NewMovement(
	object *game.Object,
	actionManager *action.Manager,
	options MovementOptions,
	checkGround *CheckGround,
) Movement {
	return Movement{
		object:        object,
		actionManager: actionManager,
		options:       options,
		checkGround:   checkGround,
	}
}

func (b Movement) Enabled() bool {
	return true
}

func (b *Movement) Update(_ *engine.Engine) error {
	// Check if the jump action started.
	if !b.actionManager.ActionStarted(input.Jump) {
		return nil
	}

	// Reset the horizontal velocity of the object when the jump action is started.
	b.object.RigidBody.Velocity.X = 0

	return nil
}

func (b *Movement) FixedUpdate(_ *engine.Engine) error {
	// Check if the rigid body is accessible.
	if b.object == nil {
		return nil
	}
	if b.object.RigidBody == nil {
		return nil
	}

	// Check if the object is in contact with the ground.
	if !b.checkGround.IsGrounded() {
		return nil
	}

	// Check if the jump action is being performed.
	if b.actionManager.Action(input.Jump) {
		return nil
	}

	// Compute the velocity to add to the object based on the left and right actions.
	velocity := vector2.Zero()
	if b.actionManager.Action(input.Left) {
		// Add velocity to the left direction.
		velocity = velocity.Add(vector2.Left())
	}
	if b.actionManager.Action(input.Right) {
		// Add velocity to the right direction.
		velocity = velocity.Add(vector2.Right())
	}

	// Reset the horizontal velocity of the object when no movement action is performed.
	if velocity.Zero() {
		b.object.RigidBody.Velocity.X = 0
		return nil
	}

	// Add the computed velocity when the movement actions are performed.
	b.object.RigidBody.AddVelocity(velocity.Mul(b.options.Speed))

	return nil
}
