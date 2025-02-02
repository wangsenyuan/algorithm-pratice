package p3437

import "strings"

func maxDistance(s string, k int) int {

	const nwse = "NSEW"
	var best int
	cnt := make([]int, 4)
	tmp := make([]int, 4)
	for _, x := range []byte(s) {
		i := strings.IndexByte(nwse, x)
		cnt[i]++
		copy(tmp, cnt)
		kk := k
		if tmp[0] < tmp[1] {
			tmp[0], tmp[1] = tmp[1], tmp[0]
		}
		u := min(kk, tmp[1])
		tmp[0] += u
		tmp[1] -= u
		kk -= u

		if tmp[2] < tmp[3] {
			tmp[2], tmp[3] = tmp[3], tmp[2]
		}
		u = min(kk, tmp[3])
		tmp[3] -= u
		tmp[2] += u
		kk -= u
		best = max(best, tmp[0]-tmp[1]+tmp[2]-tmp[3])
	}

	return best
}

func abs(num int) int {
	return max(num, -num)
}
