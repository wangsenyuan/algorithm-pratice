package p5540

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	ps := make([]Pair, len(points))

	for i := 0; i < len(points); i++ {
		ps[i] = Pair{points[i][0], points[i][1]}
	}

	sort.Sort(Pairs(ps))
	var best int
	for i := 0; i < len(ps); {
		j := i
		for j < len(ps) && ps[j].first == ps[i].first {
			j++
		}

		if j < len(ps) {
			if ps[j].first-ps[i].first > best {
				best = ps[j].first - ps[i].first
			}
		}

		i = j
	}

	return best
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
