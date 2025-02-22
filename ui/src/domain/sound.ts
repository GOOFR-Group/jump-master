import { GameObjectSound } from './game-state';

/**
 * Represents a mapped type for sound elements.
 *
 * The key represents the sound name with its corresponding `HTMLAudioElement` as value.
 */
export type SoundByName = {
	[key in GameObjectSound]: HTMLAudioElement;
};
