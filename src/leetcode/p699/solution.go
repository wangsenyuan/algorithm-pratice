package main

import "fmt"

type interval struct {
	start  int
	end    int
	height int
}

func fallingSquares(positions [][]int) []int {
	n := len(positions)
	res := make([]int, n)
	intervals := make([]interval, 0, n)
	best := 0

	for i := 0; i < n; i++ {
		cur := positions[i]
		left, side := cur[0], cur[1]
		height := side

		for _, inter := range intervals {
			if inter.end <= left || left+side <= inter.start {
				continue
			}
			if inter.height+side > height {
				height = inter.height + side
			}
		}

		intervals = append(intervals, interval{left, left + side, height})

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
