package main

func searchMatrix(matrix [][]int, target int) bool {
	var m = len(matrix)
	if m == 0 {
		return false
	}
	var n = len(matrix[0])
	if n == 0 {
		return false
	}
	var right = m
	var left = 0
	var indexy = 0
	if target > matrix[right-1][0] {
		indexy = right - 1
		left = right
	}
	for left < right {

		var mid = (left + right) / 2
		if matrix[mid][0] == target {
			return true
		}
		if left != m-1 {
			if target > matrix[left][0] && target < matrix[left+1][0] {
				indexy = left
				break
			}
		}
		if matrix[mid][0] > target {
			right = mid
		} else {
			if target > matrix[mid][0] && target < matrix[mid+1][0] {
				indexy = mid
				break
			}
			left = mid + 1
		}

	}

	if target > matrix[indexy][n-1] {
		return false
	}
	left = 0
	right = n
	for left < right {
		var mid = (right + left) / 2
		if matrix[indexy][mid] == target {
			return true
		}
		if matrix[indexy][mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return false
}

func leetcode74() {
	var matrix = [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50}}

	println(searchMatrix(matrix, 11) == false)
}
