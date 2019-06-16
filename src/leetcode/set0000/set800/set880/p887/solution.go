package p887

func projectionArea(grid [][]int) int {
	var area1, area2, area3 int

	n := len(grid)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				area1++
			}
		}
	}

	for i := 0; i < n; i++ {
		var height int
		for j := 0; j < n; j++ {
			if grid[i][j] > height {
				height = grid[i][j]
			}
		}
		area2 += height
	}

	for j := 0; j < n; j++ {
		var height int
		for i := 0; i < n; i++ {
			if grid[i][j] > height {
				height = grid[i][j]
			}
		}
		area3 += height
	}

	return area1 + area2 + area3
}
