using System;

public class MergeSortedArray {
    public void Merge(int[] nums1, int m, int[] nums2, int n) {
        int i = m - 1;
        int j = n - 1;
        int k = m + n - 1;

        while (i >= 0 && j >= 0) {
            if (nums1[i] > nums2[j]) {
                nums1[k--] = nums1[i--];
            } else {
                nums1[k--] = nums2[j--];
            }
        }

        while (j >= 0) {
            nums1[k--] = nums2[j--];
        }
    }
}

class Program {
    static void Main(string[] args) {
        int[] nums1 = new int[] {1, 2, 3, 0, 0, 0};
        int m = 3;
        int[] nums2 = new int[] {2, 5, 6};
        int n = 3;

        MergeSortedArray merger = new MergeSortedArray();
        merger.Merge(nums1, m, nums2, n);

        Console.WriteLine("Merged Array:");
        foreach (int num in nums1) {
            Console.Write(num + " ");
        }
    }
}
