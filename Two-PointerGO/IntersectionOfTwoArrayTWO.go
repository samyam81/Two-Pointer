package main

func intersect(nums1 []int, nums2 []int) []int {
    if len(nums1) == 0 || len(nums2) == 0 {
        return []int{}
    }

    freqMap := make(map[int]int)
    for _, num := range nums1 {
        freqMap[num]++
    }

    var intersection []int
    for _, num := range nums2 {
        if freqMap[num] > 0 {
            intersection = append(intersection, num)
            freqMap[num]--
        }
    }

    return intersection
}
