package p2327

import "sort"

const MOD = 1000000007

func modAdd(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

var dd = []int{-1, 0, 1, 0, -1}

func countPaths(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	cells := make([]Cell, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cells = append(cells, Cell{grid[i][j], i, j})
		}
	}

	sort.Slice(cells, func(i, j int) bool {
		return cells[i].value < cells[j].value
	})

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 1
		}
	}

	for _, cell := range cells {
		i, j := cell.i, cell.j
		for k := 0; k < 4; k++ {
			u, v := i+dd[k], j+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && grid[u][v] < grid[i][j] {
				modAdd(&dp[i][j], dp[u][v])
			}
		}
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			modAdd(&res, dp[i][j])
		}
	}
	return res
}

type Cell struct {
	value int
	i     int
	j     int
}
