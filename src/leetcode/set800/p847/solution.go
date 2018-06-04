package p847

func shortestPathLength(graph [][]int) int {
	n := len(graph)

	N := 1 << uint(n)

	INF := 1 << 20

	//f[state][i] = visit set of nodes (represented by state), and ends at i
	f := make([][]int, N)
	for i := 0; i < N; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = INF
		}
	}

	que := make([]int, N*n)

	front, tail := 0, 0

	for i := 0; i < n; i++ {
		f[1<<uint(i)][i] = 0
		que[tail] = (1<<uint(i))*n + i
		tail++
	}

	best := INF
	for front < tail {
		tt := tail
		for front < tt {
			v := que[front]
			front++
			state, i := v/n, v%n

			if state == N-1 {
				best = min(best, f[state][i])
			}

			for _, j := range graph[i] {
				next := state | (1 << uint(j))
				if f[state][i]+1 < f[next][j] {
					f[next][j] = f[state][i] + 1
					que[tail] = next*n + j
					tail++
				}
			}
		}
	}
	return best
}

func nextState(v int) int {
	t := (v | (v - 1)) + 1
	return t | ((((t & -t) / (v & -v)) >> 1) - 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
