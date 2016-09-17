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
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	width := 0
	prev := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp := dp[j+1]
			if matrix[i][j] == '1' {
				dp[j+1] = min(dp[j], dp[j+1], prev) + 1
				if dp[j+1] > width {
					width = dp[j+1]
				}
			} else {
				dp[j+1] = 0
			}
			prev = tmp
		}
	}

	return width * width
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}

	return c
}
