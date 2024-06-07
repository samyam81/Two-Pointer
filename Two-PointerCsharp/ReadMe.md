# Two-PointerCSharp

This repository contains C# solutions for various problems implemented using the Two Pointer algorithm approach.

## Prerequisites

To run or contribute to the code in this repository, you should have:

- Basic knowledge of C# programming language.
- Understanding of basic data structures such as arrays, linked lists, and dictionaries.
- Familiarity with algorithms like Two Pointer, Floyd's Tortoise and Hare, and sorting algorithms.

## Contents

1. **IntersectionOfLinkedList.cs**:
   This C# code defines a class `IntersectionOfLinkedList` containing a method `GetIntersectionNode` that finds the intersection node of two singly-linked lists. It iterates through both lists simultaneously, advancing each pointer until they either meet at an intersection node or reach the end of both lists. If there is no intersection, it returns null.

2. **LinkedListCycle.cs**:
   This C# code defines a class `LinkedListCycle` with a `Main` method that creates a linked list with a cycle and checks if it has a cycle. The `hasCycle` method implements Floyd's Tortoise and Hare algorithm to detect a cycle by using two pointers, slow and fast, which move at different speeds through the list until they either meet (indicating a cycle) or one reaches the end of the list (indicating no cycle).

3. **MergeSortList.cs**:
   This C# code defines a class `MergeSortedArray` with a method `Merge` that merges two sorted arrays `nums1` and `nums2` into `nums1` in sorted order. It iterates through both arrays from the end, comparing elements and placing them in the correct position in `nums1`. The `Main` method in the `Program` class creates two arrays, merges them using the `MergeSortedArray` class, and prints the merged array.

4. **MoveZero.cs**:
   This C# code defines a class `MoveZero` with a `Main` method that moves all zeros in an array to the end while maintaining the relative order of non-zero elements. It uses two pointers, `left` starting from the beginning and `right` starting from the end of the array, swapping elements when `arr[left]` is zero. Finally, it prints the modified array.

5. **Target.cs**:
   This C# code defines a class `Target` with a `Main` method that finds a solution to a target sum problem in a sorted array. It initializes a dictionary to store complements (the difference between the target and each element) and iterates through the array. If the complement of the current element is found in the dictionary, it returns the indices of the two elements whose sum equals the target. Otherwise, it stores the element and its index in the dictionary. Finally, it prints the elements corresponding to the indices.

6. **ThreeSum.cs**:
   This C# code defines a class `ThreeSumSolution` with a method `ThreeSum` that finds all unique triplets in an array whose sum is zero. It sorts the input array and uses a two-pointer approach to iterate through it, fixing one element and finding the other two elements such that their sum equals the negation of the fixed element. It avoids duplicate triplets by skipping identical elements during iteration. The `Main` method creates an instance of `ThreeSumSolution`, applies the method to an example array, and prints the resulting triplets.

7. **TrapWater.cs**:
   This C# code defines a class `TrapWater` with a `Main` method that calculates the total amount of water trapped between bars represented by an array. It uses three auxiliary arrays: `left` to store the maximum bar height to the left of each position, `right` to store the maximum bar height to the right of each position, and `value` to store the trapped water at each position. It iterates through the array to fill these auxiliary arrays and then calculates the trapped water by finding the minimum of the corresponding values in the `left` and `right` arrays and subtracting the height of the current bar. Finally, it sums up the trapped water and prints the total.

8. **ValidPalindrome.cs**:
   This C# code defines a class `ValidPalindrome` with a method `IsPalindrome` that determines whether a given string is a valid palindrome. It first handles cases where the string is empty or null, considering them as valid palindromes. Then, it removes non-alphanumeric characters and converts the string to lowercase using regular expressions. After preprocessing the string, it uses two pointers (`start` and `end`) to iterate through the string from both ends, comparing characters at each position. If any pair of characters doesn't match, it returns false; otherwise, it returns true if the entire string is iterated without finding any mismatch.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.
