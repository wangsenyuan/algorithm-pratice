package p2769

func checkArray(nums []int, k int) bool {
	n := len(nums)
	diff := make([]int, n+1)
	for i := 0; i < n; i++ {
		if i > 0 {
			diff[i] += diff[i-1]
		}
		x := nums[i] + diff[i]
		if x < 0 || x > 0 && i+k > n {
			return false
		}
		diff[i] -= x
		if x > 0 {
			diff[i+k] += x
		}
	}

	return true
}
