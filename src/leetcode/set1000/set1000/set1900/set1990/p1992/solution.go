package p1992

var dd = []int{-1, 0, 1, 0, -1}

func findFarmland(land [][]int) [][]int {
	m := len(land)
	n := len(land[0])

	vis := make([][]bool, m)

	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	var x2, y2 int
	var dfs func(i, j int)
	dfs = func(i, j int) {
		vis[i][j] = true
		if i > x2 {
			x2 = i
		}
		if j > y2 {
			y2 = j
		}
		for k := 0; k < 4; k++ {
			x, y := i+dd[k], j+dd[k+1]
			if x >= 0 && x < m && y >= 0 && y < n && !vis[x][y] && land[x][y] == 1 {
				dfs(x, y)
			}
		}
	}
	var res [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 1 && !vis[i][j] {
				x2, y2 = i, j
				dfs(i, j)
				res = append(res, []int{i, j, x2, y2})
			}
		}
	}
	return res
}
