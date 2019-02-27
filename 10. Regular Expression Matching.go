package main

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		} else {
			return false
		}
	}

	if len(p) == 1{
		return len(s) == 1 && (p == "." || p == s)
	}
	if p[1] != "*"[0] {
		if len(s) == 0 {
			return false
		}
		return (s[0] == p[0] || p[0] == "."[0]) && isMatch(s[1:], p[1:])
	}
	for {
		if !(len(s) != 0 && (s[0] == p[0] || p[0] == "."[0])) {
			break
		}
		if isMatch(s, p[2:]) {
			return true
		}
		s = s[1:]
	}
	return isMatch(s, p[2:])
}

func leetcode10() {
	s := "aaa"
	p := "ab*a*c*a"
	println(isMatchDP(s, p))
}