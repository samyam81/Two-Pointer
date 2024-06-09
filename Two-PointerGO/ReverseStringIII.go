package main

func reverseWords(s string) string {
    chars := []rune(s)
    
    reverse := func(start, end int) {
        for start < end {
            chars[start], chars[end] = chars[end], chars[start]
            start++
            end--
        }
    }
    
    start, end := 0, 0
    
    for end < len(chars) {
        for end < len(chars) && chars[end] != ' ' {
            end++
        }
        reverse(start, end-1)
        start = end + 1
        end = start
    }
    
    return string(chars)
}
