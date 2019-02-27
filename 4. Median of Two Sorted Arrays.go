package main



func findKth(nums1 []int, m int, nums2 []int, n int, k int) float64 {
	if m > n {
		return findKth(nums2, n, nums1, m, k)
	}

	if m == 0 {
		return float64(nums2[k - 1])
	}

	if (k == 1) {
		return float64(min(nums2[0], nums1[0]))
	}

	pa := min(int(k / 2), int(m))
	pb := k - pa
	if nums1[pa-1] < nums2[pb-1] {
		return findKth(nums1[pa:], m - pa, nums2, n, k - pa)
	} else if nums1[pa-1] > nums2[pb-1] {
		return findKth(nums1, m, nums2[pb:], n - pb, k - pb)
	} else {
		return float64(nums1[pa - 1])
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	total := m + n
	if total % 2 == 1 {
		return findKth(nums1, m, nums2, n, total / 2 + 1)
	} else {
		return (findKth(nums1, m, nums2, n, total / 2) + findKth(nums1, m, nums2, n, total / 2 + 1)) / 2
	}

	return 0
}

func leetcode4()  {
	//leetcode 4
	var nums1 = []int{1, 3}
	var nums2 = []int{2, 4}
	println(findMedianSortedArrays(nums1, nums2))
}