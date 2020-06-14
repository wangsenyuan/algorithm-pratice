package p1479

func runningSum(nums []int) []int {
	n := len(nums)

	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = nums[i]
		if i > 0 {
			res[i] += res[i-1]
		}
	}
	return res
}
