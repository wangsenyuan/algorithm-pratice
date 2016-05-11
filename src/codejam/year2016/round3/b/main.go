package main

import (
	"fmt"
	"math"
)

func main() {
	t := 0
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {
		var b int
		var m int64
		fmt.Scanf("%d %d", &b, &m)
		grid, found := play(b, m)
		if !found {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i)
			continue
		}

		fmt.Printf("Case #%d: POSSIBLE\n", i)
		outputGrid(grid, b)
	}
}

func outputGrid(grid [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d", grid[i][j])
		}
		fmt.Println()
	}
}

func play(b int, m int64) ([][]int, bool) {
	n := int64(math.Exp2(float64(b) - 2))
	if m > n {
		return nil, false
	}

	grid := make([][]int, b, b)
	for i := 0; i < b; i++ {
		grid[i] = make([]int, b, b)
	}

	_m := m
	if m == n {
		_m = m - 1
		grid[0][b-1] = 1
	}

	for _m > 0 {
		i := findLastSetBit(_m)
		grid[0][b-2-i] = 1
		_m = _m & (_m - 1)
	}

	for i := 1; i < b-1; i++ {
		for j := i + 1; j < b; j++ {
			grid[i][j] = 1
		}
	}

	return grid, true
}

func findLastSetBit(x int64) int {
	var pos uint = 0
	for ((x >> pos) & 1) == 0 {
		pos++
	}
	return int(pos)
}
