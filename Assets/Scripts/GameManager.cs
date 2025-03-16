using UnityEngine;

public class GameManager : MonoBehaviour
{
    // Player animation constants.
    public const string ANIMATION_PLAYER_IDLE = "Idle";
    public const string ANIMATION_PLAYER_WALK = "Walk";
    public const string ANIMATION_PLAYER_JUMP_HOLD = "Jump Hold";
    public const string ANIMATION_PLAYER_JUMP = "Jump";
    public const string ANIMATION_PLAYER_JUMP_FALL = "Jump Fall";
    public const string ANIMATION_PLAYER_KNOCK_BACK = "Knock Back";
    public const string ANIMATION_PLAYER_FALL = "Fall";

    // Tag constants.
    public const string TAG_PLATFORM = "Platform";

    // Action constants.
    public const string INPUT_ACTION_MOVE = "Move";
    public const string INPUT_ACTION_JUMP = "Jump";
}
