/**
 * Type of user actions.
 */
export type ActionType = 'Jump' | 'Left' | 'Right';

/**
 * Action triggered by the user.
 *
 * The first value of the tuple represents
 * the type of action triggered.
 *
 * The second value of the tuple represents
 * the current state of the action.
 */
type Action = [ActionType, boolean];

/**
 * Actions triggered by the user.
 */
export type Actions = Array<Action>;
