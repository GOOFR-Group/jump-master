using System.Collections.Generic;
using UnityEngine;

public class CheckPlatformContact : MonoBehaviour
{
    /// <summary>
    /// Defines the map of platform objects that the current object is in contact with. 
    /// The map represents the state of the contact by the platform object id.
    /// </summary>
    private readonly Dictionary<int, bool> platforms = new();

    private void OnTriggerEnter2D(Collider2D collision)
    {
        // Check if the colliding object contains the platform tag.
        // If not, there is not contact with the platform.
        if (!collision.gameObject.CompareTag(GameManager.TAG_PLATFORM))
        {
            return;
        }

        // Set the current platform as true since it is touching the object.
        int instanceID = collision.gameObject.GetInstanceID();
        platforms[instanceID] = true;
    }

    private void OnTriggerExit2D(Collider2D collision)
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
    /// Indicates if the current object is touching another with the platform tag. 
    /// </summary>
    /// <returns>True if the current object is touching a platform, otherwise false.</returns>
    public bool IsTouchingPlatform()
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
