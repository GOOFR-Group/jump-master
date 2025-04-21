using System.Collections;
using UnityEngine;

public class EndGame : MonoBehaviour
{
    [Header("Settings")]

    /// <summary>
    /// Defines the amount of time to wait before ending the game and showing the pause menu.
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
    /// Waits for a specified amount of time before ending the game and showing the pause menu.
    /// </summary>
    /// <returns>IEnumerator</returns>
    private IEnumerator DelayedEndGame(float delay)
    {
        yield return new WaitForSeconds(delay);

        // Check if the game is not paused.
        if (!GameManager.IsGamePaused())
        {
            // Pause the game and show the pause menu.
            GameManager.ToggleTimeScale();
        }
    }
}
