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

func marshalVectors2(v []vector2.Vector2) []interface{} {
	vectors := make([]interface{}, len(v))
	for i := 0; i < len(v); i++ {
		vectors[i] = marshalVector2(v[i])
	}

	return vectors
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

func marshalRenderer(renderer *game.Renderer, image interface{}, flipHorizontally interface{}) map[string]interface{} {
	if renderer == nil {
		return nil
	}

	return map[string]interface{}{
		"width":            renderer.Width,
		"height":           renderer.Height,
		"offset":           marshalVector2(renderer.Offset),
		"layer":            renderer.Layer,
		"image":            image,
		"flipHorizontally": flipHorizontally,
	}
}

func marshalMaterial(material game.Material) map[string]interface{} {
	return map[string]interface{}{
		"elasticity": material.Elasticity,
		"friction":   material.Friction,
	}
}

func marshalCollider(collider *game.Collider2D) map[string]interface{} {
	if collider == nil {
		return nil
	}

	response := map[string]interface{}{
		"isTrigger": collider.IsTrigger,
		"layer":     collider.Layer,
		"material":  marshalMaterial(collider.Material),
		"density":   collider.Density,
	}

	switch collider.Type() {
	case game.CircleType:
		circleCollider := collider.Collider().(*game.CircleCollider2D)
		response["radius"] = circleCollider.Radius
		response["pivot"] = marshalVector2(circleCollider.Pivot)

	case game.EdgeType:
		edgeCollider := collider.Collider().(*game.EdgeCollider2D)
		response["points"] = marshalVectors2(edgeCollider.Points)
		response["offset"] = marshalVector2(edgeCollider.Offset)

	case game.PolygonType:
		polygonCollider := collider.Collider().(*game.PolygonCollider2D)
		response["points"] = marshalVectors2(polygonCollider.Points)
		response["offset"] = marshalVector2(polygonCollider.Offset)
	}

	return response
}

func marshalGameObject(gameObject game.Object) map[string]interface{} {
	return map[string]interface{}{
		"id": gameObject.ID(),

		"active": gameObject.Active,
		"tag":    gameObject.Tag,

		"transform": marshalTransform(gameObject.Transform),
		"rigidBody": marshalRigidBody(gameObject.RigidBody),
		"renderer":  marshalRenderer(gameObject.Renderer, gameObject.Property(property.Image), gameObject.Property(property.FlipHorizontally)),
		"collider":  marshalCollider(gameObject.Collider),
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
