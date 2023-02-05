package p2565

func minCapability(nums []int, k int) int {
	n := len(nums)

	check := func(mid int) bool {
		var cnt int
		for i := 0; i < n; {
			if nums[i] > mid {
				i++
				continue
			}
			j := i
			for i < n && nums[i] <= mid {
				i++
			}

			cnt += count(i - j)
		}

		return cnt >= k
	}

	// 最好是窃取最小的的k个nums
	// 但是有连续的情况
	var l, r int = 0, 1e9

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

func count(n int) int {
	if n <= 2 {
		return 1
	}
	return (n + 1) / 2
}
