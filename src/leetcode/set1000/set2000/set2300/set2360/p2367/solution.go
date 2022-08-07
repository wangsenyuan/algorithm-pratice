package p2367

func arithmeticTriplets(nums []int, diff int) int {
	var res int

	n := len(nums)

	for i := 0; i < n; i++ {
		j := search(n, func(j int) bool {
			return nums[j]-nums[i] >= diff
		})
		if j < n && nums[j]-nums[i] == diff {
			k := search(n, func(k int) bool {
				return nums[k]-nums[j] >= diff
			})
			if k < n && nums[k]-nums[j] == diff {
				res++
			}
		}
	}
	return res
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
