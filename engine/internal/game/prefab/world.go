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

// NewGridObjects creates all the objects in grid for the given configuration.
func NewGridObjects(e game.Engine, config config.Grid) error {
	gameEngine := e.Engine()

	// Define the grid configuration.
	grid := vector2.Vector2{
		X: config.Width,
		Y: config.Height,
	}

	for _, gridObject := range config.Objects {
		gridObjectPosition := grid.Scale(gridObject.Position)

		for i, columns := range gridObject.Blocks {
			for j, image := range columns {
				// Ignore objects that do not contain an image.
				if len(image) == 0 {
					continue
				}

				// Compute block position relative to the first block and based on the grid dimensions.
				blockPosition := grid.Scale(vector2.Vector2{
					X: float64(j),
					Y: float64(i),
				})
				blockPosition = gridObjectPosition.Add(blockPosition)

				// Create the grid game object.
				gameObject := core.Object{
					Active: true,
					Tag:    tag.Platform,
					Transform: core.Transform2D{
						Position: blockPosition,
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
				gameObject.SetProperty(property.Image, image)

				// Check if the object needs a collider.
				if gridObject.CanCollide {
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
	}

	return nil
}
