package main

func removePalindromeSub(s string) int {
    if isPalindrome(s) {
		return 1
	}
	return 2
}

func isPalindrome(str string) bool {
    left := 0
    right := len(str) - 1
    for left < right {
        if str[left] != str[right] {
            return false
        }
        left++
        right--
    }
    return true
}
