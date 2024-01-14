import { createResource } from 'solid-js';
import { A } from '@solidjs/router';
import { loadEngine } from '../utils/engine';
import type { Version } from '../modules/version';

function GameMenu() {
	return (
		<ul>
			<li>
				<A href="/game">
					<button class="btn mx-auto">Start</button>
				</A>
			</li>
		</ul>
	);
}

function EngineVersion({ engineVersion }: { engineVersion: Version }) {
	return (
		<div class="flex items-end justify-center">
			<pre class="rounded-xl border border-neutral-950 bg-neutral-900 p-3 text-white shadow-xl">
				{JSON.stringify(engineVersion, null, 2)}
			</pre>
		</div>
	);
}

function Home() {
	const [engine] = createResource(loadEngine);

	return (
		<div class="flex h-full flex-col items-center">
			<div class="absolute -z-10 h-full w-full bg-[url('/images/forest.jpg')] brightness-50" />
			<div class="relative flex h-full flex-col justify-center gap-8 py-16">
				<div class="flex flex-1 flex-col justify-center gap-8">
					<h1 class="text-8xl font-semibold text-white">Jump Master</h1>

					{engine.state === 'pending' && <p>Loading...</p>}
					{engine.state === 'errored' && <p>An unexpected error ocurred</p>}
					{engine.state === 'ready' && <GameMenu />}
				</div>

				{engine.state === 'ready' && (
					<EngineVersion engineVersion={engine().version()} />
				)}
			</div>
		</div>
	);
}

export default Home;
