package behaviour

import (
	"github.com/goofr-group/game-engine/pkg/engine"
	"github.com/goofr-group/physics-engine/pkg/game"

	"github.com/goofr-group/jump-master/engine/internal/game/property"
)

// SoundController defines the structure of the sound controller behaviour.
type SoundController struct {
	playerObject *game.Object        // Defines the player object.
	playerSounds map[string]struct{} // Defines the current map of player sounds.
}

// NewSoundController returns a new sound controller behaviour.
func NewSoundController(
	playerObject *game.Object,
) SoundController {
	return SoundController{
		playerObject: playerObject,
		playerSounds: make(map[string]struct{}),
	}
}

func (b SoundController) Enabled() bool {
	return true
}

func (b *SoundController) Update(_ *engine.Engine) error {
	// Set the sounds property with the list of sounds for the current frame.
	b.playerObject.SetProperty(property.Sounds, b.PlayerSounds())

	// Reset the list of sounds for the next frame.
	b.playerSounds = map[string]struct{}{}

	return nil
}

// PlayerSounds returns the list of sounds associated with the player to be played in the current frame.
func (b SoundController) PlayerSounds() []string {
	sounds := make([]string, 0, len(b.playerSounds))

	for sound := range b.playerSounds {
		sounds = append(sounds, sound)
	}

	return sounds
}

// AddPlayerSound adds the given player sound to the list.
// The list of player sounds is cleared each frame.
func (b *SoundController) AddPlayerSound(sound string) {
	b.playerSounds[sound] = struct{}{}
}
