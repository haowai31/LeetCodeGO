package main

func search(nums []int, target int) int {
	var result = -1
	var left = 0
	var right = len(nums) - 1
	var middle = 0
	for {
		if left > right || left > len(nums) || right < 0 {
			break
		}
		middle = (left + right) / 2
		if nums[middle] == target {
			result = middle
			break
		}
		if nums[middle] < nums[right] {
			if target > nums[middle] && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			if target >= nums[left] && target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}

	}

	return result
}

func leetcode33() {
	var nums []int = []int{4, 5, 6, 7, 0, 1, 2}
	var target = 3
	println(search(nums, target))
}
