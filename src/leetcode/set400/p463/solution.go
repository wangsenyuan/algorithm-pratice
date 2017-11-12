package main

import "fmt"

func main() {
	grid := [][]int{[]int{0, 1, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 1, 0, 0},
		[]int{1, 1, 0, 0},
	}

	//grid := [][]int{[]int{1, 1}, []int{1, 1}}

	fmt.Println(islandPerimeter(grid))
}

var dx = []int{1, 0, -1, 0}
var dy = []int{0, 1, 0, -1}

func islandPerimeter(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	var total = 0
	var sides = 0
	var perimeter func(x int, y int)

	perimeter = func(i int, j int) {
		if grid[i][j] < 0 {
			return
		}
		total++

		grid[i][j] = -1
		for k := 0; k < 4; k++ {
			x, y := i+dx[k], j+dy[k]
			if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
				continue
			}

			sides++
			perimeter(x, y)
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				perimeter(i, j)
				return 4*total - sides
			}
		}
	}

	return 0
}
