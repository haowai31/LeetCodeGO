package main

import "math"

func myAtoi(str string) int {
	var result int
	var neg bool
	
	if len(str) == 0 {
		return 0
	}

	var index = 0
	for {
		if index > len(str) - 1{
			return 0
		}
		if string(str[index]) != " " {
			break
		}
		index += 1
	}

	if str[index] == "-"[0] {
		neg = true
		index += 1
	} else {
		if str[index] == "+"[0] {
			index += 1
		}
	}

	if index > len(str) - 1{
		return 0
	}

	if str[index] < 48 || str[index] > 57 {
		return 0
	}

	for {
		if index > len(str) - 1 {
			break
		}
		if str[index] == " "[0] {
			break
		}

		if str[index] < 48 || str[index] > 57 {
			break
		}

		if result > math.MaxInt32 {
			if neg {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		result *= 10
		result += int(str[index] - 48)
		index += 1
	}

	if neg {
		result *= -1
		if result < math.MinInt32 {
			result = math.MinInt32
		}
	} else {
		if result > math.MaxInt32 {
			result = math.MaxInt32
		}
	}

	return result
}

func leetcode8() {
	var s = "9223372036854775808"

	println(myAtoi(s))
}