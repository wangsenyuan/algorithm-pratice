package p757

import "sort"

func intersectionSizeTwo(intervals [][]int) int {
	sort.Sort(IVS(intervals))
	n := len(intervals)
	ans, max1, max2 := 0, -1, -1

	for i := 0; i < n; i++ {
		if intervals[i][0] > max1 {
			ans += 2
			max1 = intervals[i][1]
			max2 = max1 - 1
		} else if intervals[i][0] > max2 {
			ans++
			if max1 == intervals[i][1] {
				max2 = max1 - 1
			} else {
				max2 = max1
			}
			max1 = intervals[i][1]
		}
	}

	return ans
}

type IVS [][]int

func (this IVS) Len() int {
	return len(this)
}

func (this IVS) Less(i, j int) bool {
	return this[i][1] < this[j][1]
}

func (this IVS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
