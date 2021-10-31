package p2054

import "sort"

func maxTwoEvents(events [][]int) int {
	n := len(events)

	// need to find the best score before start - 1
	type Point struct {
		first, second int
		tp            int
	}

	P := make([]Point, 2*n)

	for i := 0; i < n; i++ {
		cur := events[i]
		P[2*i] = Point{cur[0], cur[2], 0}
		P[2*i+1] = Point{cur[1], cur[2], 1}
	}

	sort.Slice(P, func(i, j int) bool {
		// when equal, make start before end
		return P[i].first < P[j].first || P[i].first == P[j].first && P[i].tp < P[j].tp
	})

	var best int
	var mostValue int

	for _, cur := range P {
		if cur.tp == 0 {
			// enter
			best = max(best, mostValue+cur.second)
		} else {
			mostValue = max(mostValue, cur.second)
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
