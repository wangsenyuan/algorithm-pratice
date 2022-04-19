package p2242

type Pair struct {
	first  int
	second int
}

func maximumScore(scores []int, edges [][]int) int {
	n := len(scores)
	conn := make([][]Pair, n)

	update := func(u int, v int) {
		score := scores[v]

		if len(conn[u]) == 0 {
			conn[u] = []Pair{Pair{score, v}}
			return
		}

		var i int
		for i < len(conn[u]) && score < conn[u][i].first {
			i++
		}
		conn[u] = append(conn[u], Pair{score, v})

		if i < len(conn[u]) {
			// score >= conn[u].first
			copy(conn[u][i+1:], conn[u][i:])
			conn[u][i] = Pair{score, v}
		}
		if len(conn[u]) > 3 {
			conn[u] = conn[u][:3]
		}
	}

	for _, e := range edges {
		u, v := e[0], e[1]
		update(u, v)
		update(v, u)
	}

	best := -1

	for _, e := range edges {
		x, y := e[0], e[1]
		for i := 0; i < len(conn[x]); i++ {
			a := conn[x][i].second
			if a == y {
				continue
			}
			for j := 0; j < len(conn[y]); j++ {
				b := conn[y][j].second
				if b == x || b == a {
					continue
				}
				best = max(best, scores[x]+scores[y]+scores[a]+scores[b])
			}
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
