# Sprite Fusion

The map of the game is designed using the [Sprite Fusion editor](https://www.spritefusion.com/editor).  
The project can be found in this directory: [spritefusion.json](/engine/configs/spritefusion/spritefusion.json).  
The map is exported in JSON format and can be found in the engine configuration directory: [map.json](/engine/configs/map.json).

## Note 

The project contains a layer called "Tileset" which must always be the first in the hierarchy. This allows the exported map to generate deterministic tile IDs, which in turn allows these IDs to be mapped to their actual sprite path. The map can be found in the `tileSprites` property of the [engine configuration](/engine/configs/engine.json).