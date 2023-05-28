package p2709

func differenceOfDistinctValues(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	get := func(x, y int) int {
		a := make(map[int]int)
		for i := 1; x-i >= 0 && y-i >= 0; i++ {
			a[grid[x-i][y-i]]++
		}

		b := make(map[int]int)

		for i := 1; x+i < m && y+i < n; i++ {
			b[grid[x+i][y+i]]++
		}

		return abs(len(a) - len(b))
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = get(i, j)
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
