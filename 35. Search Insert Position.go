package main

func searchInsert(nums []int, target int) int {
	var result = 0
	var left = 0
	var right = len(nums) - 1
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			result = middle
			break
		}
		if nums[middle] > target {
			right = middle - 1
			if target > nums[right] {
				result = right + 1
				break
			}
		} else {
			left = middle + 1
			if target < nums[left] {
				result = left
				break
			}
		}
	}

	return result
}

func leetcode35() {
	var nums []int = []int{1, 3, 5, 7}
	var target = 6
	println(searchInsert(nums, target))
}
