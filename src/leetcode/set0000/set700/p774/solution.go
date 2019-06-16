package p774

import (
	"math"
	"sort"
)

func minmaxGasDist(stations []int, K int) float64 {
	sort.Ints(stations)

	n := len(stations)
	var max int

	for i := 1; i < n; i++ {
		if stations[i]-stations[i-1] > max {
			max = stations[i] - stations[i-1]
		}
	}

	check := func(d float64) bool {
		k := K
		for i := 1; i < n; i++ {
			dist := stations[i] - stations[i-1]
			fd := float64(dist)
			bd := math.Abs(d - fd)
			if bd < 0.000001 {
				continue
			}
			need := int(fd / d)
			if float64(need)*d < fd {
				need++
			}
			k -= (need - 1)
			if k < 0 {
				return false
			}
		}
		return true
	}

	left := float64(0.0)
	right := float64(max)

	for right-left > 0.000001 {
		mid := left + (right-left)/2.0
		tmp := check(mid)
		if tmp {
			right = mid
		} else {
			left = mid
		}
	}
	return left
}
