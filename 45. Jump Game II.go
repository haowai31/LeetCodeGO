package main

func jump(nums []int) int {
	var e = 0
	var max = 0
	var result = 0
	for i := 0; i < len(nums)-1; i++ {
		if max < i+nums[i] {
			max = i + nums[i]
		}
		if i == e {
			result += 1
			e = max
		}
	}
	return result
}

func leetcode45() {
	var nums = []int{2, 3, 1, 1, 4}
	println(jump(nums))
}
