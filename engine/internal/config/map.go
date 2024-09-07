package config

// Tile defines the structure of the map tile configuration.
type Tile struct {
	ID string `json:"id"` // Defines the identifier of the tile.
	X  int    `json:"x"`  // Defines the position on the x-axis of the tile on the map.
	Y  int    `json:"y"`  // Defines the position on the y-axis of the tile on the map.
}

// Layer defines the structure of the map layer configuration.
type Layer struct {
	Name     string `json:"name"`     // Defines the name of the layer.
	Tiles    []Tile `json:"tiles"`    // Defines the tiles of the layer.
	Collider bool   `json:"collider"` // Defines if the layer can collide with other dynamic objects in the world.
}

// Map defines the structure of the map configuration.
type Map struct {
	TileSize int     `json:"tileSize"`  // Defines the size of each tile.
	Width    int     `json:"mapWidth"`  // Defines the width of the map.
	Height   int     `json:"mapHeight"` // Defines the height of the map.
	Layers   []Layer `json:"layers"`    // Defines the layers of the map.
}
