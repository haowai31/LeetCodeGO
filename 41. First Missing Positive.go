package main

import "sort"

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return 1
	}
	var count = 1
	for i := range nums {
		if nums[i] <= 0 || nums[i] == count-1 {
			continue
		} else {
			if nums[i] != count {
				return count
			}
			count += 1
		}
	}

	if nums[len(nums)-1]+1 <= 0 {
		return 1
	}
	return nums[len(nums)-1] + 1
}

func leetcode41() {
	var nums = []int{1, 2, 0}
	println(firstMissingPositive(nums))
}
