# Two-Pointer Go

This repository contains GO solutions for various problems implemented using the Two Pointer algorithm approach.

## Prerequisites

To run or contribute to the code in this repository, you should have:

- Basic knowledge of GO programming language.
- Understanding of basic data structures such as arrays and HashMaps.
- Familiarity with algorithms like Two Pointer.

## Contents

1. **AssignCookies.go**:

   This Go code defines a function `findContentChildren` which takes two integer slices `g` and `s` representing the greed factors of children and the sizes of available cookies respectively. It sorts both slices in ascending order. Then, it iterates through the sizes of cookies. Within each iteration, it checks if there are still children to be satisfied (`gindex < gSzindex`) and if the current child's greed factor can be satisfied with the current cookie (`g[gindex] <= s[index]`). If so, it increments the index of satisfied children (`gindex`). Finally, it returns the count of satisfied children. The algorithm aims to maximize the number of satisfied children by assigning cookies to them in a way that their greed factors are accommodated.

2. **BadVersion.go**:

   This Go code implements a binary search algorithm to find the first occurrence of a "bad version" within a range of versions. The function `firstBadVersion` takes an integer `n`, representing the total number of versions, and returns an integer representing the index of the first bad version. 

   Within the binary search loop, it continually updates the search range (`ss81` and `kdb`) until they converge, indicating the first bad version is found. At each iteration, it calculates the middle index (`mid`) of the current search range and calls the `isBadVersion` function to determine if the version is bad. Based on the result, it updates the search range accordingly to narrow down the search space until the first bad version is found. 

   The `isBadVersion` function is declared but not implemented, and it's expected to return a boolean indicating whether a given version is bad or not.

3. **CountBinaryString.go**:

   This Go code defines a function `countBinarySubstrings` that takes a string `s` composed of '0's and '1's and counts the number of non-empty (contiguous) substrings that have equal number of consecutive '0's and '1's. 

   The function iterates through each character of the string `s`, maintaining counts of consecutive '1's (`ones`) and consecutive '0's (`zeros`). It also uses a boolean flag `flag1` to track the last character seen. Whenever it encounters a character different from the previous one, it calculates the number of valid substrings that can be formed using the current sequence of consecutive '0's and '1's, and updates the count accordingly. 

   The `min` function simply returns the minimum of two unsigned integers. 

   Overall, the algorithm efficiently counts the number of valid substrings by keeping track of consecutive sequences of '0's and '1's.

4. **InterSectionOfArrayONE.go**:

   The provided Go function `intersection` efficiently computes the intersection of two integer slices `nums1` and `nums2`. It first constructs a frequency table for elements in `nums1`, then iterates through `nums2` to identify common elements. Common elements are appended to a new slice named `New`, and their counts in the frequency table are reset to avoid duplicates. The function ultimately returns the `New` slice containing the intersection of the two input slices.

5. **IntersectionOfTwoArrayTWO.go**:

   The `intersect` function efficiently computes the intersection of two integer slices `nums1` and `nums2`. It utilizes a frequency map to track the occurrence of each element in `nums1`, then iterates through `nums2` to identify common elements. Common elements are appended to the result slice while reducing their counts in the frequency map to handle duplicates. Finally, the function returns the computed intersection slice.

6. *NimGame.go*:

   The `canWinNim` function in this Go code determines whether a player can win the game of Nim with `n` stones based on the optimal strategy. If the number of stones is divisible by 4, the function returns `false`, indicating that the player cannot win because no matter how many stones they take initially, their opponent can always take the remaining stones in a way that leaves a multiple of 4 stones for the player to take. Otherwise, it returns `true`, indicating that the player can win by taking an appropriate number of stones in the initial move to leave the opponent with a multiple of 4 stones.

