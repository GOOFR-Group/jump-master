package domain

import (
	"github.com/goofr-group/game-engine/pkg/rendering"
	"github.com/goofr-group/physics-engine/pkg/game"
)

// GameState defines the state of the game.
type GameState struct {
	GameObjects []game.Object    `json:"gameObjects"`
	Camera      rendering.Camera `json:"camera"`
}
