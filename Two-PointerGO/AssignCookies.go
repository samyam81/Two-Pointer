package main

import "sort"

func findContentChildren(g []int, s []int) int {
    sort.Ints(s)
	sort.Ints(g)

	gindex:=0
	gSzindex:=len(g)
	sSz:=len(s)

	for index := 0; index < sSz; index++ {
		if  gindex < gSzindex && g[gindex] <= s[index] {
			gindex++
		}
	}

	return gindex;
}