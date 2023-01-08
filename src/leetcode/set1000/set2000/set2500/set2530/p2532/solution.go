package p2532

import (
	"container/heap"
)

func findCrossingTime(n int, k int, time [][]int) int {

	right := make(PQ, 0, k)
	left := make(PQ, 0, k)

	for i := 0; i < k; i++ {
		w := Worker{i, time[i][0] + time[i][2]}
		heap.Push(&left, w)
	}

	events := make(Events, 0, n)

	var cur int
	var res int

	for n > 0 || events.Len() > 0 {
		// if right que not empty, let him move to left first
		// those workers create events with rightToLeft + putNew timestamp
		// and the time bridge free is sum(right_to_left) + cur
		// for thoses events <= free time, put them in a priority_queue
		// then pop less-effecient worker from the queue
		// move to left
		if (left.Len() == 0 || n == 0) && right.Len() == 0 && events.Len() > 0 {
			cur = max(cur, events[0].priority)
		}

		for events.Len() > 0 && events[0].priority <= cur {
			// cross the bridge
			e := heap.Pop(&events).(Event)
			i := e.id
			s := e.side
			w := Worker{i, time[i][0] + time[i][2]}
			if s == 0 {
				heap.Push(&left, w)
			} else {
				heap.Push(&right, w)
			}
		}

		if right.Len() > 0 {
			w := heap.Pop(&right).(Worker)
			i := w.id
			// next event is, cross the bridge, put object at new warehouse, and reach left side of the bridge
			e := Event{i, cur + time[i][2] + time[i][3], 0}
			heap.Push(&events, e)
			// right to left
			cur += time[i][2]
			res = max(res, cur)
			continue
		}

		if left.Len() > 0 {
			w := heap.Pop(&left).(Worker)
			i := w.id

			if n > 0 {
				n--
				e := Event{i, cur + time[i][0] + time[i][1], 1}
				heap.Push(&events, e)
				cur += time[i][0]
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Worker struct {
	id       int
	priority int
}

func (this Worker) LessEfficientThat(that Worker) bool {
	return this.priority > that.priority || this.priority == that.priority && this.id > that.id
}

type PQ []Worker

func (this PQ) Len() int {
	return len(this)
}

func (this PQ) Less(i, j int) bool {
	return this[i].LessEfficientThat(this[j])
}

func (this PQ) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *PQ) Push(x interface{}) {
	*this = append(*this, x.(Worker))
}

func (this *PQ) Pop() interface{} {
	old := *this
	n := len(old)
	res := old[n-1]
	*this = old[:n-1]
	return res
}

type Event struct {
	id       int
	priority int
	side     int
}

type Events []Event

func (this Events) Len() int {
	return len(this)
}

func (this Events) Less(i, j int) bool {
	return this[i].priority < this[j].priority
}

func (this Events) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *Events) Push(x interface{}) {
	*this = append(*this, x.(Event))
}

func (this *Events) Pop() interface{} {
	old := *this
	n := len(old)
	res := old[n-1]
	*this = old[:n-1]
	return res
}
