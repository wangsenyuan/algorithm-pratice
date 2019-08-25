package p1091

func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	if grid[0][0] == 1 {
		return -1
	}

	n := len(grid)

	N := n * n
	que := make([]int, N)
	dist := make([]int, N)

	dist[0] = 1

	var front int
	var end int
	que[end] = 0
	end++

	for front < end {
		cur := que[front]
		front++
		if cur == N-1 {
			return dist[cur]
		}
		x := cur / n
		y := cur % n

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				u, v := x+i, y+j
				if u == x && v == y {
					continue
				}
				if u < 0 || u >= n || v < 0 || v >= n {
					continue
				}
				if dist[u*n+v] == 0 && grid[u][v] == 0 {
					dist[u*n+v] = dist[cur] + 1
					que[end] = u*n + v
					end++
				}
			}
		}
	}

	return -1
}
