package p1040

import "sort"

const INF = 1 << 20

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	n := len(stones)
	ans := INF
	ans2 := max(stones[n-1]-stones[1], stones[n-2]-stones[0]) + 2 - n
	for i, j := 0, 0; j < n; j++ {
		for stones[j]-stones[i]+1 > n {
			i++
		}
		if j-i+1 == n-1 && stones[j]-stones[i] == n-2 {
			ans = min(ans, 2)
		} else {
			ans = min(ans, n-(j-i+1))
		}
	}

	return []int{ans, ans2}
}

func numMovesStonesII_1(stones []int) []int {
	sort.Ints(stones)

	n := len(stones)

	// n >= 3
	first := stones[0]
	last := stones[n-1]

	if last-first+1 == n {
		// all full
		return []int{0, 0}
	}

	var slide func(i int)

	ans := INF

	slide = func(i int) {
		if i < n {
			out := i

			j := sort.Search(n, func(j int) bool {
				return stones[j] > stones[i]+n-1
			})
			if j < n {
				out += n - j
			}

			if out > 1 {
				// any way can fit in
				ans = min(ans, out)
			} else if out == 1 {
				// special case
				if stones[j-1] == stones[i]+n-1 {
					// there is a gap in between
					ans = min(ans, 1)
				} else {
					ans = min(ans, 2)
				}
			}
			slide(i + 1)
		}
	}

	slide(0)

	ans2 := max(stones[n-1]-stones[1], stones[n-2]-stones[0]) + 2 - n

	return []int{ans, ans2}
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
