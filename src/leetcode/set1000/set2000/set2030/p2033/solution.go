package p2033

import "sort"

func minOperations(grid [][]int, x int) int {
	m := len(grid)
	n := len(grid[0])
	cnt := make([]int, m*n)
	r := grid[0][0] % x
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j]%x != r {
				return -1
			}
			cnt[i*n+j] = grid[i][j] / x
		}
	}
	// so final answer would be some p * x + r
	sort.Ints(cnt)
	median := cnt[m*n/2]
	var res int
	for i := 0; i < len(cnt); i++ {
		res += abs(cnt[i] - median)
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
