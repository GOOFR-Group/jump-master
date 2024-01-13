import type { Actions } from './actions';
import type { GameState } from './game-state';
import type { Version } from './version';

/**
 * Represents the type for `window.engine`.
 */
export interface Engine {
	/**
	 * Retrieves the information of the latest engine build.
	 *
	 * @returns Information of the latest engine build.
	 */
	version(): Version;

	/**
	 * Runs the next frame of the game and returns the game state.
	 *
	 * @param actions User actions in the game.
	 * @returns Game state.
	 */
	step(actions: Actions): GameState;
}