7. **RemovePalindrome.go**:

   The `removePalindromeSub` function in this Go code determines the minimum number of removals needed to make the string `s` into a palindrome. If `s` is already a palindrome, it returns 1, as no removals are required. Otherwise, it returns 2, indicating that at most two removals are needed to transform the string into a palindrome.

   The `isPalindrome` function checks whether a given string is a palindrome by comparing characters from both ends of the string towards the center. If characters from opposite ends ever differ, the function returns `false`; otherwise, it returns `true`.

8. **ReverseOnlyLetters.go**:

   The `reverseOnlyLetters` function in this Go code reverses only the letters in the given string `s`, preserving the positions of non-letter characters. It iterates through the string from both ends, swapping letters encountered on opposite sides, while leaving non-letter characters untouched. The function utilizes the `unicode.IsLetter` function to determine whether a character is a letter or not. Finally, it returns the modified string with letters reversed and non-letter characters unchanged.

9. **ReserveStringII.go**:

   The `reverseStr` function in this Go code reverses substrings of length `k` within the given string `s`, alternating between each `2k` segment. It iterates through the string by `2k` increments, and within each segment, it reverses the substring of length `k`. This is achieved by swapping characters from both ends until reaching the middle of the substring. Finally, it returns the modified string with substrings of length `k` reversed.

10. **ReserveStringIII.go**:

    The `reverseWords` function in this Go code reverses the order of words in a given string `s`. It works by first converting the string to a slice of runes for in-place manipulation. Then, it defines a closure `reverse` to reverse a portion of the string between two indices. 

    The function iterates through the string, locating word boundaries delimited by spaces. For each word, it calls the `reverse` closure to reverse the characters in place. Finally, it returns the modified string with the words reversed while maintaining their relative positions.

11. **ReverseVowel.go**:

    The `reverseVowels` function in this Go code reverses the positions of vowels within a given string `s`. It first converts the string into a slice of runes for in-place manipulation. Then, it iterates through the string from both ends, locating vowels using the `isVowel` function.

    If both the character at the left and right positions are vowels, it swaps them. Otherwise, if the character at the left position is not a vowel, it advances the left pointer, and if the character at the right position is not a vowel, it decrements the right pointer.

    Finally, it returns the modified string with the positions of vowels reversed while preserving the positions of non-vowel characters.

12. **RotateList.go**:

    The `rotateRight` function in this Go code performs a right rotation on a linked list `head` by `k` positions. It first handles base cases where the list is empty or has only one node. Then, it calculates the length of the list and finds the node `temp` that will be the new head after rotation.

    Next, it divides the list into two parts at node `temp`, reverses each part separately, and then reverses the entire list to obtain the final rotated list. Finally, it returns the head of the rotated list.

    The `reverse` function reverses a linked list and returns the new head of the reversed list.

13. **SortParityONE.go**:

    The `sortArrayByParity` function in this Go code sorts an array of integers `nums` such that all even integers appear before all odd integers. It uses a two-pointer approach, where the left pointer (`left`) starts at the beginning of the array, and the right pointer (`right`) starts at the end of the array.

    While the left pointer is less than the right pointer, it checks if the integer at the left pointer is odd and the integer at the right pointer is even. If so, it swaps them to maintain the desired order. Then, it moves the left pointer to the right until it encounters an odd integer and moves the right pointer to the left until it encounters an even integer.

    Finally, it returns the modified `nums` array with even integers appearing before odd integers while maintaining the relative order within each category.

14. **StrobogrammaticNumber.go**:

    The `isStrobogrammatic` function in this Go code checks whether a given number `num` is strobogrammatic, meaning it reads the same when rotated 180 degrees. It uses a map `strobogrammaticMap` to define pairs of valid strobogrammatic digits.

    The function initializes two pointers, `left` pointing to the beginning of the string `num`, and `right` pointing to the end. It iterates through the string from both ends, comparing the digits at corresponding positions using the `strobogrammaticMap`. If any pair of digits does not match the strobogrammatic mapping, the function returns `false`. Otherwise, if all pairs match, it returns `true`, indicating that the number is strobogrammatic.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.

