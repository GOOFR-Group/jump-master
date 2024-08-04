package prefab

import (
	"fmt"

	"github.com/goofr-group/game-engine/pkg/rendering"
	"github.com/goofr-group/go-math/rotation/matrix"
	"github.com/goofr-group/go-math/vector2"
	core "github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/config"
	"github.com/goofr-group/jump-master/engine/internal/game"
	"github.com/goofr-group/jump-master/engine/internal/game/property"
	"github.com/goofr-group/jump-master/engine/internal/game/tag"
)

// NewMap creates all the objects in the map for the given configuration.
func NewMap(e game.Engine, config config.Map, tileSprites map[string]string) error {
	gameEngine := e.Engine()

	// Define the grid configuration.
	grid := vector2.Vector2{
		X: float64(config.TileSize),
		Y: float64(config.TileSize),
	}

	for _, layer := range config.Layers {
		for _, tile := range layer.Tiles {
			// Create the grid game object.
			gameObject := core.Object{
				Active: true,
				Tag:    tag.Platform,
				Transform: core.Transform2D{
					Position: grid.Scale(vector2.Vector2{
						X: float64(tile.X),
						Y: float64(tile.Y),
					}),
					Rotation: matrix.Identity(),
					Scale:    vector2.One(),
				},
				RigidBody: &core.RigidBody2D{
					BodyType:           core.BodyStatic,
					CollisionDetection: core.DiscreteDetection,
					Interpolation:      core.NoneInterpolation,
				},
				Renderer: &core.Renderer{
					Width:  grid.X,
					Height: grid.Y,
					Offset: vector2.Vector2{
						X: -grid.X / 2,
						Y: -grid.Y / 2,
					},
					Layer: rendering.DefaultRenderLayer,
				},
			}

			// Set the image path of the object.
			image := tileSprites[tile.ID]
			gameObject.SetProperty(property.Image, image)

			// Check if the object needs a collider.
			if layer.Collider {
				collider := core.NewBoxCollider(grid, vector2.Vector2{
					X: -grid.X / 2,
					Y: -grid.Y / 2,
				})
				gameObject.Collider = &collider
			}

			// Add the grid object to the game engine.
			err := gameEngine.CreateGameObject(&gameObject, nil)
			if err != nil {
				return fmt.Errorf("failed to create grid game object: %w", err)
			}
		}
	}

	return nil
}
