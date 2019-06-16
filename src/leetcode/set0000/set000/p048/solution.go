package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{8, 9, 4},
		[]int{7, 6, 5},
	}

	rotate(matrix)

	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}
}

func rotate(matrix [][]int) {
	var fn func(width int, i int)
	fn = func(width int, i int) {
		if width < 2 {
			return
		}
		fn(width-2, i+1)

		for k := 0; k < width-1; k++ {
			j := i + k
			x := matrix[i][j]
			matrix[i][j] = matrix[i+width-1-k][i]
			matrix[i+width-1-k][i] = matrix[i+width-1][i+width-1-k]
			matrix[i+width-1][i+width-1-k] = matrix[j][i+width-1]
			matrix[j][i+width-1] = x
		}
	}
	fn(len(matrix), 0)
}
