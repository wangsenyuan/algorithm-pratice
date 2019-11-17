package p1260

func shiftGrid(grid [][]int, k int) [][]int {
	n := len(grid)
	m := len(grid[0])
	row := make([]int, n*m)

	for i := 0; i < n; i++ {
		copy(row[i*m:], grid[i])
	}

	k = k % (n * m)

	swap(row, 0, n*m)
	swap(row, 0, k)
	swap(row, k, n*m)

	ng := make([][]int, n)

	for i := 0; i < n; i++ {
		ng[i] = make([]int, m)
		copy(ng[i], row[i*m:(i+1)*m])
	}

	return ng
}

func swap(arr []int, start, end int) {
	for i, j := start, end-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
