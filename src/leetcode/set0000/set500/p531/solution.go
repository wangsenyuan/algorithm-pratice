package p531

func findLonelyPixel(picture [][]byte) int {
	n := len(picture)
	if n == 0 {
		return 0
	}

	m := len(picture[0])
	if m == 0 {
		return 0
	}

	row := make([]int, n)
	col := make([]int, m)

	for i := 0; i < len(picture); i++ {
		for j := 0; j < len(picture[i]); j++ {
			if picture[i][j] == 'B' {
				row[i]++
				col[j]++
			}
		}
	}

	var res int

	for i := 0; i < len(picture); i++ {
		for j := 0; j < len(picture[i]); j++ {
			if picture[i][j] == 'B' && row[i] == 1 && col[j] == 1 {
				res++
			}
		}
	}

	return res
}
