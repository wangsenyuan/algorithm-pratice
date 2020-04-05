package p1405

import (
	"bytes"
	"sort"
)

func longestDiverseString(a int, b int, c int) string {
	// x, y, z

	ps := []*Pair{&Pair{a, 0}, &Pair{b, 1}, &Pair{c, 2}}
	n := a + b + c
	arr := make([]*Pair, n)
	var i int
	for i < n {
		sort.Sort(Pairs(ps))
		if ps[0].first == 0 {
			break
		}
		arr[i] = &Pair{1, ps[0].second}
		ps[0].first--

		i++

		if i == n {
			break
		}

		if ps[1].first == 0 {
			break
		}

		arr[i] = &Pair{1, ps[1].second}
		ps[1].first--
		i++
	}

	n = i
	if ps[0].first > 0 {
		for i := 0; i < n && ps[0].first > 0; i++ {
			if arr[i].second == ps[0].second {
				arr[i].first++
				ps[0].first--
			}
		}
	}
	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		a := byte('a' + arr[i].second)
		for arr[i].first > 0 {
			buf.WriteByte(a)
			arr[i].first--
		}
	}
	return buf.String()
}

type Pair struct {
	first, second int
}

type Pairs []*Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first > ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
