package p2187

import "sort"

func minimumTime(time []int, totalTrips int) int64 {
	// sum(ans / time[i]) >= totalTrips
	sort.Ints(time)
	n := len(time)
	total := int64(totalTrips)
	check := func(ans int64) bool {
		var trips int64

		for i := 0; i < n && trips < total; i++ {
			cur := ans / int64(time[i])
			if cur == 0 {
				break
			}
			trips += cur
		}
		return trips >= total
	}

	var left int64
	var right int64 = 1 << 60

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
