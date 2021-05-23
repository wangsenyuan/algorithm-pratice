package p1870

import "math"

const eps = 1e-8
const MAX = 10000001

func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	check := func(speed int) bool {
		var tot float64
		for i := 0; i < n-1; i++ {
			d := dist[i]
			x := float64(d) / float64(speed)
			y := math.Floor(x)
			if x-y > eps {
				y++
			}
			tot += y
		}

		tot += float64(dist[n-1]) / float64(speed)

		return tot <= hour
	}

	left := 1
	right := MAX

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if right == MAX {
		return -1
	}
	return right
}
