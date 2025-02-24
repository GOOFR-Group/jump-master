import { createResource } from 'solid-js';
import { loadEngine } from '../../utils/engine';
import GameMenu from './menu';
import EngineVersion from './version';

function Home() {
	const [engine] = createResource(loadEngine);

	return (
		<div class="flex h-full flex-col items-center">
			<div class="absolute -z-10 h-full w-full bg-[url('/images/forest.jpg')] brightness-50" />
			<div class="relative flex h-full flex-col justify-center gap-8 py-16">
				<div class="flex flex-1 flex-col justify-center gap-8">
					<h1 class="text-8xl font-semibold text-white">Jump Master</h1>

					{engine.state === 'pending' && <p>Loading...</p>}
					{engine.state === 'errored' && <p>An unexpected error occurred</p>}
					{engine.state === 'ready' && <GameMenu />}
				</div>

				{import.meta.env.DEV && engine.state === 'ready' && (
					<EngineVersion engineVersion={engine().version()} />
				)}
			</div>
		</div>
	);
}

export default Home;
