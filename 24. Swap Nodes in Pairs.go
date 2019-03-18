package main

func swapPairs(head *ListNode) *ListNode {
	var tmp = &ListNode{}
	tmp.Next = head
	result := tmp

	for {
		if tmp.Next == nil || tmp.Next.Next == nil {
			break
		}
		next := tmp.Next
		tmp.Next = tmp.Next.Next
		next.Next = tmp.Next.Next
		tmp.Next.Next = next
		tmp = tmp.Next
		tmp = tmp.Next
	}
	return result.Next
}

func leetcode24() {
	//leetcode24
	var lists = &ListNode{}
	l1 := lists
	l1.Val = 1
	l1.Next = &ListNode{}
	l1.Next.Val = 4
	l1.Next.Next = &ListNode{}
	l1.Next.Next.Val = 5
	swapPairs(l1)
}
