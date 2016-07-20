package main

import "fmt"

func main() {
	r := uniquePaths(100, 100)
	fmt.Println(r)
}

func uniquePaths(m int, n int) int {
	fx := make([][]int, m)
	for i := range fx {
		fx[i] = make([]int, n)
	}

	fx[0][0] = 1

	for j := 1; j < n; j++ {
		fx[0][j] = 1
	}

	for i := 1; i < m; i++ {
		fx[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			fx[i][j] += fx[i-1][j]

			fx[i][j] += fx[i][j-1]
		}
	}

	return fx[m-1][n-1]
}
