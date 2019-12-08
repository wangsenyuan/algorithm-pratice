package p1282

import "sort"

func groupThePeople(groupSizes []int) [][]int {
	n := len(groupSizes)

	ps := make([]Pair, n)

	for i := 0; i < n; i++ {
		ps[i] = Pair{groupSizes[i], i}
	}

	sort.Sort(Pairs(ps))

	belongTo := make([]int, n)

	var j int
	var i int
	var g int
	for i < n {
		if i-j == ps[j].first {
			g++
			j = i
		}

		belongTo[ps[i].second] = g
		i++
	}
	g++
	ans := make([][]int, g)
	for i := 0; i < g; i++ {
		ans[i] = make([]int, 0, 10)
	}

	for i := 0; i < n; i++ {
		ans[belongTo[i]] = append(ans[belongTo[i]], i)
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
