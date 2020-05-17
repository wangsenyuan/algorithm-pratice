package p1449

import "sort"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var res int

	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			res++
		}
	}
	return res
}

func busyStudent1(startTime []int, endTime []int, queryTime int) int {
	n := len(startTime)
	events := make([]Pair, 2*n+1)
	for i := 0; i < n; i++ {
		events[2*i] = Pair{startTime[i], 1}
		events[2*i+1] = Pair{endTime[i], -1}
	}
	events[2*n] = Pair{queryTime, 0}

	sort.Sort(Pairs(events))

	var cnt int

	for i := 0; i < len(events); i++ {
		evt := events[i]
		if evt.second == 0 {
			return cnt
		}
		cnt += evt.second
	}
	return 0
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first || this[i].first == this[j].first && this[i].second > this[j].second
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
