package p826

import "sort"

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	m := len(difficulty)
	n := len(worker)

	ps := make([]Pair, m)

	for i := 0; i < m; i++ {
		ps[i] = Pair{difficulty[i], profit[i]}
	}

	sort.Sort(PS(ps))

	maxProfits := make([]int, m)
	maxProfits[0] = ps[0].profit
	for i := 1; i < m; i++ {
		maxProfits[i] = maxProfits[i-1]
		if ps[i].profit > maxProfits[i] {
			maxProfits[i] = ps[i].profit
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		afford := worker[i]
		k := sort.Search(m, func(j int) bool {
			return ps[j].difficulty > afford
		})
		if k > 0 {
			ans += maxProfits[k-1]
		}
	}

	return ans
}

type Pair struct {
	difficulty, profit int
}

type PS []Pair

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i].difficulty < ps[j].difficulty || (ps[i].difficulty == ps[j].difficulty && ps[i].profit < ps[j].profit)
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
