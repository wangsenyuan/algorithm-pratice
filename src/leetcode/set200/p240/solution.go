package main

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	i := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	})

	if i <= 0 {
		return false
	}

	for k := 0; k < i; k++ {
		j := sort.SearchInts(matrix[k], target)
		if j >= 0 && j < len(matrix[k]) && matrix[k][j] == target {
			return true
		}
	}

	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1

	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}

	return false
}
