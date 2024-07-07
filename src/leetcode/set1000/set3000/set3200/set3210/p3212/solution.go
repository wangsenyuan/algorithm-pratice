package p3212

func numberOfSubmatrices(grid [][]byte) int {
	// cnt[i][j] = 0的那些
	n := len(grid)
	m := len(grid[0])
	cnt := make([][]int, n)
	has := make([][]bool, n)

	for i := 0; i < n; i++ {
		cnt[i] = make([]int, m)
		has[i] = make([]bool, m)

		for j := 0; j < m; j++ {
			var add int
			if grid[i][j] == 'X' {
				add++
			} else if grid[i][j] == 'Y' {
				add--
			}
			cnt[i][j] = add
			if add == 1 {
				has[i][j] = true
			}
			if i > 0 {
				cnt[i][j] += cnt[i-1][j]
				has[i][j] = has[i][j] || has[i-1][j]
			}
			if j > 0 {
				cnt[i][j] += cnt[i][j-1]
				has[i][j] = has[i][j] || has[i][j-1]
			}
			if i > 0 && j > 0 {
				cnt[i][j] -= cnt[i-1][j-1]
			}
		}
	}

	var res int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if cnt[i][j] == 0 && has[i][j] {
				res++
			}
		}
	}
	return res
}
