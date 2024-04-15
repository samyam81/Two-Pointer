using System;

namespace kobbie
{
    class MoveZero
    {
        static void Main()
        {
            int[] arr = { 0, 1, 2, 0, 0, 6 };
            int left = 0;
            int right = arr.Length - 1;
            
            while (left <= right)
            {
                if (arr[left] == 0)
                {
                    int temp = arr[left];
                    arr[left] = arr[right];
                    arr[right] = temp;
                    --right;
                }
                left++;
            }
            
            for (int i = 0; i < arr.Length; i++)
            {
                Console.Write(arr[i] + " ");
            }
        }
    }
}
