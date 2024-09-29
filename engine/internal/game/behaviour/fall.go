package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/config"
	"github.com/goofr-group/jump-master/engine/internal/game/animation"
	"github.com/goofr-group/jump-master/engine/internal/game/sound"
)

// Fall defines the structure of the fall behaviour.
type Fall struct {
	object *game.Object
	config config.Fall

	checkGround     *CheckGround
	animator        *Animator
	soundController *SoundController

	timer float64 // Defines the timer that captures the amount of time the object is falling.
}

// NewFall returns a new fall behaviour with the given configuration.
func NewFall(
	object *game.Object,
	config config.Fall,
	checkGround *CheckGround,
	animator *Animator,
	soundController *SoundController,
) Fall {
	return Fall{
		object:          object,
		config:          config,
		checkGround:     checkGround,
		animator:        animator,
		soundController: soundController,
	}
}

func (b Fall) Enabled() bool {
	return true
}

func (b *Fall) Update(e *engine.Engine) error {
	time := e.Time()

	// Check if the rigid body is accessible.
	if b.object == nil {
		return nil
	}
	if b.object.RigidBody == nil {
		return nil
	}

	// Check if the object is falling.
	if b.object.RigidBody.Velocity.Y < -Epsilon && !b.checkGround.TouchingGround() {
		// If it is falling, update the timer.
		b.timer += time.DeltaTime
		return nil
	}

	// Check if the object was falling for longer than the allowed duration.
	if b.timer > b.config.AllowedDuration {
		b.object.RigidBody.Velocity.X = 0
		b.animator.SetAnimation(animation.Fall)
		b.soundController.AddPlayerSound(sound.Fall)
	}

	// Reset the timer.
	b.timer = 0

	return nil
}
