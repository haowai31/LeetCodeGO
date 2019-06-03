package main

func mySqrt(x int) int {
	var result = 0
	if x == 0 {
		return 0
	}
	for {
		if result*result >= x {
			if result*result == x {
				break
			} else {
				result -= 1
				break
			}
		}
		result += 1
	}
	return result
}
func leetcode69() {
	var x = 8
	println(mySqrt(x))
}
