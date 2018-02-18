package p787

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	oo := 1 << 30

	graph := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make(map[int]int)
	}

	for _, flight := range flights {
		u, v, c := flight[0], flight[1], flight[2]
		graph[u][v] = c
	}

	// each city may be in the que (at most) K+1 times
	que := make([]int, n*(K+2))
	head, tail := 0, 0
	que[tail] = src
	tail++
	ans := oo
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, K+2)
		for j := 0; j <= K+1; j++ {
			dp[i][j] = oo
		}
	}
	dp[src][0] = 0
	flags := make([]int, n)
	for i := 0; i < n; i++ {
		flags[i] = -1
	}
	var stops int
	for head < tail && stops <= K {
		tmp := tail
		for head < tmp {
			v := que[head]
			head++
			for w, c := range graph[v] {
				dp[w][stops+1] = min(dp[w][stops+1], dp[v][stops]+c)
				if w == dst {
					ans = min(ans, dp[v][stops]+c)
				}
				if flags[w] != stops+1 {
					que[tail] = w
					tail++
					flags[w] = stops + 1
				}
			}
		}
		stops++
	}

	if ans == oo {
		return -1
	}
	return ans
}

func findCheapestPrice1(n int, flights [][]int, src int, dst int, K int) int {
	// always stop at dst
	K++
	oo := 1 << 30
	//dp[x][y] = arrives at city x, and take y stops
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = oo
		}
	}
	dp[src][0] = 0

	graph := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make(map[int]int)
	}

	for _, flight := range flights {
		u, v, c := flight[0], flight[1], flight[2]
		graph[u][v] = c
	}

	for k := 0; k < K; k++ {
		for i := 0; i < n; i++ {
			if dp[i][k] < oo {
				for w, c := range graph[i] {
					dp[w][k+1] = min(dp[w][k+1], dp[i][k]+c)
				}
			}
		}
	}

	ans := oo
	for k := 0; k <= K; k++ {
		ans = min(ans, dp[dst][k])
	}
	if ans == oo {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
