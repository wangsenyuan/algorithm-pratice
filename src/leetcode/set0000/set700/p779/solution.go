package p779

func kthGrammar(N int, K int) int {
	var dfs func(n, k int) int

	dfs = func(n, k int) int {
		if k == 0 {
			return 0
		}

		if n == 0 {
			return 0
		}

		tmp := dfs(n-1, k/2)
		if k%2 == 1 {
			return 1 - tmp
		}
		return tmp
	}

	return dfs(N-1, K-1)
}
