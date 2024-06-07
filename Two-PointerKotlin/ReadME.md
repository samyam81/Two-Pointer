# Two-Pointer Kotlin

This repository contains Kotlin solutions for various problems implemented using the Two Pointer algorithm approach.

## Prerequisites

To run or contribute to the code in this repository, you should have:

- Basic knowledge of Kotlin programming language.
- Understanding of basic data structures such as arrays and HashMaps.
- Familiarity with algorithms like Two Pointer.

## Contents

1. **boat.kt**: 
   This code defines a function `numRescueBoats` that takes an array of people's weights and a weight limit for each boat. It sorts the array in ascending order, then uses a two-pointer approach to iterate through the array from both ends, attempting to fit as many people as possible onto each boat while respecting the weight limit. It returns the minimum number of boats needed to rescue all the people.

2. **FirstIndex.kt**:
   This code defines a function `strStr` that implements a simple algorithm to find the index of the first occurrence of a substring (needle) within a string (haystack). It iterates through the haystack string, checking if each substring starting at the current index matches the needle. If a match is found, it returns the index of the start of the matching substring; otherwise, it returns -1.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.
