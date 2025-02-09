/**
 * Represents a 2D vector, point or position.
 */
export interface Point {
	x: number;
	y: number;
}

/**
 * Represents spacial information of a game object.
 */
export interface Transform {
	/**
	 * Position of the game object in world space.
	 */
	position: Point;

	/**
	 * Rotation of the game object in world space.
	 */
	rotation: number;

	/**
	 * Scale of the game object in world space.
	 */
	scale: Point;
}

/**
 * Represents the kinematics related data of a game object.
 */
export interface RigidBody {
	/**
	 * Defines the type of the rigid body.
	 */
	bodyType: number;

	/**
	 * Used by the physics engine to check if two objects have collided.
	 */
	collisionDetection: number;

	/**
	 * Used to estimate the position of the rigid body between physics updates.
	 */
	interpolation: number;

	/**
	 * Indicates whether to calculate the mass from the Collider2D's density and area.
	 */
	autoMass: boolean;

	/**
	 * Defines the mass of a RigidBody.
	 */
	mass: number;

	/**
	 * Defines the intensity of the force acting opposite to a RigidBody's linear motion.
	 */
	drag: number;

	/**
	 * Defines the linear velocity of a RigidBody.
	 */
	velocity: Point;

	/**
	 * Defines the intensity of the angular force acting opposite to a RigidBody's rotational motion.
	 */
	angularDrag: number;

	/**
	 * Defines the angular velocity of a RigidBody.
	 */
	angularVelocity: number;

	/**
	 * Defines how much gravity affects a RigidBody.
	 */
	gravityScale: number;
}

/**
 * Represents simple rendering information of a game object.
 */
export interface Renderer {
	/**
	 * Layer where the game object is rendered.
	 */
	layer: string;

	/**
	 * Width of the game object render in pixels.
	 */
	width: number;

	/**
	 * Height of the game object render in pixels.
	 */
	height: number;

	/**
	 * Offset of the game object from its transform position.
	 */
	offset: Point;

	/**
	 * Image asset of the game object.
	 */
	image: string | null;

	/**
	 * Indicates whether the player should be flipped horizontally or not.
	 */
	flipHorizontally: boolean | null;
}

/**
 * Represents all properties defined in object, as well as any dynamic properties.
 */
export interface GameObject {
	/**
	 * Identifier of this game object.
	 * No game object can share the same id in the same game world.
	 */
	id: number;

	/**
	 * Active determines if this game object is active.
	 * A deactivated game object will not be physically simulated
	 * or interact with other game objects.
	 */
	active: boolean;

	/**
	 * Determines the tag or name of this game object.
	 * Multiple game objects can share the same tag.
	 */
	tag: string;

	/**
	 * Spacial information of a game object.
	 */
	transform: Transform;

	/**
	 * Kinematic information of a game object.
	 */
	rigidBody: RigidBody | null;

	/**
	 * Simple rendering information of a game object.
	 */
	renderer: Renderer | null;
}

/**
 * Represents a rendering viewport.
 * Given a viewport dimensions, it assists in transforming a game object's
 * world coordinates to screen coordinates.
 */
export interface Camera {
	/**
	 * Camera's world position.
	 */
	position: Point;

	/**
	 * Camera's world rotation.
	 */
	rotation: number;

	/**
	 * Camera's scale that's going to be applied in transformations.
	 */
	scale: Point;

	/**
	 * Width of the camera in pixels.
	 */
	width: number;

	/**
	 * Height of the camera in pixels.
	 */
	height: number;

	/**
	 * Amount of pixels to which an in-world unit maps.
	 */
	ppu: number;
}

/**
 * Represents the state of the game.
 * Includes the camera and the game objects in the world.
 *
 * If an error occurs, `error` will contain an error message.
 */
export interface GameState {
	/**
	 * Error message.
	 */
	error: string | null;

	/**
	 * Game objects in the world.
	 */
	gameObjects: GameObject[];

	/**
	 * Viewpoint through which the player views the game world.
	 */
	camera: Camera;
}
