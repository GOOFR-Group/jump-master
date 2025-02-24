import { createSignal, onCleanup, onMount } from 'solid-js';
import type { ActionType } from '../../domain/actions';

/**
 * Handles the keyboard events to determine the game actions to perform.
 * @returns Actions signal.
 */
function useActions() {
	const [actions, setActions] = createSignal<Record<ActionType, boolean>>({
		Left: false,
		Right: false,
		Jump: false,
	});

	/**
	 * Handles keydown event and updates the signal.
	 * @param e Keydown event.
	 */
	function handleKeyDown(e: KeyboardEvent) {
		switch (e.code) {
			case 'ArrowLeft':
			case 'KeyA':
				setActions(prev => ({
					...prev,
					Left: true,
				}));
				break;
			case 'ArrowRight':
			case 'KeyD':
				setActions(prev => ({
					...prev,
					Right: true,
				}));
				break;
			case 'Space':
				setActions(prev => ({
					...prev,
					Jump: true,
				}));
				break;
		}
	}

	/**
	 * Handles keyup event and updates the signal.
	 * @param e Keyup event.
	 */
	function handleKeyUp(e: KeyboardEvent) {
		switch (e.code) {
			case 'ArrowLeft':
			case 'KeyA':
				setActions(prev => ({
					...prev,
					Left: false,
				}));
				break;
			case 'ArrowRight':
			case 'KeyD':
				setActions(prev => ({
					...prev,
					Right: false,
				}));
				break;
			case 'Space':
				setActions(prev => ({
					...prev,
					Jump: false,
				}));
				break;
		}
	}

	onMount(() => {
		window.addEventListener('keydown', handleKeyDown);
		window.addEventListener('keyup', handleKeyUp);

		onCleanup(() => {
			window.removeEventListener('keydown', handleKeyDown);
			window.removeEventListener('keyup', handleKeyUp);
		});
	});

	return actions;
}

export default useActions;
