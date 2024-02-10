package config

import "github.com/goofr-group/go-math/vector2"

// Object defines the structure of the object configuration.
type Object struct {
	Position       vector2.Vector2 `json:"position"`       // Defines the position of the object.
	ColliderSize   vector2.Vector2 `json:"colliderSize"`   // Defines the collider size of the object.
	ColliderOffset vector2.Vector2 `json:"colliderOffset"` // Defines the collider offset of the object.
	RendererSize   vector2.Vector2 `json:"rendererSize"`   // Defines the renderer size of the object.
	Mass           float64         `json:"mass"`           // Defines the mass of the object.
	Drag           float64         `json:"drag"`           // Defines the drag of the object.
}

// Movement defines the structure of the movement configuration.
type Movement struct {
	Speed float64 `json:"speed"` // Defines the movement speed.
}

// Jump defines the structure of the jump configuration.
type Jump struct {
	Impulse       float64 `json:"impulse"`       // Defines the base impulse of the jump to accumulate each second the jump action is performed.
	MinImpulse    float64 `json:"minImpulse"`    // Defines the minimum impulse of the jump.
	MaxImpulse    float64 `json:"maxImpulse"`    // Defines the maximum impulse of the jump.
	DiagonalAngle float64 `json:"diagonalAngle"` // Defines the angle in degrees to apply when jumping left or right.
}

// Animator defines the structure of the animator configuration.
type Animator struct {
	Repeat   bool     `json:"repeat"`   // Defines if the frames should loop.
	Duration float64  `json:"duration"` // Defines the duration in seconds of each frame.
	Frames   []string `json:"frames"`   // Defines the images to display per frame.
}

// Animations defines the type of the animations.
type Animations map[string]Animator

// Player defines the structure of the player configuration.
type Player struct {
	Object     Object     `json:"object"`     // Object configurations.
	Movement   Movement   `json:"movement"`   // Movement behaviour configurations.
	Jump       Jump       `json:"jump"`       // Jump behaviour configurations.
	Animations Animations `json:"animations"` // Animation configurations.
}
