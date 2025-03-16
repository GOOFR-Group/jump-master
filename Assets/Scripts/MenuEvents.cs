using UnityEngine;
using UnityEngine.InputSystem;
using UnityEngine.SceneManagement;
using UnityEngine.UIElements;

public class MenuEvents : MonoBehaviour
{
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

    private void Awake()
    {
        UIDocument uiDocument = GetComponent<UIDocument>();

        startButton = uiDocument.rootVisualElement.Q<Button>("StartButton");

        space = uiDocument.rootVisualElement.Q<VisualElement>("Space");
        left = uiDocument.rootVisualElement.Q<VisualElement>("Left");
        right = uiDocument.rootVisualElement.Q<VisualElement>("Right");

        spaceLabel = uiDocument.rootVisualElement.Q<Label>("SpaceLabel");
        leftLabel = uiDocument.rootVisualElement.Q<Label>("LeftLabel");
        rightLabel = uiDocument.rootVisualElement.Q<Label>("RightLabel");

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
            space.AddToClassList("space-pressed");
            spaceLabel.AddToClassList("control-label-pressed");
        }
        else
        {
            space.RemoveFromClassList("space-pressed");
            spaceLabel.RemoveFromClassList("control-label-pressed");
        }

        if (leftAction.IsPressed())
        {
            left.AddToClassList("left-pressed");
            leftLabel.AddToClassList("control-label-pressed");
        }
        else
        {
            left.RemoveFromClassList("left-pressed");
            leftLabel.RemoveFromClassList("control-label-pressed");
        }

        if (rightAction.IsPressed())
        {
            right.AddToClassList("right-pressed");
            rightLabel.AddToClassList("control-label-pressed");
        }
        else
        {
            right.RemoveFromClassList("right-pressed");
            rightLabel.RemoveFromClassList("control-label-pressed");
        }
    }

    private void OnStartButtonClick(ClickEvent clickEvent)
    {
        SceneManager.LoadScene(GameManager.SCENE_MAIN);
    }
}
