package p2517

import "sort"

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)

	n := len(price)

	check := func(d int) bool {

		cnt := 1
		prev := price[0]

		for i := 1; i < n && cnt < k; i++ {
			if price[i]-prev >= d {
				cnt++
				prev = price[i]
			}
		}

		return cnt >= k
	}

	var l, r = 0, price[n-1] - price[0] + 1

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r - 1
}
