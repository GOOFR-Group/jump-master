import type { Animator } from '../../../domain/animations';

function loadAnimation(src: string) {
	const image = new Image();

	return new Promise<HTMLImageElement>(res => {
		image.onload = e => {
			res(e.currentTarget as HTMLImageElement);
		};
		image.style.imageRendering = 'pixelated';
		image.src = src;
	});
}

export async function loadPlayerAnimator() {
	const playerAnimator: Animator['player'] = {
		idle: {},
		walk: {},
		jump: {},
		'jump-hold': {},
		'jump-fall': {},
	};

	playerAnimator.idle['0.png'] = await loadAnimation(
		'/animations/player/idle/0.png',
	);
	playerAnimator.idle['1.png'] = await loadAnimation(
		'/animations/player/idle/1.png',
	);

	playerAnimator.walk['0.png'] = await loadAnimation(
		'/animations/player/walk/0.png',
	);
	playerAnimator.walk['1.png'] = await loadAnimation(
		'/animations/player/walk/1.png',
	);
	playerAnimator.walk['2.png'] = await loadAnimation(
		'/animations/player/walk/2.png',
	);
	playerAnimator.walk['3.png'] = await loadAnimation(
		'/animations/player/walk/3.png',
	);

	playerAnimator['jump-hold']['0.png'] = await loadAnimation(
		'/animations/player/jump-hold/0.png',
	);

	playerAnimator.jump['0.png'] = await loadAnimation(
		'/animations/player/jump/0.png',
	);
	playerAnimator.jump['1.png'] = await loadAnimation(
		'/animations/player/jump/1.png',
	);

	playerAnimator['jump-fall']['0.png'] = await loadAnimation(
		'/animations/player/jump-fall/0.png',
	);
	playerAnimator['jump-fall']['1.png'] = await loadAnimation(
		'/animations/player/jump-fall/1.png',
	);
	playerAnimator['jump-fall']['2.png'] = await loadAnimation(
		'/animations/player/jump-fall/2.png',
	);
	playerAnimator['jump-fall']['3.png'] = await loadAnimation(
		'/animations/player/jump-fall/3.png',
	);

	return playerAnimator;
}

export async function loadAnimator() {
	const playerAnimator = await loadPlayerAnimator();

	return {
		player: playerAnimator,
	};
}
