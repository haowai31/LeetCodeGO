package main

func min(m, n int) int {
	if m < n {
		return m
	} else {
		return n
	}
}
func abs(m int) int {
	if m < 0 {
		m = -m
	}
	return m
}

type ListNode struct {
	Val  int
	Next *ListNode
}
