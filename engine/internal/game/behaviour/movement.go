package behaviour

import (
	"math"

	"github.com/goofr-group/game-engine/pkg/action"
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/go-math/mathf"
	"github.com/goofr-group/go-math/vector2"
	"github.com/goofr-group/physics-engine/pkg/game"

	input "github.com/goofr-group/jump-master/engine/internal/game/action"
	"github.com/goofr-group/jump-master/engine/internal/game/animation"
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
	animator    *Animator

	jumpAction bool
}

// NewMovement returns a new movement behaviour with the given options.
func NewMovement(
	object *game.Object,
	actionManager *action.Manager,
	options MovementOptions,
	checkGround *CheckGround,
	animator *Animator,
) Movement {
	return Movement{
		object:        object,
		actionManager: actionManager,
		options:       options,
		checkGround:   checkGround,
		animator:      animator,
	}
}

func (b Movement) Enabled() bool {
	return true
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
	if !b.checkGround.IsGrounded() || b.object.RigidBody.Velocity.Y > mathf.Epsilon64*100 {
		return nil
	}

	// Check if the jump action is being performed.
	if b.jumpAction {
		if b.actionManager.Action(input.Left) {
			b.object.Transform.Scale.X = -math.Abs(b.object.Transform.Scale.X)
		}
		if b.actionManager.Action(input.Right) {
			b.object.Transform.Scale.X = math.Abs(b.object.Transform.Scale.X)
		}
		return nil
	}

	// Compute the velocity to add to the object based on the left and right actions.
	velocity := vector2.Zero()
	if b.actionManager.Action(input.Left) {
		// Add velocity to the left direction.
		velocity = velocity.Add(vector2.Left())
		b.object.Transform.Scale.X = -math.Abs(b.object.Transform.Scale.X)
	}
	if b.actionManager.Action(input.Right) {
		// Add velocity to the right direction.
		velocity = velocity.Add(vector2.Right())
		b.object.Transform.Scale.X = math.Abs(b.object.Transform.Scale.X)
	}

	// Reset the horizontal velocity of the object when no movement action is performed.
	if velocity.Zero() {
		b.object.RigidBody.Velocity.X = 0
		b.animator.SetAnimation(animation.Idle)
		return nil
	}

	// Add the computed velocity when the movement actions are performed.
	b.object.RigidBody.AddVelocity(velocity.Mul(b.options.Speed))
	b.animator.SetAnimation(animation.Walk)

	return nil
}

func (b *Movement) Update(_ *engine.Engine) error {
	// Update the jump action.
	b.jumpAction = b.actionManager.Action(input.Jump)

	return nil
}
