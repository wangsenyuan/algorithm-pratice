package p2908

const inf = 1 << 30

func minimumSum(nums []int) int {
	n := len(nums)
	L := make([]int, n)

	L[0] = nums[0]

	for i := 1; i < n; i++ {
		L[i] = min(L[i-1], nums[i])
	}
	var best = inf
	R := nums[n-1]

	for i := n - 2; i > 0; i-- {
		if L[i-1] < nums[i] && nums[i] > R {
			best = min(best, L[i-1]+nums[i]+R)
		}
		R = min(R, nums[i])
	}

	if best == inf {
		return -1
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
