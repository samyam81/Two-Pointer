package main

func sortArrayByParity(nums []int) []int {
	left:=0
	right:=len(nums)-1

	for left<right{
		if nums[left]%2>nums[right]%2 {
			temp:=nums[left]
			nums[left]=nums[right]
			nums[right]=temp
		}
		if nums[left]%2==0 {
			left++
		}

		if nums[right]%2!=0 {
			right--
		}
	}

	return nums
}
