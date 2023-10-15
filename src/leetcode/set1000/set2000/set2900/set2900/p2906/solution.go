package p2906

func constructProductMatrix(grid [][]int) [][]int {

	n := len(grid)
	m := len(grid[0])

	pref := make([]int, n*m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cur := i*m + j
			if cur > 0 {
				pref[cur] = mul(pref[cur-1], grid[i][j])
			} else {
				pref[cur] = grid[i][j] % mod
			}
		}
	}
	suf := 1
	ans := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		ans[i] = make([]int, m)
		for j := m - 1; j >= 0; j-- {
			cur := i*m + j
			if cur > 0 {
				ans[i][j] = mul(suf, pref[cur-1])
			} else {
				ans[i][j] = suf
			}
			suf = mul(suf, grid[i][j])
		}
	}

	return ans
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

const mod = 12345
