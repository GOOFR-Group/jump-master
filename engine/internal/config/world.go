package config

import "github.com/goofr-group/go-math/vector2"

// Physics defines the structure of the physics configuration.
type Physics struct {
	UpdateRate float64         `json:"updateRate"` // Defines the physics update rate in seconds.
	Gravity    vector2.Vector2 `json:"gravity"`    // Defines the gravity of the game world.
}

// Camera defines the structure of the camera configuration.
type Camera struct {
	Width  float64 `json:"width"`  // Defines the width of the camera.
	Height float64 `json:"height"` // Defines the height of the camera.
	PPU    float64 `json:"ppu"`    // Defines pixels per game world unit.
}

// GridObject defines the structure of the grid object configuration.
type GridObject struct {
	Position   vector2.Vector2 `json:"position"`   // Defines the position of the object in the grid. The first element in blocks represents the origin. The others blocks are added from the left to right, top to bottom.
	CanCollide bool            `json:"canCollide"` // Defines if the object can collide with other dynamic objects.
	Blocks     [][]string      `json:"blocks"`     // Defines the matrix representation of each block in the object.
}

// Grid defines the structure of the grid configuration.
type Grid struct {
	Width   float64      `json:"width"`   // Defines the width of each block in the grid.
	Height  float64      `json:"height"`  // Defines the height of each block in the grid.
	Objects []GridObject `json:"objects"` // Defines the objects in the grid.
}

// World defines the structure of the world configuration.
type World struct {
	Physics Physics `json:"physics"` // Defines the physics of the game world.
	Camera  Camera  `json:"camera"`  // Defines the camera of the game world.
	Grid    Grid    `json:"grid"`    // Defines the grid of the game world.
}
