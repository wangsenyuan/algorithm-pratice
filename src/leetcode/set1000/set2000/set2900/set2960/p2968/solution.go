package p2968

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	freq := make([]int, n*n+1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			freq[grid[i][j]]++
		}
	}
	// sum = sum2 - b + a
	// a - b = sum - sum2
	res := []int{0, 0}
	for i := 1; i <= n*n; i++ {
		if freq[i] == 2 {
			res[0] = i
		} else if freq[i] == 0 {
			res[1] = i
		}
	}

	return res
}
