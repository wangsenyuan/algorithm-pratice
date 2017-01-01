package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		t--
		var n int
		fmt.Scanf("%d", &n)
		grid := make([][]byte, n)
		var s string
		for i := 0; i < n; i++ {
			fmt.Scanf("%s", &s)
			grid[i] = []byte(s)
		}

		res := findPossibleMirrorCells(n, grid)

		fmt.Println(res)

	}
}

func findPossibleMirrorCells(n int, grid [][]byte) int {
	mirrors := make([][]int, n)
	for i := 0; i < n; i++ {
		mirrors[i] = make([]int, n)
	}

	cnt := 0
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == '#' {
				continue
			}

			if i == n-1 || (mirrors[i + 1][j]&1 == 1) {
				mirrors[i][j] |= 1
			}

			if j == n-1 || (mirrors[i][j + 1]&2 == 2) {
				mirrors[i][j] |= 2
			}

			if mirrors[i][j] == 3 {
				cnt ++
			}
		}
	}
	return cnt
}
