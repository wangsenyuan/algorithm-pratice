package p928

func minMalwareSpread(graph [][]int, initial []int) int {
	n := len(graph)

	spe := make([]bool, n)
	for i := 0; i < len(initial); i++ {
		spe[initial[i]] = true
	}

	var dfs func(node int, removed int)
	vis := make([]bool, n)

	dfs = func(node int, removed int) {
		if !vis[node] {
			vis[node] = true
			for i := 0; i < n; i++ {
				if i != node && i != removed && graph[node][i] == 1 && !vis[i] {
					dfs(i, removed)
				}
			}
		}
	}
	var ans int
	infected := n

	for i := 0; i < len(initial); i++ {
		x := initial[i]
		for j := 0; j < n; j++ {
			vis[j] = false
		}
		for j := 0; j < len(initial); j++ {
			y := initial[j]
			if x == y || vis[y] {
				continue
			}
			dfs(y, x)
		}
		var cnt int
		for j := 0; j < n; j++ {
			if vis[j] {
				cnt++
			}
		}
		if cnt < infected || (cnt == infected && x < ans) {
			infected = cnt
			ans = x
		}
	}

	return ans
}
