package main


func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left := 0
	right := 0
	var set  = map[uint8]int{}
	set[s[0]] = 0
	result := 1
	for {
		if right == len(s) - 1 {
			break
		}
		right += 1
		location, ok := set[s[right]]
		if (ok) {
			for i:=left;i<location;i++ {
				delete(set, s[i])
			}
			left = location + 1
		}
		set[s[right]] = right

		if result < right - left + 1 {
			result = right - left + 1
		}
	}
	return result
}


func leetcode3() {
	//leetcode 3
	leetcode3s := "pwwkew"
	println(lengthOfLongestSubstring(leetcode3s))
}