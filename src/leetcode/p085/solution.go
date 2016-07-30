package main

import "fmt"

func main() {
	matrix := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}

	fmt.Println(maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	area := 0

	right := make([]int, len(matrix[0]))
	left := make([]int, len(matrix[0]))
	height := make([]int, len(matrix[0]))

	for i := range right {
		right[i] = len(matrix[0])
	}

	for i := 0; i < len(matrix); i++ {
		l, r := 0, len(matrix[i])

		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}

		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				left[j] = max(left[j], l)
			} else {
				left[j] = 0
				l = j + 1
			}
		}

		for j := len(matrix[i]) - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				right[j] = min(right[j], r)
			} else {
				right[j] = len(matrix[i])
				r = j
			}
		}

		for j := 0; j < len(matrix[i]); j++ {
			area = max(area, height[j]*(right[j]-left[j]))
		}
	}

	return area
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
