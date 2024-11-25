package p3364

func minimumSumSubarray(nums []int, l int, r int) int {
	res := 1 << 30

	n := len(nums)
	for i := 0; i < n; i++ {
		var sum int

		for j := i; j < n; j++ {
			sum += nums[j]
			if j-i+1 > r {
				break
			}
			if j-i+1 >= l && sum > 0 {
				res = min(res, sum)
			}
		}
	}
	if res == 1<<30 {
		return -1
	}
	return res
}
