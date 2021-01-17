package p1726

var dd = []int{-1, 0, 1, 0, -1}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	rows := len(grid)
	cols := len(grid[0])

	dp := make([][][][][]int, rows)
	degree := make([][][][][]int, rows)
	for cx := 0; cx < rows; cx++ {
		dp[cx] = make([][][][]int, cols)
		degree[cx] = make([][][][]int, cols)
		for cy := 0; cy < cols; cy++ {
			dp[cx][cy] = make([][][]int, rows)
			degree[cx][cy] = make([][][]int, rows)
			for mx := 0; mx < rows; mx++ {
				dp[cx][cy][mx] = make([][]int, cols)
				degree[cx][cy][mx] = make([][]int, cols)
				for my := 0; my < cols; my++ {
					dp[cx][cy][mx][my] = make([]int, 2)
					degree[cx][cy][mx][my] = make([]int, 2)
				}
			}
		}
	}

	getNeighbors := func(x, y, jumps int) [][]int {
		res := make([][]int, 0, 4*jumps)
		res = append(res, []int{x, y})
		for k := 0; k < 4; k++ {
			u, v := x, y
			for i := 1; i <= jumps; i++ {
				u += dd[k]
				v += dd[k+1]
				if u < 0 || u >= rows || v < 0 || v >= cols || grid[u][v] == '#' {
					break
				}
				res = append(res, []int{u, v})
			}
		}
		return res
	}

	var cx0, cy0, mx0, my0, fx0, fy0 int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'F' {
				fx0, fy0 = i, j
			}
			if grid[i][j] == 'C' {
				cx0, cy0 = i, j
			}
			if grid[i][j] == 'M' {
				mx0, my0 = i, j
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '#' {
				continue
			}
			for u := 0; u < rows; u++ {
				for v := 0; v < cols; v++ {
					if grid[u][v] == '#' {
						continue
					}
					degree[i][j][u][v][0] = len(getNeighbors(i, j, catJump))
					degree[i][j][u][v][1] = len(getNeighbors(u, v, mouseJump))
				}
			}
		}
	}

	que := make([]State, 0, rows*cols*rows*cols*2)

	enqueue := func(cx, cy, mx, my, op int) {
		que = append(que, State{cx, cy, mx, my, op})
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != '#' && grid[i][j] != 'F' {
				dp[i][j][i][j][0] = 1
				enqueue(i, j, i, j, 0)
				dp[i][j][i][j][1] = -1
				enqueue(i, j, i, j, 1)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != '#' && grid[i][j] != 'F' {
				dp[fx0][fy0][i][j][1] = -1
				enqueue(fx0, fy0, i, j, 1)
				dp[i][j][fx0][fy0][0] = -1
				enqueue(i, j, fx0, fy0, 0)
			}
		}
	}

	var pos int

	for pos < len(que) {
		cur := que[pos]
		pos++
		if cur.op == 0 {
			neighbors := getNeighbors(cur.mx, cur.my, mouseJump)
			for _, tmp := range neighbors {
				x, y := tmp[0], tmp[1]
				degree[cur.cx][cur.cy][x][y][cur.op^1]--
				if dp[cur.cx][cur.cy][x][y][cur.op^1] == 0 {
					if dp[cur.cx][cur.cy][cur.mx][cur.my][cur.op] == -1 {
						dp[cur.cx][cur.cy][x][y][cur.op^1] = 1
						enqueue(cur.cx, cur.cy, x, y, cur.op^1)
					} else if degree[cur.cx][cur.cy][x][y][cur.op^1] == 0 {
						dp[cur.cx][cur.cy][x][y][cur.op^1] = -1
						enqueue(cur.cx, cur.cy, x, y, cur.op^1)
					}
				}
			}
		} else {
			neighbors := getNeighbors(cur.cx, cur.cy, catJump)

			for _, tmp := range neighbors {
				x, y := tmp[0], tmp[1]
				degree[x][y][cur.mx][cur.my][cur.op^1]--
				if dp[x][y][cur.mx][cur.my][cur.op^1] == 0 {
					if dp[cur.cx][cur.cy][cur.mx][cur.my][cur.op] == -1 {
						dp[x][y][cur.mx][cur.my][cur.op^1] = 1
						enqueue(x, y, cur.mx, cur.my, cur.op^1)
					} else if degree[x][y][cur.mx][cur.my][cur.op^1] == 0 {
						dp[x][y][cur.mx][cur.my][cur.op^1] = -1
						enqueue(x, y, cur.mx, cur.my, cur.op^1)
					}
				}
			}
		}
	}

	return dp[cx0][cy0][mx0][my0][1] == 1
}

type State struct {
	cx, cy, mx, my, op int
}
