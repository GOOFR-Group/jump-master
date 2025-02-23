import type { Accessor } from 'solid-js';

function SoundButton({
	muted,
	onClick,
}: {
	muted: Accessor<boolean>;
	onClick: (e: MouseEvent & { currentTarget: HTMLButtonElement }) => void;
}) {
	/**
	 * Handles keydown event performed on the button.
	 * @param e Keydown event.
	 */
	function handleKeyDown(
		e: KeyboardEvent & { currentTarget: HTMLButtonElement },
	) {
		e.preventDefault();
		e.currentTarget.blur();
	}

	return (
		<div class="flex items-center gap-2 text-4xl text-white drop-shadow-xl">
			Sound:
			<button onClick={onClick} onKeyDown={handleKeyDown}>
				<img src={`/images/controls/audio/${muted() ? 'off' : 'on'}.png`} />
			</button>
		</div>
	);
}

export default SoundButton;
