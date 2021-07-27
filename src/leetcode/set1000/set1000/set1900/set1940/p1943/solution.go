package p1943

import "sort"

func splitPainting(segments [][]int) [][]int64 {
	n := len(segments)
	events := make([]Event, 2*n)
	for i := 0; i < n; i++ {
		seg := segments[i]
		events[2*i] = Event{seg[0], 1, seg[2]}
		events[2*i+1] = Event{seg[1], -1, seg[2]}
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].pos < events[j].pos || events[i].pos == events[j].pos && events[i].tp < events[j].tp
	})
	res := make([][]int64, 0, n)
	var pos int
	var sum int64
	for i := 0; i < 2*n; i++ {
		cur := events[i]
		if cur.tp == -1 {
			if pos < cur.pos {
				res = append(res, []int64{int64(pos), int64(cur.pos), sum})
			}
			sum -= int64(cur.color)
			pos = cur.pos
		} else {
			if pos > 0 {
				if pos < cur.pos && sum > 0 {
					res = append(res, []int64{int64(pos), int64(cur.pos), sum})
				}
			}
			sum += int64(cur.color)
			pos = cur.pos
		}
	}
	return res
}

type Event struct {
	pos   int
	tp    int
	color int
}
