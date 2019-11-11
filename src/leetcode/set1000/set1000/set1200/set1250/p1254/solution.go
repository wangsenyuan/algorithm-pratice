package p1254

func closedIsland(grid [][]int) int {
	// 0 is island

	m := len(grid)
	n := len(grid[0])
	if m <= 2 || n <= 2 {
		return 0
	}

	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = -1
		}
	}

	que := make([]int, m*n)

	dd := []int{-1, 0, 1, 0, -1}

	bfs := func(x, y int) bool {
		var end int
		que[end] = x*n + y
		end++
		dist[x][y] = 0
		closed := true

		var front int

		for front < end {
			cur := que[front]
			front++
			a, b := cur/n, cur%n

			for i := 0; i < 4; i++ {
				c, d := a+dd[i], b+dd[i+1]
				if c < 0 || c >= m || d < 0 || d >= n {
					closed = false
					continue
				}

				if grid[c][d] == 0 && dist[c][d] < 0 {
					// island
					que[end] = c*n + d
					end++
					dist[c][d] = dist[a][b] + 1
				}
			}
		}
		return closed
	}

	var res int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && dist[i][j] < 0 {
				closed := bfs(i, j)
				if closed {
					res++
				}
			}
		}
	}

	return res
}
