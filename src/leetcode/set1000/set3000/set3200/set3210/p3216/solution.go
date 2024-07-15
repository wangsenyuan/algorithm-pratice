package p3216

import "sort"

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	sort.Ints(horizontalCut)
	sort.Ints(verticalCut)
	i, j := m-1, n-1
	var ans int
	cnt := make([]int, 2)
	for i > 0 || j > 0 {
		if j == 0 || i > 0 && horizontalCut[i-1] > verticalCut[j-1] {
			ans += horizontalCut[i-1] * (cnt[0] + 1)
			cnt[1]++
			i--
		} else {
			ans += verticalCut[j-1] * (cnt[1] + 1)
			cnt[0]++
			j--
		}
	}
	return int64(ans)
}
