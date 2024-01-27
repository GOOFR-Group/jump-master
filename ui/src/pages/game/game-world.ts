import type { Actions } from '../../domain/actions';
import type { Engine } from '../../domain/engine';
import type {
	Camera,
	GameObject,
	Renderer,
	Transform,
} from '../../domain/game-state';

class GameWorldDebugTools {
	#ctx: CanvasRenderingContext2D;

	/**
	 * Initializes game world debug tools.
	 * @param ctx Canvas 2D context.
	 */
	constructor(ctx: CanvasRenderingContext2D) {
		this.#ctx = ctx;
	}

	draw(transform: Transform, renderer: Renderer) {
		this.#ctx.fillStyle = 'red';
		this.#ctx.font = 'bold 1rem monospace';
		this.#ctx.rotate(transform.rotation);
		this.#ctx.scale(transform.scale.x, transform.scale.y);

		const objectInfo = [
			`H: ${renderer.height} W: ${renderer.width}`,
			`X: ${transform.position.x.toFixed(3)} Y: ${transform.position.y.toFixed(
				3,
			)}`,
			`Sx: ${transform.scale.x} Sy: ${transform.scale.y}`,
			`R: ${(transform.rotation * (180 / Math.PI)).toFixed(3)}`,
		];
		const lineHeight = 16;

		// Draw game object info
		for (let i = 0; i < objectInfo.length; i++) {
			this.#ctx.fillText(
				objectInfo[i],
				renderer.offset.x,
				renderer.offset.y - (objectInfo.length - 1 - i) * lineHeight,
			);
		}
	}
}

/**
 * Represents the game world.
 */
class GameWorld {
	#ctx: CanvasRenderingContext2D;

	#engine: Engine;

	#debugTools: GameWorldDebugTools;

	/**
	 * Initializes a game world.
	 * @param ctx Canvas 2D context.
	 * @param engine Game engine.
	 */
	constructor(ctx: CanvasRenderingContext2D, engine: Engine) {
		this.#ctx = ctx;
		this.#engine = engine;
		this.#debugTools = new GameWorldDebugTools(this.#ctx);
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
			const { transform, renderer } = gameObject;

			if (renderer) {
				this.#ctx.beginPath();
				this.#ctx.save();
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

				if (gameObject.tag === 'Player') {
					this.#debugTools.draw(transform, renderer);
				}

				this.#ctx.restore();
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
