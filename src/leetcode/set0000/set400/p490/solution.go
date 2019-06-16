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
	dest := []int{3, 2}

	fmt.Println(hasPath(maze, start, dest))
}

func hasPath(maze [][]int, start []int, destination []int) bool {

	n := len(maze)
	m := len(maze[0])

	dirs := [][]int{{0, 1}, {-1, 0}, {1, 0}, {0, -1}}; //r, u, d, l}

	var move func(x, y, dir int, checked [][]bool) bool

	move = func(x, y, d int, checked [][]bool) bool {
		r, c := x, y

		if d >= 0 {
			for r >= 0 && r < n && c >= 0 && c < m && maze[r][c] == 0 {
				r += dirs[d][0]
				c += dirs[d][1]
			}

			r -= dirs[d][0]
			c -= dirs[d][1]
		}

		if r == destination[0] && c == destination[1] {
			return true
		}

		if checked[r][c] {
			return false
		}

		checked[r][c] = true

		for i := 0; i < 4; i++ {
			if i == d || 3-d == i {
				continue
			}
			rr, cc := r + dirs[i][0], c + dirs[i][1]
			if rr >= 0 && rr < n && cc >= 0 && cc < m && maze[rr][cc] == 0 {
				if move(r, c, i, checked) {
					return true
				}
			}
		}

		return false
	}

	checked := make([][]bool, n)
	for i := 0; i < n; i++ {
		checked[i] = make([]bool, m)
	}

	return move(start[0], start[1], -1, checked)
}
