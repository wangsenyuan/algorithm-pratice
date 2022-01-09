package p2134

func minSwaps(nums []int) int {
	n := len(nums)
	sum := make([]int, n)

	for i := 0; i < len(nums); i++ {
		sum[i] = nums[i]
		if i > 0 {
			sum[i] += sum[i-1]
		}
	}
	cnt := sum[n-1]
	if cnt <= 1 {
		return 0
	}
	best := n
	for i := 0; i < n; i++ {
		//if we get [i...i+cnt-1] as 1
		var tmp int
		if i+cnt-1 < n {
			tmp = sum[i+cnt-1]
		} else {
			tmp = sum[n-1]
			tmp += sum[(i+cnt-1)%n]
		}
		if i > 0 {
			tmp -= sum[i-1]
		}
		best = min(best, cnt-tmp)
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
