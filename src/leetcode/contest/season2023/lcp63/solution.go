package lcp63

func ballGame(num int, plate []string) [][]int {
	n := len(plate)
	m := len(plate[0])
	dist := make([][][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = make([]int, 4)
			for k := 0; k < 4; k++ {
				dist[i][j][k] = -1
			}
		}
	}
	var inf = num + 1
	var dfs func(u, v, w, steps int) int
	dfs = func(u, v, w, steps int) int {
		if steps > num {
			return inf
		}

		if dist[u][v][w] > 0 {
			return dist[u][v][w]
		}

		if plate[u][v] == 'O' {
			// a hole
			return 0
		}

		if dist[u][v][w] == 0 {
			// a loop back
			return inf
		}
		dist[u][v][w] = 0

		if w == 0 && u-1 >= 0 {
			// go up
			if plate[u-1][v] == '.' {
				dist[u][v][w] = 1 + dfs(u-1, v, 0, steps+1)
			} else if plate[u-1][v] == 'E' {
				// turn right
				dist[u][v][w] = 1 + dfs(u-1, v, 1, steps+1)
			} else {
				// W
				dist[u][v][w] = 1 + dfs(u-1, v, 3, steps+1)
			}
		} else if w == 1 && v+1 < m {
			// go right
			if plate[u][v+1] == '.' {
				dist[u][v][w] = 1 + dfs(u, v+1, 1, steps+1)
			} else if plate[u][v+1] == 'E' {
				dist[u][v][w] = 1 + dfs(u, v+1, 2, steps+1)
			} else {
				dist[u][v][w] = 1 + dfs(u, v+1, 0, steps+1)
			}
		} else if w == 2 && u+1 < n {
			// go down
			if plate[u+1][v] == '.' {
				dist[u][v][w] = 1 + dfs(u+1, v, 2, steps+1)
			} else if plate[u+1][v] == 'E' {
				dist[u][v][w] = 1 + dfs(u+1, v, 3, steps+1)
			} else {
				dist[u][v][w] = 1 + dfs(u+1, v, 1, steps+1)
			}
		} else if w == 3 && v-1 >= 0 {
			if plate[u][v-1] == '.' {
				dist[u][v][w] = 1 + dfs(u, v-1, 3, steps+1)
			} else if plate[u][v-1] == 'E' {
				dist[u][v][w] = 1 + dfs(u, v-1, 0, steps+1)
			} else {
				dist[u][v][w] = 1 + dfs(u, v-1, 2, steps+1)
			}
		} else {
			// go out
			dist[u][v][w] = inf
		}

		return dist[u][v][w]
	}

	var res [][]int

	for j := 1; j+1 < m; j++ {
		if plate[0][j] == '.' {
			tmp := dfs(0, j, 2, 0)
			if tmp <= num {
				res = append(res, []int{0, j})
			}
		}
		if n > 1 && plate[n-1][j] == '.' {
			tmp := dfs(n-1, j, 0, 0)
			if tmp <= num {
				res = append(res, []int{n - 1, j})
			}
		}
	}
	for i := 1; i+1 < n; i++ {
		if plate[i][0] == '.' {
			tmp := dfs(i, 0, 1, 0)
			if tmp <= num {
				res = append(res, []int{i, 0})
			}
		}
		if m > 1 && plate[i][m-1] == '.' {
			tmp := dfs(i, m-1, 3, 0)
			if tmp <= num {
				res = append(res, []int{i, m - 1})
			}
		}
	}

	return res
}

func ballGame1(num int, plate []string) [][]int {
	n := len(plate)
	m := len(plate[0])
	// 0 表示面向上方给自移动
	// 1 表示面向右方格子移动
	// 2 表示面向下方格子移动
	// 3 表示面向左方格子移动
	// 当下个格子是个空格子时，直接进入下个格子，且朝向不变
	// 当下个格子是W是，进入下个格子，且朝向按顺时针方向变化90度
	// 当下个格子是E时，逆时针变化90度
	dist := make([]int, n*m*4)

	que := make([]int, n*m*4)

	var front, end int
	pushEnd := func(i, j, k int) {
		que[end] = i*m*4 + j*4 + k
		end++
	}

	pop := func() (i int, j int, k int) {
		u := que[front]
		front++
		i = u / (4 * m)
		j = (u % (4 * m)) / 4
		k = u % 4
		return
	}

	update := func(i, j, k, d int) {
		u := i*4*m + j*4 + k
		if dist[u] < 0 {
			dist[u] = d
			pushEnd(i, j, k)
		}
	}

	getValue := func(i int, j int) byte {
		return plate[i][j]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 4; k++ {
				dist[i*4*m+j*4+k] = -1
				if getValue(i, j) == 'O' {
					update(i, j, k, 0)
				}
			}
		}
	}

	for front < end {
		u, v, a := pop()
		cur := dist[u*4*m+v*4+a]
		if a == 0 && u > 0 {
			if getValue(u-1, v) == '.' {
				update(u-1, v, 0, cur+1)
			} else if getValue(u-1, v) == 'W' {
				// towards right now
				update(u-1, v, 1, cur+1)
			} else if getValue(u-1, v) == 'E' {
				// t-left
				update(u-1, v, 3, cur+1)
			}
		}
		if a == 1 && v+1 < m {
			if getValue(u, v+1) == '.' {
				update(u, v+1, 1, cur+1)
			} else if getValue(u, v+1) == 'W' {
				// t-down
				update(u, v+1, 2, cur+1)
			} else if getValue(u, v+1) == 'E' {
				update(u, v+1, 0, cur+1)
			}
		}
		if a == 2 && u+1 < n {
			if getValue(u+1, v) == '.' {
				update(u+1, v, 2, cur+1)
			}
			if getValue(u+1, v) == 'W' {
				update(u+1, v, 3, cur+1)
			}
			if getValue(u+1, v) == 'E' {
				update(u+1, v, 1, cur+1)
			}
		}
		if a == 3 && v-1 >= 0 {
			// t-right
			if getValue(u, v-1) == '.' {
				update(u, v-1, 3, cur+1)
			}
			if getValue(u, v-1) == 'W' {
				update(u, v-1, 0, cur+1)
			}
			if getValue(u, v-1) == 'E' {
				update(u, v-1, 2, cur+1)
			}
		}
	}

	var res [][]int
	for j := 1; j+1 < m; j++ {
		// top line, needs to go up
		u := 0*4*m + j*4 + 0
		if getValue(0, j) == '.' && dist[u] > 0 && dist[u] <= num {
			res = append(res, []int{0, j})
		}
		if n > 1 {
			u = (n-1)*4*m + j*4 + 2
			// bottom line, needs to go up
			if getValue(n-1, j) == '.' && dist[u] > 0 && dist[u] <= num {
				res = append(res, []int{n - 1, j})
			}
		}
	}

	for i := 1; i+1 < n; i++ {
		// left line
		if getValue(i, 0) == '.' && dist[i*4*m+0*4+3] > 0 && dist[i*4*m+0*4+3] <= num {
			res = append(res, []int{i, 0})
		}
		if m > 1 {
			if getValue(i, m-1) == '.' && dist[i*4*m+(m-1)*4+1] > 0 && dist[i*4*m+(m-1)*4+1] <= num {
				res = append(res, []int{i, m - 1})
			}
		}
	}

	return res
}
