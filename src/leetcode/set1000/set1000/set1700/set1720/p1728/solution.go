package p1728

import "sort"

func largestSubmatrix(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	// m * n <= 100000
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			grid[i][j] = matrix[i][j]
			if i > 0 && matrix[i][j] == 1 {
				grid[i][j] += grid[i-1][j]
			}
		}
	}
	var res int
	row := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			row[j] = grid[i][j]
		}
		sort.Ints(row)
		for j := 0; j < n; j++ {
			tmp := row[j] * (n - j)
			res = max(res, tmp)
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
