import type { ImageBySource } from '../../../domain/image';
import { loadImage } from '../../../utils/image';

/**
 * Sources of the animation images used in the game world.
 */
const ANIMATION_IMAGE_SOURCES = {
	PLAYER: {
		IDLE: {
			0: '/animations/player/idle/0.png',
			1: '/animations/player/idle/1.png',
		},
		WALK: {
			0: '/animations/player/walk/0.png',
			1: '/animations/player/walk/1.png',
			2: '/animations/player/walk/2.png',
			3: '/animations/player/walk/3.png',
		},
		JUMP_HOLD: {
			0: '/animations/player/jump-hold/0.png',
		},
		JUMP: {
			0: '/animations/player/jump/0.png',
			1: '/animations/player/jump/1.png',
		},
		JUMP_FALL: {
			0: '/animations/player/jump-fall/0.png',
			1: '/animations/player/jump-fall/1.png',
			2: '/animations/player/jump-fall/2.png',
			3: '/animations/player/jump-fall/3.png',
		},
	},
} as const;

/**
 * Loads the player animator.
 *
 * Pre-loads the images for the different animations of the player.
 *
 * @returns Player animator.
 */
export async function loadPlayerAnimator(): Promise<ImageBySource> {
	return {
		// Load idle images
		[ANIMATION_IMAGE_SOURCES.PLAYER.IDLE[0]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.IDLE[0],
		),
		[ANIMATION_IMAGE_SOURCES.PLAYER.IDLE[1]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.IDLE[1],
		),

		// Load walk images
		[ANIMATION_IMAGE_SOURCES.PLAYER.WALK[0]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.WALK[0],
		),
		[ANIMATION_IMAGE_SOURCES.PLAYER.WALK[1]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.WALK[1],
		),
		[ANIMATION_IMAGE_SOURCES.PLAYER.WALK[2]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.WALK[2],
		),
		[ANIMATION_IMAGE_SOURCES.PLAYER.WALK[3]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.WALK[3],
		),

		// Load jump hold image
		[ANIMATION_IMAGE_SOURCES.PLAYER.JUMP_HOLD[0]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.JUMP_HOLD[0],
		),

		// Load jump images
		[ANIMATION_IMAGE_SOURCES.PLAYER.JUMP[0]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.JUMP[0],
		),
		[ANIMATION_IMAGE_SOURCES.PLAYER.JUMP[1]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.JUMP[1],
		),

		// Load jump fall images
		[ANIMATION_IMAGE_SOURCES.PLAYER.JUMP_FALL[0]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.JUMP_FALL[0],
		),
		[ANIMATION_IMAGE_SOURCES.PLAYER.JUMP_FALL[1]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.JUMP_FALL[1],
		),
		[ANIMATION_IMAGE_SOURCES.PLAYER.JUMP_FALL[2]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.JUMP_FALL[2],
		),
		[ANIMATION_IMAGE_SOURCES.PLAYER.JUMP_FALL[3]]: await loadImage(
			ANIMATION_IMAGE_SOURCES.PLAYER.JUMP_FALL[3],
		),
	};
}

/**
 * Loads the animator.
 * @returns Animator.
 */
export async function loadAnimator() {
	const playerAnimator = await loadPlayerAnimator();

	return playerAnimator;
}