package lcp56

var dd = []int{-1, 0, 1, 0, -1}

func conveyorBelt(matrix []string, start []int, end []int) int {
	n := len(matrix)
	m := len(matrix[0])

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = -1
		}
	}

	que := make([]int, m*n*4)
	var head, tail int

	var visit func(x int, y int)

	visit = func(x int, y int) {
		que[tail] = x*m + y
		tail++
		d := getDirection(matrix[x][y])
		u, v := x+dd[d], y+dd[d+1]
		if u >= 0 && u < n && v >= 0 && v < m && dist[u][v] < 0 {
			dist[u][v] = dist[x][y]
			visit(u, v)
		}
	}
	dist[start[0]][start[1]] = 0
	visit(start[0], start[1])

	for dist[end[0]][end[1]] < 0 {
		tmp_tail := tail
		for head < tmp_tail {
			cur := que[head]
			head++
			x, y := cur/m, cur%m
			for i := 0; i < 4; i++ {
				u, v := x+dd[i], y+dd[i+1]
				if u >= 0 && u < n && v >= 0 && v < m && dist[u][v] < 0 {
					dist[u][v] = dist[x][y] + 1
					visit(u, v)
				}
			}
		}
	}

	return dist[end[0]][end[1]]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const INF = 1 << 30

func getDirection(x byte) int {
	if x == '^' {
		return 0
	}
	if x == '>' {
		return 1
	}
	if x == 'v' {
		return 2
	}
	return 3
}
