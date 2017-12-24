package p753

import "strconv"

func crackSafe(n int, k int) string {
	N := pow(k, n)

	visited := make(map[int]bool)

	base := 1
	for i := 1; i < n; i++ {
		base *= 10
	}

	var dfs func(prev int) bool
	var res string
	dfs = func(prev int) bool {
		if len(visited) == N {
			return true
		}
		for i := 0; i < k; i++ {
			cur := (prev%base)*10 + i
			if !visited[cur] {
				visited[cur] = true
				res += strconv.Itoa(i)
				if dfs(cur) {
					return true
				}
				delete(visited, cur)
				res = res[:len(res)-1]
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		res += strconv.Itoa(0)
	}
	visited[0] = true
	dfs(0)
	return res
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	c := pow(a, b/2)
	d := c * c
	if b%2 == 1 {
		return a * d
	}
	return d
}
