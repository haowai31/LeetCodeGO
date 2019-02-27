package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0{
		return true
	}

	var p []int
	y := x
	for {
		if y == 0{
			break
		}
		p = append(p, y % 10)
		y = y / 10
	}
	left, right := 0, len(p) - 1
	for {
		if left > right {
			return true
		}
		if p[left] != p[right] {
			return false
		}
		left += 1
		right -= 1
	}
}

func leetcode9() {
	x := 12
	println(isPalindrome(x))
}