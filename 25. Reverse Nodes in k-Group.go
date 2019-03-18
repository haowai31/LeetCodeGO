package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	var result = &ListNode{}
	var link = result
	link.Next = head
	for {
		var tmp = []*ListNode{}
		onece := link
		for i := 0; i < k; i++ {
			if onece.Next != nil {
				tmp = append(tmp, onece.Next)
				onece = onece.Next
			} else {
				return result.Next
			}
		}
		last := tmp[k-1].Next
		link.Next = tmp[k-1]
		for i := k - 1; i > 0; i-- {
			tmp[i].Next = tmp[i-1]
		}
		tmp[0].Next = last
		link = tmp[0]
	}
}
func leetcode25() {
	//leetcode25
	l1 := &ListNode{}
	l1.Val = 1
	l1.Next = &ListNode{}
	l1.Next.Val = 4
	l1.Next.Next = &ListNode{}
	l1.Next.Next.Val = 5
	l1.Next.Next.Next = &ListNode{}
	l1.Next.Next.Next.Val = 6

	reverseKGroup(l1, 2)
}
