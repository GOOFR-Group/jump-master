/**
 * Loads an image.
 * @param src Image source.
 * @returns Image.
 */
export function loadImage(src: string) {
	const image = new Image();

	return new Promise<HTMLImageElement>(res => {
		image.onload = () => res(image);
		image.src = src;
	});
}
