package main

import "math"
func reverseStr(s string, k int) string {
    ch := []rune(s)
	for i := 0; i < len(s)-1; i=i+(k*2) {
		left:=i
		right:=int(math.Min(float64(i+k-1),float64(len(s)-1)))
		for left<right {
			temp:=ch[left]
			ch[left]=ch[right]
			left++
			ch[right]=temp
			right--
		}
	}

	return string(ch)
}