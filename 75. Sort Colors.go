package main

func sortColors(nums []int) {
	var left = 0
	var right = len(nums)
	if right == 0 {
		return
	}
	var index = 0
	for index < right {
		if nums[index] == 2 {
			nums[index], nums[right-1] = nums[right-1], nums[index]
			right -= 1
		} else if nums[index] == 1 {
			index += 1
		} else if nums[index] == 0 {
			nums[index], nums[left] = nums[left], nums[index]
			left += 1
			index += 1
		}
	}
}

func leetcode75() {
	var nums = []int{2, 0, 2, 2, 2, 0}
	sortColors(nums)
}
