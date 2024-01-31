package config

import (
	"encoding/json"
	"fmt"

	"github.com/goofr-group/jump-master/engine"
)

const (
	// pathPlayerConfig defines the path of the player configuration.
	pathPlayerConfig = "configs/player.jsonc"
)

// LoadPlayer loads the player configuration.
func LoadPlayer() (Player, error) {
	data, err := engine.ConfigsFS.ReadFile(pathPlayerConfig)
	if err != nil {
		return Player{}, fmt.Errorf("failed to read file: %w", err)
	}

	var player Player
	err = json.Unmarshal(data, &player)
	if err != nil {
		return Player{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return player, nil
}
