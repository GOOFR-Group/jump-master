using System.Collections.Generic;
using UnityEngine;

public class CheckGround : MonoBehaviour
{
    /// <summary>
    /// Defines the map of ground objects that the current object is in contact with. The map represents the
    /// state of the contact by the ground object id.
    /// </summary>
    private readonly Dictionary<int, bool> grounds = new Dictionary<int, bool>();

    /// <summary>
    /// Indicates if the current object is touching the ground. The ground is represented by any object with the
    /// platform tag that is in contact with the feet of this object.
    /// </summary>
    /// <returns>True if the current object is touching the ground, otherwise false.</returns>
    public bool IsTouchingGround()
    {
        foreach (var touching in grounds.Values)
        {
            if (touching)
            {
                return true;
            }
        }

        return false;
    }

    private void OnTriggerEnter2D(Collider2D collision)
    {
        // Check if the colliding object contains the platform tag.
        // If not, there is not contact with the ground.
        if (!collision.gameObject.CompareTag(GameManager.TAG_PLATFORM))
        {
            return;
        }

        // Set the current ground as true since it is touching the object.
        int instanceID = collision.gameObject.GetInstanceID();
        grounds[instanceID] = true;
    }

    private void OnTriggerExit2D(Collider2D collision)
    {
        // Check if the colliding object contains the platform tag.
        if (!collision.gameObject.CompareTag(GameManager.TAG_PLATFORM))
        {
            return;
        }

        // Set the current ground as false since it is not touching the object anymore.
        int instanceID = collision.gameObject.GetInstanceID();
        grounds[instanceID] = false;
    }
}