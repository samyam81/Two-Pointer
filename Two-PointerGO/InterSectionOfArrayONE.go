package main

func intersection(nums1 []int, nums2 []int) []int {
    if len(nums1)==0 || len(nums2)==0 {
		return make([]int, 0)
	}

	Old:=make([]int, 1000)
	for i:= range nums1{
		Old[nums1[i]]++
	}

	New:=make([]int,0)
	for i:= range nums2 {
		if Old[nums2[i]]>0 {
			New = append(New, nums2[i])
			Old[nums2[i]]=0
		}
	}
	return New
}