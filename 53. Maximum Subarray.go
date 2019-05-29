package main

import "math"

func maxSubArray(nums []int) int {
	var result = 0
	var sum = math.MinInt32

	for i := range nums {
		result = result + nums[i]

		if sum < result {
			sum = result
		}
		if result < 0 {
			result = 0
		}
	}
	return sum
}

func leetcode53() {
	var nums = []int{-2}
	println(maxSubArray(nums))
}
