package main

import "fmt"

var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1, 0}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1, 0}

func imageSmoother(M [][]int) [][]int {
	m := len(M)
	if m == 0 {
		return nil
	}

	n := len(M[0])

	if n == 0 {
		return nil
	}

	smooth := func(i, j int) int {
		sum, cnt := 0, 0
		for k := 0; k < len(dx); k++ {
			x, y := i+dx[k], j+dy[k]

			if x >= 0 && x < m && y >= 0 && y < n {
				sum += M[x][y]
				cnt++
			}
		}
		return sum / cnt
	}

	R := make([][]int, m)
	for i := 0; i < m; i++ {
		R[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			R[i][j] = smooth(i, j)
		}
	}

	return R
}

func main() {
	M := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	fmt.Println(imageSmoother(M))
}
