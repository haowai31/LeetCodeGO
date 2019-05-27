package main

import "sort"

func permuteUnique(nums []int) [][]int {
	var result = [][]int{}
	sort.Ints(nums)
	permuteUniqueHelper(nums, []int{}, &result)
	return result
}

func permuteUniqueHelper(nums []int, now []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, now)
		return
	}
	for i, v := range nums {
		if i != 0 {
			if nums[i] == nums[i-1] {
				continue
			}
		}
		tmp := append([]int{}, nums...)
		tmpv := append([]int{}, now...)
		permuteUniqueHelper(append(tmp[:i], tmp[i+1:]...), append(tmpv, v), result)
	}
}

func leetcode47() {
	var nums = []int{1, 1, 3}
	println(len(permuteUnique(nums)))
}
