package p802

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	degrees := make([]int, n)

	rg := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		rg[i] = make(map[int]bool)
	}

	for i, row := range graph {
		for _, j := range row {
			degrees[i]++
			rg[j][i] = true
		}
	}

	que := make([]int, n)
	head, tail := 0, 0
	for i := 0; i < n; i++ {
		if degrees[i] == 0 {
			que[tail] = i
			tail++
		}
	}
	can := make([]bool, n)
	for head < tail {
		tt := tail
		for head < tt {
			v := que[head]
			can[v] = true
			head++
			for w := range rg[v] {
				degrees[w]--
				if degrees[w] == 0 {
					que[tail] = w
					tail++
				}
			}
		}
	}

	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if can[i] {
			res = append(res, i)
		}
	}

	return res
}

func eventualSafeNodes1(graph [][]int) []int {
	n := len(graph)
	visited := make([]int, n)
	path := make([]bool, n)
	res := make([]bool, n)
	var visit func(i int, label int) int

	// 0 not in loop, but can't reach terminal; 1 can reach terminal; 2 in loop
	visit = func(v int, label int) int {
		visited[v] = label
		path[v] = true
		if len(graph[v]) == 0 {
			path[v] = false
			// safe
			res[v] = true
			return 1
		}

		var canReachTerminal bool
		var inLoop bool
		for _, w := range graph[v] {
			if path[w] {
				inLoop = true
				continue
			}
			tmp := visit(w, label)
			if tmp == 2 {
				path[v] = false
				inLoop = true
				continue
			}

			if tmp == 1 {
				canReachTerminal = true
			}
		}
		path[v] = false

		if inLoop {
			return 2
		}

		if canReachTerminal {
			res[v] = true
			return 1
		}

		return 0
	}

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			visit(i, i+1)
		}
	}
	ans := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if res[i] {
			ans = append(ans, i)
		}
	}

	return ans
}
