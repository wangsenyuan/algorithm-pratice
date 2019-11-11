package p1251

func oddCells(n int, m int, indices [][]int) int {
	row := make([]int, n)
	col := make([]int, m)

	for _, index := range indices {
		x, y := index[0], index[1]
		row[x]++
		col[y]++
	}

	var x, y int

	for i := 0; i < n; i++ {
		x += row[i] & 1
	}

	for j := 0; j < m; j++ {
		y += col[j] & 1
	}

	return x*(m-y) + y*(n-x)
}
