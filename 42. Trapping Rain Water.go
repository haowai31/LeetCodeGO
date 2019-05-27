package main

func trap(height []int) int {
	var result = 0
	var top = 0
	var topindex = 0
	var leasttop = 0
	for i := range height {
		if top < height[i] {
			top = height[i]
			topindex = i
		}
	}
	for i := 0; i < topindex; i++ {
		if height[i] > leasttop {
			leasttop = height[i]
		} else {
			if height[i] < leasttop {
				result += leasttop - height[i]
			}
		}
	}

	leasttop = 0
	for i := len(height) - 1; i > topindex; i-- {
		if height[i] > leasttop {
			leasttop = height[i]
		} else {
			if height[i] < leasttop {
				result += leasttop - height[i]
			}
		}
	}

	return result
}

func leetcode42() {
	var height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	println(trap(height))
}
