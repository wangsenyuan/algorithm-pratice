package main

import "fmt"

func main() {
	/*dungeon := [][]int{
		[]int{-2, -3, 3},
		[]int{-5, -10, 1},
		[]int{10, 30, -5},
	}*/
	dungeon := [][]int{
		[]int{0, -3},
	}

	fmt.Println(calculateMinimumHP(dungeon))
}

func calculateMinimumHP(dungeon [][]int) int {
	n := len(dungeon)
	if n == 0 {
		return 0
	}

	m := len(dungeon[0])

	if m == 0 {
		return 0
	}

	left := make([]int, m)

	left[m-1] = max(1, 1-dungeon[n-1][m-1])

	for j := m - 2; j >= 0; j-- {
		left[j] = max(1, left[j+1]-dungeon[n-1][j])

	}

	for i := n - 2; i >= 0; i-- {
		row := dungeon[i]
		for j := m - 1; j >= 0; j-- {
			left[j] = left[j] - row[j]
			if j < m-1 && left[j+1]-row[j] < left[j] {
				left[j] = left[j+1] - row[j]
			}
			if left[j] < 1 {
				left[j] = 1
			}
		}
	}

	return left[0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
