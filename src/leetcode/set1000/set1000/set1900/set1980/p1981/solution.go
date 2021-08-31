package p1981

import "sort"

/**
You are given an m x n integer matrix mat and an integer target.

Choose one integer from each row in the matrix such that the absolute difference between target and the sum of the chosen elements is minimized.

Return the minimum absolute difference.

The absolute difference between two numbers a and b is the absolute value of a - b.
m, n <= 70
1 <= mat[i][j] <= 70
target <= 800
*/
func minimizeTheDifference(mat [][]int, target int) int {
	// m * n * mat[i][j] <= 70 * 70 * 70
	m := len(mat)
	n := len(mat[0])
	//需要找到一个column_sum >= target
	arr := make([]int, 0, target+1)
	bak := make([]int, 0, (target+1)*n)

	arr = append(arr, 0)

	for i := 0; i < m; i++ {
		bak = bak[:0]
		sort.Ints(mat[i])
		for j := 0; j < n; j++ {
			if j > 0 && mat[i][j] == mat[i][j-1] {
				continue
			}
			for k := 0; k < len(arr); k++ {
				bak = append(bak, arr[k]+mat[i][j])
				if arr[k]+mat[i][j] >= target {
					break
				}
			}
		}
		sort.Ints(bak)
		arr = arr[:0]
		for u := 1; u <= len(bak); u++ {
			if u == len(bak) || bak[u] > bak[u-1] {
				arr = append(arr, bak[u-1])
				if bak[u-1] >= target {
					break
				}
			}
		}
	}
	x := sort.SearchInts(arr, target)
	if x < len(arr) {
		if x > 0 && arr[x]-target > target-arr[x-1] {
			return target - arr[x-1]
		}
		return arr[x] - target
	}
	return target - arr[x-1]
}
