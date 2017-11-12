package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(findDiagonalOrder(matrix))
}

func findDiagonalOrder(matrix [][]int) []int {
	m := len(matrix)

	if m == 0 {
		return nil
	}

	n := len(matrix[0])

	res := make([]int, m*n)

	dirs := [][]int{{-1, 1}, {1, -1}}
	d := 0
	row, col := 0, 0
	for i := 0; i < m*n; i++ {
		res[i] = matrix[row][col]
		row += dirs[d][0]
		col += dirs[d][1]

		if row >= m {
			row = m - 1
			col += 2
			d = 1 - d
		}

		if col >= n {
			col = n - 1
			row += 2
			d = 1 - d
		}

		if row < 0 {
			row = 0
			d = 1 - d
		}
		if col < 0 {
			col = 0
			d = 1 - d
		}
	}

	return res
}
