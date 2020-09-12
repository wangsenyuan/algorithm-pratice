package e

func chaseGame(edges [][]int, startA int, startB int) int {
	startA--
	startB--
	n := len(edges)
	G := getGraph(n, edges)
	degree := make([]int, n)
	for _, edge := range edges {
		degree[edge[0]-1]++
		degree[edge[1]-1]++
	}
	que := make([]int, n)
	bfs := func(x int, dist []int) {
		var end int
		que[end] = x
		dist[x] = 0
		end++
		var front int
		for front < end {
			u := que[front]
			front++
			for _, v := range G[u] {
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[end] = v
					end++
				}
			}
		}
	}

	leet := make([]int, n)
	code := make([]int, n)
	for i := 0; i < n; i++ {
		leet[i] = -1
		code[i] = -1
	}

	bfs(startA, leet)
	bfs(startB, code)

	var cnt = n
	var end int
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			que[end] = i
			end++
			cnt--
		}
	}
	var front int
	for front < end {
		u := que[front]
		front++
		for _, v := range G[u] {
			degree[v]--
			if degree[v] == 1 {
				cnt--
				que[end] = v
				end++
			}
		}
	}
	var ans = 1
	if cnt == 3 {
		for i := 0; i < n; i++ {
			if leet[i] > code[i]+1 {
				ans = max(ans, leet[i])
			}
		}
	} else {
		for i := 0; i < n; i++ {
			if leet[i] > code[i]+1 {
				if degree[i] > 1 {
					return -1
				} else {
					ans = max(ans, leet[i])
				}
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getGraph(n int, E [][]int) [][]int {
	degree := make([]int, n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		degree[u]++
		degree[v]++
	}
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, degree[i])
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	return adj
}
