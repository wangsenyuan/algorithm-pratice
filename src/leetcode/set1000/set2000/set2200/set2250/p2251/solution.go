package p2251

import "sort"

func fullBloomFlowers(flowers [][]int, persons []int) []int {
	events := make([]Event, 2*len(flowers)+len(persons))
	var p int
	for i := 0; i < len(flowers); i++ {
		events[p] = Event{flowers[i][0], 0, i}
		p++
		events[p] = Event{flowers[i][1], 2, i}
		p++
	}

	for i, cur := range persons {
		events[p] = Event{cur, 1, i}
		p++
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Less(events[j])
	})

	ans := make([]int, len(persons))
	var cnt int
	for i := 0; i < len(events); i++ {
		if events[i].kind == 0 {
			cnt++
		} else if events[i].kind == 2 {
			cnt--
		} else {
			ans[events[i].index] = cnt
		}
	}
	return ans
}

type Event struct {
	time  int
	kind  int
	index int
}

func (this Event) Less(that Event) bool {
	return this.time < that.time || this.time == that.time && this.kind < that.kind
}
