package p1162

var dd = []int{-1, 0, 1, 0, -1}

func maxDistance(grid [][]int) int {
	N := len(grid)
	dist := make([][]int, N)
	que := make([]int, N*N)
	var end int
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if grid[i][j] == 1 {
				que[end] = i*N + j
				end++
				dist[i][j] = 0
			} else {
				dist[i][j] = -1
			}
		}
	}
	if end == 0 || end == N*N {
		// no land or all land
		return -1
	}

	var front int
	for front < end {
		cur := que[front]
		front++
		x, y := cur/N, cur%N
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < N && v >= 0 && v < N && dist[u][v] < 0 {
				dist[u][v] = dist[x][y] + 1
				que[end] = u*N + v
				end++
			}
		}
	}
	var ans int

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if dist[i][j] > ans {
				ans = dist[i][j]
			}
		}
	}
	return ans
}
