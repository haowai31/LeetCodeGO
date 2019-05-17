package main

func searchRange(nums []int, target int) []int {
	var result []int = []int{-1, -1}
	var left = 0
	var right = len(nums) - 1
	var index = -1
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			index = mid
			break
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if index != -1 {
		left = index
		for {
			if left < 0 {
				break
			}
			if nums[left] == target {
				result[0] = left
			}
			left -= 1
		}
		right = index
		for {
			if right >= len(nums) {
				break
			}
			if nums[right] == target {
				result[1] = right
			}
			right += 1
		}
	}

	return result
}

func leetcode34() {
	var nums []int = []int{5, 7, 7, 8, 8, 10}
	var target = 8
	println(searchRange(nums, target))
}
