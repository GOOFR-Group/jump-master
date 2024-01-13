/**
 * Represents the information about the latest engine build.
 */
export interface Version {
	/**
	 * Git tag.
	 */
	tag: string | null;

	/**
	 * Git commit.
	 */
	commit: string | null;

	/**
	 * Go version.
	 */
	go: string | null;

	/**
	 * Timestamp of build.
	 */
	build: string | null;
}
