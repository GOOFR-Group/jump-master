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

/**
 * Represents the game world.
 */
class GameWorld {
	#ctx: CanvasRenderingContext2D;

	#engine: Engine;

	#animator: ImageBySource;

	/**
	 * Initializes a game world.
	 * @param ctx Canvas 2D context.
	 * @param engine Game engine.
	 */
	constructor(
		ctx: CanvasRenderingContext2D,
		engine: Engine,
		animator: ImageBySource,
	) {
		this.#ctx = ctx;
		this.#engine = engine;
		this.#animator = animator;
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

		gameObjects.sort(
			(a, b) =>
				GameObjectTagOrder.indexOf(a.tag as GameObjectTag) -
				GameObjectTagOrder.indexOf(b.tag as GameObjectTag),
		);

		for (const gameObject of gameObjects) {
			const { transform, renderer, tag } = gameObject;

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
}

export default GameWorld;
