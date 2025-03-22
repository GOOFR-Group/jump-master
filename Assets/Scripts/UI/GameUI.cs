using UnityEngine;
using UnityEngine.Audio;
using UnityEngine.InputSystem;
using UnityEngine.UIElements;

public class GameUI : MonoBehaviour
{
    [Header("Required Components")]

    [SerializeField] private AudioMixer audioMixerSfx;

    // Input actions.
    private InputAction pauseAction;

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

        pauseAction = InputSystem.actions.FindActionMap(GameManager.INPUT_ACTION_MAP_MENU).FindAction(GameManager.INPUT_ACTION_MENU_PAUSE);

        pausedPanel = uiDocument.rootVisualElement.Q<VisualElement>(PAUSED_PANEL_NAME);

        pauseResumeButton = uiDocument.rootVisualElement.Q<Button>(PAUSE_RESUME_BUTTON_NAME);
        soundButton = uiDocument.rootVisualElement.Q<Button>(SOUND_BUTTON_NAME);
    }

    private void OnEnable()
    {
        pauseAction.started += TogglePauseResumeCallback;
        pauseResumeButton.RegisterCallback<ClickEvent>(OnPauseResumeButtonClick);
        soundButton.RegisterCallback<ClickEvent>(OnSoundButtonClick);
    }

    private void OnDisable()
    {
        pauseAction.started -= TogglePauseResumeCallback;
        pauseResumeButton.UnregisterCallback<ClickEvent>(OnPauseResumeButtonClick);
        soundButton.UnregisterCallback<ClickEvent>(OnSoundButtonClick);
    }

    /// <summary>
    /// Pauses and resumes the game.
    /// </summary>
    private void TogglePauseResume()
    {
        pauseResumeButton.ToggleInClassList(PAUSE_RESUME_PAUSED_CLASS_NAME);
        pausedPanel.ToggleInClassList(PAUSED_PANEL_REVEAL_CLASS_NAME);
        GameManager.ToggleTimeScale();
    }

    /// <summary>
    /// Pauses and resumes the game when the pause/resume action button is clicked.
    /// </summary>
    /// <param name="context">Callback context.</param>
    private void TogglePauseResumeCallback(InputAction.CallbackContext context)
    {
        TogglePauseResume();
    }

    /// <summary>
    /// Pauses and resumes the game when the pause/resume UI button is clicked.
    /// </summary>
    /// <param name="clickEvent">Click event.</param>
    private void OnPauseResumeButtonClick(ClickEvent clickEvent)
    {
        TogglePauseResume();
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
