package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		segs := make([][]int, n)
		for i := 0; i < n; i++ {
			segs[i] = readNNums(reader, 4)
		}
		m := readNum(reader)
		queries := readNNums(reader, m)
		res := solve(segs, queries)
		s := fmt.Sprintf("%v", res)
		s = s[1 : len(s)-1]
		buf.WriteString(fmt.Sprintf("%s\n", s))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const inf = 1 << 31

func solve(segs [][]int, queries []int) []int {
	n := len(segs)
	m := len(queries)
	events := make([]Event, 0, 2*n+m)
	for id, seg := range segs {
		l, b := seg[0], seg[3]
		events = append(events, Event{l, id, 0})
		events = append(events, Event{b, id, 2})
	}

	for i, x := range queries {
		events = append(events, Event{x, i, 1})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].pos > events[j].pos || events[i].pos == events[j].pos && events[i].tp > events[j].tp
	})

	ans := make([]*Item, n)
	for i := 0; i < n; i++ {
		ans[i] = new(Item)
		ans[i].id = i
		ans[i].priority = segs[i][3]
	}

	res := make([]int, m)

	pq := make(PriorityQueue, 0, n)

	for _, evt := range events {
		id := evt.id
		if evt.tp == 2 {
			if pq.Len() > 0 {
				ans[id].priority = max(ans[id].priority, pq[0].priority)
			}
			heap.Push(&pq, ans[id])
		} else if evt.tp == 0 {
			// leave
			pq.update(ans[id], inf)
			heap.Pop(&pq)
		} else {
			// query
			res[id] = queries[id]
			if pq.Len() > 0 {
				res[id] = max(res[id], pq[0].priority)
			}
		}
	}

	return res
}

func solve1(segs [][]int, queries []int) []int {
	n := len(segs)
	ss := make([]Segment, n)
	for i := 0; i < n; i++ {
		ss[i] = Segment{segs[i][0], segs[i][1], segs[i][2], segs[i][3], i}
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].l < ss[j].l || ss[i].l == ss[j].l && ss[i].b < ss[j].b
	})

	set := NewUFSet(n)
	// 这个地方联合的不大对，如果到了区间i, a[i] <= b[j] 且 l[j] <= b[i],那么从i，就可以到j
	for i := 0; i < n; {
		j := i
		b := ss[j].b
		for i < n && ss[i].l <= b {
			set.Union(ss[j].id, ss[i].id)
			b = max(ss[i].b, b)
			i++
		}
	}
	// 最远能够到达的b的位置
	far := make([]int, n)
	for i := 0; i < n; i++ {
		j := set.Find(ss[i].id)
		far[j] = max(far[j], ss[i].b)
	}

	events := make([]Event, 0, 2*n+len(queries))
	for i, cur := range segs {
		l, r := cur[0], cur[3]
		events = append(events, Event{l, i, 0})
		events = append(events, Event{r, i, 2})
	}

	for i := 0; i < len(queries); i++ {
		events = append(events, Event{queries[i], i + n, 1})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].pos > events[j].pos || events[i].pos == events[j].pos && events[i].tp > events[j].tp
	})

	res := make([]int, len(queries))

	items := make([]*Item, n)
	for i := 0; i < n; i++ {
		items[i] = new(Item)
		items[i].id = i
		items[i].priority = far[set.Find(i)]
	}

	pq := make(PriorityQueue, 0, n)

	for _, evt := range events {
		if evt.tp == 2 {
			// r event
			heap.Push(&pq, items[evt.id])
		} else if evt.tp == 0 {
			// l event
			pq.update(items[evt.id], inf)
			heap.Pop(&pq)
		} else {
			// query event
			x := evt.pos
			i := evt.id - n
			res[i] = x
			if pq.Len() > 0 {
				res[i] = max(res[i], pq[0].priority)
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

type Segment struct {
	l, r int
	a, b int
	id   int
}

type Event struct {
	pos int
	id  int
	tp  int
}

type UFSet struct {
	arr  []int
	cnt  []int
	size int
}

func NewUFSet(n int) UFSet {
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	return UFSet{arr, cnt, n}
}

func (uf *UFSet) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UFSet) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}

	if uf.cnt[pa] < uf.cnt[pb] {
		pa, pb = pb, pa
	}
	uf.cnt[pa] += uf.cnt[pb]
	uf.arr[pb] = pa
	uf.size--
	return true
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

func (pq *PriorityQueue) update(it *Item, p int) {
	it.priority = p
	heap.Fix(pq, it.index)
}
