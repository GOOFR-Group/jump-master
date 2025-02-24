import { useNavigate } from '@solidjs/router';
import { createSignal } from 'solid-js';

function GameMenu() {
	const [startPressed, setStartPressed] = createSignal(false);
	const navigate = useNavigate();

	return (
		<ul>
			<li class="flex items-center justify-center">
				<button onClick={() => navigate('/play')}>
					<img
						src={`images/controls/start/${
							startPressed() ? 'pressed' : 'default'
						}.png`}
						onMouseEnter={() => setStartPressed(true)}
						onMouseLeave={() => setStartPressed(false)}
					/>
				</button>
			</li>
		</ul>
	);
}

export default GameMenu;
