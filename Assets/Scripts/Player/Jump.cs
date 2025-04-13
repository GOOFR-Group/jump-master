using System.Collections.Generic;
using System.Linq;
using UnityEngine;
using UnityEngine.InputSystem;

public class Jump : MonoBehaviour
{
    [Header("Settings")]

    /// <summary>
    /// Defines the base impulse of the jump to accumulate each second the jump action is performed.
    /// </summary>
    [SerializeField] private float impulse;

    /// <summary>
    /// Defines the minimum impulse of the jump.
    /// </summary>
    [SerializeField] private float minImpulse;

    /// <summary>
    /// Defines the maximum impulse of the jump.
    /// </summary>
    [SerializeField] private float maxImpulse;

    /// <summary>
    /// Defines the minimum angle in degrees to apply when jumping left or right.
    /// </summary>
    [SerializeField] private float minDiagonalAngle;

    /// <summary>
    /// Defines the maximum angle in degrees to apply when jumping left or right.
    /// </summary>
    [SerializeField] private float maxDiagonalAngle;

    /// <summary>
    /// Defines the buffer, in seconds, of actions to be considered before the jump.
    /// </summary>
    [SerializeField] private float actionBufferTimeBeforeJump;

    /// <summary>
    /// Defines the buffer, in seconds, of actions to be considered after the jump.
    /// </summary>
    [SerializeField] private float actionBufferTimeAfterJump;

    [Header("Required Components")]

    [SerializeField] private CheckPlatformContact checkGround;
    [SerializeField] private AudioClip audioClip;
    private SpriteRenderer spriteRenderer;
    private Animator animator;
    private new Rigidbody2D rigidbody2D;
    private AudioSource audioSource;

    /// <summary>
    ///  Defines the previously used jump impulse.
    /// </summary>
    private float usedImpulse;

    /// <summary>
    ///  Defines the current accumulated jump impulse.
    /// </summary>
    private float accumulatedImpulse;

    /// <summary>
    /// Defines if the object is able to jump.
    /// </summary>
    private bool canJump;

    /// <summary>
    /// Defines the current buffer, in seconds, of actions to be considered before the jump.
    /// </summary>
    private float currentActionBufferTimeBeforeJump;

    /// <summary>
    /// Defines the current buffer, in seconds, of actions to be considered after the jump.
    /// </summary>
    private float currentActionBufferTimeAfterJump;

    /// <summary>
    /// Defines the action buffer, in frames, to be considered before the jump action is performed.
    /// </summary>
    private readonly Queue<float> actionBufferBeforeJump = new();

    /// <summary>
    /// Defines the action buffer, in frames, to be considered after the jump action is performed.
    /// </summary>
    private Queue<float> actionBufferAfterJump = new();

    // Input actions.
    private InputAction moveAction;
    private InputAction jumpAction;

    private void Awake()
    {
        spriteRenderer = GetComponent<SpriteRenderer>();
        animator = GetComponent<Animator>();
        rigidbody2D = GetComponent<Rigidbody2D>();
        audioSource = GetComponent<AudioSource>();

        moveAction = InputSystem.actions.FindActionMap(GameManager.INPUT_ACTION_MAP_PLAYER).FindAction(GameManager.INPUT_ACTION_PLAYER_MOVE);
        jumpAction = InputSystem.actions.FindActionMap(GameManager.INPUT_ACTION_MAP_PLAYER).FindAction(GameManager.INPUT_ACTION_PLAYER_JUMP);
    }

    private void Start()
    {
        usedImpulse = 0;
        accumulatedImpulse = 0;
        canJump = false;

        currentActionBufferTimeBeforeJump = actionBufferTimeBeforeJump;
        currentActionBufferTimeAfterJump = 0;
    }

