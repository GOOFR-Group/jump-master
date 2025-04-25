using System.Collections;
using UnityEngine;

public class EndGame : MonoBehaviour
{
    [Header("Settings")]

    /// <summary>
    /// Defines the amount of time to wait before calling the end game function.
    /// </summary>
    [SerializeField] private float delay;

    private void OnTriggerEnter2D(Collider2D collision)
    {
        // Check if the colliding object contains the player tag.
        if (!collision.gameObject.CompareTag(GameManager.TAG_PLAYER))
        {
            return;
        }

        // Start the coroutine to end the game after a delay.
        StartCoroutine(DelayedEndGame(delay));
    }

    /// <summary>
    /// Waits for a specified amount of time before ending the game.
    /// </summary>
    /// <param name="delay">The amount of time to wait before ending the game.</param>
    /// <returns>IEnumerator</returns>
    private IEnumerator DelayedEndGame(float delay)
    {
        yield return new WaitForSeconds(delay);

        GameManager.EndGame();
    }
}
