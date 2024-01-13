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
			gitTag:    js.Null(),
			gitCommit: js.Null(),
			build:     js.Null(),
		}

		if len(buildGoVersion) != 0 {
			response[goVersion] = buildGoVersion
		}
		if len(buildTag) != 0 {
			response[gitTag] = buildTag
		}
		if len(buildCommit) != 0 {
			response[gitCommit] = buildCommit
		}
		if len(buildTime) != 0 {
			response[build] = buildTime
		}

		return response
	})
}
