//go:build js && wasm

package main

import (
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

	// Keep context open.
	<-make(chan struct{})
}
