package p1263

var dd = []int{-1, 0, 1, 0, -1}

func minPushBox(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	N := n * m

	vis := make([]int, N)
	playerQue := make([]int, N)

	canReachAt := func(a, b, c, d, x, y int, label int) bool {
		// c, d is the destination, x, y is the block position
		// a, b is the start position
		if c < 0 || c >= n || d < 0 || d >= m {
			return false
		}
		if grid[c][d] == '#' {
			return false
		}

		var end int
		playerQue[end] = a*m + b
		end++

		vis[a*m+b] = label

		var front int

		for front < end {
			cur := playerQue[front]
			front++
			u, v := cur/m, cur%m

			if u == c && v == d {
				return true
			}

			for k := 0; k < 4; k++ {
				p, q := u+dd[k], v+dd[k+1]
				if p < 0 || p == n || q < 0 || q == m {
					continue
				}
				if grid[p][q] == '#' {
					continue
				}
				if p == x && q == y {
					continue
				}
				if vis[p*m+q] != label {
					vis[p*m+q] = label
					playerQue[end] = p*m + q
					end++
				}
			}

		}
		return false
	}

	//(a, b, c, d) (player position & block position)
	dp := make([][]int, N)

	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = -1
		}
	}

	sx, sy := findPosition(grid, 'S')
	bx, by := findPosition(grid, 'B')
	dp[sx*m+sy][bx*m+by] = 0

	que := make([]int, N*N)
	var end int
	que[end] = calIndex(sx*m+sy, bx*m+by, N)
	end++

	var front int
	var label int
	for front < end {
		cur := que[front]
		front++

		pp, bb := cur/N, cur%N
		sx, sy = pp/m, pp%m
		bx, by = bb/m, bb%m

		for k := 0; k < 4; k++ {
			u, v := bx+dd[k], by+dd[k+1]

			label++

			if canReachAt(sx, sy, u, v, bx, by, label) {
				uu, vv := -1, -1
				//can push
				if k == 0 {
					// stand at up, then push down
					if bx+1 < n && grid[bx+1][by] != '#' {
						// can push down
						uu = bx + 1
						vv = by
					}
				} else if k == 1 {
					// stand at right, then push left
					if by-1 >= 0 && grid[bx][by-1] != '#' {
						uu = bx
						vv = by - 1
					}
				} else if k == 2 {
					// stand at down, then push up
					if bx-1 >= 0 && grid[bx-1][by] != '#' {
						uu = bx - 1
						vv = by
					}
				} else {
					if by+1 < m && grid[bx][by+1] != '#' {
						uu = bx
						vv = by + 1
					}
				}

				if uu < 0 || vv < 0 {
					continue
				}
				if dp[bx*m+by][uu*m+vv] < 0 {
					dp[bx*m+by][uu*m+vv] = dp[pp][bb] + 1
					ii := calIndex(bx*m+by, uu*m+vv, N)
					que[end] = ii
					end++
				}
			}
		}
	}

	tx, ty := findPosition(grid, 'T')

	res := -1

	for k := 0; k < 4; k++ {
		u, v := tx+dd[k], ty+dd[k+1]

		if u < 0 || u == n || v < 0 || v == m || grid[u][v] == '#' {
			continue
		}

		if dp[u*m+v][tx*m+ty] > 0 && (res < 0 || res > dp[u*m+v][tx*m+ty]) {
			res = dp[u*m+v][tx*m+ty]
		}
	}

	return res
}

func calIndex(a, b, n int) int {
	return a*n + b
}

func findPosition(grid [][]byte, target byte) (x int, y int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == target {
				x = i
				y = j
				return
			}
		}
	}
	return
}
