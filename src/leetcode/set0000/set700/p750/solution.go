package p750

func countCornerRectangles(grid [][]int) int {
	n := len(grid)
	if n <= 1 {
		return 0
	}

	m := len(grid[0])
	if m <= 1 {
		return 0
	}
	var ans int
	rows := make([][]int, n)
	cols := make([][]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				if rows[i] == nil {
					rows[i] = make([]int, 0, m)
				}
				if cols[j] == nil {
					cols[j] = make([]int, 0, n)
				}

				for _, a := range rows[i] {
					for _, b := range cols[j] {
						if grid[b][a] == 1 {
							ans++
						}
					}
				}

				rows[i] = append(rows[i], j)
				cols[j] = append(cols[j], i)
			}
		}
	}

	return ans
}
