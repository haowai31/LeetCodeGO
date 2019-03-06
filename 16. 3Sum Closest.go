package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var result = math.MaxInt32
	var sum = 0

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			tmp := nums[i] + nums[left] + nums[right]
			if target == tmp {
				return tmp
			} else {
				x := abs(target - tmp)
				if x < result {
					result = x
					sum = tmp
				}
				if target > tmp {
					left += 1
				} else {
					right -= 1
				}
			}
		}
	}
	return sum
}

func leetcode16() {
	//leetcode16
	var nums = []int{-1, 2, 1, 4}
	println(threeSumClosest(nums, 1))
}
