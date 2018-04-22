package p822

import "sort"

func flipgame(fronts []int, backs []int) int {
	n := len(fronts)
	ps := make(PS, n)
	rps := make(PS, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{fronts[i], backs[i]}
		rps[i] = Pair{backs[i], fronts[i]}
	}

	x := find(n, ps)
	y := find(n, rps)
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	return min(x, y)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func find(n int, ps PS) int {
	sort.Sort(ps)

	for i := 0; i < n; {
		back := ps[i].back
		j := i
		can := true
		for j < n && ps[j].back == back {
			if ps[j].front == back {
				can = false
			}
			j++
		}

		if can {
			return back
		}
		i = j
	}
	return 0
}

type Pair struct {
	front, back int
}

type PS []Pair

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i].back < ps[j].back
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
