package main

func rotateRight(head *ListNode, k int) *ListNode {
	var result = &ListNode{}
	result.Next = head
	if head == nil {
		return head
	}
	var count = 1
	for head.Next != nil {
		count += 1
		head = head.Next
	}
	if count <= 1 {
		return head
	}
	k = k % count
	if k == 0 {
		return result.Next
	}

	var right = &ListNode{}
	right = result.Next
	for x := 0; x < (count - k - 1); x++ {
		right = right.Next
	}
	left := result.Next
	result.Next = right.Next
	head.Next = left
	right.Next = nil
	return result.Next
}

func leetcode61() {
	l1 := &ListNode{}
	l1.Val = 1
	l1.Next = &ListNode{}
	l1.Next.Val = 4
	l1.Next.Next = &ListNode{}
	l1.Next.Next.Val = 5
	l1.Next.Next.Next = &ListNode{}
	l1.Next.Next.Next.Val = 6

	println(rotateRight(l1, 0))
}
