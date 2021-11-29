package p2089

func getAverages(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	var sum int
	for j := 0; j < 2*k && j < n; j++ {
		sum += nums[j]
	}

	for i, j := 0, 2*k; j < n; i, j = i+1, j+1 {
		sum += nums[j]
		res[j-k] = sum / (j - i + 1)
		sum -= nums[i]
	}

	return res
}
