package main

import "strconv"

func addBinary(a string, b string) string {
	var step = 0
	if len(a) > len(b) {
		a, b = b, a
	}
	if len(b) == 0 {
		return ""
	}
	var index = 0
	for index < len(b) {
		var tmp = 0
		if index < len(a) {
			tmp += int(a[len(a)-index-1]) - 48
		}
		tmp += step + int(b[len(b)-index-1]) - 48
		step = tmp / 2
		tmp = tmp % 2
		b = b[:(len(b)-1-index)] + strconv.Itoa(tmp) + b[(len(b)-index):]
		index += 1
	}
	if step == 1 {
		b = "1" + b
	}
	return b

}

func leetcode67() {
	var a = "11"
	var b = "1101"
	println(addBinary(a, b))
}
