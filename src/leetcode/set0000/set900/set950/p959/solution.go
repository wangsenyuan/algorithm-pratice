package p959

func regionsBySlashes(grid []string) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}

	labels := make([][][]int, n)
	for i := 0; i < n; i++ {
		labels[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			labels[i][j] = make([]int, 4)
			for k := 0; k < 4; k++ {
				labels[i][j][k] = -1
			}
		}
	}

	var dfs func(i, j, k int, label int)
	dfs = func(i, j, k int, label int) {
		if i < 0 || i >= n || j < 0 || j >= m {
			return
		}
		if labels[i][j][k] == label {
			return
		}
		labels[i][j][k] = label

		if grid[i][j] == ' ' {
			for x := 1; x < 4; x++ {
				kk := (x + k) % 4
				dfs(i, j, kk, label)
			}

			if k == 0 {
				//up
				dfs(i-1, j, 2, label)
			} else if k == 1 {
				dfs(i, j+1, 3, label)
			} else if k == 2 {
				dfs(i+1, j, 0, label)
			} else {
				dfs(i, j-1, 1, label)
			}
			return
		}
		if grid[i][j] == '/' {
			if k == 0 {
				dfs(i, j, 3, label)
				dfs(i-1, j, 2, label)
			} else if k == 1 {
				dfs(i, j, 2, label)
				dfs(i, j+1, 3, label)
			} else if k == 2 {
				dfs(i, j, 1, label)
				dfs(i+1, j, 0, label)
			} else {
				dfs(i, j, 0, label)
				dfs(i, j-1, 1, label)
			}
			return
		}
		//grid[i][j] == \
		if k == 0 {
			dfs(i, j, 1, label)
			dfs(i-1, j, 2, label)
		} else if k == 1 {
			dfs(i, j, 0, label)
			dfs(i, j+1, 3, label)
		} else if k == 2 {
			dfs(i, j, 3, label)
			dfs(i+1, j, 0, label)
		} else {
			dfs(i, j, 2, label)
			dfs(i, j-1, 1, label)
		}
	}
	var label int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 4; k++ {
				if labels[i][j][k] < 0 {
					dfs(i, j, k, label)
					label++
				}
			}
		}
	}
	return label
}
