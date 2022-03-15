package p2200

func findKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	var res []int

	for i := 0; i < n; i++ {
		if nums[i] == key {
			a := max(0, i-k)
			if len(res) > 0 {
				a = max(a, res[len(res)-1]+1)
			}
			b := min(n-1, i+k)
			for j := a; j <= b; j++ {
				res = append(res, j)
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
