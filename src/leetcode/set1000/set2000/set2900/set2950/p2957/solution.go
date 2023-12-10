package p2957

func numberOfSets(n int, maxDistance int, roads [][]int) int {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = 1 << 30
		}
		dist[i][i] = 0
	}

	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		dist[u][v] = min(dist[u][v], w)
		dist[v][u] = min(dist[v][u], w)
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	var res int
	buf := make([]int, n)
	for state := 0; state < 1<<n; state++ {
		var p int
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 0 {
				buf[p] = i
				p++
			}
			for j := 0; j < n; j++ {
				dp[i][j] = 1 << 30
			}
		}

		for i := 0; i < p; i++ {
			u := buf[i]
			for j := 0; j < p; j++ {
				v := buf[j]
				dp[u][v] = dist[u][v]
			}
			//dp[u][u] = 0
		}

		for k := 0; k < p; k++ {
			w := buf[k]
			for i := 0; i < p; i++ {
				u := buf[i]
				for j := 0; j < p; j++ {
					v := buf[j]
					dp[u][v] = min(dp[u][v], dp[u][w]+dp[w][v])
				}
			}
		}
		ok := true
		for i := 0; i < p && ok; i++ {
			for j := 0; j < p && ok; j++ {
				u, v := buf[i], buf[j]
				if dp[u][v] > maxDistance {
					ok = false
				}
			}
		}
		if ok {
			res++
		}
	}

	return res
}
