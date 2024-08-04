package config

import (
	"encoding/json"
	"fmt"

	"github.com/goofr-group/jump-master/engine"
)

const (
	// pathPlayerConfig defines the path of the player configuration.
	pathPlayerConfig = "configs/player.json"
	// pathEngineConfig defines the path of the engine configuration.
	pathEngineConfig = "configs/engine.json"

	failedToReadFile  = "failed to read file"
	failedToUnmarshal = "failed to unmarshal"
)

// LoadPlayer loads the player configuration.
func LoadPlayer() (Player, error) {
	data, err := engine.ConfigsFS.ReadFile(pathPlayerConfig)
	if err != nil {
		return Player{}, fmt.Errorf("%s: %w", failedToReadFile, err)
	}

	var config Player

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Player{}, fmt.Errorf("%s: %w", failedToUnmarshal, err)
	}

	return config, nil
}

// LoadEngine loads the engine configuration.
func LoadEngine() (Engine, error) {
	data, err := engine.ConfigsFS.ReadFile(pathEngineConfig)
	if err != nil {
		return Engine{}, fmt.Errorf("%s: %w", failedToReadFile, err)
	}

	var config Engine

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Engine{}, fmt.Errorf("%s: %w", failedToUnmarshal, err)
	}

	return config, nil
}
