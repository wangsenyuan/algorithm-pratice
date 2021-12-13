package p2039

import "math"

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	n := len(nums2)

	check := func(x int64) bool {
		var cnt int64
		for i := 0; i < len(nums1) && cnt < k; i++ {
			a := nums1[i]
			if a == 0 {
				if x >= 0 {
					cnt += int64(n)
				}
				continue
			}
			if a > 0 {
				j := search(n, func(j int) bool {
					return int64(a)*int64(nums2[j]) > x
				})
				cnt += int64(j)
			} else {
				// a < 0
				j := search(n, func(j int) bool {
					return int64(a)*int64(nums2[j]) <= x
				})
				cnt += int64(n - j)
			}
		}

		return cnt >= k
	}

	var l, r int64 = math.MinInt64 >> 2, math.MaxInt64 >> 2

	for l < r {
		mid := l + (r-l)/2

		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

func search(n int, fn func(int) bool) int {
	l, r := 0, n

	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
