using UnityEngine;
using UnityEngine.Audio;

public class AudioManager : MonoBehaviour
{
    [Header("Required Components")]

    [SerializeField] private AudioMixer audioMixerSfx;

    /// <summary>
    /// The volume level that indicates the audio is un-muted.
    /// /// This is typically set to 0 dB, which is the default volume level for audio mixers.
    /// </summary>
    private const float UN_MUTE_VOLUME = 0;

    /// <summary>
    /// The volume level that indicates the audio is muted.
    /// This is typically set to -80 dB, which is a common value for audio mixers to represent silence.
    /// </summary>
    private const float MUTE_VOLUME = -80;

    // Audio mixer.
    private const string AUDIO_MIXER_MASTER_VOLUME = "Volume";

    /// <summary>
    /// Checks if the audio is muted by checking the audio mixer volume.
    /// </summary>
    /// <returns>Returns true if the volume is -80, otherwise false.</returns>
    public bool IsAudioMuted()
    {
        return audioMixerSfx.GetFloat(AUDIO_MIXER_MASTER_VOLUME, out float volume) && volume == MUTE_VOLUME;
    }

    /// <summary>
    /// Toggles the audio mixer volume between muted and un-muted states.
    /// </summary>
    public void ToggleAudioMixerVolume()
    {
        if (IsAudioMuted())
        {
            audioMixerSfx.SetFloat(AUDIO_MIXER_MASTER_VOLUME, UN_MUTE_VOLUME);
        }
        else
        {
            audioMixerSfx.SetFloat(AUDIO_MIXER_MASTER_VOLUME, MUTE_VOLUME);
        }
    }
}
