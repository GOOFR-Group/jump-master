package config

import (
	"encoding/json"
	"fmt"

	"github.com/goofr-group/jump-master/engine"
)

const (
	// pathPlayerConfig defines the path of the player configuration.
	pathPlayerConfig = "configs/player.json"
	// pathWorldConfig defines the path of the world configuration.
	pathWorldConfig = "configs/world.json"
)

// LoadPlayer loads the player configuration.
func LoadPlayer() (Player, error) {
	data, err := engine.ConfigsFS.ReadFile(pathPlayerConfig)
	if err != nil {
		return Player{}, fmt.Errorf("failed to read file: %w", err)
	}

	var config Player
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Player{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return config, nil
}

// LoadWorld loads the world configuration.
func LoadWorld() (World, error) {
	data, err := engine.ConfigsFS.ReadFile(pathWorldConfig)
	if err != nil {
		return World{}, fmt.Errorf("failed to read file: %w", err)
	}

	var config World
	err = json.Unmarshal(data, &config)
	if err != nil {
		return World{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return config, nil
}
