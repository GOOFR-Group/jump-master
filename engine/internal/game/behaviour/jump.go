package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/action"
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/go-math/mathf"
	"github.com/goofr-group/go-math/rotation/matrix"
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
		options:       options,
		checkGround:   checkGround,

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

	// Check if the object can jump.
	if !b.canJump {
		return nil
	}

	// Compute the diagonal angle based on the fraction of accumulated impulse.
	// Subtract the diagonal angle from 90 because the velocity vector already starts at 90º degrees.
	diagonalAngle := 90 - b.options.DiagonalAngle
	diagonalAngle *= b.accumulatedImpulse / b.options.MaxImpulse

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
		b.accumulatedImpulse += b.options.Impulse * b.options.ImpulseMultiplier
		// Ensure that the accumulated impulse is not greater than the maximum defined.
		b.accumulatedImpulse = mathf.Min(b.accumulatedImpulse, b.options.MaxImpulse)

		// Reset the horizontal velocity of the object when the jump action is being performed.
		b.object.RigidBody.Velocity.X = 0
	}

	// Check if the jump action was released.
	if b.actionManager.ActionEnded(input.Jump) {
		// The object is able to jump
		b.canJump = true
	}

	return nil
}
