package c

import "sort"

func getTriggerTime(increase [][]int, requirements [][]int) []int {
	n := len(increase)

	cs := make([]Pair, n+1)
	rs := make([]Pair, n+1)
	hs := make([]Pair, n+1)
	cs[0] = Pair{0, 0}
	rs[0] = Pair{0, 0}
	hs[0] = Pair{0, 0}
	var a, b, c int
	for i := 0; i < n; i++ {
		a += increase[i][0]
		b += increase[i][1]
		c += increase[i][2]
		cs[i+1] = Pair{a, i + 1}
		rs[i+1] = Pair{b, i + 1}
		hs[i+1] = Pair{c, i + 1}
	}

	find := func(req []int) int {
		i := sort.Search(n+1, func(i int) bool {
			return cs[i].v >= req[0]
		})

		if i > n {
			return -1
		}

		j := sort.Search(n+1, func(j int) bool {
			return rs[j].v >= req[1]
		})

		if j > n {
			return -1
		}

		k := sort.Search(n+1, func(k int) bool {
			return hs[k].v >= req[2]
		})

		if k > n {
			return -1
		}

		return max3(cs[i].w, rs[j].w, hs[k].w)
	}

	res := make([]int, len(requirements))

	for i := 0; i < len(requirements); i++ {
		res[i] = find(requirements[i])
	}

	return res
}

type Pair struct {
	v int
	w int
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}
