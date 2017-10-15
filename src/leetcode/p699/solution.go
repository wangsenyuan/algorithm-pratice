package main

import "fmt"

func fallingSquares(positions [][]int) []int {
	n := len(positions)
	res := make([]int, n)

	prev := make(map[int][]int)

	best := 0

	for i := 0; i < n; i++ {
		cur := positions[i]
		left, side := cur[0], cur[1]
		height := side
		for pos, sq := range prev {
			a, b := sq[0], sq[1]
			if pos+a <= left || pos >= left+side {
				continue
			}
			if b+side > height {
				height = b + side
			}
		}
		prev[left] = []int{side, height}

		if height > best {
			best = height
		}
		res[i] = best
	}

	return res
}

func main() {
	fmt.Println(fallingSquares([][]int{{1, 2}, {2, 3}, {6, 1}}))
	fmt.Println(fallingSquares([][]int{{100, 100}, {200, 100}}))
}
