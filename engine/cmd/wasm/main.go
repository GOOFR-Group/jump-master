//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/goofr-group/jump-master/engine/internal/app"
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
	// Set up engine.
	app := app.New()

	// Set up WASM API.
	js.Global().Set(entryPoint, make(map[string]interface{}))
	module := js.Global().Get(entryPoint)
	module.Set(methodVersion, jsVersion(Version, GitCommit, GoVersion, Build))
	module.Set(methodStep, jsStep(app))

	// Set up game world.
	err := app.StartGameWorld()
	if err != nil {
		err = fmt.Errorf("failed to start game world: %w", err)
		js.Global().Get("console").Call("error", err.Error())
		panic(err)
	}

	// Keep context open.
	<-make(chan struct{})
}
