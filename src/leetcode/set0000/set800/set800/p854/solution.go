package p854

import (
	"math"
)

func kSimilarity(A string, B string) int {

	var loop func(a, b string, idx int) int

	cache := make(map[string]int)

	loop = func(a, b string, idx int) int {
		if idx >= len(a) {
			return 0
		}
		c := a + "-" + b
		if _, found := cache[c]; !found {
			var ans int

			if a[idx] == b[idx] {
				ans = loop(a, b, idx+1)
			} else {
				ans = math.MaxInt32
				for i := idx + 1; i < len(a); i++ {
					if a[i] == b[idx] {
						x := a[:idx] + b[idx:idx+1] + a[idx+1:i] + a[idx:idx+1] + a[i+1:]
						ans = min(ans, loop(x, b, idx+1)+1)
					}
				}

				if ans2, found2 := cache[c]; !found2 || ans2 > ans {
					cache[c] = ans
				}
			}
		}

		return cache[c]
	}

	return loop(A, B, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
