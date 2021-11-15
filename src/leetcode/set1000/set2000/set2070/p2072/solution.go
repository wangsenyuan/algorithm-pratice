package p2072

import "sort"

func timeRequiredToBuy(tickets []int, k int) int {
	ps := make([]Pair, len(tickets))

	for i, tic := range tickets {
		ps[i] = Pair{tic, i}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
	})
	n := len(ps)
	var res int
	var prev int
	for i := 0; i < n; i++ {
		cur := ps[i]
		delta := cur.first - prev
		if cur.second != k {
			res += delta * (n - i)
		} else {
			res += (delta - 1) * (n - i)
			for j := i + 1; j < n; j++ {
				if ps[j].second < cur.second {
					res++
				}
			}
			res++
			break
		}

		prev = cur.first
	}
	return res
}

type Pair struct {
	first, second int
}
