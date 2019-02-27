package main

func maxArea(height []int) int {
	var left, right = 0, len(height)-1
	var result = 0
	for {
		if left > right {
			break
		}
		var tmp int
		if height[left] < height[right] {
			tmp = height[left]
		} else {
			tmp = height[right]
		}
		var sum = (right - left) * tmp
		if result < sum {
			result = sum
		}
		if height[left] < height[right] {
			left += 1
			continue
		}else {
			right -= 1
			continue
		}
	}
	return result
}

func leetcode11() {
	//leetcode11
	var x = []int{1,8,6,2,5,4,8,3,7}
	println(maxArea((x)))
}