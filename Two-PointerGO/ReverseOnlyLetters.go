package main

import "unicode"

func reverseOnlyLetters(s string) string {
    var left int = 0
    var right int = len(s) - 1
    charArray := []rune(s)

    for left < right {
        if unicode.IsLetter(charArray[left]) && unicode.IsLetter(charArray[right]) {
            temp := charArray[left]
            charArray[left] = charArray[right]
            charArray[right] = temp
            left++
            right--
        } else if !unicode.IsLetter(charArray[left]) {
            left++
        } else if !unicode.IsLetter(charArray[right]) {
            right--
        }
    }

    return string(charArray)
}
