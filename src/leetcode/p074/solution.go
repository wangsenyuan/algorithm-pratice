package main

import (
	"fmt"
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	i := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	})

	if i == 0 || i > len(matrix) {
		return false
	}

	j := sort.Search(len(matrix[i-1]), func(j int) bool {
		return matrix[i-1][j] >= target
	})

	if j < 0 || j >= len(matrix[i-1]) || matrix[i-1][j] != target {
		return false
	}

	return true
}

func main() {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}

	fmt.Println(searchMatrix(matrix, 3))

	matrix = [][]int{[]int{1}}

	fmt.Println(searchMatrix(matrix, 2))
}
