package main


func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	bytes := make([]byte, len(s))
	index := 0
	cycleLen := numRows * 2 - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j + i < len(s); j = j + cycleLen {
			bytes[index] = s[i + j]
			index++
			if (i != 0 && i != numRows - 1 && j + cycleLen - i < len(s)) {
				bytes[index] = s[j + cycleLen - i]
				index++
			}
		}
	}
	return string(bytes)
}

func leetcode6()  {
	var s = "PAYPALISHIRING"
	var numRows = 3
	println(convert(s, numRows))
}