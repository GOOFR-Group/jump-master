import { A } from '@solidjs/router';

function GameMenu() {
	return (
		<ul>
			<li>
				<A href="/play">
					<button class="btn mx-auto">Start</button>
				</A>
			</li>
		</ul>
	);
}

export default GameMenu;
