package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/action"
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/go-math/mathf"
	"github.com/goofr-group/go-math/rotation/matrix"
	"github.com/goofr-group/go-math/vector2"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/config"
	input "github.com/goofr-group/jump-master/engine/internal/game/action"
	"github.com/goofr-group/jump-master/engine/internal/game/animation"
)

// Jump defines the structure of the jump behaviour.
type Jump struct {
	object        *game.Object
	actionManager *action.Manager
	config        config.Jump

	checkGround *CheckGround
	animator    *Animator

	accumulatedImpulse float64 // Defines the current accumulated jump impulse.
	canJump            bool    // Defines if the object is able to jump.
}

// NewJump returns a new jump behaviour with the given configuration.
func NewJump(
	object *game.Object,
	actionManager *action.Manager,
	checkGround *CheckGround,
	animator *Animator,
	config config.Jump,
) Jump {
	return Jump{
		object:        object,
		actionManager: actionManager,
		config:        config,
		checkGround:   checkGround,
		animator:      animator,

		accumulatedImpulse: 0,
		canJump:            false,
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

	// Check if the object is falling.
	if b.object.RigidBody.Velocity.Y < -Epsilon && !b.checkGround.IsGrounded() {
		b.animator.SetAnimation(animation.JumpFall)
	}

	// Check if the object can jump.
	if !b.canJump {
		return nil
	}

	// Compute the diagonal angle based on the fraction of accumulated impulse.
	// Subtract the diagonal angle from 90 because the velocity vector already starts at 90º degrees.
	diagonalAngle := 90 - b.config.DiagonalAngle
	diagonalAngle *= b.accumulatedImpulse / b.config.MaxImpulse

	// Compute the jump rotation based on the left and right actions.
	rotation := matrix.Identity()
	leftRotation := matrix.FromEuler(diagonalAngle)
	if b.actionManager.Action(input.Left) {
		rotation = rotation.Mul(leftRotation)
	}
	if b.actionManager.Action(input.Right) {
		rightRotation := leftRotation.Transpose()
		rotation = rotation.Mul(rightRotation)
	}

	// Apply the jump velocity based on the computed rotation and accumulated impulse.
	velocity := vector2.Up()
	velocity = rotation.RotateVector(velocity)
	velocity = velocity.Mul(b.accumulatedImpulse)

	b.object.RigidBody.AddVelocity(velocity)
	b.animator.SetAnimation(animation.Jump)

	// Reset the accumulated impulse and jump flag.
	b.accumulatedImpulse = 0
	b.canJump = false

	return nil
}

func (b *Jump) Update(_ *engine.Engine) error {
	// Check if the rigid body is accessible.
	if b.object == nil {
		return nil
	}
	if b.object.RigidBody == nil {
		return nil
	}

	// Check if the object is in contact with the ground.
	if !b.checkGround.IsGrounded() {
		b.accumulatedImpulse = 0
		return nil
	}

	// Check if the jump action is being performed.
	if b.actionManager.Action(input.Jump) {
		// Apply the impulse multiplier.
		b.accumulatedImpulse += b.config.Impulse * b.config.ImpulseMultiplier
		// Ensure that the accumulated impulse is not greater than the maximum defined.
		b.accumulatedImpulse = mathf.Min(b.accumulatedImpulse, b.config.MaxImpulse)

		// Reset the horizontal velocity of the object when the jump action is being performed.
		b.object.RigidBody.Velocity.X = 0
		b.animator.SetAnimation(animation.JumpHold)
	}

	// Check if the jump action was released.
	if b.actionManager.ActionEnded(input.Jump) {
		// The object is able to jump
		b.canJump = true
	}

	return nil
}
