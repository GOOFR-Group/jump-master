export type ImageFrame = `${number}.png`;

export type AnimationFrameImage = Record<ImageFrame, HTMLImageElement>;

export interface Animator {
	player: {
		idle: AnimationFrameImage;
		walk: AnimationFrameImage;
		jump: AnimationFrameImage;
		'jump-hold': AnimationFrameImage;
		'jump-fall': AnimationFrameImage;
	};
}

export type GameObjectAnimator = keyof Animator;

export type GameObjectAnimation = keyof Animator[keyof Animator];
