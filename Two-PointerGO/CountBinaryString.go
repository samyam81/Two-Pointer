package main

func countBinarySubstrings(s string) int {
    var n int = len(s)
    var count uint = 0
    var ones uint = 0
    var zeros uint = 0
    flag1 := false

    for i := 0; i < n; i++ {
        if s[i] == '1' {
            ones++
            if flag1 && (i == n-1 || s[i+1] == '0') {
                count += min(ones, zeros)
                flag1 = false
                zeros = 0
            }
        } else {
            zeros++
            flag1 = true
            if flag1 && (i == n-1 || s[i+1] == '1') {
                count += min(ones, zeros)
                ones = 0
            }
        }
    }
    return int(count)
}

func min(a, b uint) uint {
    if a < b {
        return a
    }
    return b
}
