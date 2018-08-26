package p892

var dd = []int{-1, 0, 1, 0, -1}

func surfaceArea(grid [][]int) int {
	R := len(grid)
	if R == 0 {
		return 0
	}
	C := len(grid[0])
	if C == 0 {
		return 0
	}

	vis := make([][]bool, R)
	for i := 0; i < R; i++ {
		vis[i] = make([]bool, C)
	}
	que := make([]int, R*C)
	bfs := func(x, y int) int {
		var ans int
		vis[x][y] = true
		head, tail := 0, 0
		que[tail] = x*C + y
		tail++
		for head < tail {
			cur := que[head]
			head++
			i, j := cur/C, cur%C

			for k := 0; k < 4; k++ {
				a, b := i+dd[k], j+dd[k+1]
				if a < 0 || a >= R || b < 0 || b >= C {
					ans += grid[i][j]
					continue
				}

				if vis[a][b] {
					//a, b is checked
					if grid[a][b] < grid[i][j] {
						ans += grid[i][j] - grid[a][b]
					}
					continue
				}

				if grid[a][b] <= grid[i][j] {
					vis[a][b] = true
					ans += grid[i][j] - grid[a][b]
					que[tail] = a*C + b
					tail++
				}
			}
		}
		return ans
	}

	var ans int
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] > 0 {
				ans += 2
			}
			if !vis[i][j] {
				ans += bfs(i, j)
			}
		}
	}
	return ans
}
