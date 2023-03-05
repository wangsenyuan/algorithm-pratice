package p2580

import "sort"

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	var cnt int

	pos := ranges[0][1]

	for i := 1; i <= len(ranges); i++ {
		if i == len(ranges) || ranges[i][0] > pos {
			cnt++
			if i < len(ranges) {
				pos = ranges[i][1]
			}
		} else {
			pos = max(pos, ranges[i][1])
		}
	}

	return pow(2, cnt)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

const MOD = 1e9 + 7

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}
