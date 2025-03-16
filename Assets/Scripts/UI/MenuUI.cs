using UnityEngine;
using UnityEngine.InputSystem;
using UnityEngine.SceneManagement;
using UnityEngine.UIElements;

public class MenuEvents : MonoBehaviour
{
    // UI elements.
    private Button startButton;

    private VisualElement space;
    private VisualElement left;
    private VisualElement right;

    private Label spaceLabel;
    private Label leftLabel;
    private Label rightLabel;

    // Input actions.
    private InputAction leftAction;
    private InputAction rightAction;
    private InputAction jumpAction;

    // UI element names.
    private const string START_BUTTON_NAME = "StartButton";
    private const string SPACE_NAME = "Space";
    private const string LEFT_NAME = "Left";
    private const string RIGHT_NAME = "Right";
    private const string SPACE_LABEL_NAME = "SpaceLabel";
    private const string LEFT_LABEL_NAME = "LeftLabel";
    private const string RIGHT_LABEL_NAME = "RightLabel";

    // UI class names.
    private const string SPACE_PRESSED_CLASS_NAME = "space-pressed";
    private const string LEFT_PRESSED_CLASS_NAME = "left-pressed";
    private const string RIGHT_PRESSED_CLASS_NAME = "right-pressed";
    private const string CONTROL_LABEL_PRESSED_CLASS_NAME = "control-label-pressed";

    private void Awake()
    {
        UIDocument uiDocument = GetComponent<UIDocument>();

        startButton = uiDocument.rootVisualElement.Q<Button>(START_BUTTON_NAME);

        space = uiDocument.rootVisualElement.Q<VisualElement>(SPACE_NAME);
        left = uiDocument.rootVisualElement.Q<VisualElement>(LEFT_NAME);
        right = uiDocument.rootVisualElement.Q<VisualElement>(RIGHT_NAME);

        spaceLabel = uiDocument.rootVisualElement.Q<Label>(SPACE_LABEL_NAME);
        leftLabel = uiDocument.rootVisualElement.Q<Label>(LEFT_LABEL_NAME);
        rightLabel = uiDocument.rootVisualElement.Q<Label>(RIGHT_LABEL_NAME);

        leftAction = InputSystem.actions.FindActionMap(GameManager.INPUT_ACTION_MAP_MENU).FindAction(GameManager.INPUT_ACTION_MENU_LEFT);
        rightAction = InputSystem.actions.FindActionMap(GameManager.INPUT_ACTION_MAP_MENU).FindAction(GameManager.INPUT_ACTION_MENU_RIGHT);
        jumpAction = InputSystem.actions.FindAction(GameManager.INPUT_ACTION_PLAYER_JUMP);
    }

    private void OnEnable()
    {
        startButton.RegisterCallback<ClickEvent>(OnStartButtonClick);
    }

    private void OnDisable()
    {
        startButton.UnregisterCallback<ClickEvent>(OnStartButtonClick);
    }

    private void Update()
    {
        if (jumpAction.IsPressed())
        {
            space.AddToClassList(SPACE_PRESSED_CLASS_NAME);
            spaceLabel.AddToClassList(CONTROL_LABEL_PRESSED_CLASS_NAME);
        }
        else
        {
            space.RemoveFromClassList(SPACE_PRESSED_CLASS_NAME);
            spaceLabel.RemoveFromClassList(CONTROL_LABEL_PRESSED_CLASS_NAME);
        }

        if (leftAction.IsPressed())
        {
            left.AddToClassList(LEFT_PRESSED_CLASS_NAME);
            leftLabel.AddToClassList(CONTROL_LABEL_PRESSED_CLASS_NAME);
        }
        else
        {
            left.RemoveFromClassList(LEFT_PRESSED_CLASS_NAME);
            leftLabel.RemoveFromClassList(CONTROL_LABEL_PRESSED_CLASS_NAME);
        }

        if (rightAction.IsPressed())
        {
            right.AddToClassList(RIGHT_PRESSED_CLASS_NAME);
            rightLabel.AddToClassList(CONTROL_LABEL_PRESSED_CLASS_NAME);
        }
        else
        {
            right.RemoveFromClassList(RIGHT_PRESSED_CLASS_NAME);
            rightLabel.RemoveFromClassList(CONTROL_LABEL_PRESSED_CLASS_NAME);
        }
    }

    /// <summary>
    /// Loads main game scene when the start menu button is clicked.
    /// </summary>
    /// <param name="clickEvent">Click event.</param>
    private void OnStartButtonClick(ClickEvent clickEvent)
    {
        SceneManager.LoadScene(GameManager.SCENE_MAIN);
    }
}
