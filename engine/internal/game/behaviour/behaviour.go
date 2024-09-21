package behaviour

import "math"

const (
	// Epsilon defines the epsilon used in the behaviours.
	Epsilon = 0.1
)

// easeOutSine calculates the easing out sine function for a given amount t.
// The parameter t represents the absolute progress of the animation in the bounds of 0 (beginning of the animation) and
// 1 (end of animation).
//
// It is used to create a smooth animation transition.
func easeOutSine(t float64) float64 {
	return math.Sin(t * math.Pi / 2)
}
