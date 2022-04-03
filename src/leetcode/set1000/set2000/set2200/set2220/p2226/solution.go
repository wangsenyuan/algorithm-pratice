package p2226

const MAX_X = 10000010

func maximumCandies(candies []int, k int64) int {
	n := len(candies)
	check := func(took int) bool {
		if took == 0 {
			return true
		}
		var cnt int64
		for i := 0; i < n; i++ {
			cnt += int64(candies[i]) / int64(took)
		}
		return cnt >= k
	}

	lo, hi := 0, MAX_X

	for lo < hi {
		mid := (lo + hi) / 2
		if !check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi - 1
}
