# Sprite Fusion

The map of the game is designed using the [Sprite Fusion editor](https://www.spritefusion.com/editor).  
The project can be found in this directory: [Jump_Master.json](/engine/configs/spritefusion/Jump_Master.json).  
The map is exported in JSON format and can be found in the engine configuration directory: [map.json](/engine/configs/map.json).

## Layers

### Props-Background

The objects in this layer are displayed behind the player object.

### Props-Foreground

The objects in this layer are displayed in front of the player object.

### Platforms

The objects in the layers that contain the substring `Platform` are used by the physics engine to collide with the player object.  

## Note 

The project contains a layer called "Tileset" which must always be the first in the hierarchy. This allows the exported map to generate deterministic tile IDs, which in turn allows these IDs to be mapped to their actual sprite path. The map can be found in the `tileSprites` property of the [engine configuration](/engine/configs/engine.json).