import useActions from '../play/use-actions';

function Controls() {
	const actions = useActions();

	return (
		<div class="flex flex-col items-center gap-8">
			<h2 class="drop-shadow-text text-4xl text-white">Controls</h2>
			<div class="flex items-center justify-center gap-2">
				<div class="relative">
					<div
						class="drop-shadow-text absolute bottom-full left-1/2 -translate-x-1/2 text-white opacity-0 transition-all"
						classList={{ 'opacity-100 -translate-y-1': actions().Jump }}
					>
						Jump
					</div>
					<img
						src={`images/controls/spacebar/${
							actions().Jump ? 'pressed' : 'default'
						}.png`}
					/>
				</div>

				<div class="relative">
					<div
						class="drop-shadow-text absolute bottom-full left-1/2 -translate-x-1/2 text-white opacity-0 transition-all"
						classList={{ 'opacity-100 -translate-y-1': actions().Left }}
					>
						Left
					</div>
					<img
						src={`images/controls/arrow-left/${
							actions().Left ? 'pressed' : 'default'
						}.png`}
					/>
				</div>

				<div class="relative">
					<div
						class="drop-shadow-text absolute bottom-full left-1/2 -translate-x-1/2 text-white opacity-0 transition-all"
						classList={{ 'opacity-100 -translate-y-1': actions().Right }}
					>
						Right
					</div>
					<img
						src={`images/controls/arrow-right/${
							actions().Right ? 'pressed' : 'default'
						}.png`}
					/>
				</div>
			</div>
		</div>
	);
}

export default Controls;
