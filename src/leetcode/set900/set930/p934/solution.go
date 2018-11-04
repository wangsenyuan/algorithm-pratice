package p934

import "math"

func shortestBridge(A [][]int) int {
	n := len(A)
	m := len(A[0])

	dd := []int{-1, 0, 1, 0, -1}

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = -1
		}
	}

	bfs := func(x, y int) int {
		best := math.MaxInt32
		dist[x][y] = 0
		que := make([]int, 2*n*m)
		front, end := n*m, n*m
		que[end] = x*m + y
		end++
		for front < end {
			cur := que[front]
			front++
			a, b := cur/m, cur%m

			for k := 0; k < 4; k++ {
				c, d := a+dd[k], b+dd[k+1]
				if c >= 0 && c < n && d >= 0 && d < m {
					if A[a][b] == 1 && A[c][d] == 1 {
						// same comp
						if dist[c][d] == -1 {
							dist[c][d] = 0
							que[front-1] = c*m + d
							front--
						}
					} else if A[c][d] == 0 {
						if dist[c][d] == -1 {
							dist[c][d] = dist[a][b] + 1
							que[end] = c*m + d
							end++
						}
					} else if A[c][d] == 1 && dist[c][d] == -1 {
						// another island, no need to put it in the que
						if best > dist[a][b] {
							best = dist[a][b]
						}
					}
				}
			}
		}
		return best
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 1 {
				return bfs(i, j)
			}
		}
	}

	return 0
}
