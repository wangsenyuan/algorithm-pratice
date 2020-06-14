package p1480

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	cnt := make(map[int]int)

	for _, num := range arr {
		cnt[num]++
	}

	xx := make([]int, 0, len(cnt))

	for _, v := range cnt {
		xx = append(xx, v)
	}

	sort.Ints(xx)

	var i int
	var kk int

	for i < len(xx) && kk+xx[i] <= k {
		kk += xx[i]
		i++
	}

	return len(xx) - i
}