    private void FixedUpdate()
    {
        // Check if the object is falling.
        if (rigidbody2D.linearVelocity.y < -Mathf.Epsilon && !checkGround.IsTouchingPlatform())
        {
            animator.Play(GameManager.ANIMATION_PLAYER_JUMP_FALL);
        }

        // Check if the object can jump.
        if (!canJump)
        {
            return;
        }

        // Get action from the buffer.
        // Start by searching for the action in the buffer after the jump because it is the most recent.
        // If no action is found, search in the buffer before the jump.
        float action = 0;

        foreach (var a in actionBufferAfterJump.Reverse())
        {
            if (a != 0)
            {
                action = a;

                // Override sprite flip when an action after the jump is considered.
                if (action < 0)
                {
                    spriteRenderer.flipX = true;
                }
                else if (action > 0)
                {
                    spriteRenderer.flipX = false;
                }

                break;
            }
        }

        if (action == 0)
        {
            foreach (var a in actionBufferBeforeJump.Reverse())
            {
                if (a != 0)
                {
                    action = a;
                    break;
                }
            }
        }

        // Jump only if an action is taken within the expected buffer.
        if (action == 0 && currentActionBufferTimeAfterJump > 0)
        {
            return;
        }

        // Limit the minimum jump impulse.
        accumulatedImpulse = Mathf.Max(accumulatedImpulse, minImpulse);

        // Compute the jump rotation based on the left and right actions.
        Vector2 direction = Vector2.up;

        // Check that the left or right action is being performed.
        if (action != 0)
        {
            // Compute the diagonal angle based on the fraction of accumulated impulse.
            float fraction = accumulatedImpulse / maxImpulse;
            float diagonalAngle = Mathf.Lerp(minDiagonalAngle, maxDiagonalAngle, fraction);

            // Compute the rotated direction from the right side.
            Quaternion rotation = Quaternion.Euler(0, 0, diagonalAngle);
            direction = rotation * Vector2.right;

            // If the object is performing a left action, invert the direction vector.
            if (action < 0)
            {
                direction.x = -direction.x;
            }
        }

        // Apply the jump velocity based on the computed rotation and accumulated impulse.
        usedImpulse = accumulatedImpulse;
        Vector2 velocity = direction * accumulatedImpulse;

        rigidbody2D.AddForce(velocity, ForceMode2D.Impulse);
        animator.Play(GameManager.ANIMATION_PLAYER_JUMP);
        audioSource.clip = audioClip;
        audioSource.Play();

        // Reset the accumulated impulse and jump flag.
        accumulatedImpulse = 0;
        canJump = false;

        // Reset the buffer.
        currentActionBufferTimeAfterJump = 0;
        actionBufferAfterJump.Clear();
    }

    private void Update()
    {
        // Check if the game is paused.
        if (Time.timeScale == 0)
        {
            return;
        }

        // Update current timers.
        if (currentActionBufferTimeBeforeJump > 0)
        {
            currentActionBufferTimeBeforeJump -= Time.deltaTime;
        }
        if (currentActionBufferTimeAfterJump > 0)
        {
            currentActionBufferTimeAfterJump -= Time.deltaTime;
        }

        // Save actions in the buffer.
        float moveValue = moveAction.ReadValue<float>();

        actionBufferBeforeJump.Enqueue(moveValue);
        if (currentActionBufferTimeBeforeJump <= 0)
        {
            actionBufferBeforeJump.Dequeue();
        }

        if (currentActionBufferTimeAfterJump > 0)
        {
            actionBufferAfterJump.Enqueue(moveValue);
        }

        // Check if the object is in contact with the ground and if the fall animation has already ended.
        AnimatorStateInfo currentAnimatorStateInfo = animator.GetCurrentAnimatorStateInfo(0);

        if (!checkGround.IsTouchingPlatform() ||
        (currentAnimatorStateInfo.IsName(GameManager.ANIMATION_PLAYER_FALL) && currentAnimatorStateInfo.normalizedTime < 1))
        {
            accumulatedImpulse = 0;
            return;
        }

        // Check if the jump action is being performed.
        if (jumpAction.IsPressed())
        {
            // Apply the impulse multiplier and ensure that the accumulated impulse is not greater than the maximum defined.
            accumulatedImpulse += impulse * Time.deltaTime;
            accumulatedImpulse = Mathf.Min(accumulatedImpulse, maxImpulse);

            // Reset the horizontal velocity of the object when the jump action is being performed.
            rigidbody2D.linearVelocityX = 0;
            animator.Play(GameManager.ANIMATION_PLAYER_JUMP_HOLD);
        }

        // Check if the jump action was released.
        if (jumpAction.WasReleasedThisFrame())
        {
            // The object is able to jump
            canJump = true;

            // Reset the action buffer after the jump.
            currentActionBufferTimeAfterJump = actionBufferTimeAfterJump;
        }
    }

    /// <summary>
    /// Indicates the previously used jump impulse.
    /// </summary>
    /// <returns>The previously used jump impulse.</returns>
    public float UsedImpulse()
    {
        return usedImpulse;
    }

    /// <summary>
    /// Indicates the maximum jump impulse.
    /// </summary>
    /// <returns>The maximum jump impulse.</returns>
    public float MaxImpulse()
    {
        return maxImpulse;
    }
}
