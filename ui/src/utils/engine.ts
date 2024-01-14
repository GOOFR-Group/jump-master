import type { Engine } from '../modules/engine';

declare class Go {
	importObject: WebAssembly.Imports;
	run(instance: WebAssembly.Instance): Promise<void>;
}

declare global {
	interface Window {
		engine: Engine;
	}
}

const WASM_FILE = 'engine.wasm';

/**
 * Loads WebAssembly binary of game engine.
 * @returns Game engine object.
 */
export async function loadEngine() {
	if (!WebAssembly.instantiateStreaming) {
		// polyfill
		WebAssembly.instantiateStreaming = async (resp, importObject) => {
			const source = await (await resp).arrayBuffer();
			return await WebAssembly.instantiate(source, importObject);
		};
	}

	const go = new Go();
	const module = await WebAssembly.instantiateStreaming(
		fetch(WASM_FILE),
		go.importObject,
	);

	go.run(module.instance);

	return window.engine;
}
