package p3224

import "slices"

func maximumScore1(grid [][]int) int64 {
	n := len(grid)
	if n == 1 {
		return 0
	}
	if n == 2 {
		res := []int{0, 0}
		for i := 0; i < n; i++ {
			res[0] += grid[i][0]
			res[1] += grid[i][1]
		}
		return int64(max(res[0], res[1]))
	}
	m := n + 1
	dp := make([]int, m*m)
	// first two cols
	var sum int
	for a := 0; a <= n; a++ {
		if a > 0 {
			sum += grid[a-1][1]
		}
		tmp := sum
		for b := 0; b <= n; b++ {
			if b > 0 && b <= a {
				tmp -= grid[b-1][1]
			}
			if b > a {
				tmp += grid[b-1][0]
			}
			if b > 0 {
				tmp += grid[b-1][2]
			}
			dp[a*m+b] = tmp
		}
	}
	fp := make([]int, m*m)
	for j := 2; j < n; j++ {
		for a := 0; a <= n; a++ {
			for b := 0; b <= n; b++ {
				var add int
				for c := 0; c <= n; c++ {
					if c > 0 && c <= b {
						// 这个格子被涂成了黑色，且它之前被计算了
						add -= grid[c-1][j]
					}
					if c > max(a, b) {
						add += grid[c-1][j-1]
					}
					if j+1 < n && c > 0 {
						add += grid[c-1][j+1]
					}
					fp[b*m+c] = max(fp[b*m+c], dp[a*m+b]+add)
				}
			}
		}
		copy(dp, fp)
		clear(fp)
	}

	return int64(slices.Max(dp))
}

func maximumScore(grid [][]int) (ans int64) {
	n := len(grid)
	colSum := make([][]int64, n)
	for j := range colSum {
		colSum[j] = make([]int64, n+1)
		for i, row := range grid {
			colSum[j][i+1] = colSum[j][i] + int64(row[j])
		}
	}

	f := make([][][2]int64, n)
	for j := range f {
		f[j] = make([][2]int64, n+1)
	}
	for j := 0; j < n-1; j++ {
		// 用前缀最大值优化
		preMax := f[j][0][1] - colSum[j][0]
		for pre := 1; pre <= n; pre++ {
			f[j+1][pre][0] = max(f[j][pre][0], preMax+colSum[j][pre])
			f[j+1][pre][1] = f[j+1][pre][0]
			preMax = max(preMax, f[j][pre][1]-colSum[j][pre])
		}

		// 用后缀最大值优化
		sufMax := f[j][n][0] + colSum[j+1][n]
		for pre := n - 1; pre > 0; pre-- {
			f[j+1][pre][0] = max(f[j+1][pre][0], sufMax-colSum[j+1][pre])
			sufMax = max(sufMax, f[j][pre][0]+colSum[j+1][pre])
		}

		// 单独计算 pre=0 的状态
		f[j+1][0][0] = sufMax                      // 无需考虑 f[j][0][0]，因为不能连续三列全白
		f[j+1][0][1] = max(f[j][0][0], f[j][n][0]) // 第 j 列要么全白，要么全黑
	}

	for _, row := range f[n-1] {
		ans = max(ans, row[0])
	}
	return ans
}
