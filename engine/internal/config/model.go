package config

import "github.com/goofr-group/jump-master/engine/internal/game/behaviour"

// Player defines the structure of the player configuration.
type Player struct {
	Movement   behaviour.MovementOptions            `json:"movement"`
	Jump       behaviour.JumpOptions                `json:"jump"`
	Animations map[string]behaviour.AnimatorOptions `json:"animations"`
}
