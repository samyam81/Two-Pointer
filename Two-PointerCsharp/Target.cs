using System;
using System.Collections.Generic;

public class Target
{
    public static void Main(string[] args)
    {
        int[] arr = { 1, 2, 4, 8, 16 };
        int target = 20;
        Target t = new Target();
        int[] resultIndices = t.FindSolution(arr, target);
        foreach (int index in resultIndices)
        {
            Console.Write(arr[index] + " ");
        }
    }

    public int[] FindSolution(int[] arr, int target)
    {
        Dictionary<int, int> comp = new Dictionary<int, int>();

        for (int i = 0; i < arr.Length; i++)
        {
            int com = target - arr[i];
            if (comp.ContainsKey(com))
            {
                return new int[] { comp[com], i };
            }
            comp[arr[i]] = i;
        }
        return new int[0];
    }
}
