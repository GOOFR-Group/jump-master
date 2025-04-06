using UnityEngine;

public class CameraController : MonoBehaviour
{
    [Header("Settings")]

    /// <summary>
    /// Defines the speed of the animation transition.
    /// </summary>
    [SerializeField] private float transitionSpeed;

    [Header("Required Components")]

    private new Camera camera;
    private GameObject playerGameObject;

    /// <summary>
    /// Defines the current amount, in the range [0; 1], that has been transitioned from the previousPosition to the 
    /// currentPosition.
    /// </summary>
    private float currentTransition;

    /// <summary>
    /// Defines the camera initial position.
    /// </summary>
    private Vector3 initialPosition;

    /// <summary>
    /// Defines the previous position of the camera.
    /// </summary>
    private Vector3 previousPosition;

    /// <summary>
    /// Defines the current position of the camera. 
    /// It represents the target position of the current level.
    /// </summary>
    private Vector3 currentPosition;

    private void Awake()
    {
        camera = GetComponent<Camera>();
        playerGameObject = GameObject.FindGameObjectWithTag(GameManager.TAG_PLAYER);
    }

    private void Start()
    {
        // Check that the camera transition speed is valid.
        if (transitionSpeed <= 0)
        {
            // Invalid speed, defaults to 1.
            transitionSpeed = 1;
        }

        // Save the camera initial position.
        initialPosition = transform.position;

        // Initialize the previous and current positions.
        previousPosition = initialPosition;
        currentPosition = initialPosition;
    }

    private void Update()
    {
        if (playerGameObject == null)
        {
            Debug.LogError("null player game object");
            return;
        }

        // Get the player bounds.
        if (!playerGameObject.TryGetComponent<Collider2D>(out var playerCollider))
        {
            Debug.LogError("null player collider");
            return;
        }

        // Compute the camera position based on the player minimum bound.
        float cameraHeight = camera.orthographicSize * 2f;
        int level = (int)((playerCollider.bounds.min.y - initialPosition.y + cameraHeight * 0.5) / cameraHeight);

        // newPosition represents the new position of the camera considering the current level of the player.
        Vector3 newPosition = new(transform.position.x, cameraHeight * level + initialPosition.y, initialPosition.z);


        // Check whether the camera position has changed.
        if (newPosition != currentPosition)
        {
            // Update the previous and current positions.
            previousPosition = transform.position;
            currentPosition = newPosition;

            // Reset the animation transition.
            currentTransition = 0;
        }

        // If the transition has been completed, avoid unnecessary computation.
        if (currentTransition >= 1)
        {
            return;
        }

        // Compute the new animation transition.
        currentTransition = Mathf.Clamp(currentTransition + Time.deltaTime * transitionSpeed, 0, 1);

        // Apply the current transition using an easing function.
        transform.position = Vector3.Lerp(previousPosition, currentPosition, Easing.EaseOutSine(currentTransition));
    }
}
