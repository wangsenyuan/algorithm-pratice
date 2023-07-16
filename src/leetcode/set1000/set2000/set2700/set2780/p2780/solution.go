package p2780

import "sort"

func maximumBeauty(nums []int, k int) int {
	type Event struct {
		pos int
		tp  int
	}
	evts := make([]Event, len(nums)*2)
	for i, num := range nums {
		evts[i*2] = Event{num - k, 0}
		evts[i*2+1] = Event{num + k, 1}
	}

	sort.Slice(evts, func(i, j int) bool {
		return evts[i].pos < evts[j].pos || evts[i].pos == evts[j].pos && evts[i].tp < evts[j].tp
	})

	var open int
	var ans int
	for _, evt := range evts {
		if evt.tp == 0 {
			open++
		} else {
			open--
		}
		if ans < open {
			ans = open
		}
	}
	return ans
}
