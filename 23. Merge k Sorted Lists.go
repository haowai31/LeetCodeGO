package main

func mergeKListsInsert(listfirst []int, x int) {

}

func insertlistnode(listfirst []int, target int) int {
	var result = 0
	flag := 0
	for j := 0; j < len(listfirst); j++ {
		if target <= listfirst[j] {
			result = j
			flag = 1
			break
		}
	}
	if flag == 0 {
		result = len(listfirst)
	}
	return result
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head = &ListNode{}
	var result = head
	var listfirst = []int{}
	var listfirstpoint = []int{}
	var pointlist = []*ListNode{}
	for count, i := range lists {
		pointlist = append(pointlist, i)

		if i != nil {
			if len(listfirst) == 0 {
				listfirst = append(listfirst, i.Val)
				listfirstpoint = append(listfirstpoint, count)
			} else {
				point := insertlistnode(listfirst, i.Val)
				if point != len(listfirstpoint) {
					listfirst = append(listfirst[:point+1], listfirst[point:]...)
					listfirst[point] = i.Val
					listfirstpoint = append(listfirstpoint[:point+1], listfirstpoint[point:]...)
					listfirstpoint[point] = count
				} else {
					listfirst = append(listfirst, i.Val)
					listfirstpoint = append(listfirstpoint, count)
				}
			}
		}
	}

	for {
		if len(listfirst) == 0 {
			break
		}
		var tmp = ListNode{}
		tmp.Val = listfirst[0]
		head.Next = &tmp
		head = head.Next
		tar := pointlist[listfirstpoint[0]]
		if tar.Next != nil {
			tar = tar.Next
			pointlist[listfirstpoint[0]] = tar
			target := tar.Val
			point := insertlistnode(listfirst, target)
			if point != len(listfirstpoint) {
				listfirst = append(listfirst[:point+1], listfirst[point:]...)
				listfirst[point] = target
				listfirstpoint = append(listfirstpoint[:point+1], listfirstpoint[point:]...)
				listfirstpoint[point] = listfirstpoint[0]
			} else {
				listfirst = append(listfirst, target)
				listfirstpoint = append(listfirstpoint, listfirstpoint[0])
			}
		}
		listfirst = listfirst[1:]
		listfirstpoint = listfirstpoint[1:]
	}

	return (result).Next
}
func leetcode23() {
	//leetcode23
	var lists = []*ListNode{}
	l1 := ListNode{}
	l1.Val = 1
	l1.Next = &ListNode{}
	l1.Next.Val = 4
	l1.Next.Next = &ListNode{}
	l1.Next.Next.Val = 5

	l2 := ListNode{}
	l2.Val = 1
	l2.Next = &ListNode{}
	l2.Next.Val = 3
	l2.Next.Next = &ListNode{}
	l2.Next.Next.Val = 4

	l3 := ListNode{}
	l3.Val = 2
	l3.Next = &ListNode{}
	l3.Next.Val = 6
	lists = append(lists, &l1)
	lists = append(lists, &l2)
	lists = append(lists, &l3)
	mergeKLists(lists)
}
