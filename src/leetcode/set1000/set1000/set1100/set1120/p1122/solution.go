package p1122

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	ii := make(map[int]int)

	for i := 0; i < len(arr2); i++ {
		ii[arr2[i]] = i
	}

	ps := make(Ps, len(arr1))

	for i := 0; i < len(arr1); i++ {
		if j, found := ii[arr1[i]]; found {
			ps[i] = Pair{j, arr1[i]}
		} else {
			ps[i] = Pair{len(arr1), arr1[i]}
		}
	}

	sort.Sort(ps)

	res := make([]int, len(arr1))

	for i := 0; i < len(arr1); i++ {
		res[i] = ps[i].second
	}

	return res
}

type Pair struct {
	first  int
	second int
}

type Ps []Pair

func (this Ps) Len() int {
	return len(this)
}

func (this Ps) Less(i, j int) bool {
	return this[i].first < this[j].first || (this[i].first == this[j].first && this[i].second < this[j].second)
}

func (this Ps) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
