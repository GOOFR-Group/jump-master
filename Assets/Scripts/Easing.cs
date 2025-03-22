using UnityEngine;

public class Easing
{
    /// <summary>
    /// Calculates the easing out sine function for a given amount t.
    /// </summary>
    /// <param name="t">Represents the absolute progress of the animation in the bounds of 0 (beginning of the animation) 
    /// and 1 (end of animation).</param>
    /// <returns>The output of the function.</returns>
    public static float EaseOutSine(float t)
    {
        return Mathf.Sin(t * Mathf.PI / 2);
    }
}
