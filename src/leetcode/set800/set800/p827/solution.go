package p827

var dir = []int{-1, 0, 1, 0, -1}

func largestIsland(grid [][]int) int {
	cache := make(map[int]int)

	n := len(grid)

	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	que := make([]int, n*n)
	preProcess := func(x, y, color int) {
		front, tail := 0, 0
		que[tail] = x*n + y
		tail++
		board[x][y] = color
		for front < tail {
			tt := tail
			for front < tt {
				i, j := que[front]/n, que[front]%n
				front++
				for k := 0; k < 4; k++ {
					a, b := i+dir[k], j+dir[k+1]
					if a >= 0 && a < n && b >= 0 && b < n && grid[a][b] == 1 && board[a][b] == 0 {
						board[a][b] = color
						que[tail] = a*n + b
						tail++
					}
				}
			}
		}
		cache[color] = tail
	}
	var ans int
	var color int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 0 && grid[i][j] == 1 {
				color++
				preProcess(i, j, color)
				if cache[color] > ans {
					ans = cache[color]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				//switch i, j
				ns := make(map[int]int)
				if i > 0 {
					if board[i-1][j] > 0 {
						ns[board[i-1][j]] = cache[board[i-1][j]]
					}
				}

				if i < n-1 {
					if board[i+1][j] > 0 {
						ns[board[i+1][j]] = cache[board[i+1][j]]
					}
				}

				if j > 0 {
					if board[i][j-1] > 0 {
						ns[board[i][j-1]] = cache[board[i][j-1]]
					}
				}

				if j < n-1 {
					if board[i][j+1] > 0 {
						ns[board[i][j+1]] = cache[board[i][j+1]]
					}
				}

				var tmp int
				for _, v := range ns {
					tmp += v
				}
				if tmp+1 > ans {
					ans = tmp + 1
				}
			}
		}
	}

	return ans
}
