package p2594

import "sort"

func repairCars(ranks []int, cars int) int64 {
	n := len(ranks)
	t := int64(cars) * int64(cars) * int64(ranks[0])

	sort.Ints(ranks)

	check := func(x int64) bool {
		r := cars
		var cnt int
		for i := 0; i < n && cnt < cars; i++ {
			m := search(r+1, func(j int) bool {
				return int64(ranks[i])*int64(j)*int64(j) > x
			}) - 1
			cnt += m
			r = m
		}
		return cnt >= cars
	}

	return search(t, check)
}

type Num interface {
	int | int64
}

func search[T Num](n T, f func(T) bool) T {
	var l, r T = 0, n
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
