import { GameObjectSound } from '../../../domain/game-state';

/**
 * Sources of the sounds used in the game world.
 */
export const SOUND_SOURCES = {
	[GameObjectSound.JUMP]: new Audio('sounds/jump.ogg'),
	[GameObjectSound.KNOCK_BACK]: new Audio('sounds/knockBack.ogg'),
	[GameObjectSound.LANDING]: new Audio('sounds/landing.ogg'),
	[GameObjectSound.FALL]: new Audio('sounds/fall.ogg'),
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
