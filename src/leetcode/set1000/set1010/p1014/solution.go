package p1014

func shipWithinDays(weights []int, D int) int {
	n := len(weights)
	check := func(W int) bool {
		var sum int
		res := 1
		for i := 0; i < n; i++ {
			if weights[i] > W {
				return false
			}
			sum += weights[i]
			if sum > W {
				res++
				sum = weights[i]
			}
		}
		return res <= D
	}

	left := 1
	right := 1 << 30
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
