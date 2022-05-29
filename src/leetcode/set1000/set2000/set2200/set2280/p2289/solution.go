package p2289

var dd = []int{-1, 0, 1, 0, -1}

func minimumObstacles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	que := make([]int, 2*m*n)
	front, end := m*n, m*n

	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = -1
		}
	}
	dist[0][0] = 0
	que[end] = 0
	end++

	for front < end {
		cur := que[front]
		front++
		x, y := cur/n, cur%n
		if x == m-1 && y == n-1 {
			break
		}

		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && dist[u][v] < 0 {
				if grid[u][v] == 1 {
					dist[u][v] = dist[x][y] + 1
					que[end] = u*n + v
					end++
				} else {
					dist[u][v] = dist[x][y]
					front--
					que[front] = u*n + v
				}
			}
		}
	}

	return dist[m-1][n-1]
}
