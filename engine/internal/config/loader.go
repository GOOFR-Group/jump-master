package config

import (
	"encoding/json"
	"fmt"

	"github.com/goofr-group/jump-master/engine"
)

const (
	// pathEngineConfig defines the path of the engine configuration.
	pathEngineConfig = "configs/engine.json"
	// pathPlayerConfig defines the path of the player configuration.
	pathPlayerConfig = "configs/player.json"
	// pathMapConfig defines the path of the map configuration.
	pathMapConfig = "configs/map.json"

	failedToReadFile  = "failed to read file"
	failedToUnmarshal = "failed to unmarshal"
)

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

// LoadMap loads the map configuration.
func LoadMap() (Map, error) {
	data, err := engine.ConfigsFS.ReadFile(pathMapConfig)
	if err != nil {
		return Map{}, fmt.Errorf("%s: %w", failedToReadFile, err)
	}

	var config Map

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Map{}, fmt.Errorf("%s: %w", failedToUnmarshal, err)
	}

	return config, nil
}
