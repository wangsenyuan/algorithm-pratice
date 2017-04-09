package main

import "fmt"

func main() {
	wall := [][]int{
		{1, 2, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{2, 4},
		{3, 1, 2},
		{1, 3, 1, 1},
	}

	fmt.Println(leastBricks(wall))
}

func leastBricks(wall [][]int) int {
	edges := make(map[int]int)

	for _, row := range wall {
		prev := 0
		for i, brick := range row {
			prev += brick
			if i < len(row)-1 {
				edges[prev]++
			}
		}
	}

	var maxCount = 0
	for _, count := range edges {
		if maxCount < count {
			maxCount = count
		}
	}
	return len(wall) - maxCount
}
