package p3281

var dx = []int{-2, -2, -1, 1, 2, 2, 1, -1}
var dy = []int{-1, 1, 2, 2, 1, -1, -2, -2}

func maxMoves(kx int, ky int, positions [][]int) int {
	n := len(positions)
	// n <= 15
	// dp[state][x][y]表示在马在位置x, y时，剩余state时的最大结果
	// (x, y)可以使用position的下标代替，因为每次都停在那里
	dist := make([][]int, 50)
	for i := 0; i < 50; i++ {
		dist[i] = make([]int, 50)
	}

	que := make([]int, 50*50)

	bfs := func(s int) []int {
		for i := 0; i < 50; i++ {
			for j := 0; j < 50; j++ {
				dist[i][j] = -1
			}
		}
		start := positions[s]
		dist[start[1]][start[0]] = 0
		var head, tail int
		que[head] = start[1]*50 + start[0]
		head++
		for tail < head {
			r, c := que[tail]/50, que[tail]%50
			tail++
			for j := 0; j < 8; j++ {
				u, v := r+dx[j], c+dy[j]
				if u >= 0 && u < 50 && v >= 0 && v < 50 && dist[u][v] < 0 {
					dist[u][v] = dist[r][c] + 1
					que[head] = u*50 + v
					head++
				}
			}
		}
		ans := make([]int, n)

		for i, cur := range positions {
			ans[i] = dist[cur[1]][cur[0]]
		}

		return ans
	}

	move := make([][]int, n)
	kk := make([]int, n)
	for i := 0; i < n; i++ {
		move[i] = bfs(i)
		kk[i] = dist[ky][kx]
	}

	if n == 1 {
		return kk[0]
	}

	N := 1 << n

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	var play func(state int, i int) int

	play = func(state int, i int) (ans int) {
		if state == N-1 {
			return 0
		}

		if dp[state][i] >= 0 {
			return dp[state][i]
		}

		defer func() {
			dp[state][i] = ans
		}()

		// 当前都是alice的操作

		for j := 0; j < n; j++ {
			if (state>>j)&1 == 0 {
				// alice 选择j
				tmp := move[i][j]
				next := state | (1 << j)
				// 现在轮到bob去选择
				if next < N-1 {
					tmp2 := 1 << 30
					for k := 0; k < n; k++ {
						// bob选择k
						if (next>>k)&1 == 0 {
							tmp2 = min(tmp2, move[j][k]+play(next|(1<<k), k))
						}
					}
					ans = max(ans, tmp+tmp2)
				} else {
					ans = max(ans, tmp)
				}

			}
		}

		return ans
	}

	var ans int

	for i := 0; i < n; i++ {
		// 如果选择i
		tmp := kk[i]
		tmp2 := 1 << 30
		for j := 0; j < n; j++ {
			if i != j {
				// bob 的选择
				tmp2 = min(tmp2, move[i][j]+play((1<<i)|(1<<j), j))
			}
		}
		ans = max(ans, tmp+tmp2)
	}

	return ans
}
