//go:build js && wasm

package main

import "syscall/js"

const (
	goVersion = "go"
	gitTag    = "tag"
	gitCommit = "commit"
	build     = "build"
)

func jsVersion(buildGoVersion, buildTag, buildCommit, buildTime string) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response := map[string]interface{}{
			goVersion: js.Null(),
			gitCommit: js.Null(),
			gitTag:    js.Null(),
			build:     js.Null(),
		}

		if buildGoVersion != "" {
			response[goVersion] = buildGoVersion
		}
		if buildTag != "" {
			response[gitTag] = buildTag
		}
		if buildCommit != "" {
			response[gitCommit] = buildCommit
		}
		if buildTime != "" {
			response[build] = buildTime
		}

		return response
	})
}
