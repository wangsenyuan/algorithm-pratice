package lcp78

const inf = 1 << 30

func rampartDefensiveLine(rampart [][]int) int {
	n := len(rampart)
	// 扩张后的位置，必须包含原来的位置

	check := func(x int) bool {
		// can expand x units or not
		prev := rampart[0][1]
		for i := 1; i+1 < n; i++ {
			if prev > rampart[i][0] {
				return false
			}
			dist := rampart[i][0] - prev
			if dist >= x {
				prev = rampart[i][1]
				continue
			}
			// dist <= x
			prev = rampart[i][1] + x - dist
		}
		return prev <= rampart[n-1][0]
	}

	l, r := 0, inf
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
