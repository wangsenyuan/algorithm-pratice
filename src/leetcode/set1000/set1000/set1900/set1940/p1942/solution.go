package p1942

import "sort"

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	cnt := make([]int, 4*n)

	var init func(i int, l, r int)
	init = func(i int, l, r int) {
		if l == r {
			cnt[i] = 1
			return
		}
		mid := (l + r) / 2
		init(2*i, l, mid)
		init(2*i+1, mid+1, r)
		cnt[i] = cnt[2*i] + cnt[2*i+1]
	}

	init(1, 0, n-1)

	var get func(i int, l, r int, xth int) int
	get = func(i int, l, r int, xth int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if cnt[2*i] >= xth {
			return get(2*i, l, mid, xth)
		}
		return get(2*i+1, mid+1, r, xth-cnt[2*i])
	}

	var set func(i int, l, r int, x int, v int)

	set = func(i int, l, r int, x int, v int) {
		if l == r {
			cnt[i] += v
			return
		}
		mid := (l + r) / 2
		if x <= mid {
			set(2*i, l, mid, x, v)
		} else {
			set(2*i+1, mid+1, r, x, v)
		}
		cnt[i] = cnt[2*i] + cnt[2*i+1]
	}

	events := make([]Event, 2*n)

	for i := 0; i < n; i++ {
		cur := times[i]
		events[2*i] = Event{cur[0], 0, i}
		events[2*i+1] = Event{cur[1], -1, i}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].time < events[j].time || events[i].time == events[j].time && events[i].tp < events[j].tp
	})

	seat := make([]int, n)

	for _, event := range events {
		if event.tp == 0 {
			// take a seat
			x := get(1, 0, n-1, 1)
			if event.index == targetFriend {
				return x
			}
			seat[event.index] = x
			set(1, 0, n-1, x, -1)
		} else {
			set(1, 0, n-1, seat[event.index], 1)
		}
	}
	return -1
}

type Event struct {
	time  int
	tp    int
	index int
}
