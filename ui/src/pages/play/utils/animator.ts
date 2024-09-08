import type { ImageBySource } from '../../../domain/image';
import { loadImage } from '../../../utils/image';

/**
 * Sources of the animation images used in the game world.
 */
const ANIMATION_IMAGE_SOURCES = [
	// Fence
	'/images/fence/0.png',
	'/images/fence/1.png',
	'/images/fence/2.png',
	'/images/fence/3.png',
	'/images/fence/4.png',
	'/images/fence/5.png',

	// Plant
	'/images/plant/0.png',
	'/images/plant/1.png',
	'/images/plant/2.png',
	'/images/plant/3.png',
	'/images/plant/4.png',

	// Platform bush
	'/images/platform/bush/0.png',
	'/images/platform/bush/1.png',
	'/images/platform/bush/2.png',
	'/images/platform/bush/3.png',
	'/images/platform/bush/4.png',
	'/images/platform/bush/5.png',

	// Platform forest dirt
	'/images/platform/forest/dirt/0.png',
	'/images/platform/forest/dirt/1.png',
	'/images/platform/forest/dirt/2.png',
	'/images/platform/forest/dirt/3.png',
	'/images/platform/forest/dirt/4.png',
	'/images/platform/forest/dirt/5.png',
	'/images/platform/forest/dirt/6.png',
	'/images/platform/forest/dirt/7.png',

	// Platform forest grass
	'/images/platform/forest/grass/0.png',
	'/images/platform/forest/grass/1.png',
	'/images/platform/forest/grass/2.png',
	'/images/platform/forest/grass/3.png',
	'/images/platform/forest/grass/4.png',
	'/images/platform/forest/grass/5.png',

	// Platform grass
	'/images/platform/grass/0.png',
	'/images/platform/grass/1.png',
	'/images/platform/grass/2.png',
	'/images/platform/grass/3.png',
	'/images/platform/grass/4.png',
	'/images/platform/grass/5.png',
	'/images/platform/grass/6.png',
	'/images/platform/grass/7.png',
	'/images/platform/grass/8.png',

	// Platform jungle dirt
	'/images/platform/jungle/dirt/0.png',
	'/images/platform/jungle/dirt/1.png',
	'/images/platform/jungle/dirt/2.png',
	'/images/platform/jungle/dirt/3.png',
	'/images/platform/jungle/dirt/4.png',
	'/images/platform/jungle/dirt/5.png',
	'/images/platform/jungle/dirt/6.png',

	// Platform jungle grass
	'/images/platform/jungle/grass/0.png',
	'/images/platform/jungle/grass/1.png',
	'/images/platform/jungle/grass/2.png',
	'/images/platform/jungle/grass/3.png',
	'/images/platform/jungle/grass/4.png',
	'/images/platform/jungle/grass/5.png',

	// Platform water
	'/images/platform/water/0.png',
	'/images/platform/water/1.png',

	// Player fall
	'/images/player/fall/0.png',
	'/images/player/fall/1.png',
	'/images/player/fall/2.png',
	'/images/player/fall/3.png',

	// Player idle
	'/images/player/idle/0.png',
	'/images/player/idle/1.png',

	// Player jump
	'/images/player/jump/0.png',

	// Player jump fall
	'/images/player/jump-fall/0.png',
	'/images/player/jump-fall/1.png',

	// Player jump hold
	'/images/player/jump-hold/0.png',

	// Player knock-back
	'/images/player/knock-back/0.png',

	// Player walk
	'/images/player/walk/0.png',
	'/images/player/walk/1.png',
	'/images/player/walk/2.png',
	'/images/player/walk/3.png',

	// Rock
	'/images/rock/0.png',
	'/images/rock/1.png',

	// Sign
	'/images/sign/0.png',
	'/images/sign/1.png',
	'/images/sign/2.png',
	'/images/sign/3.png',
	'/images/sign/4.png',
	'/images/sign/5.png',
];

/**
 * Loads the animator.
 * @returns Animator.
 */
export async function loadAnimator() {
	const animator: ImageBySource = {};
	const imagePromises = [];

	for (const imgSrc of ANIMATION_IMAGE_SOURCES) {
		imagePromises.push(loadImage(imgSrc));
	}

	const imageElements = await Promise.all(imagePromises);

	for (let i = 0; i < ANIMATION_IMAGE_SOURCES.length; i++) {
		const imgSrc = ANIMATION_IMAGE_SOURCES[i];
		animator[imgSrc] = imageElements[i];
	}

	return animator;
}
