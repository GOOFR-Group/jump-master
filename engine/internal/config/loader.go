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
)

// LoadEngine loads the engine configuration.
func LoadEngine() (Engine, error) {
	return loadConfig[Engine](pathEngineConfig)
}

// LoadPlayer loads the player configuration.
func LoadPlayer() (Player, error) {
	return loadConfig[Player](pathPlayerConfig)
}

// LoadMap loads the map configuration.
func LoadMap() (Map, error) {
	return loadConfig[Map](pathMapConfig)
}

// loadConfig returns the configuration in the specified path.
func loadConfig[T any](path string) (T, error) {
	var config T

	data, err := engine.ConfigsFS.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return config, nil
}
