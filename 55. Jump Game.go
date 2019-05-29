package main

func canJump(nums []int) bool {
	var last = 0
	var i = 0
	for i = 0; i <= last && i < len(nums); i++ {
		if last < nums[i]+i {
			last = nums[i] + i
		}
	}
	return i == len(nums)
}

func leetcode55() {
	var nums = []int{2, 3, 1, 1, 4}
	println(canJump(nums))
}
