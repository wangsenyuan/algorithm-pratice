package p2221

func triangularSum(nums []int) int {
	n := len(nums)

	arr := make([]int, n)

	for n > 1 {
		for i := 0; i+1 < n; i++ {
			arr[i] = (nums[i] + nums[i+1]) % 10
		}
		copy(nums[:n-1], arr[:n-1])
		n--
	}

	return nums[0]
}
