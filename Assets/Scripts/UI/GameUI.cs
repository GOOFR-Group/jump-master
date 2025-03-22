using UnityEngine;
using UnityEngine.Audio;
using UnityEngine.UIElements;

public class GameUI : MonoBehaviour
{
    [Header("Required Components")]

    [SerializeField] private AudioMixer audioMixerSfx;

    // UI elements.
    private VisualElement pausedPanel;
    private Button pauseResumeButton;
    private Button soundButton;

    // UI element names.
    private const string PAUSED_PANEL_NAME = "PausedPanel";
    private const string PAUSE_RESUME_BUTTON_NAME = "PauseResumeButton";
    private const string SOUND_BUTTON_NAME = "SoundButton";

    // UI class names.
    private const string PAUSED_PANEL_REVEAL_CLASS_NAME = "paused-panel-reveal";
    private const string PAUSE_RESUME_PAUSED_CLASS_NAME = "pause-resume-button-paused";
    private const string SOUND_MUTED_CLASS_NAME = "sound-button-muted";

    private void Awake()
    {
        UIDocument uiDocument = GetComponent<UIDocument>();

        pausedPanel = uiDocument.rootVisualElement.Q<VisualElement>(PAUSED_PANEL_NAME);

        pauseResumeButton = uiDocument.rootVisualElement.Q<Button>(PAUSE_RESUME_BUTTON_NAME);
        soundButton = uiDocument.rootVisualElement.Q<Button>(SOUND_BUTTON_NAME);
    }

    private void OnEnable()
    {
        pauseResumeButton.RegisterCallback<ClickEvent>(OnPauseResumeButtonClick);
        soundButton.RegisterCallback<ClickEvent>(OnSoundButtonClick);
    }

    private void OnDisable()
    {
        pauseResumeButton.UnregisterCallback<ClickEvent>(OnPauseResumeButtonClick);
        soundButton.UnregisterCallback<ClickEvent>(OnSoundButtonClick);
    }

    /// <summary>
    /// Pauses and resumes the game when the pause/resume button is clicked.
    /// </summary>
    /// <param name="clickEvent">Click event.</param>
    private void OnPauseResumeButtonClick(ClickEvent clickEvent)
    {
        pauseResumeButton.ToggleInClassList(PAUSE_RESUME_PAUSED_CLASS_NAME);
        pausedPanel.ToggleInClassList(PAUSED_PANEL_REVEAL_CLASS_NAME);
        GameManager.ToggleTimeScale();
    }

    /// <summary>
    /// Toggles game sounds on/off when the sound button is clicked.
    /// </summary>
    /// <param name="clickEvent">Click event.</param>
    private void OnSoundButtonClick(ClickEvent clickEvent)
    {
        soundButton.ToggleInClassList(SOUND_MUTED_CLASS_NAME);

        if (audioMixerSfx.GetFloat(GameManager.AUDIO_MIXER_MASTER_VOLUME, out float volume) && volume == 0)
        {
            audioMixerSfx.SetFloat(GameManager.AUDIO_MIXER_MASTER_VOLUME, -80);
        }
        else
        {
            audioMixerSfx.ClearFloat(GameManager.AUDIO_MIXER_MASTER_VOLUME);
        }
    }
}
