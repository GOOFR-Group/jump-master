using UnityEngine;
using UnityEngine.InputSystem;
using UnityEngine.SceneManagement;
using UnityEngine.UIElements;

public class GameUI : MonoBehaviour
{
    [Header("Required Components")]

    [SerializeField] private AudioManager audioManager;
    [SerializeField] private TimeManager timeManager;

    // Input actions.
    private InputAction pauseAction;

    // UI elements.
    private VisualElement pausedPanel;
    private Button pauseResumeButton;
    private Button soundButton;
    private Button restartButton;
    private Button exitToMenuButton;
    private Label timerLabel;

    // UI element names.
    private const string PAUSED_PANEL_NAME = "PausedPanel";
    private const string PAUSE_RESUME_BUTTON_NAME = "PauseResumeButton";
    private const string SOUND_BUTTON_NAME = "SoundButton";
    private const string RESTART_BUTTON_NAME = "RestartButton";
    private const string EXIT_TO_MENU_BUTTON_NAME = "ExitToMenuButton";
    private const string TIMER_LABEL_NAME = "TimerLabel";

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
        restartButton = uiDocument.rootVisualElement.Q<Button>(RESTART_BUTTON_NAME);
        exitToMenuButton = uiDocument.rootVisualElement.Q<Button>(EXIT_TO_MENU_BUTTON_NAME);
        timerLabel = uiDocument.rootVisualElement.Q<Label>(TIMER_LABEL_NAME);
    }

    private void OnEnable()
    {
        GameManager.OnTimeScaleToggled += TogglePauseMenu;

        pauseAction.started += TogglePauseResumeCallback;
        pauseResumeButton.RegisterCallback<ClickEvent>(OnPauseResumeButtonClick);
        soundButton.RegisterCallback<ClickEvent>(OnSoundButtonClick);
        restartButton.RegisterCallback<ClickEvent>(OnRestartButtonClick);
        exitToMenuButton.RegisterCallback<ClickEvent>(OnExitToMenuButtonClick);
    }

    private void OnDisable()
    {
        GameManager.OnTimeScaleToggled -= TogglePauseMenu;

        pauseAction.started -= TogglePauseResumeCallback;
        pauseResumeButton.UnregisterCallback<ClickEvent>(OnPauseResumeButtonClick);
        soundButton.UnregisterCallback<ClickEvent>(OnSoundButtonClick);
        restartButton.UnregisterCallback<ClickEvent>(OnRestartButtonClick);
        exitToMenuButton.UnregisterCallback<ClickEvent>(OnExitToMenuButtonClick);
    }

    private void Start()
    {
        // Ensure the volume is set correctly at the start of the game.
        UpdateSoundButton();
    }

    private void LateUpdate()
    {
        // Update the timer label with the current timer value.
        UpdateTimer();
    }

    /// <summary>
    /// Toggles the pause menu visibility.
    /// </summary>
    private void TogglePauseMenu()
    {
        pauseResumeButton.ToggleInClassList(PAUSE_RESUME_PAUSED_CLASS_NAME);
        pausedPanel.ToggleInClassList(PAUSED_PANEL_REVEAL_CLASS_NAME);
    }

    /// <summary>
    /// Synchronizes the sound button state with the audio mixer volume.
    /// </summary>
    private void UpdateSoundButton()
    {
        if (audioManager.IsAudioMuted())
        {
            soundButton.AddToClassList(SOUND_MUTED_CLASS_NAME);
        }
        else
        {
            soundButton.RemoveFromClassList(SOUND_MUTED_CLASS_NAME);
        }
    }

    /// <summary>
    /// Updates the timer label with the current timer value.
    /// </summary>
    private void UpdateTimer()
    {
        timerLabel.text = TimeManager.FormatTime(timeManager.Timer);
    }

    /// <summary>
    /// Pauses and resumes the game when the pause/resume action button is clicked.
    /// </summary>
    /// <param name="context">Callback context.</param>
    private void TogglePauseResumeCallback(InputAction.CallbackContext context)
    {
        GameManager.ToggleTimeScale();
    }

    /// <summary>
    /// Pauses and resumes the game when the pause/resume UI button is clicked.
    /// </summary>
    /// <param name="clickEvent">Click event.</param>
    private void OnPauseResumeButtonClick(ClickEvent clickEvent)
    {
        GameManager.ToggleTimeScale();
    }

    /// <summary>
    /// Toggles game sounds on/off when the sound button is clicked.
    /// </summary>
    /// <param name="clickEvent">Click event.</param>
    private void OnSoundButtonClick(ClickEvent clickEvent)
    {
        audioManager.ToggleAudioMixerVolume();
        UpdateSoundButton();
    }

    /// <summary>
    /// Restarts the game when the restart button is clicked.
    /// </summary>
    /// <param name="clickEvent">Click event.</param>
    private void OnRestartButtonClick(ClickEvent clickEvent)
    {
        // Restart the game by reloading the current scene.
        SceneManager.LoadScene(SceneManager.GetActiveScene().name);

        // Ensure the game is unpaused when restarting.
        Time.timeScale = 1;
    }

    /// <summary>
    /// Exits to the menu when the exit to menu button is clicked.
    /// </summary>
    /// <param name="clickEvent">Click event.</param>
    private void OnExitToMenuButtonClick(ClickEvent clickEvent)
    {
        // Load the menu scene.
        SceneManager.LoadScene(GameManager.SCENE_MENU);
    }
}
