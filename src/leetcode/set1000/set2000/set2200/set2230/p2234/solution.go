package p2234

import "sort"

const INF = 1 << 20

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {

	sort.Ints(flowers)
	reverse(flowers)
	// 假设前x是perfect的，剩余的都不是perfect的
	n := len(flowers)

	suf := make([]int64, n+1)

	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + int64(flowers[i])
	}

	checkAvg := func(avg int, hand int64, x int) bool {
		j := sort.Search(n, func(j int) bool {
			return flowers[j] < avg
		})

		if j == n {
			return true
		}
		if j < x {
			j = x
		}

		return suf[j]+hand >= int64(avg)*int64(n-j)
	}

	getMin := func(hand int64, x int) int {
		l, r := 1, INF

		for l < r {
			mid := (l + r) / 2
			if !checkAvg(mid, hand, x) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return r - 1
	}

	hand := newFlowers
	var best int64
	var x int
	for x < n {
		if flowers[x] >= target {
			x++
			best = int64(x) * int64(full)
			continue
		}
		tmp := int64(x) * int64(full)
		avg := getMin(hand, x)
		if avg >= target {
			avg = target - 1
		}
		tmp += int64(avg) * int64(partial)

		if best < tmp {
			best = tmp
		}

		if flowers[x] < target {
			// good enough
			hand -= (int64(target) - int64(flowers[x]))
			if hand < 0 {
				break
			}
		}
		x++
	}

	if hand >= 0 && best < int64(x)*int64(full) {
		best = int64(x) * int64(full)
	}

	return best
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
