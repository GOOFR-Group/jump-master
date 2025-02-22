import { GameObjectSound } from '../../../domain/game-state';
import type { SoundByName } from '../../../domain/sound';

/**
 * Sources of the sounds used in the game world.
 */
const SOUND_SOURCES: Record<GameObjectSound, string> = {
	[GameObjectSound.JUMP]: 'sounds/jump.ogg',
	[GameObjectSound.JUMP_HOLD]: 'sounds/jumpHold.ogg',
	[GameObjectSound.KNOCK_BACK]: 'sounds/knockBack.ogg',
	[GameObjectSound.LANDING]: 'sounds/jump.ogg',
	[GameObjectSound.FALL]: 'sounds/fall.wav',
};

/**
 * Plays the given sound.
 * @returns Sounds by name.
 */
export function loadSounds() {
	return Object.fromEntries(
		Object.entries(SOUND_SOURCES).map(([key, value]) => [
			key,
			new Audio(value),
		]),
	) as SoundByName;
}

/**
 * Plays the given sound.
 * @param audio Audio element.
 */
export function playSound(audio: HTMLAudioElement) {
	audio.currentTime = 0;
	audio.play();
}
