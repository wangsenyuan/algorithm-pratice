package p1854

import "sort"

func maximumPopulation(logs [][]int) int {
	n := len(logs)
	events := make([][]int, 2*n)

	for i := 0; i < n; i++ {
		log := logs[i]
		events[2*i] = []int{log[0], 1}
		events[2*i+1] = []int{log[1], -1}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0] || events[i][0] == events[j][0] && events[i][1] < events[j][1]
	})
	var ans int
	var best int
	var sum int
	for _, cur := range events {
		sum += cur[1]
		if sum > best {
			best = sum
			ans = cur[0]
		}
	}
	return ans
}
