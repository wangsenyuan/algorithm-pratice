package main

import "fmt"

func main() {
	/*matrix := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}*/

	/*matrix := [][]byte{
		[]byte{'0', '0'},
		[]byte{'0', '0'},
	}*/
	matrix := [][]byte{
		[]byte{'1'},
	}

	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	m, n := size(matrix)
	if m == 0 {
		return 0
	}

	fill := make([][]int, m+1)
	fill[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		fill[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			fill[i+1][j+1] = fill[i+1][j] + fill[i][j+1] - fill[i][j]
			if matrix[i][j] == '1' {
				fill[i+1][j+1]++
			}
		}
	}

	check := func(x int) bool {
		for i := 0; i+x <= m; i++ {
			for j := 0; j+x <= n; j++ {
				area := fill[i+x][j+x] - fill[i+x][j] - fill[i][j+x] + fill[i][j]
				if area == x*x {
					return true
				}
			}
		}
		return false
	}

	lp := 0
	up := m
	if n < m {
		up = n
	}

	res := 0
	for lp <= up {
		mid := lp + (up-lp)/2
		if check(mid) {
			res = mid
			lp = mid + 1
		} else {
			up = mid - 1
		}
	}
	return res * res
}

func size(matrix [][]byte) (int, int) {
	m := len(matrix)
	if m == 0 {
		return 0, 0
	}

	n := len(matrix[0])
	if n == 0 {
		return 0, 0
	}

	return m, n
}
