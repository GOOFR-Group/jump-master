import { createResource } from 'solid-js';
import { loadEngine } from '../utils/engine';

function Game() {
	const [engine] = createResource(loadEngine);

	return (
		<div class="flex h-full w-full flex-col items-center justify-center">
			<div class="absolute -z-10 h-full w-full bg-[url('/images/forest.jpg')] brightness-50" />
			{engine.state === 'pending' && <p>Loading...</p>}
			{engine.state === 'errored' && <p>An unexpected error ocurred</p>}
			{engine.state === 'ready' && <p>Loaded engine</p>}
		</div>
	);
}

export default Game;
