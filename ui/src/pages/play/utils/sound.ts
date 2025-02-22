import { GameObjectSound } from '../../../domain/game-state';
import type { SoundByName } from '../../../domain/sound';

/**
 * Sources of the sounds used in the game world.
 */
const SOUND_SOURCES: SoundByName = {
	[GameObjectSound.JUMP]: new Audio('sounds/jump.ogg'),
	[GameObjectSound.JUMP_HOLD]: new Audio('sounds/jumpHold.ogg'),
	[GameObjectSound.KNOCK_BACK]: new Audio('sounds/knockBack.ogg'),
	[GameObjectSound.LANDING]: new Audio('sounds/jump.ogg'),
	[GameObjectSound.FALL]: new Audio('sounds/fall.wav'),
};

/**
 * Loads the game sounds.
 * @returns Sounds by name.
 */
export function loadSounds() {
	return SOUND_SOURCES;
}

/**
 * Plays the given sound.
 * @param audio Audio element.
 */
export async function playSound(audio: HTMLAudioElement) {
	audio.currentTime = 0;
	try {
		await audio.play();
	} catch {
		// Ignore errors.
	}
}
