import { A } from '@solidjs/router';
import { createSignal } from 'solid-js';

function GameMenu() {
	const [startPressed, setStartPressed] = createSignal(false);

	return (
		<ul>
			<li class="flex items-center justify-center">
				<A href="/play">
					<img
						src={`/images/controls/start/${
							startPressed() ? 'pressed' : 'default'
						}.png`}
						onMouseEnter={() => setStartPressed(true)}
						onMouseLeave={() => setStartPressed(false)}
					/>
				</A>
			</li>
		</ul>
	);
}

export default GameMenu;
