package p5610

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	var sum int
	for i := 0; i < n; i++ {
		sum += nums[i]
		res[i] = (i+1)*nums[i] - sum
	}
	sum = 0
	for i := n - 1; i >= 0; i-- {
		sum += nums[i]
		res[i] += sum - (n-i)*nums[i]
	}

	return res
}
