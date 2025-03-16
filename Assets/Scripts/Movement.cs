using UnityEngine;
using UnityEngine.InputSystem;

public class Movement : MonoBehaviour
{
    [Header("Settings")]

    /// <summary>
    /// Defines the movement speed multiplier.
    /// </summary>
    [SerializeField] private float speed;

    [Header("Required Components")]

    [SerializeField] private CheckPlatformContact checkGround;
    private SpriteRenderer spriteRenderer;
    private new Rigidbody2D rigidbody2D;

    // Action status.
    private bool leftActionStatus;
    private bool rightActionStatus;
    private bool jumpActionStatus;

    // Input actions.
    private InputAction moveAction;
    private InputAction jumpAction;

    private void Awake()
    {
        spriteRenderer = GetComponent<SpriteRenderer>();
        rigidbody2D = GetComponent<Rigidbody2D>();

        moveAction = InputSystem.actions.FindActionMap(GameManager.INPUT_ACTION_MAP_PLAYER).FindAction(GameManager.INPUT_ACTION_PLAYER_MOVE);
        jumpAction = InputSystem.actions.FindActionMap(GameManager.INPUT_ACTION_MAP_PLAYER).FindAction(GameManager.INPUT_ACTION_PLAYER_JUMP);
    }

    private void FixedUpdate()
    {
        // Check if the object is in contact with the ground.
        if (!checkGround.IsTouchingPlatform() || rigidbody2D.linearVelocityY > Mathf.Epsilon)
        {
            return;
        }

        // Compute the direction to add to the object based on the left and right actions.
        int direction = 0;
        if (leftActionStatus)
        {
            // Add the left direction.
            direction -= 1;
            spriteRenderer.flipX = true;
        }
        if (rightActionStatus)
        {
            // Add the right direction.
            direction += 1;
            spriteRenderer.flipX = false;
        }

        // Reset the horizontal velocity of the object 
        // when no movement action is performed or the jump action is being performed.
        if (direction == 0 || jumpActionStatus)
        {
            rigidbody2D.linearVelocityX = 0;
            return;
        }

        // Add the computed velocity when the movement actions are performed.
        rigidbody2D.linearVelocityX = direction * speed * Time.fixedDeltaTime;
    }

    private void Update()
    {
        // Update the input actions.
        float moveValue = moveAction.ReadValue<float>();

        leftActionStatus = moveValue < 0;
        rightActionStatus = moveValue > 0;
        jumpActionStatus = jumpAction.IsPressed();
    }
}