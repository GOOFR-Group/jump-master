package behaviour

import (
	"math"

	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/go-math/mathf"
	"github.com/goofr-group/go-math/rotation/matrix"
	"github.com/goofr-group/go-math/vector2"
	"github.com/goofr-group/physics-engine/pkg/collision"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/config"
	"github.com/goofr-group/jump-master/engine/internal/game/animation"
	"github.com/goofr-group/jump-master/engine/internal/game/sound"
	"github.com/goofr-group/jump-master/engine/internal/game/tag"
)

// KnockBack defines the structure of the Knock-back behaviour.
type KnockBack struct {
	object *game.Object
	config config.KnockBack

	checkGround     *CheckGround
	checkCeiling    *CheckCeiling
	jump            *Jump
	animator        *Animator
	soundController *SoundController

	// platforms defines the map of platform objects that the current object is in contact with. The map represents the
	// state of the contact by the platform object id.
	platforms map[int64]bool
	// previousVelocity defines the velocity value from the previous physics update. Used to check if the object was
	// falling in a straight line before colliding.
	previousVelocity vector2.Vector2
}

// NewKnockBack returns a new knock-back behaviour with the given configuration.
func NewKnockBack(
	object *game.Object,
	config config.KnockBack,
	checkGround *CheckGround,
	checkCeiling *CheckCeiling,
	jump *Jump,
	animator *Animator,
	soundController *SoundController,
) KnockBack {
	return KnockBack{
		object:          object,
		config:          config,
		checkGround:     checkGround,
		checkCeiling:    checkCeiling,
		jump:            jump,
		animator:        animator,
		soundController: soundController,

		platforms: make(map[int64]bool),
	}
}

func (b KnockBack) Enabled() bool {
	return true
}

func (b *KnockBack) FixedUpdate(_ *engine.Engine) error {
	// Check if the rigid body is accessible.
	if b.object == nil {
		return nil
	}
	if b.object.RigidBody == nil {
		return nil
	}

	// Update the previous velocity.
	b.previousVelocity = b.object.RigidBody.Velocity

	return nil
}

func (b *KnockBack) OnCollisionEnter(e *engine.Engine, otherID int64, manifold collision.Manifold) error {
	// Check if the rigid body is accessible.
	if b.object == nil {
		return nil
	}
	if b.object.RigidBody == nil {
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

	// alreadyInContact defines if the object is already in contact with a platform.
	alreadyInContact := b.PlatformContact()

	// Set the current platform as true since it is touching the object.
	b.platforms[otherID] = true

	// Check if the object is already in contact with a platform, if it is on the ground or in touch with the ceiling.
	// There is also no need to apply knock-back when the velocity vector of the object represents a 90º angle (object
	// falling in a straight line).
	if alreadyInContact || b.checkGround.TouchingGround() || b.checkCeiling.TouchingCeiling() ||
		mathf.Approximately(math.Abs(vector2.Up().Dot(b.previousVelocity.Normalized())), 1) {
		return nil
	}

	// Get the collision contact point.
	var contactPoint vector2.Vector2
	for _, cp := range manifold.ContactPoints {
		contactPoint = cp.Position
	}

	// Compute the diagonal angle based on the fraction of used jump impulse.
	diagonalAngle := b.config.DiagonalAngle
	diagonalAngle *= b.jump.UsedImpulse() / b.jump.MaxImpulse()

	// Compute the rotated direction from the right side.
	rotation := matrix.FromEuler(diagonalAngle)
	direction := rotation.RotateVector(vector2.Right())

	// If the object is in the left side of the collision, invert the direction vector.
	if b.object.Transform.Position.X < contactPoint.X {
		direction.X = -direction.X
	}

	// Apply the knock-back velocity based on the computed rotation and impulse.
	velocity := direction.Mul(b.config.Impulse)
	b.object.RigidBody.AddAcceleration(velocity)
	b.animator.SetAnimation(animation.KnockBack)
	b.soundController.AddPlayerSound(sound.KnockBack)

	return nil
}

func (b *KnockBack) OnCollisionExit(e *engine.Engine, otherID int64, _ collision.Manifold) error {
	// Get the colliding object.
	otherObject := e.World().GetGameObjectByID(otherID)
	if otherObject == nil {
		return nil
	}

	// Check if the colliding object contains the platform tag.
	if otherObject.Tag != tag.Platform {
		return nil
	}

	// Set the current platform as false since it is not touching the object anymore.
	b.platforms[otherID] = false

	return nil
}

// PlatformContact returns true if the current object is touching any platform. The platform is represented by any
// object with the Platform tag.
func (b KnockBack) PlatformContact() bool {
	for _, touching := range b.platforms {
		if touching {
			return true
		}
	}

	return false
}
