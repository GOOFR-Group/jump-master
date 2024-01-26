package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/action"
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/go-math/vector2"
	"github.com/goofr-group/physics-engine/pkg/game"

	input "github.com/goofr-group/jump-master/engine/internal/game/action"
)

type MovementOptions struct {
	Speed float64
}

type Movement struct {
	object        *game.Object
	actionManager *action.Manager

	MovementOptions
}

func NewMove(
	object *game.Object,
	actionManager *action.Manager,
	options MovementOptions,
) Movement {
	return Movement{
		object:          object,
		actionManager:   actionManager,
		MovementOptions: options,
	}
}

func (b Movement) Enabled() bool {
	return true
}

func (b *Movement) FixedUpdate(_ *engine.Engine) error {
	if b.object == nil {
		return nil
	}
	if b.object.RigidBody == nil {
		return nil
	}

	force := vector2.Zero()
	if b.actionManager.Action(input.Left) {
		force = force.Add(vector2.Left())
	}
	if b.actionManager.Action(input.Right) {
		force = force.Add(vector2.Right())
	}

	b.object.RigidBody.AddImpulse(force.Mul(b.MovementOptions.Speed))

	return nil
}
