package main

import "fmt"

func main() {
	grid := [][]byte{[]byte("0X00"), []byte("X0YX"), []byte("0X00")}
	fmt.Println(maxKilledEnemies(grid))
}

func maxKilledEnemies(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	a, b, c, d := makeBoard(m, n), makeBoard(m, n), makeBoard(m, n), makeBoard(m, n)
	check1(grid, a, b)
	check2(grid, c, d)

	mx := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' && a[i][j]+b[i][j]+c[i][j]+d[i][j] > mx {
				mx = a[i][j] + b[i][j] + c[i][j] + d[i][j]
			}
		}
	}
	return mx
}

func makeBoard(m, n int) [][]int {
	board := make([][]int, m, m)
	for i := range board {
		board[i] = make([]int, n, n)
	}
	return board
}

func check1(grid [][]byte, f [][]int, g [][]int) {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'X' {
				if i > 0 {
					f[i][j] = f[i-1][j] + 1
				} else {
					f[i][j] = 1
				}
				if j > 0 {
					g[i][j] = g[i][j-1] + 1
				} else {
					g[i][j] = 1
				}
			} else if grid[i][j] == '0' {
				if i > 0 {
					f[i][j] = f[i-1][j]
				}
				if j > 0 {
					g[i][j] = g[i][j-1]
				}
			}
		}
	}
}

func check2(grid [][]byte, f [][]int, g [][]int) {
	m := len(grid)
	n := len(grid[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 'X' {
				if i < m-1 {
					f[i][j] = f[i+1][j] + 1
				} else {
					f[i][j] = 1
				}
				if j < n-1 {
					g[i][j] = g[i][j+1] + 1
				} else {
					g[i][j] = 1
				}
			} else if grid[i][j] == '0' {
				if i < m-1 {
					f[i][j] = f[i+1][j]
				}
				if j < n-1 {
					g[i][j] = g[i][j+1]
				}
			}
		}
	}
}
