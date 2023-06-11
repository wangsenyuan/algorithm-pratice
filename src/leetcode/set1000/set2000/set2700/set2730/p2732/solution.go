package p2732

func goodSubsetofBinaryMatrix(grid [][]int) []int {
	n := len(grid)
	m := len(grid[0])

	for i := 0; i < n; i++ {
		ok := true
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				ok = false
				break
			}
		}
		if ok {
			return []int{i}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ok := true

			for k := 0; k < m; k++ {
				if grid[i][k]+grid[j][k] == 2 {
					ok = false
					break
				}
			}
			if ok {
				return []int{i, j}
			}
		}
	}
	return nil
}
