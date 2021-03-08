package p131

func partition(s string) [][]string {
	n := len(s)
	can := make([][]bool, n)
	for i := 0; i < n; i++ {
		can[i] = make([]bool, n+1)
		can[i][0] = true
		can[i][1] = true
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := i + l - 1
			if s[i] == s[j] && can[i+1][l-2] {
				can[i][l] = true
			}
		}
	}

	var res [][]string

	var dfs func(i int, m int)

	cur := make([]string, n)

	dfs = func(i int, m int) {
		if i == n {
			tmp := make([]string, m)
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		for l := 1; i+l <= n; l++ {
			if can[i][l] {
				cur[m] = s[i : i+l]
				dfs(i+l, m+1)
			}
		}
	}

	dfs(0, 0)

	return res
}
