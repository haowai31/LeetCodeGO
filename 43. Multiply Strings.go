package main

import "strconv"

func multiply(num1 string, num2 string) string {
	var result = ""
	if num2 == "0" || num1 == "0" {
		return "0"
	}
	if len(num2) == 1 {
		var step = 0
		for j := range num1 {
			var tmp = (int(num1[len(num1)-j-1])-48)*(int(num2[0])-48) + step
			result = strconv.Itoa(tmp%10) + result
			step = tmp / 10
		}
		if step != 0 {
			result = strconv.Itoa(step) + result
		}
		return result
	}

	for i := range num2 {
		var tmp = multiply(num1, string(num2[len(num2)-i-1]))
		for j := 0; j < i; j++ {
			tmp += "0"
		}
		result = stringplus(tmp, result)
	}

	return result
}

func stringplus(num1 string, num2 string) string {
	var index = 0
	var step = 0
	if len(num1) > len(num2) {
		var tmp = num1
		num1 = num2
		num2 = tmp
	}
	index = len(num2) - 1
	for {
		if index < 0 {
			break
		}
		var tmp = int(num2[index]) + step - 48
		if len(num2)-index-1 < len(num1) {
			tmp += int(num1[len(num1)-len(num2)+index]) - 48
		}
		if tmp > 9 {
			tmp -= 10
			step = 1
		} else {
			step = 0
		}
		num2 = num2[:index] + strconv.Itoa(tmp) + num2[index+1:]
		index -= 1
	}
	if step != 0 {
		num2 = "1" + num2
	}
	return num2
}

func leetcode43() {
	var num1 = "123"
	var num2 = "456"
	println(multiply(num1, num2))
}
