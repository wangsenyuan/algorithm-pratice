package p793

func preimageSizeFZF(K int) int {
	if K == 0 {
		return 5
	}
	return findLeastNum(K+1) - findLeastNum(K)
}

func findLeastNum(n int) int {
	if n == 1 {
		return 5
	}
	check := func(p int) bool {
		tmp, count, f := p, 0, 5
		for f <= tmp {
			count += tmp / f
			f = f * 5
		}
		return count >= n
	}

	left, right := 0, 5*n
	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
