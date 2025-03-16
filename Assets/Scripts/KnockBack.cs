using System.Collections.Generic;
using UnityEngine;

public class KnockBack : MonoBehaviour
{
    [Header("Settings")]

    /// <summary>
    /// Defines the impulse of the knock-back.
    /// </summary>
    [SerializeField] private float impulse;

    /// <summary>
    /// Defines the angle in degrees to apply when there is a knock-back.
    /// </summary>
    [SerializeField] private float diagonalAngle;

    [Header("Required Components")]

    [SerializeField] private CheckPlatformContact checkCeiling;
    [SerializeField] private CheckPlatformContact checkGround;
    private Animator animator;
    private new Rigidbody2D rigidbody2D;
    private Jump jump;

    /// <summary>
    /// Defines the map of platform objects that the current object is in contact with. 
    /// The map represents the state of the contact by the platform object id.
    /// </summary>
    private readonly Dictionary<int, bool> platforms = new();

    /// <summary>
    /// Defines the velocity value from the previous physics update. 
    /// Used to check if the object was falling in a straight line before colliding.
    /// </summary>
    private Vector2 previousVelocity;

    private void Awake()
    {
        animator = GetComponent<Animator>();
        rigidbody2D = GetComponent<Rigidbody2D>();
        jump = GetComponent<Jump>();
    }

    private void FixedUpdate()
    {
        // Update the previous velocity.
        previousVelocity = rigidbody2D.linearVelocity;
    }

    void OnCollisionEnter2D(Collision2D collision)
    {
        // Check if the colliding object contains the platform tag.
        if (!collision.gameObject.CompareTag(GameManager.TAG_PLATFORM))
        {
            return;
        }

        // alreadyInContact defines if the object is already in contact with a platform.
        bool alreadyInContact = PlatformContact();

        // Set the current platform as true since it is touching the object.
        int instanceID = collision.gameObject.GetInstanceID();
        platforms[instanceID] = true;

        // Check if the object is already in contact with a platform, if it is on the ground or in touch with the ceiling.
        // There is also no need to apply knock-back when the velocity vector of the object represents a 90º angle (object
        // falling in a straight line).
        if (alreadyInContact || checkGround.IsTouchingPlatform() || checkCeiling.IsTouchingPlatform() ||
            Mathf.Approximately(Mathf.Abs(Vector2.Dot(Vector2.up, previousVelocity.normalized)), 1))
        {
            return;
        }

        // Get the collision contact point.
        Vector2 contactPoint = collision.GetContact(0).point;

        // Compute the diagonal angle based on the fraction of used jump impulse.
        float diagonalAngle = this.diagonalAngle;
        diagonalAngle *= jump.UsedImpulse() / jump.MaxImpulse();

        // Compute the rotated direction from the right side.
        Quaternion rotation = Quaternion.Euler(0, 0, diagonalAngle);
        Vector2 direction = rotation * Vector2.right;

        // If the object is in the left side of the collision, invert the direction vector.
        if (transform.position.x < contactPoint.x)
        {
            direction.x = -direction.x;
        }

        // Apply the knock-back velocity based on the computed rotation and impulse.
        Vector2 velocity = direction * impulse;
        rigidbody2D.AddForce(velocity, ForceMode2D.Impulse);
        animator.Play(GameManager.ANIMATION_PLAYER_KNOCK_BACK);
    }

    private void OnCollisionExit2D(Collision2D collision)
    {
        // Check if the colliding object contains the platform tag.
        if (!collision.gameObject.CompareTag(GameManager.TAG_PLATFORM))
        {
            return;
        }

        // Set the current platform as false since it is not touching the object anymore.
        int instanceID = collision.gameObject.GetInstanceID();
        platforms[instanceID] = false;
    }

    /// <summary>
    /// Indicates if the current object is touching any platform. 
    /// The platform is represented by any object with the platform tag.
    /// </summary>
    /// <returns>True if the current object is touching any platform.</returns>
    private bool PlatformContact()
    {
        foreach (var touching in platforms.Values)
        {
            if (touching)
            {
                return true;
            }
        }

        return false;
    }
}
