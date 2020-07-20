package p1519

func countSubTrees(n int, edges [][]int, labels string) []int {

	degree := make([]int, n)

	for _, e := range edges {
		u, v := e[0], e[1]
		degree[u]++
		degree[v]++
	}

	conn := make([][]int, n)
	for i := 0; i < n; i++ {
		conn[i] = make([]int, 0, degree[i])
	}

	for _, e := range edges {
		u, v := e[0], e[1]
		conn[u] = append(conn[u], v)
		conn[v] = append(conn[v], u)
	}

	var dfs func(p, u int) []int
	ans := make([]int, n)
	dfs = func(p, u int) []int {
		cnt := make([]int, 26)
		cnt[int(labels[u]-'a')]++
		for _, v := range conn[u] {
			if v == p {
				continue
			}
			tmp := dfs(u, v)
			for i := 0; i < 26; i++ {
				cnt[i] += tmp[i]
			}
		}
		ans[u] = cnt[int(labels[u]-'a')]
		return cnt
	}

	dfs(-1, 0)

	return ans
}
