package p1929

var dd = []int{-1, 0, 1, 0, -1}

func nearestExit(maze [][]byte, entrance []int) int {
	m := len(maze)
	n := len(maze[0])
	que := make([]int, m*n)

	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = -1
		}
	}
	var front, end int
	dist[entrance[0]][entrance[1]] = 0
	que[end] = entrance[0]*n + entrance[1]
	end++

	best := -1

	for front < end {
		cur := que[front]
		front++
		x, y := cur/n, cur%n
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && maze[u][v] == '.' && dist[u][v] < 0 {
				dist[u][v] = dist[x][y] + 1
				if u == 0 || u == m-1 || v == 0 || v == n-1 {
					if best < 0 || best > dist[u][v] {
						best = dist[u][v]
					}
				}
				que[end] = u*n + v
				end++
			}
		}
	}
	return best
}
