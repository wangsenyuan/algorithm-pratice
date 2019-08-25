package p1138

import (
	"bytes"
)

func alphabetBoardPath(target string) string {
	board := []string{
		"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z",
	}
	pos := make([][]int, 26)
	for i := 0; i < len(board); i++ {
		row := board[i]
		for j := 0; j < len(row); j++ {
			x := int(row[j] - 'a')
			pos[x] = []int{i, j}
		}
	}

	var buf bytes.Buffer
	var x, y int
	for u := 0; u < len(target); u++ {
		c := target[u] - 'a'
		i, j := pos[c][0], pos[c][1]
		if x < 5 {
			for j < y {
				buf.WriteByte('L')
				y--
			}
			for j > y {
				buf.WriteByte('R')
				y++
			}

			for i < x {
				buf.WriteByte('U')
				x--
			}
			for i > x {
				buf.WriteByte('D')
				x++
			}
		} else {
			for i < x {
				buf.WriteByte('U')
				x--
			}

			for j > y {
				buf.WriteByte('R')
				y++
			}
		}
		buf.WriteByte('!')
	}
	return buf.String()
}
