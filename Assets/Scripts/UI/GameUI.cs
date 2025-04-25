using System.Collections.Generic;
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
    private VisualElement gamePanel;
    private VisualElement pausedPanel;
    private VisualElement endGamePanel;
    private Button pauseResumeButton;
    private Button soundButton;
    private List<Button> restartButtons;
    private List<Button> exitToMenuButtons;
    private List<Label> timerLabels;

    // UI element names.
    private const string GAME_PANEL_NAME = "GamePanel";
    private const string PAUSED_PANEL_NAME = "PausedPanel";
    private const string END_GAME_PANEL_NAME = "EndGamePanel";
    private const string PAUSE_RESUME_BUTTON_NAME = "PauseResumeButton";
    private const string SOUND_BUTTON_NAME = "SoundButton";
    private const string RESTART_BUTTON_NAME = "RestartButton";
    private const string EXIT_TO_MENU_BUTTON_NAME = "ExitToMenuButton";
    private const string TIMER_LABEL_NAME = "TimerLabel";

    // UI class names.
    private const string HIDDEN_CLASS_NAME = "hidden";
    private const string PAUSE_RESUME_PAUSED_CLASS_NAME = "pause-resume-button-paused";
    private const string SOUND_MUTED_CLASS_NAME = "sound-button-muted";

    private void Awake()
    {
        UIDocument uiDocument = GetComponent<UIDocument>();

        pauseAction = InputSystem.actions.FindActionMap(GameManager.INPUT_ACTION_MAP_MENU).FindAction(GameManager.INPUT_ACTION_MENU_PAUSE);

        gamePanel = uiDocument.rootVisualElement.Q<VisualElement>(GAME_PANEL_NAME);
        pausedPanel = uiDocument.rootVisualElement.Q<VisualElement>(PAUSED_PANEL_NAME);
        endGamePanel = uiDocument.rootVisualElement.Q<VisualElement>(END_GAME_PANEL_NAME);

        pauseResumeButton = uiDocument.rootVisualElement.Q<Button>(PAUSE_RESUME_BUTTON_NAME);
        soundButton = uiDocument.rootVisualElement.Q<Button>(SOUND_BUTTON_NAME);
        restartButtons = uiDocument.rootVisualElement.Query<Button>(RESTART_BUTTON_NAME).ToList();
        exitToMenuButtons = uiDocument.rootVisualElement.Query<Button>(EXIT_TO_MENU_BUTTON_NAME).ToList();
        timerLabels = uiDocument.rootVisualElement.Query<Label>(TIMER_LABEL_NAME).ToList();
    }

    private void OnEnable()
    {
        GameManager.OnTimeScaleToggled += OnPauseMenuToggled;
        GameManager.OnGameEnded += OnEndGame;

        pauseAction.started += OnPauseActionStarted;
        pauseResumeButton.RegisterCallback<ClickEvent>(OnPauseResumeButtonClick);
        soundButton.RegisterCallback<ClickEvent>(OnSoundButtonClick);
        restartButtons.ForEach(button => button.RegisterCallback<ClickEvent>(OnRestartButtonClick));
        exitToMenuButtons.ForEach(button => button.RegisterCallback<ClickEvent>(OnExitToMenuButtonClick));
    }

    private void OnDisable()
    {
        GameManager.OnTimeScaleToggled -= OnPauseMenuToggled;
        GameManager.OnGameEnded -= OnEndGame;

        pauseAction.started -= OnPauseActionStarted;
        pauseResumeButton.UnregisterCallback<ClickEvent>(OnPauseResumeButtonClick);
        soundButton.UnregisterCallback<ClickEvent>(OnSoundButtonClick);
        restartButtons.ForEach(button => button.UnregisterCallback<ClickEvent>(OnRestartButtonClick));
        exitToMenuButtons.ForEach(button => button.UnregisterCallback<ClickEvent>(OnExitToMenuButtonClick));
    }

    private void Start()
    {
        gamePanel.RemoveFromClassList(HIDDEN_CLASS_NAME);
        pausedPanel.AddToClassList(HIDDEN_CLASS_NAME);
        endGamePanel.AddToClassList(HIDDEN_CLASS_NAME);

        // Ensure the volume is set correctly at the start of the game.
        UpdateSoundButton();
    }

    private void LateUpdate()
    {
        // Update the timer label with the current timer value.
        UpdateTimer();
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
        foreach (Label timerLabel in timerLabels)
        {
            timerLabel.text = TimeManager.FormatTime(timeManager.Timer);
        }
    }

    /// <summary>
    /// Toggles the pause menu visibility.
    /// </summary>
    private void OnPauseMenuToggled()
    {
        pauseResumeButton.ToggleInClassList(PAUSE_RESUME_PAUSED_CLASS_NAME);
        pausedPanel.ToggleInClassList(HIDDEN_CLASS_NAME);
    }

    /// <summary>
    /// Displays the end game panel and hides the game panel.
    /// </summary>
    private void OnEndGame()
    {
        gamePanel.AddToClassList(HIDDEN_CLASS_NAME);
        endGamePanel.RemoveFromClassList(HIDDEN_CLASS_NAME);

        // Ensure the game is paused when the game ends.
        if (!GameManager.IsGamePaused())
        {
            GameManager.ToggleTimeScale();
        }
    }

    /// <summary>
    /// Pauses and resumes the game when the pause action button is pressed.
    /// </summary>
    /// <param name="context">Callback context.</param>
    private void OnPauseActionStarted(InputAction.CallbackContext context)
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
        if (GameManager.IsGamePaused())
        {
            GameManager.ToggleTimeScale();
        }
    }

    /// <summary>
    /// Exits to the menu when the exit to menu button is clicked.
    /// </summary>
    /// <param name="clickEvent">Click event.</param>
    private void OnExitToMenuButtonClick(ClickEvent clickEvent)
    {
        // Load the menu scene.
        SceneManager.LoadScene(GameManager.SCENE_MENU);

        // Ensure the game is unpaused when exiting to the menu.
        if (GameManager.IsGamePaused())
        {
            GameManager.ToggleTimeScale();
        }
    }
}
