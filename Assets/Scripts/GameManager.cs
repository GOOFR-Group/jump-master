using UnityEngine;

public class GameManager : MonoBehaviour
{
    // Scene constants.
    public const string SCENE_MAIN = "MainScene";

    // Player animation constants.
    public const string ANIMATION_PLAYER_IDLE = "Idle";
    public const string ANIMATION_PLAYER_WALK = "Walk";
    public const string ANIMATION_PLAYER_JUMP_HOLD = "Jump Hold";
    public const string ANIMATION_PLAYER_JUMP = "Jump";
    public const string ANIMATION_PLAYER_JUMP_FALL = "Jump Fall";
    public const string ANIMATION_PLAYER_KNOCK_BACK = "Knock Back";
    public const string ANIMATION_PLAYER_FALL = "Fall";

    // Tag constants.
    public const string TAG_PLAYER = "Player";
    public const string TAG_PLATFORM = "Platform";

    // Action constants.
    public const string INPUT_ACTION_MAP_PLAYER = "Player";
    public const string INPUT_ACTION_MAP_MENU = "Menu";

    public const string INPUT_ACTION_PLAYER_MOVE = "Move";
    public const string INPUT_ACTION_PLAYER_JUMP = "Jump";

    public const string INPUT_ACTION_MENU_LEFT = "Left";
    public const string INPUT_ACTION_MENU_RIGHT = "Right";
    public const string INPUT_ACTION_MENU_JUMP = "Jump";
    public const string INPUT_ACTION_MENU_PAUSE = "Pause";

    // Audio constants.
    public const string AUDIO_MIXER_MASTER_VOLUME = "Volume";

    /// <summary>
    /// Toggles the time scale to pause or resume the game.
    /// </summary>
    public static void ToggleTimeScale()
    {
        if (Time.timeScale == 0)
        {
            Time.timeScale = 1;
        }
        else
        {
            Time.timeScale = 0;
        }
    }
}
