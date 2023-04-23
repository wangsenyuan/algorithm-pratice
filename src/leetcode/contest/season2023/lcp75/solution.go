package lcp75

var dd []int = []int{-1, 0, 1, 0, -1}

func challengeOfTheKeeper(maze []string) int {
	n := len(maze)
	m := len(maze[0])
	// n <= 200, m <= 200,
	// 如果 player 必经之路，可以将其传送到trap的区域，答案为-1

	s := make([]int, 2)
	t := make([]int, 2)

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if maze[i][j] == 'S' {
				s = []int{i, j}
			} else if maze[i][j] == 'T' {
				t = []int{i, j}
			}
		}
	}

	que := make([]int, n*m)

	bfs := func(x, y int) {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				dist[i][j] = -1
			}
		}
		var front, end int
		que[end] = x*m + y
		end++
		dist[x][y] = 0
		for front < end {
			cur := que[front]
			front++
			u, v := cur/m, cur%m
			for i := 0; i < 4; i++ {
				a, b := u+dd[i], v+dd[i+1]
				if a >= 0 && a < n && b >= 0 && b < m && maze[a][b] != '#' && dist[a][b] < 0 {
					dist[a][b] = dist[u][v] + 1
					que[end] = a*m + b
					end++
				}
			}
		}
	}

	bfs(t[0], t[1])

	if dist[s[0]][s[1]] < 0 {
		return -1
	}

	findMirror := func(x int, y int) int {
		a, b := n-1-x, y
		c, d := x, m-1-y
		var tmp int
		if maze[a][b] != '#' {
			tmp = dist[a][b]
		}
		if maze[c][d] != '#' {
			if tmp >= 0 && (dist[c][d] < 0 || dist[c][d] > tmp) {
				tmp = dist[c][d]
			}
		}
		return tmp
	}

	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i == s[0] && j == s[1] || i == t[0] && j == t[1] {
				// no mirror at start
				continue
			}
			if maze[i][j] == '#' {
				board[i][j] = -1
				continue
			}
			board[i][j] = findMirror(i, j)
		}
	}

	bfs2 := func(x int, y int, limit int) bool {
		// can player reach in limit
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				dist[i][j] = -1
			}
		}
		var front, end int
		que[end] = x*m + y
		end++
		dist[x][y] = 0

		for front < end {
			cur := que[front]
			front++
			u, v := cur/m, cur%m
			for k := 0; k < 4; k++ {
				i, j := u+dd[k], v+dd[k+1]
				if i >= 0 && i < n && j >= 0 && j < m && board[i][j] >= 0 && board[i][j] <= limit && dist[i][j] < 0 {
					que[end] = i*m + j
					end++
					dist[i][j] = dist[u][v] + 1
				}
			}
		}

		return dist[t[0]][t[1]] >= 0
	}

	l, r := 0, 1<<30

	if !bfs2(s[0], s[1], r) {
		return -1
	}

	for l < r {
		mid := (l + r) / 2
		if bfs2(s[0], s[1], mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
