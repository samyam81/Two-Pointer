package main

import "strings"

func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, r)
}

func reverseVowels(s string) string {
	charArray := []rune(s)

	left := 0
	right := len(s) - 1

	for left < right {
		if isVowel(charArray[left]) && isVowel(charArray[right]) {
			temp := charArray[left]
			charArray[left] = charArray[right]
			charArray[right] = temp
			left++
			right--
		} else if !isVowel(charArray[left]) {
			left++
		} else if !isVowel(charArray[right]) {
			right--
		}
	}

	return string(charArray)
}

