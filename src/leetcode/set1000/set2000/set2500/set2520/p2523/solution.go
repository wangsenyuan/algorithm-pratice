package p2523

func closestPrimes(left int, right int) []int {
	flag := make([]bool, right+1)

	// n := right + 1
	var ans []int

	check := func(nums1, nums2 int) bool {
		if len(ans) == 0 {
			return true
		}
		if nums2-nums1 < ans[1]-ans[0] {
			return true
		}
		return false
	}

	prev := -1
	for i := 2; i <= right; i++ {
		if !flag[i] {
			if prev >= left {
				if check(prev, i) {
					ans = []int{prev, i}
				}
			}
			prev = i
			for j := i * i; j <= right; j += i {
				flag[j] = true
			}
		}
	}

	if len(ans) == 0 {
		return []int{-1, -1}
	}

	return ans
}
