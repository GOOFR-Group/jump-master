//go:build js && wasm

package main

import (
	"errors"
	"fmt"
	"syscall/js"

	"github.com/goofr-group/jump-master/engine/internal/app"
	"github.com/goofr-group/jump-master/engine/internal/domain"
)

type stepRequest struct {
	Actions map[string]bool
}

// jsStep runs a step of the game engine.
func jsStep(app *app.App) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var gameState domain.GameState
		err := func() error {
			if len(args) != 1 {
				return errors.New("unexpected number of arguments in step")
			}

			request, err := unmarshalStepRequest(args[0])
			if err != nil {
				return fmt.Errorf("failed to unmarshal step request: %w", err)
			}

			gameState, err = app.GameStep(request.Actions)
			if err != nil {
				return fmt.Errorf("failed to perform game step: %w", err)
			}

			return nil
		}()

		return marshalStepResponse(gameState, err)
	})
}

// unmarshalStepRequest deserializes the step request.
func unmarshalStepRequest(value js.Value) (stepRequest, error) {
	if value.IsNull() || value.Type() != js.TypeObject {
		return stepRequest{}, errors.New("unexpected object type")
	}

	request := stepRequest{
		Actions: make(map[string]bool, value.Length()),
	}

	for i := 0; i < value.Length(); i++ {
		action := value.Index(i)
		if action.Length() != 2 {
			return stepRequest{}, errors.New("unexpected action length")
		}

		k := action.Index(0)
		if k.Type() != js.TypeString {
			return stepRequest{}, errors.New("unexpected action key")
		}

		v := action.Index(1)
		if v.Type() != js.TypeBoolean {
			return stepRequest{}, errors.New("unexpected action value")
		}

		request.Actions[k.String()] = v.Bool()
	}

	return request, nil
}

// marshalStepResponse serializes the step response and returns a javascript object with the game state information.
func marshalStepResponse(gameState domain.GameState, err error) map[string]interface{} {
	response := map[string]interface{}{
		"error":       nil,
		"gameObjects": marshalGameObjects(gameState.GameObjects),
		"camera":      marshalCamera(gameState.Camera),
	}

	if err != nil {
		response["error"] = err.Error()
	}

	return response
}
