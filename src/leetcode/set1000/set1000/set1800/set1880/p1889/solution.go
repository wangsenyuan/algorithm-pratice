package p1889

import "sort"

const MOD = 1000000007

func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)

	sum := make([]int64, len(packages)+1)

	for i := 0; i < len(packages); i++ {
		sum[i+1] = sum[i] + int64(packages[i])
	}

	var best int64 = -1

	for i := 0; i < len(boxes); i++ {
		sort.Ints(boxes[i])
		var cur int64
		var prev int
		for j := 0; j < len(boxes[i]); j++ {
			k := search(len(packages), func(k int) bool {
				return packages[k] > boxes[i][j]
			})
			if k == 0 {
				continue
			}
			cur += int64(k-prev)*int64(boxes[i][j]) - (sum[k] - sum[prev])
			prev = k
		}
		if prev < len(packages) {
			continue
		}
		if best < 0 || cur < best {
			best = cur
		}
	}

	if best < 0 {
		return -1
	}

	return int(best % MOD)
}

func search(n int, fn func(i int) bool) int {
	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
