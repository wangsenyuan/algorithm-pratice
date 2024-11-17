package p3355

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	add := make([]int, n+1)

	for _, query := range queries {
		l, r := query[0], query[1]
		add[l]++
		add[r+1]--
	}

	for i := 1; i < n; i++ {
		add[i] += add[i-1]
	}

	for i := range n {
		if add[i] < nums[i] {
			return false
		}
	}
	return true
}
