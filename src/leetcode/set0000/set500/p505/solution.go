package main

import "fmt"

func main() {
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}

	start := []int{0, 4}
	dest := []int{4, 4}

	fmt.Println(shortestDistance(maze, start, dest))
}

func shortestDistance(maze [][]int, start []int, destination []int) int {

	n := len(maze)
	m := len(maze[0])

	MAX_SD := n*m + 1
	sd := MAX_SD

	dirs := [][]int{{0, 1}, {-1, 0}, {1, 0}, {0, -1}}; //r, u, d, l}

	var move func(x, y, dir, dist int, cache [][]int)

	move = func(x, y, d, dist int, cache [][]int) {
		if dist > sd || dist > cache[x][y] {
			return
		}

		r, c := x, y

		if d >= 0 {
			for r >= 0 && r < n && c >= 0 && c < m && maze[r][c] == 0 {
				cache[r][c] = min(cache[r][c], dist)
				r += dirs[d][0]
				c += dirs[d][1]
				dist++
			}

			r -= dirs[d][0]
			c -= dirs[d][1]
			dist--
		}

		if r == destination[0] && c == destination[1] {
			if dist < sd {
				sd = dist
			}
			return
		}

		for i := 0; i < 4; i++ {
			if i == d || 3-d == i {
				continue
			}
			rr, cc := r + dirs[i][0], c + dirs[i][1]
			if rr >= 0 && rr < n && cc >= 0 && cc < m && maze[rr][cc] == 0 {
				move(r, c, i, dist, cache)
			}
		}

	}

	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = MAX_SD
		}
	}

	move(start[0], start[1], -1, 0, cache)

	if sd == MAX_SD {
		return -1
	}
	return sd
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
