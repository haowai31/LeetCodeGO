package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return pow50(x, n)
}

func pow50(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	v := pow50(x, n/2)
	v = v * v
	if n%2 == 1 {
		v = v * x
	}
	return v
}

func leetcode50() {
	var x = 2.10000
	var n = 3
	println(myPow(x, n))
}
