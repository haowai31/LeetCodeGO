package main

import "sort"

func threeSum_SLOW(nums []int) [][]int {
	var result = [][]int{}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := 0 - nums[i] - nums[j]
			if target < nums[j] {
				continue
			}
			for k := j + 1; k < len(nums) && target >= nums[k]; k++ {
				if target == nums[k] {
					tmp := []int{nums[i], nums[j], nums[k]}
					result = append(result, tmp)
					break
				}
			}
		}
	}
	return result
}
