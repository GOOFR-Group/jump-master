import type { Actions } from '../../domain/actions';
import type { Engine } from '../../domain/engine';
import type { Camera, GameObject } from '../../domain/game-state';

/**
 * Represents the game world.
 */
class GameWorld {
	#ctx: CanvasRenderingContext2D;

	#engine: Engine;

	/**
	 * Initializes a game world.
	 * @param ctx Canvas 2D context.
	 * @param engine Game engine.
	 */
	constructor(ctx: CanvasRenderingContext2D, engine: Engine) {
		this.#ctx = ctx;
		this.#engine = engine;
	}

	/**
	 * Draws the object visible to the camera.
	 * @param gameObjects Game objects to draw.
	 * @param camera Game camera.
	 */
	#draw(gameObjects: GameObject[], camera: Camera) {
		this.#ctx.canvas.width = camera.width;
		this.#ctx.canvas.height = camera.height;

		for (const gameObject of gameObjects) {
			const transform = gameObject.transform;

			const renderer = gameObject.renderer;
			if (renderer) {
				this.#ctx.beginPath();
				this.#ctx.save()
				this.#ctx.translate(transform.position.x, transform.position.y);
				this.#ctx.scale(transform.scale.x, transform.scale.y);
				this.#ctx.rotate(transform.rotation);
				this.#ctx.rect(
					renderer.offset.x,
					renderer.offset.y,
					renderer.width,
					renderer.height,
				);
				this.#ctx.fillStyle = '#fbbf24';
				this.#ctx.fill();
				this.#ctx.restore()
				this.#ctx.closePath();
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
