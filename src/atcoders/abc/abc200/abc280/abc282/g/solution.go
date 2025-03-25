package main

import "fmt"

func main() {
	var n, k, p int
	fmt.Scan(&n, &k, &p)

	res := solve(n, k, p)

	fmt.Println(res)
}

func solve(n int, k int, mod int) int {
	add := func(nums ...int) int {
		res := 0
		for _, num := range nums {
			res += num
			if res >= mod {
				res -= mod
			}
		}
		return res
	}

	dp := make([][][]int, k+1)
	ndp := make([][][]int, k+1)
	for x := 0; x <= k; x++ {
		dp[x] = make([][]int, n)
		ndp[x] = make([][]int, n)
		for i := 0; i < n; i++ {
			dp[x][i] = make([]int, n)
			ndp[x][i] = make([]int, n)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[0][i][j] = 1
		}
	}

	ct := make([][][]int, k+1)
	for x := 0; x <= k; x++ {
		ct[x] = make([][]int, n+1)
		for i := 0; i <= n; i++ {
			ct[x][i] = make([]int, n+1)
		}
	}

	get := func(id int, lx int, rx int, ly int, ry int) int {
		if lx >= rx || ly >= ry {
			return 0
		}

		return add(ct[id][rx][ry], mod-ct[id][lx][ry], mod-ct[id][rx][ly], ct[id][lx][ly])
	}

	for t := 1; t < n; t++ {
		for s := 0; s <= min(t, k); s++ {
			for i := 0; i < n+1-t; i++ {
				for j := 0; j < n+1-t; j++ {
					ct[s][i+1][j+1] = add(ct[s][i][j+1], ct[s][i+1][j], mod-ct[s][i][j], dp[s][i][j])
				}
			}
		}

		for ns := 0; ns <= min(t, k); ns++ {
			for i := 0; i < n-t; i++ {
				for j := 0; j < n-t; j++ {
					ndp[ns][i][j] = add(get(ns, 0, i+1, j+1, n+1-t), get(ns, i+1, n+1-t, 0, j+1))
					if ns > 0 {
						ndp[ns][i][j] = add(ndp[ns][i][j], get(ns-1, 0, i+1, 0, j+1), get(ns-1, i+1, n+1-t, j+1, n+1-t))
					}
				}
			}
		}

		dp, ndp = ndp, dp
	}

	return dp[k][0][0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
