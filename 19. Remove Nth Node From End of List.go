package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h1 := head
	h2 := head
	count := 0
	for count < n {
		h1 = h1.Next
		count += 1
	}

	if h1 == nil {
		return head.Next
	}

	for h1.Next != nil {
		h1 = h1.Next
		h2 = h2.Next
	}
	h2.Next = h2.Next.Next

	return head
}
func leetcode19() {
	//leetcode19
	var first = ListNode{}
	var first2 = ListNode{}
	var first3 = ListNode{}
	var first4 = ListNode{}
	var first5 = ListNode{}
	first.Val = 1
	first2.Val = 2
	first3.Val = 3
	first4.Val = 4
	first5.Val = 5
	first.Next = &first2
	first2.Next = &first3
	first3.Next = &first4
	first4.Next = &first5

	removeNthFromEnd(&first, 2)
}
