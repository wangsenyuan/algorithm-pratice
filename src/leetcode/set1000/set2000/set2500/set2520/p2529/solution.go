package p2529

func maximumCount(nums []int) int {
	n := len(nums)
	i := search(n, func(i int) bool {
		return nums[i] > 0
	})
	x := n - i
	y := search(n, func(j int) bool {
		return nums[j] >= 0
	})
	return max(x, y)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
