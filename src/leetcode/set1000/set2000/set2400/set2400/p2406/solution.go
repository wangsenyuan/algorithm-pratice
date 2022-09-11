package p2406

import "sort"

func minGroups(intervals [][]int) int {

	n := len(intervals)

	events := make([]Pair, 2*n)

	for i := 0; i < n; i++ {
		events[2*i] = Pair{intervals[i][0], 1}
		events[2*i+1] = Pair{intervals[i][1] + 1, -1}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].first < events[j].first ||
			events[i].first == events[j].first && events[i].second < events[j].second
	})

	var cnt int
	var best int

	for _, evt := range events {
		cnt += evt.second
		best = max(best, cnt)
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}
