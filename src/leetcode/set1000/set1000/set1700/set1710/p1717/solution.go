package p1717

func constructDistancedSequence(n int) []int {
	m := 2*(n-1) + 1
	arr := make([]int, m)
	vis := make([]bool, n+1)
	var dfs func(i int) bool

	dfs = func(i int) bool {
		if i == m {
			return true
		}
		if arr[i] > 0 {
			return dfs(i + 1)
		}
		// try put som at i
		for cur := n; cur > 0; cur-- {
			if vis[cur] {
				continue
			}
			if cur > 1 && (i+cur >= m || arr[i+cur] > 0) {
				continue
			}

			vis[cur] = true
			arr[i] = cur
			if cur > 1 {
				arr[i+cur] = cur
			}
			if dfs(i + 1) {
				return true
			}
			vis[cur] = false
			arr[i] = 0
			if cur > 1 {
				arr[i+cur] = 0
			}
		}
		return false
	}
	dfs(0)

	return arr
}
