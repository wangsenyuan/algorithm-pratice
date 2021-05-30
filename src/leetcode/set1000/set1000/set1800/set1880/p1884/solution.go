package p1884

import (
	"container/heap"
)

func assignTasks(servers []int, tasks []int) []int {
	pq := make(Servers, len(servers))
	for i := 0; i < len(servers); i++ {
		cur := new(Server)
		cur.pos = i
		cur.weight = servers[i]
		cur.index = i
		pq[i] = cur
	}

	heap.Init(&pq)
	ans := make([]int, len(tasks))
	events := make(Events, 0, len(tasks))
	var time int
	var i int
	for i < len(tasks) {
		time = max(time, i)
		// 所有在time之前free的server，重新放入pq
		for len(events) > 0 && events[0].freeAt <= time {
			cur := heap.Pop(&events).(*Event)
			sv := new(Server)
			sv.pos = cur.server
			sv.weight = servers[cur.server]
			heap.Push(&pq, sv)
		}
		if len(pq) == 0 {
			//  no free servers, wait to another free event to happen
			time = events[0].freeAt
			continue
		}
		sv := heap.Pop(&pq).(*Server)
		ans[i] = sv.pos
		evt := new(Event)
		evt.server = sv.pos
		evt.freeAt = time + tasks[i]
		heap.Push(&events, evt)
		i++
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Server struct {
	weight int
	pos    int
	index  int
}

type Servers []*Server

func (ss Servers) Len() int {
	return len(ss)
}

func (ss Servers) Less(i, j int) bool {
	return ss[i].weight < ss[j].weight || ss[i].weight == ss[j].weight && ss[i].pos < ss[j].pos
}

func (ss Servers) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
	ss[i].index = i
	ss[j].index = j
}

func (ss *Servers) Push(item interface{}) {
	n := len(*ss)
	s := item.(*Server)
	s.index = n
	*ss = append(*ss, s)
}

func (ss *Servers) Pop() interface{} {
	old := *ss
	n := len(old)
	res := old[n-1]
	res.index = -1
	*ss = old[:n-1]
	return res
}

type Event struct {
	server int
	freeAt int
	index  int
}

type Events []*Event

func (ss Events) Len() int {
	return len(ss)
}

func (ss Events) Less(i, j int) bool {
	return ss[i].freeAt < ss[j].freeAt
}

func (ss Events) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
	ss[i].index = i
	ss[j].index = j
}

func (ss *Events) Push(item interface{}) {
	n := len(*ss)
	s := item.(*Event)
	s.index = n
	*ss = append(*ss, s)
}

func (ss *Events) Pop() interface{} {
	old := *ss
	n := len(old)
	res := old[n-1]
	res.index = -1
	*ss = old[:n-1]
	return res
}
