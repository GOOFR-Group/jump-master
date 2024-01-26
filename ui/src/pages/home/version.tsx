import type { Version } from '../../domain/version';

function EngineVersion({ engineVersion }: { engineVersion: Version }) {
	return (
		<div class="flex items-end justify-center">
			<pre class="rounded-xl border border-neutral-950 bg-neutral-900 p-3 text-white shadow-xl">
				{JSON.stringify(engineVersion, null, 2)}
			</pre>
		</div>
	);
}

export default EngineVersion;
