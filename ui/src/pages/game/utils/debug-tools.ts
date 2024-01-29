import type { GameObject, Point } from '../../../domain/game-state';

/**
 * Represents a set of utility functions for drawing information
 * about a game object on a 2D canvas.
 */
class DebugTools {
	static readonly #FILL_STYLE = 'red';
	static readonly #FONT = 'bold 1rem monospace';
	static readonly #OFFSET = { x: 42, y: 0 };

	/**
	 * Renders multiline text.
	 * @param ctx Canvas 2D context.
	 * @param multilineText Text to render.
	 * @param offset Offset from the origin where the text is rendered.
	 * @param lineHeight Height of each line.
	 */
	static #renderMultilineText(
		ctx: CanvasRenderingContext2D,
		multilineText: string[],
		offset: Point = { x: 0, y: 0 },
		lineHeight = 16,
	) {
		// Draw game object info
		for (let i = 0; i < multilineText.length; i++) {
			ctx.fillText(
				multilineText[i],
				offset.x,
				offset.y - (multilineText.length - 1 - i) * lineHeight,
			);
		}
	}

	/**
	 * Draws debug information about a game object.
	 * @param ctx Canvas 2D context.
	 * @param gameObject Game object.
	 * @param debugPosition Position where the debug information will be rendered.
	 */
	static drawGameObjectInfo(
		ctx: CanvasRenderingContext2D,
		gameObject: GameObject,
		debugPosition: Point,
	) {
		ctx.beginPath();
		ctx.save();

		ctx.fillStyle = DebugTools.#FILL_STYLE;
		ctx.font = DebugTools.#FONT;

		const { transform, rigidBody, renderer } = gameObject;

		ctx.translate(
			debugPosition.x + DebugTools.#OFFSET.x,
			debugPosition.y + DebugTools.#OFFSET.y,
		);
		ctx.rotate(transform.rotation);

		const info = [
			`X: ${transform.position.x.toFixed(3)} Y: ${transform.position.y.toFixed(
				3,
			)}`,
			`Sx: ${transform.scale.x} Sy: ${transform.scale.y}`,
			`R: ${(transform.rotation * (180 / Math.PI)).toFixed(3)}`,
		];

		if (rigidBody) {
			const { mass, velocity, angularVelocity, drag, angularDrag } = rigidBody;
			info.push(
				`M: ${mass}`,
				`Vx: ${velocity.x} Vy: ${velocity.y}`,
				`AV: ${angularVelocity}`,
				`D: ${drag}`,
				`AD: ${angularDrag}`,
			);
		}

		if (renderer) {
			const { height, width } = renderer;
			info.push(`H: ${height} W: ${width}`);
		}

		DebugTools.#renderMultilineText(ctx, info);

		ctx.restore();
		ctx.closePath();
	}
}

export default DebugTools;
