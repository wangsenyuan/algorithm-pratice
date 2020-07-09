package main

func main() {

}

func solve(n int, m int, E [][]int) (bool, string) {
	if n&1 == 1 {
		n++
	}
	degree := make([]int, n)
	rd := make([]int, n)
	for i := 0; i < m; i++ {
		u, v := E[i][0], E[i][1]
		u--
		v--
		// there will be edges between u ^ 1 & v, u & v ^ 1
		degree[u^1]++
		degree[v^1]++
		rd[u]++
		rd[v]++
	}

	adj := make([][]int, n)
	radj := make([][]int, n)

	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, degree[i])
		radj[i] = make([]int, 0, rd[i])
	}
	for i := 0; i < m; i++ {
		u, v := E[i][0], E[i][1]
		u--
		v--
		// there will be edges between u ^ 1 & v, u & v ^ 1
		adj[u^1] = append(adj[u^1], v)
		adj[v^1] = append(adj[v^1], u)
		radj[u] = append(radj[u], v^1)
		radj[v] = append(radj[v], u^1)
	}

	vis := make([]bool, n)
	who := make([]int, n)
	time := new(int)
	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs1(radj, vis, who, time, i)
		}
	}

	comp := make([]int, n)
	for i := 0; i < n; i++ {
		comp[i] = -1
	}

	for i := 0; i < n; i++ {
		if comp[who[i]] == -1 {
			dfs2(adj, comp, who[i], who[i])
		}
	}

	for i := 0; i < n; i += 2 {
		if comp[i] == comp[i^1] {
			return false, ""
		}
	}

	for i := 0; i < n; i++ {
		vis[i] = false
	}

	ans := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		if !vis[who[i]] {
			var val = ans[comp[i]] >= 0
			dfs3(adj, comp, vis, ans, who[i], val)
		}
	}

	res := make([]byte, n)

	for i := 0; i < n; i++ {
		if ans[comp[i]] > 0 {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}
	return true, string(res)
}

func dfs1(graph [][]int, vis []bool, who []int, cur *int, u int) {
	vis[u] = true

	for _, v := range graph[u] {
		if !vis[v] {
			dfs1(graph, vis, who, cur, v)
		}
	}
	who[*cur] = u
	*cur++
}

func dfs2(graph [][]int, comp []int, u int, c int) {
	comp[u] = c

	for _, v := range graph[u] {
		if comp[v] == -1 {
			dfs2(graph, comp, v, c)
		}
	}
}

func dfs3(graph [][]int, comp []int, vis []bool, ans []int, u int, val bool) {
	vis[u] = true
	c := comp[u^1]
	if val {
		ans[c] = -1
	} else {
		ans[c] = 1
	}

	for _, v := range graph[u] {
		if !vis[v] {
			dfs3(graph, comp, vis, ans, v, val)
		}
	}

}
