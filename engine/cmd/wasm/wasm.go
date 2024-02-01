//go:build js && wasm

package main

import (
	"github.com/goofr-group/game-engine/pkg/rendering"
	"github.com/goofr-group/go-math/vector2"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/game/property"
)

func marshalVector2(v vector2.Vector2) map[string]interface{} {
	return map[string]interface{}{
		"x": v.X,
		"y": v.Y,
	}
}

func marshalTransform(transform game.Transform2D) map[string]interface{} {
	return map[string]interface{}{
		"position": marshalVector2(transform.Position),
		"rotation": transform.Rotation.Radians(),
		"scale":    marshalVector2(transform.Scale),
	}
}

func marshalRigidBody(rigidBody *game.RigidBody2D) map[string]interface{} {
	if rigidBody == nil {
		return nil
	}

	return map[string]interface{}{
		"bodyType":           int(rigidBody.BodyType),
		"collisionDetection": int(rigidBody.CollisionDetection),
		"interpolation":      int(rigidBody.Interpolation),
		"autoMass":           rigidBody.AutoMass,
		"mass":               rigidBody.Mass,
		"drag":               rigidBody.Drag,
		"velocity":           marshalVector2(rigidBody.Velocity),
		"angularDrag":        rigidBody.AngularDrag,
		"angularVelocity":    rigidBody.AngularVelocity,
		"gravityScale":       rigidBody.GravityScale,
	}
}

func marshalRenderer(renderer *game.Renderer, image interface{}) map[string]interface{} {
	if renderer == nil {
		return nil
	}

	return map[string]interface{}{
		"width":  renderer.Width,
		"height": renderer.Height,
		"offset": marshalVector2(renderer.Offset),
		"layer":  renderer.Layer,
		"image":  image,
	}
}

func marshalGameObject(gameObject game.Object) map[string]interface{} {
	return map[string]interface{}{
		"id": gameObject.ID(),

		"active": gameObject.Active,
		"tag":    gameObject.Tag,

		"transform": marshalTransform(gameObject.Transform),
		"rigidBody": marshalRigidBody(gameObject.RigidBody),
		"renderer":  marshalRenderer(gameObject.Renderer, gameObject.Property(property.Image)),
	}
}

func marshalGameObjects(gameObjects []game.Object) []interface{} {
	objects := make([]interface{}, len(gameObjects))
	for i := 0; i < len(gameObjects); i++ {
		objects[i] = marshalGameObject(gameObjects[i])
	}

	return objects
}

func marshalCamera(camera rendering.Camera) map[string]interface{} {
	return map[string]interface{}{
		"position": marshalVector2(camera.Position),
		"rotation": camera.Rotation.Radians(),
		"scale":    marshalVector2(camera.Scale),

		"width":  camera.PixelWidth,
		"height": camera.PixelHeight,
		"ppu":    camera.PixelsPerUnit,
	}
}
