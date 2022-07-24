package p2352

import "bytes"

func equalPairs(grid [][]int) int {
	n := len(grid)
	row := make(map[string]int)

	for i := 0; i < n; i++ {
		var buf bytes.Buffer
		for j := 0; j < n; j++ {
			buf.WriteRune(rune(grid[i][j]))
		}
		row[buf.String()]++
	}

	var res int
	for j := 0; j < n; j++ {
		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			buf.WriteRune(rune(grid[i][j]))
		}
		s := buf.String()
		res += row[s]
	}

	return res
}
