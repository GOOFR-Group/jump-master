import { createResource, onCleanup, onMount } from 'solid-js';
import { loadEngine } from '../../utils/engine';
import type { Engine } from '../../domain/engine';
import GameWorld from './game-world';
import type { Actions } from '../../domain/actions';
import useActions from './use-actions';
import type { ImageBySource } from '../../domain/image';
import { loadAnimator } from './utils/animator';
import { loadSounds } from './utils/sound';
import type { SoundByName } from '../../domain/sound';

function Canvas({
	engine,
	animator,
	sounds,
}: {
	engine: Engine;
	animator: ImageBySource;
	sounds: SoundByName;
}) {
	let canvas!: HTMLCanvasElement;
	const actions = useActions();

	onMount(() => {
		const ctx = canvas.getContext('2d');

		if (!ctx) {
			console.error('Could not retrieve canvas 2d context');
			return;
		}

		const gameWorld = new GameWorld(ctx, engine, animator, sounds);

		let frame = requestAnimationFrame(step);

		function step() {
			frame = requestAnimationFrame(step);
			gameWorld.step(Object.entries(actions()) as Actions);
		}

		onCleanup(() => cancelAnimationFrame(frame));
	});

	return <canvas ref={canvas} class="border-2 border-blue-400" />;
}

function Play() {
	const [engine] = createResource(loadEngine);
	const [animator] = createResource(loadAnimator);

	return (
		<div class="flex h-full w-full flex-col items-center justify-center">
			<div class="absolute -z-10 h-full w-full bg-[url('/images/forest.jpg')] brightness-50" />
			{engine.state === 'pending' ||
				(animator.state === 'pending' && <p>Loading...</p>)}
			{engine.state === 'errored' ||
				(animator.state === 'errored' && <p>An unexpected error ocurred</p>)}
			{engine.state === 'ready' && animator.state === 'ready' && (
				<Canvas engine={engine()} animator={animator()} sounds={loadSounds()} />
			)}
		</div>
	);
}

export default Play;
