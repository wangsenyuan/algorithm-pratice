package main

import (
	"math"
	"fmt"
)

func main() {
	matrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	res := updateMatrix(matrix)
	for _, row := range res {
		fmt.Println(row)
	}
}

func updateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	if n == 0 {
		return nil
	}
	m := len(matrix[0])
	if m == 0 {
		return nil
	}
	que := make([]int, m*n)
	p := 0

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
		for j := 0; j < m; j++ {
			res[i][j] = math.MaxInt32
			if matrix[i][j] == 0 {
				res[i][j] = 0
				que[p] = i*m + j
				p++
			}
		}
	}
	dx := []int{-1, 0, 1, 0, -1}

	q := 0
	for q < p {
		x := que[q]
		q++
		i, j := x/m, x%m

		for k := 0; k < 4; k++ {
			a, b := i+dx[k], j+dx[k+1]
			if a < 0 || a >= n || b < 0 || b >= m || matrix[a][b] == 0 {
				continue
			}
			if res[a][b] > res[i][j]+1 {
				res[a][b] = res[i][j] + 1
				que[p] = a*m + b
				p++
			}
		}

	}
	return res
}
