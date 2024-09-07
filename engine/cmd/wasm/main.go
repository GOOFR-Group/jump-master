//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/goofr-group/jump-master/engine/internal/app"
	"github.com/goofr-group/jump-master/engine/internal/config"
)

const (
	// entryPoint defines the name of the entry point object.
	entryPoint = "engine"

	// Name of the methods within the entry point object.
	methodVersion = "version"
	methodStep    = "step"
)

// Build metadata to be set on compile-time.
var (
	GitCommit string
	Version   string
	GoVersion string
	Build     string
)

// main entry point for the application to register our engine API into a JavaScript global context.
func main() {
	// Load configurations.
	engineConfig, err := config.LoadEngine()
	if err != nil {
		err = fmt.Errorf("failed to load engine configuration: %w", err)
		panic(err)
	}

	playerConfig, err := config.LoadPlayer()
	if err != nil {
		err = fmt.Errorf("failed to load player configuration: %w", err)
		panic(err)
	}

	mapConfig, err := config.LoadMap()
	if err != nil {
		err = fmt.Errorf("failed to load map configuration: %w", err)
		panic(err)
	}

	// Set up engine.
	app := app.New(engineConfig, playerConfig, mapConfig)

	// Set up WASM API.
	js.Global().Set(entryPoint, make(map[string]interface{}))
	module := js.Global().Get(entryPoint)
	module.Set(methodVersion, jsVersion(GoVersion, Version, GitCommit, Build))
	module.Set(methodStep, jsStep(app))

	// Set up game world.
	err = app.StartGameWorld()
	if err != nil {
		err = fmt.Errorf("failed to start game world: %w", err)
		panic(err)
	}

	// Keep context open.
	<-make(chan struct{})
}
