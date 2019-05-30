package main

func plusOne(digits []int) []int {
	var result = []int{}
	if len(digits) == 0 {
		return result
	}
	var step = 1
	for i := range digits {
		var tmp = digits[len(digits)-1-i] + step
		step = tmp / 10
		tmp = tmp % 10
		result = append([]int{tmp}, result...)
	}
	if step == 1 {
		result = append([]int{1}, result...)
	}
	return result
}

func leetcode66() {
	var digits = []int{}
	println(plusOne(digits))
}
