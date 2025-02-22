import { GameObjectSound } from '../../../domain/game-state';

/**
 * Sources of the sounds used in the game world.
 */
export const SOUND_SOURCES = {
	[GameObjectSound.JUMP]: new Audio('sounds/jump.ogg'),
	[GameObjectSound.JUMP_HOLD]: new Audio('sounds/.ogg'),
	[GameObjectSound.KNOCK_BACK]: new Audio('sounds/.ogg'),
	[GameObjectSound.LANDING]: new Audio('sounds/landing.ogg'),
	[GameObjectSound.FALL]: new Audio('sounds/.wav'),
};

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
