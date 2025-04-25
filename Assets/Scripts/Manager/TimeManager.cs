using UnityEngine;

public class TimeManager : MonoBehaviour
{
    /// <summary>
    /// The time elapsed since the start of the game in seconds.
    /// This value is updated every frame unless the game is paused.
    /// </summary>
    public float Timer { get; private set; }

    private void Start()
    {
        Timer = 0;
    }

    private void Update()
    {
        // Check if the game is paused.
        if (GameManager.IsGamePaused())
        {
            return;
        }

        // Update the timer by adding the time since the last frame.
        Timer += Time.deltaTime;
    }

    /// <summary>
    /// Formats the given time in seconds into a string representation of minutes and seconds.
    /// </summary>
    /// <param name="time">Time in seconds.</param>
    /// <returns>Formatted time string in "MM:SS" format.</returns>
    public static string FormatTime(float time)
    {
        int minutes = Mathf.FloorToInt(time / 60);
        int seconds = Mathf.FloorToInt(time % 60);
        return string.Format("{0:00}:{1:00}", minutes, seconds);
    }
}
