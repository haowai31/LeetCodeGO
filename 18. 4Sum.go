package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	var result [][]int = [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1
			store_left := -1
			for left < right {
				if store_left != -1 {
					for nums[left] == nums[store_left] {
						left += 1
					}
					if left >= right {
						break
					}
				}
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					tmp := []int{nums[i], nums[j], nums[left], nums[right]}
					result = append(result, tmp)
					store_left = left
					left += 1
					right -= 1
				} else {
					if sum > target {
						right -= 1
					} else {
						left += 1
					}
				}
			}
		}
	}
	return result
}
func leetcode18() {
	//leetcode18
	var nums = []int{2, 0, 3, 0, 1, 2, 4}
	var target = 7
	println(fourSum(nums, target))
}
