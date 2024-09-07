package config

import "github.com/goofr-group/go-math/vector2"

// Physics defines the structure of the physics configuration.
type Physics struct {
	UpdateRate float64         `json:"updateRate"` // Defines the physics update rate in seconds.
	Gravity    vector2.Vector2 `json:"gravity"`    // Defines the gravity of the game world.
}

// Camera defines the structure of the camera configuration.
type Camera struct {
	Position vector2.Vector2 `json:"position"` // Defines the position of the camera.
	Width    float64         `json:"width"`    // Defines the width of the camera.
	Height   float64         `json:"height"`   // Defines the height of the camera.
	PPU      float64         `json:"ppu"`      // Defines pixels per game world unit.
}

// Engine defines the structure of the engine configuration.
type Engine struct {
	Physics     Physics           `json:"physics"`     // Defines the physics of the game engine.
	Camera      Camera            `json:"camera"`      // Defines the camera of the game engine.
	TileSprites map[string]string `json:"tileSprites"` // Defines the sprites of the map tileset per tile id.
}
