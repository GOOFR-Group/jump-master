import type { Actions } from '../../domain/actions';
import type { ImageBySource } from '../../domain/image';
import type { Engine } from '../../domain/engine';
import {
	type Camera,
	type GameObject,
	type Point,
} from '../../domain/game-state';
import DebugTools from './utils/debug-tools';
import { GameObjectTag, GameObjectTagOrder } from '../../domain/tag';
import type { SoundByName } from '../../domain/sound';
import { playSound } from './utils/sound';

/**
 * Represents the game world.
 */
class GameWorld {
	#ctx: CanvasRenderingContext2D;

	#engine: Engine;

	#animator: ImageBySource;

	#sounds: SoundByName;

	#muted: boolean;

	/**
	 * Initializes a game world.
	 * @param ctx Canvas 2D context.
	 * @param engine Game engine.
	 * @param animator Game animator.
	 * @param sounds Game sounds.
	 * @param muted Indicates if game sounds are muted.
	 */
	constructor(
		ctx: CanvasRenderingContext2D,
		engine: Engine,
		animator: ImageBySource,
		sounds: SoundByName,
		muted: boolean,
	) {
		this.#ctx = ctx;
		this.#engine = engine;
		this.#animator = animator;
		this.#sounds = sounds;
		this.#muted = muted;
	}

	/**
	 * Draws an image at the position relative to the canvas context origin.
	 * @param src Image source.
	 * @param offset Image offset.
	 * @param width Image width.
	 * @param height Image height.
	 */
	#drawImage(src: string, offset: Point, width: number, height: number) {
		const img = this.#animator[src];
		this.#ctx.drawImage(img, offset.x, offset.y, width, height);
	}

	/**
	 * Draws the object visible to the camera.
	 * @param gameObjects Game objects to draw.
	 * @param camera Game camera.
	 */
	#draw(gameObjects: GameObject[], camera: Camera) {
		this.#ctx.canvas.width = camera.width;
		this.#ctx.canvas.height = camera.height;

		// Sort game objects based on the configured tag order.
		gameObjects.sort(
			(a, b) =>
				GameObjectTagOrder.indexOf(a.tag) - GameObjectTagOrder.indexOf(b.tag),
		);

		for (const gameObject of gameObjects) {
			const { transform, renderer, tag, sounds } = gameObject;

			if (!this.#muted) {
				for (const sound of sounds) {
					const audio = this.#sounds[sound];
					playSound(audio);
				}
			}

			if (!renderer) {
				continue;
			}

			this.#ctx.beginPath();
			this.#ctx.save();

			this.#ctx.translate(transform.position.x, transform.position.y);
			this.#ctx.scale(transform.scale.x, -transform.scale.y);
			this.#ctx.rotate(-transform.rotation);

			if (renderer.flipHorizontally) {
				this.#ctx.scale(-1, 1);
			}

			if (renderer.image) {
				this.#drawImage(
					renderer.image,
					renderer.offset,
					renderer.width,
					renderer.height,
				);
			} else {
				this.#ctx.fillStyle = '#fbbf24';
				this.#ctx.fillRect(
					renderer.offset.x,
					renderer.offset.y,
					renderer.width,
					renderer.height,
				);
			}

			this.#ctx.restore();
			this.#ctx.closePath();

			// Draw debug information of player object
			if (import.meta.env.DEV && tag === GameObjectTag.PLAYER) {
				DebugTools.drawGameObjectInfo(
					this.#ctx,
					gameObject,
					transform.position,
				);
			}
		}
	}

	/**
	 * Performs a game step for the given actions.
	 * @param actions Actions to perform.
	 */
	step(actions: Actions) {
		const { error, gameObjects, camera } = this.#engine.step(actions);

		if (error) {
			console.error(error);
			return;
		}

		this.#draw(gameObjects, camera);
	}

	/**
	 * Controls whether game sounds are played or not.
	 * @returns True if game sounds are muted, false otherwise.
	 */
	get muted() {
		return this.#muted;
	}

	/**
	 * Updates the game sound muted state.
	 * When muted is true, no sounds will be played.
	 * @param value True to mute all game sounds, false to enable them.
	 */
	set muted(value: boolean) {
		this.#muted = value;
	}
}

export default GameWorld;
