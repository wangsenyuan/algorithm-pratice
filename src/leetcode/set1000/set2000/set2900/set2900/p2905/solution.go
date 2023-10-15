package p2905

import "sort"

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	n := len(nums)
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{nums[i], i}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first < ps[j].first
	})

	ends := []int{n, -1}

	add := func(i int) {
		ends[0] = min(ends[0], i)
		ends[1] = max(ends[1], i)
	}

	for l, r := 0, 0; r < n; r++ {
		for l <= r && ps[r].first-ps[l].first >= valueDifference {
			add(ps[l].second)
			l++
		}
		if ends[0] < n && abs(ends[0]-ps[r].second) >= indexDifference {
			return []int{ends[0], ps[r].second}
		}
		if ends[1] >= 0 && abs(ends[1]-ps[r].second) >= indexDifference {
			return []int{ends[1], ps[r].second}
		}
	}

	return []int{-1, -1}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
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

type Pair struct {
	first  int
	second int
}
