package p2895

const inf = 1 << 30

func minOperations(s1 string, s2 string, x int) int {
	if s1 == s2 {
		return 0
	}
	var p []int

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			p = append(p, i)
		}
	}
	if len(p)%2 != 0 {
		return -1
	}
	f0, f1 := 0, x
	for i := 1; i < len(p); i++ {
		f0, f1 = f1, min(f1+x, f0+(p[i]-p[i-1])*2)
	}

	return f1 / 2
}

func minOperations1(s1 string, s2 string, x int) int {
	n := len(s1)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 2)
			for k := 0; k < 2; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(i int, j int, f int) int

	dfs = func(i int, j int, f int) int {
		if i < 0 {
			if j == 0 && f == 0 {
				return 0
			}
			return inf
		}
		if dp[i][j][f] >= 0 {
			return dp[i][j][f]
		}
		if s1[i] == s2[i] == (f == 0) {
			return dfs(i-1, j, 0)
		}
		res := min(dfs(i-1, j+1, 0)+x, dfs(i-1, j, 1)+1)
		if j > 0 {
			res = min(res, dfs(i-1, j-1, 0))
		}
		dp[i][j][f] = res
		return res
	}

	res := dfs(n-1, 0, 0)
	if res < inf {
		return res
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
