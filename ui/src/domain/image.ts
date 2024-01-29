/**
 * Represents a mapped type for image elements.
 *
 * The key represents the image path with its corresponding `HTMLImageElement` as value.
 */
export interface ImageBySource {
	[src: string]: HTMLImageElement;
}
