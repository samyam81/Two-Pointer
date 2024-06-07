package main

func isStrobogrammatic(num string) bool {
    strobogrammaticMap := map[byte]byte{
        '0': '0',
        '1': '1',
        '6': '9',
        '8': '8',
        '9': '6',
    }

    left, right := 0, len(num)-1
    for left <= right {
        if strobogrammaticMap[num[left]] != num[right] {
            return false
        }
        left++
        right--
    }
    return true
}