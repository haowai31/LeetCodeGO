package main

import "math"

func divide(dividend int, divisor int) int {
	result := 0
	time := 1
	sum := 0
	flag := 1
	if dividend < 0 {
		dividend = -dividend
		flag = -flag
	}
	if divisor < 0 {
		divisor = -divisor
		flag = -flag
	}
	for sum < dividend {
		sum += time * divisor
		time *= 2
	}
	sum = 0
	for time > 1 {
		time /= 2
		if sum+time*divisor <= dividend {
			sum += time * divisor
			result += time
		}
	}
	result *= flag
	if result > math.MaxInt32 || result < math.MinInt32 {
		return math.MaxInt32
	}

	return result
}
func leetcode29() {
	//leetcode29
	dividend := 10
	divisor := -3
	println(divide(dividend, divisor))
}
