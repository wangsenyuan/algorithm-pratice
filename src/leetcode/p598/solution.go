package main

import "fmt"

func main() {
	fmt.Println(maxCount(3, 3, [][]int{{2, 2}, {3, 3}}))
}

func maxCount(m int, n int, ops [][]int) int {
	x, y := m, n

	for _, op := range ops {
		if op[0] < x {
			x = op[0]
		}
		if op[1] < y {
			y = op[1]
		}
	}

	return x * y
}
