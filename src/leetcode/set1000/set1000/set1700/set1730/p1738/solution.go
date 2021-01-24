package p1738

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	nums := make([]int, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				nums[i*n+j] = matrix[i][j]
				continue
			}
			if i > 0 {
				matrix[i][j] ^= matrix[i-1][j]
			}
			if j > 0 {
				matrix[i][j] ^= matrix[i][j-1]
			}
			if i > 0 && j > 0 {
				matrix[i][j] ^= matrix[i-1][j-1]
			}
			nums[i*n+j] = matrix[i][j]
		}
	}

	sort.Ints(nums)

	return nums[m*n-k]
}
