/**
 * Defines the game object tags.
 */
export enum GameObjectTag {
	PLAYER = 'Player',
	PLATFORM = 'Platform',
	PROPS_BACKGROUND = 'Props-Background',
	PROPS_FOREGROUND = 'Props-Foreground',
}

/**
 * Represents the order of game objects by tag to be drawn.
 */
export const GameObjectTagOrder = [
	GameObjectTag.PLATFORM,
	GameObjectTag.PROPS_BACKGROUND,
	GameObjectTag.PLAYER,
	GameObjectTag.PROPS_FOREGROUND,
];
