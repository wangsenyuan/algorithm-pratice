package p1000

import "fmt"

var dd = []int{-1, 0, 1}

func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
	row := make(map[int]int)
	col := make(map[int]int)
	dia1 := make(map[int]int)
	dia2 := make(map[int]int)

	grid := make(map[string]int)

	cal1 := func(x, y int) int {
		return x + y
	}

	cal2 := func(x, y int) int {
		return N - 1 - x + y
	}

	for _, lamp := range lamps {
		x, y := lamp[0], lamp[1]
		row[x]++
		col[y]++
		dia1[cal1(x, y)]++
		dia2[cal2(x, y)]++
		grid[fmt.Sprintf("%d %d", x, y)]++
	}

	ans := make([]int, len(queries))

	for i, query := range queries {
		x, y := query[0], query[1]
		if row[x] > 0 || col[y] > 0 || dia1[cal1(x, y)] > 0 || dia2[cal2(x, y)] > 0 {
			ans[i] = 1

			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					u, v := x+dd[j], y+dd[k]
					key := fmt.Sprintf("%d %d", u, v)
					if grid[key] > 0 {
						delete(grid, key)
						row[u]--
						col[v]--
						dia1[cal1(u, v)]--
						dia2[cal2(u, v)]--
					}
				}
			}
		}

	}

	return ans
}
