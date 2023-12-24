package p2974

import "sort"

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	sort.Ints(hFences)
	sort.Ints(vFences)

	widths := make(map[int]int)

	for i := 0; i < len(vFences); i++ {
		x := vFences[i]
		widths[x-1]++
		for j := 0; j < i; j++ {
			widths[x-vFences[j]]++
		}
		widths[n-x]++
	}
	widths[n-1]++

	ans := 0

	for i := 0; i < len(hFences); i++ {
		x := hFences[i]
		h := x - 1
		if widths[h] > 0 {
			ans = max(ans, h*h)
		}
		for j := 0; j < i; j++ {
			h = x - hFences[j]
			if widths[h] > 0 {
				ans = max(ans, h*h)
			}
		}
		h = m - x
		if widths[h] > 0 {
			ans = max(ans, h*h)
		}
	}
	if widths[m-1] > 0 {
		ans = max(ans, (m-1)*(m-1))
	}

	if ans == 0 {
		return -1
	}
	return ans % (1e9 + 7)
}
