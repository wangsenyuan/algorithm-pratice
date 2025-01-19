package p3427

func subarraySum(nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}

	var res int

	for i := range n {
		j := max(0, i-nums[i])
		res += sum[i+1] - sum[j]
	}

	return res
}
