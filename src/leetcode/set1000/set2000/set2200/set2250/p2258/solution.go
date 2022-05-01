package p2258

const INF = 1000000000

var dd = []int{-1, 0, 1, 0, -1}

func maximumMinutes(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	fire_at := make([][]int, m)
	for i := 0; i < m; i++ {
		fire_at[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fire_at[i][j] = -1
		}
	}

	que := make([]int, m*n)
	var front, end int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				que[end] = i*n + j
				fire_at[i][j] = 0
				end++
			}
		}
	}

	for front < end {
		i, j := que[front]/n, que[front]%n
		front++
		for k := 0; k < 4; k++ {
			x, y := i+dd[k], j+dd[k+1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 0 && fire_at[x][y] < 0 {
				fire_at[x][y] = fire_at[i][j] + 1
				que[end] = x*n + y
				end++
			}
		}
	}

	safe := make([][]int, m)
	for i := 0; i < m; i++ {
		safe[i] = make([]int, n)
	}

	check := func(time int) bool {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				safe[i][j] = -1
			}
		}

		var front, end int
		safe[0][0] = time
		que[end] = 0
		end++

		for front < end {
			cur := que[front]
			front++
			i, j := cur/n, cur%n

			for k := 0; k < 4; k++ {
				u, v := i+dd[k], j+dd[k+1]
				if u >= 0 && u < m && v >= 0 && v < n && grid[u][v] == 0 &&
					safe[u][v] < 0 &&
					(fire_at[u][v] < 0 || fire_at[u][v] > safe[i][j]+1 ||
						u == m-1 && v == n-1 && fire_at[u][v] == safe[i][j]+1) {
					safe[u][v] = safe[i][j] + 1
					que[end] = u*n + v
					end++
				}
			}
		}
		return safe[m-1][n-1] > 0
	}

	if check(INF) {
		return INF
	}

	left, right := 0, INF
	for left < right {
		mid := (left + right) / 2
		if !check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right - 1
}
