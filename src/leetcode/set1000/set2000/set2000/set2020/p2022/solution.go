package p2022

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	grid := make([][]int, m)
	for r := 0; r < m; r++ {
		grid[r] = make([]int, n)
		copy(grid[r], original[r*n:(r+1)*n])
	}
	return grid
}
