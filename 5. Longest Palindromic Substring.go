package main


func longestPalindrome(s string) string {

	var tmp string = "#"
	var p = []int{0}
	if len(s) == 0 {
		return ""
	}

	for i := 0;i<len(s);i++ {
		tmp += string(s[i]) + "#"
		p = append(p, 0)
		p = append(p, 0)
	}
	tmp = "%" + tmp
	p = append(p, 0)

	var mx = 0
	var id = 0
	var maxid = 0
	var max = 0
	for i := 1;i<len(tmp);i++ {
		if i < mx {
			p[i] = min(mx - i, p[2*id - i])
		} else {
			p[i] = 1
		}
		for {
			if i + p[i] >= len(tmp) {
				break
			}
			if tmp[i-p[i]] != tmp[i+p[i]] {
				break
			}
			p[i] ++
		}
		if i + p[i] > mx {
			mx = i + p[i]
			id = i
			if max < p[i] - 1 {
				max = p[i] - 1
				maxid = i
			}
		}
	}

	return s[(maxid-max+1)/2-1:(maxid + max)/2]
}


func leetcode5() {
	//leetcode 5
	var s string = "abba"
	println(longestPalindrome(s))
}