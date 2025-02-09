package p3446

import "sort"

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			arr[i-j] = append(arr[i-j], grid[i][j])
		}
	}

	for i := 0; i < n; i++ {
		sort.Ints(arr[i])
		reverse(arr[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			grid[i][j] = arr[i-j][j]
		}
	}

	for i := 0; i < n; i++ {
		arr[i] = arr[i][:0]
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			arr[j-i] = append(arr[j-i], grid[i][j])
		}
	}
	for i := 0; i < n; i++ {
		sort.Ints(arr[i])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			grid[i][j] = arr[j-i][i]
		}
	}
	return grid
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
