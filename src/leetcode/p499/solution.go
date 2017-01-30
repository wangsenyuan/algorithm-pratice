package main

import (
	"fmt"
	"math"
)

func main() {
	maze := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}
	ball := []int{0, 4}
	hole := []int{2, 0}

	fmt.Println(findShortestWay(maze, ball, hole))
}

func findShortestWay(maze [][]int, ball []int, hole []int) string {
	n := len(maze)
	cache := make([][]int, n)
	m := len(maze[0])

	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = math.MaxInt32
		}
	}

	dirs := []string{"r", "u", "d", "l"}
	moves := [][]int{{0, 1}, {-1, 0}, {1, 0}, {0, -1}}; //r, u, d, l}

	var move func(i, j, dir, cnt int, path string)
	var minCnt = math.MaxInt32
	var minPath = ""

	move = func(row, col, dir, cnt int, path string) {
		if cnt > minCnt || cnt > cache[row][col] {
			return
		}
		r, c := row, col
		if dir >= 0 {
			path += dirs[dir]
			for r >= 0 && r < n && c >= 0 && c < m && maze[r][c] == 0 {
				cache[r][c] = min(cache[r][c], cnt)
				if r == hole[0] && c == hole[1] {
					if cnt == minCnt && path < minPath {
						minPath = path
					} else if cnt < minCnt {
						minCnt = cnt
						minPath = path
					}
					return
				}

				r += moves[dir][0]
				c += moves[dir][1]
				cnt++
			}
			r -= moves[dir][0]
			c -= moves[dir][1]
			cnt--
		}

		for dx := 0; dx < 4; dx++ {
			if dir >= 0 && (dx == dir || 3-dx == dir) {
				continue
			}
			nr := r + moves[dx][0]
			nc := c + moves[dx][1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && maze[nr][nc] == 0 {
				move(r, c, dx, cnt, path)
			}
		}
	}

	move(ball[0], ball[1], -1, 0, "")
	if len(minPath) == 0 {
		return "impossible"
	}
	return minPath
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
