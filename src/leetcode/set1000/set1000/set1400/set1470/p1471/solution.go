package p1471

import "sort"

func getStrongest(arr []int, k int) []int {
	n := len(arr)

	sort.Ints(arr)

	x := arr[(n-1)/2]

	tmp := make([]Pair, n)

	for i := 0; i < n; i++ {
		if arr[i] < x {
			tmp[i] = Pair{x - arr[i], arr[i]}
		} else {
			tmp[i] = Pair{arr[i] - x, arr[i]}
		}
	}

	sort.Sort(Pairs(tmp))

	res := make([]int, k)

	for i := 0; i < k; i++ {
		res[i] = tmp[i].second
	}

	return res
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps Pairs) Less(i, j int) bool {
	a, b := ps[i], ps[j]
	return a.first > b.first || a.first == b.first && a.second > b.second
}
