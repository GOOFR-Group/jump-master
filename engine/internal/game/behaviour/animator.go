package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/config"
	"github.com/goofr-group/jump-master/engine/internal/game/property"
)

// Animator defines the structure of the animator behaviour.
type Animator struct {
	object     *game.Object
	animations config.Animations

	currentAnimation string  // Defines the current animation key.
	currentFrame     int     // Defines the frame of the current animation.
	currentTimer     float64 // Defines the timer of the current animation frame.
}

// NewAnimator returns a new animator behaviour with the given animations.
func NewAnimator(
	object *game.Object,
	animations config.Animations,
) Animator {
	return Animator{
		object:     object,
		animations: animations,
	}
}

func (b Animator) Enabled() bool {
	return true
}

func (b *Animator) Update(e *engine.Engine) error {
	// Check if the object is accessible.
	if b.object == nil {
		return nil
	}

	// Get the current animation configuration.
	animatorConfigs, ok := b.animations[b.currentAnimation]
	if !ok {
		return nil
	}

	// Update the current object image property.
	if b.currentFrame < len(animatorConfigs.Frames) {
		frame := animatorConfigs.Frames[b.currentFrame]
		b.object.SetProperty(property.Image, frame)
	}

	// Update the current timer.
	b.currentTimer -= e.Time().DeltaTime
	if b.currentTimer >= 0 {
		return nil
	}

	// Compute next frame.
	nextFrame := (b.currentFrame + 1) % len(animatorConfigs.Frames)
	if nextFrame == 0 && !animatorConfigs.Repeat {
		nextFrame = b.currentFrame
	}

	// Update the current frame of the animation.
	b.currentFrame = nextFrame
	b.currentTimer = animatorConfigs.Duration

	return nil
}

// SetAnimation updates the current animation being displayed.
func (b *Animator) SetAnimation(animation string) {
	if b.animations == nil {
		return
	}

	// Check if the animation is already playing.
	if b.currentAnimation == animation {
		return
	}

	// Get the given animation.
	animatorConfigs, ok := b.animations[animation]
	if !ok {
		return
	}

	// Update the current animation.
	b.currentAnimation = animation
	b.currentFrame = 0
	b.currentTimer = animatorConfigs.Duration
}
