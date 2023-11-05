package p2926

func findChampion(grid [][]int) int {
	n := len(grid)

	deg := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// i beats j
				deg[j]++
			}
		}
	}
	winer := -1
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			if winer >= 0 {
				return -1
			}
			winer = i
		}
	}

	return winer
}
