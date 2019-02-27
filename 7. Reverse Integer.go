package main

import "math"

func reverse(x int) int {
	result := 0
	var neg bool
	if x < 0 {
		neg = true
		x = x * -1
	}
	for {
		if x == 0 {
			break
		}
		result *= 10
		result += x % 10
		x /= 10
	}
	if result > math.MaxInt32 {
		result = 0
	}
	if neg {
		result *= -1
	}

	return result
}

func leetcode7() {
	x := 123
	println(reverse(x))
}
