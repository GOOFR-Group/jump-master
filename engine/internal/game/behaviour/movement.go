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
}

// NewMovement returns a new movement behaviour with the given options.
func NewMovement(
	object *game.Object,
	actionManager *action.Manager,
	options MovementOptions,
) Movement {
	return Movement{
		object:        object,
		actionManager: actionManager,
		options:       options,
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

	// TODO: Update the current movement behaviour to appear more responsive. Also add better documentation.
	force := vector2.Zero()
	if b.actionManager.Action(input.Left) {
		force = force.Add(vector2.Left())
	}
	if b.actionManager.Action(input.Right) {
		force = force.Add(vector2.Right())
	}

	b.object.RigidBody.AddImpulse(force.Mul(b.options.Speed))

	return nil
}
