package p2545

import "sort"

func sortTheStudents(score [][]int, k int) [][]int {
	type Pair struct {
		first  int
		second int
	}

	n := len(score)
	sk := make([]Pair, n)
	for i := 0; i < n; i++ {
		sk[i] = Pair{score[i][k], i}
	}

	sort.Slice(sk, func(i, j int) bool {
		return sk[i].first > sk[j].first
	})
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = score[sk[i].second]
	}

	return res
}
