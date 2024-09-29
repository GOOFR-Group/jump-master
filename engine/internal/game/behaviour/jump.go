package behaviour

import (
	"math"

	"github.com/goofr-group/game-engine/pkg/action"
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/go-math/mathf"
	"github.com/goofr-group/go-math/rotation/matrix"
	"github.com/goofr-group/go-math/vector2"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/config"
	input "github.com/goofr-group/jump-master/engine/internal/game/action"
	"github.com/goofr-group/jump-master/engine/internal/game/animation"
	"github.com/goofr-group/jump-master/engine/internal/game/sound"
)

// Jump defines the structure of the jump behaviour.
type Jump struct {
	object        *game.Object
	actionManager *action.Manager
	config        config.Jump

	checkGround     *CheckGround
	animator        *Animator
	soundController *SoundController

	usedImpulse        float64 // Defines the previously used jump impulse.
	accumulatedImpulse float64 // Defines the current accumulated jump impulse.
	canJump            bool    // Defines if the object is able to jump.
}

// NewJump returns a new jump behaviour with the given configuration.
func NewJump(
	object *game.Object,
	actionManager *action.Manager,
	config config.Jump,
	checkGround *CheckGround,
	animator *Animator,
	soundController *SoundController,
) Jump {
	return Jump{
		object:          object,
		actionManager:   actionManager,
		config:          config,
		checkGround:     checkGround,
		animator:        animator,
		soundController: soundController,

		usedImpulse:        0,
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
	if b.object.RigidBody.Velocity.Y < -Epsilon && !b.checkGround.TouchingGround() {
		b.animator.SetAnimation(animation.JumpFall)
	}

	// Check if the object can jump.
	if !b.canJump {
		return nil
	}

	// Limit the minimum jump impulse.
	b.accumulatedImpulse = math.Max(b.accumulatedImpulse, b.config.MinImpulse)

	// Compute the jump rotation based on the left and right actions.
	direction := vector2.Up()

	leftAction := b.actionManager.Action(input.Left)
	rightAction := b.actionManager.Action(input.Right)

	// Check that the left or right action is being performed. Also check that both actions are not being performed at
	// the same time.
	if (leftAction || rightAction) && !(leftAction && rightAction) {
		// Compute the diagonal angle based on the fraction of accumulated impulse.
		diagonalAngle := b.config.DiagonalAngle
		diagonalAngle *= b.accumulatedImpulse / b.config.MaxImpulse

		// Compute the rotated direction from the right side.
		rotation := matrix.FromEuler(diagonalAngle)
		direction = rotation.RotateVector(vector2.Right())

		// If the object is performing a left action, invert the direction vector.
		if leftAction {
			direction.X = -direction.X
		}
	}

	// Apply the jump velocity based on the computed rotation and accumulated impulse.
	b.usedImpulse = b.accumulatedImpulse
	velocity := direction.Mul(b.accumulatedImpulse)

	b.object.RigidBody.AddAcceleration(velocity)
	b.animator.SetAnimation(animation.Jump)
	b.soundController.AddPlayerSound(sound.Jump)

	// Reset the accumulated impulse and jump flag.
	b.accumulatedImpulse = 0
	b.canJump = false

	return nil
}

func (b *Jump) Update(e *engine.Engine) error {
	time := e.Time()

	// Check if the rigid body is accessible.
	if b.object == nil {
		return nil
	}
	if b.object.RigidBody == nil {
		return nil
	}

	// Check if the object is in contact with the ground and if the fall animation has already ended.
	if !b.checkGround.TouchingGround() ||
		(b.animator.Animation() == animation.Fall && !b.animator.AnimationEnded()) {
		b.accumulatedImpulse = 0
		return nil
	}

	// Check if the jump action is being performed.
	if b.actionManager.Action(input.Jump) {
		// Apply the impulse multiplier and ensure that the accumulated impulse is not greater than the maximum defined.
		b.accumulatedImpulse += b.config.Impulse * time.FixedDeltaTime
		b.accumulatedImpulse = mathf.Min(b.accumulatedImpulse, b.config.MaxImpulse)

		// Reset the horizontal velocity of the object when the jump action is being performed.
		b.object.RigidBody.Velocity.X = 0
		b.animator.SetAnimation(animation.JumpHold)
		b.soundController.AddPlayerSound(sound.JumpHold)
	}

	// Check if the jump action was released.
	if b.actionManager.ActionEnded(input.Jump) {
		// The object is able to jump
		b.canJump = true
	}

	return nil
}

// UsedImpulse returns the previously used jump impulse.
func (b Jump) UsedImpulse() float64 {
	return b.usedImpulse
}

// MaxImpulse returns the maximum jump impulse.
func (b Jump) MaxImpulse() float64 {
	return b.config.MaxImpulse
}
