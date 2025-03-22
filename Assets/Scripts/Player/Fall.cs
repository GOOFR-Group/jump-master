using UnityEngine;

public class Fall : MonoBehaviour
{
    [Header("Settings")]

    /// <summary>
    /// Defines the amount of time possible to be in the air until it is considered a fall when touching the ground.
    /// </summary>
    [SerializeField] private float allowedDuration;

    [Header("Required Components")]

    [SerializeField] private CheckPlatformContact checkGround;
    [SerializeField] private AudioClip audioClipFall;
    [SerializeField] private AudioClip audioClipLanding;
    private Animator animator;
    private new Rigidbody2D rigidbody2D;
    private AudioSource audioSource;

    /// <summary>
    /// Defines the timer that captures the amount of time the object is falling.
    /// </summary>
    private float timer;

    private void Awake()
    {
        animator = GetComponent<Animator>();
        rigidbody2D = GetComponent<Rigidbody2D>();
        audioSource = GetComponent<AudioSource>();
    }

    private void Update()
    {
        // Check if the object is falling.
        if (rigidbody2D.linearVelocity.y < -Mathf.Epsilon && !checkGround.IsTouchingPlatform())
        {
            // If it is falling, update the timer.
            timer += Time.deltaTime;
            return;
        }

        // Check if the object was falling for longer than the allowed duration.
        if (timer > allowedDuration)
        {
            rigidbody2D.linearVelocityX = 0;
            animator.Play(GameManager.ANIMATION_PLAYER_FALL);
            audioSource.clip = audioClipFall;
            audioSource.Play();
        }
        else if (timer > 0 && checkGround.IsTouchingPlatform())
        {
            audioSource.clip = audioClipLanding;
            audioSource.Play();
        }

        // Reset the timer.
        timer = 0;
    }
}
